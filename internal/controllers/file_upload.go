package controllers

import (
	"context"
	"fmt"
	"net/http"
	"news/internal/config"
	myerrors "news/internal/errors"
	"news/internal/logger"
	"news/internal/models"
	"news/internal/request"
	"news/internal/response"
	"news/internal/services"
	"slices"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type FileUploadController struct {
	service *services.FileUploadService
}

// NewFileUploadController Конструктор контролера новин.
func NewFileUploadController(s *services.FileUploadService) FileUploadController {
	return FileUploadController{
		service: s,
	}
}

// GetFileUploads Обробник список файлів.
// GetFileUploads godoc
//	@Summary		список файлів
//	@Description	список файлів
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/files [get]
func (controller *FileUploadController) GetFileUploads(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	files, err := controller.service.List(ctx)
	if err != nil {
		r := response.NewResponse(fiber.StatusServiceUnavailable, err.Error(), nil)
		return c.Status(fiber.StatusServiceUnavailable).JSON(r)
	}
	return c.JSON(response.NewResponse(fiber.StatusOK, "", files))
}

// GetFileUpload Обробник інформація про файл.
// GetFileUpload godoc
//	@Summary		Інформація про файл
//	@Description	Інформація про файл
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id файла"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/files/{id} [get]
func (controller *FileUploadController) GetFileUpload(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		r := response.NewResponse(fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	exists, err := controller.service.Exists(ctx, id)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	if !exists {
		r := response.NewResponse(fiber.StatusNotFound, myerrors.ResourceNotFound, nil)
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	fileUpload, err := controller.service.One(ctx, id)
	if err != nil {
		r := response.NewResponse(fiber.StatusServiceUnavailable, err.Error(), nil)
		return c.Status(fiber.StatusServiceUnavailable).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", fileUpload))
}

// AddFileUpload Обробник створення файла.
// AddFileUpload godoc
//	@Summary		Створення файла
//	@Description	Створення файла
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/files [post]
func (controller *FileUploadController) AddFileUpload(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	file, err := c.FormFile("file")
	if err != nil {
		logger.Log().Info(err)
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	var fileRequest request.FileUploadRequest
	if err := c.BodyParser(&fileRequest); err != nil {
		logger.Log().Info(err)
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	allowedTypes := []string{"image/jpeg", "image/png"}
	if !slices.Contains(allowedTypes, file.Header.Get("Content-Type")) {
		logger.Log().Info(myerrors.WrongFileFormat)
		r := response.NewResponse(fiber.StatusBadRequest, myerrors.WrongFileFormat, nil)
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	if err := fileRequest.Validate(); err != nil {
		logger.Log().Info(strings.Join(err, ";"))
		r := response.NewResponse(fiber.StatusBadRequest, strings.Join(err, ";"), nil)
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	destination := fmt.Sprintf(config.NewEnv().UploadPath+"%s", file.Filename)
	if err := c.SaveFile(file, destination); err != nil {
		logger.Log().Info(err)
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	var fileUploadDto models.FileUploadDto
	fileRequest.File = file
	fileRequest.Fill(&fileUploadDto)
	dto, err := controller.service.Create(ctx, fileUploadDto)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}

// DeleteFileUpload Обробник видалення файла.
// DeleteFileUpload godoc
//	@Summary		видалення файла
//	@Description	видалення файла
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id файла"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/files/{id} [delete]
func (controller *FileUploadController) DeleteFileUpload(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		r := response.NewResponse(fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	exists, err := controller.service.Exists(ctx, id)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	if !exists {
		r := response.NewResponse(fiber.StatusNotFound, myerrors.ResourceNotFound, nil)
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	var dto = &models.FileUploadDto{ID: id}
	dto, err = controller.service.Delete(ctx, dto)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(http.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}
