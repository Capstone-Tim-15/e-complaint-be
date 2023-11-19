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
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/labstack/echo/v4"
)

type ComplaintController interface {
	CreateComplaintController(ctx echo.Context) error
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

	// user := ctx.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// ID := claims["id"].(string)

	complaintCreateRequest.User_ID = "dadas"

	fileHeader, err := ctx.FormFile("attachment")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Missing attachment"))
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error opening file"))
	}
	defer file.Close()

	cldService, err := cloudinary.NewFromURL(fmt.Sprintf(os.Getenv("CLOUDINARY_URL"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET")))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error initializing Cloudinary"))
	}

	uploadParams := uploader.UploadParams{}
	resp, err := cldService.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error uploading file to Cloudinary"))
	}

	complaintCreateRequest.Attachment = resp.SecureURL

	result, err := c.ComplaintService.CreateComplaint(ctx, complaintCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "error when creating user") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Create Complaint Error"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign Up Error"))
	}

	response := res.ComplaintDomaintoComplaintResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Complaint", response))
}
