package controller

import (
	"context"
	"ecomplaint/model/web"
	"ecomplaint/service"
	"ecomplaint/utils/helper"
	res "ecomplaint/utils/response"
	"fmt"
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
	GetComplaintsController(ctx echo.Context) error
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
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Missing attachment"))
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error opening file"))
	}
	defer file.Close()

	cldService, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error initializing Cloudinary"))
	}

	uploadParams := uploader.UploadParams{}
	resp, err := cldService.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error uploading file to Cloudinary"))
	}

	fmt.Println(resp.SecureURL)

	fileName := path.Base(resp.SecureURL)

	complaintCreateRequest.ImageUrl = fileName

	result, err := c.ComplaintService.CreateComplaint(ctx, complaintCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "error when creating complaint") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Create Complaint Error"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign Up Error"))
	}

	response := res.ComplaintDomaintoComplaintResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Complaint", response))
}

func (c *ComplaintControllerImpl) GetComplaintController(ctx echo.Context) error {
	idQuery := ctx.QueryParam("id")

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

func (c *ComplaintControllerImpl) GetComplaintsController(ctx echo.Context) error {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize := 10

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
