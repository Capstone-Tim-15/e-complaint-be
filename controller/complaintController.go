package controller

import (
	"context"
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	res "ecomplaint/utils/response"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ComplaintController interface {
	CreateComplaintController(ctx echo.Context) error
	GetComplaintController(ctx echo.Context) error
	GetComplaintsByStatusSolved(ctx echo.Context) error
	GetComplaintsController(ctx echo.Context) error
	UpdateComplaintController(ctx echo.Context) error
	DeleteComplaintController(ctx echo.Context) error
}

type ComplaintControllerImpl struct {
	ComplaintService service.ComplaintService
}

func NewComplaintController(ComplaintService service.ComplaintService) *ComplaintControllerImpl {
	return &ComplaintControllerImpl{ComplaintService: ComplaintService}
}

func (c *ComplaintControllerImpl) CreateComplaintController(ctx echo.Context) error {
	complaintCreateRequest := web.ComplaintCreateRequest{}
	err := ctx.Bind(&complaintCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := claims["id"].(string)

	complaintCreateRequest.User_ID = ID

	fileHeader, err := ctx.FormFile("attachment")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Missing Attachment"))
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error Opening File"))
	}
	defer file.Close()

	cldService, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error Initializing Cloudinary"))
	}

	uploadParams := uploader.UploadParams{}
	resp, err := cldService.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error Uploading File to Cloudinary"))
	}

	fileName := path.Base(resp.SecureURL)

	complaintCreateRequest.ImageUrl = fileName

	complaintCreateRequest.Status = "SEND"

	result, err := c.ComplaintService.CreateComplaint(ctx, complaintCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "error when creating complaint") {
			return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Complaint Failed"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Complaint Error"))
	}

	response := res.ComplaintDomaintoComplaintResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Complaint", response))
}

func (c *ComplaintControllerImpl) GetComplaintController(ctx echo.Context) error {
	idQuery := ctx.QueryParam("id")
	statusQuery := ctx.QueryParam("status")
	categoryQuery := ctx.QueryParam("category")
	limitQuery, _ := strconv.Atoi(ctx.QueryParam("limit"))

	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := claims["id"].(string)

	if idQuery != "" {
		result, err := c.ComplaintService.FindById(ctx, idQuery)
		if err != nil {
			if strings.Contains(err.Error(), "complaint not found") {
				return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Complaint Not Found"))
			}

			return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Complaint By ID Error"))
		}

		response := res.FindComplaintDomaintoComplaintResponse(result)

		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Complaint By Id", response))
	}

	if statusQuery != "" {
		result, err := c.ComplaintService.FindByStatusUser(ctx, statusQuery, ID)
		if err != nil {
			if strings.Contains(err.Error(), "complaint not found") {
				return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Complaint Not Found"))
			}

			return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Complaint By Status Error"))
		}

		response := res.FindStatusComplaintDomaintoComplaintResponse(result)

		return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Complaint By Status", response))
	}

	result, totalCount, err := c.ComplaintService.FindByCategory(ctx, categoryQuery, int64(limitQuery))
	if err != nil {
		if strings.Contains(err.Error(), "complaint not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Complaint Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Complaint By Category Error"))
	}

	response := res.FindCategoryComplaintDomaintoComplaintResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponsePage("Successfully Get Complaint Data", 0, limitQuery, totalCount, response))
}

func (c *ComplaintControllerImpl) GetComplaintsByStatusSolved(ctx echo.Context) error {
	statusQuery := "SOLVED"
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}
	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	result, totalCount, err := c.ComplaintService.FindByStatus(ctx, statusQuery, page, limit)
	if err != nil {
		if strings.Contains(err.Error(), "complaint not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Complaint Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Complaint By Status Solved Error"))
	}

	response := res.FindStatusComplaintDomaintoComplaintResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponsePage("Successfully Get Complaint By Status Solved", page, limit, totalCount, response))
}

func (c *ComplaintControllerImpl) GetComplaintsController(ctx echo.Context) error {
	statusQuery := ctx.QueryParam("status")
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize := 10

	if statusQuery == "" {
		result, totalCount, err := c.ComplaintService.FindAll(ctx, page, pageSize)
		if err != nil {
			if strings.Contains(err.Error(), "complaints not found") {
				return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Complaints Not Found"))
			}
	
			return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Complaints Data Error"))
		}
	
		response := res.ConvertComplaintResponse(result)
	
		return ctx.JSON(http.StatusOK, helper.SuccessResponsePage("Successfully Get Complaint Data", page, pageSize, totalCount, response))
	}

	result, totalCount, err := c.ComplaintService.FindByStatus(ctx, statusQuery, page, pageSize)
	if err != nil {
		if strings.Contains(err.Error(), "complaint not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Complaint Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Complaint By Status Error"))
	}

	response := res.FindStatusComplaintDomaintoComplaintResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponsePage("Successfully Get Complaint By Status", page, pageSize, totalCount, response))
}

func (c *ComplaintControllerImpl) UpdateComplaintController(ctx echo.Context) error {
	complaintUpdateRequest := web.ComplaintUpdateRequest{}
	err := ctx.Bind(&complaintUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	complaint_id := ctx.QueryParam("complaint_id")

	result, err := c.ComplaintService.UpdateComplaint(ctx, complaint_id, complaintUpdateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "complaint not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Complaint Not Found"))
		}

		if strings.Contains(err.Error(), "error when updating complaint") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Update Complaint Error"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Complaint Error"))
	}

	response := res.ComplaintDomaintoComplaintResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Update Complaint", response))
}

func (c *ComplaintControllerImpl) DeleteComplaintController(ctx echo.Context) error {
	complaint_id := ctx.QueryParam("complaint_id")

	err := c.ComplaintService.DeleteComplaint(ctx, complaint_id)
	if err != nil {
		if strings.Contains(err.Error(), "complaint not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Complaint Not Found"))
		}

		if strings.Contains(err.Error(), "error when deleting complaint") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Delete Complaint Error"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Complaint Error"))
	}

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Delete Complaint", nil))
}
