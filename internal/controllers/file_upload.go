package controllers

import (
	"context"
	"net/http"
	myerrors "news/internal/errors"
	"news/internal/logger"
	"news/internal/request"
	"news/internal/response"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FileUploadController Контролер завантаження файлів.
type FileUploadController struct {
	service FilesService
}

// NewFileUploadController Конструктор контролера новин.
func NewFileUploadController(s FilesService) FileUploadController {
	return FileUploadController{
		service: s,
	}
}

// GetFileUploads Обробник список файлів.
// GetFileUploads godoc
//
//	@Summary		список файлів
//	@Description	список файлів
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.DocGetFileUploads200
//	@Failure		500	{object}	response.DocGetFileUpload500
//	@Router			/files [get]
func (controller *FileUploadController) GetFileUploads(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	files, err := controller.service.List(ctx)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	return c.JSON(response.NewResponse(fiber.StatusOK, "", files))
}

// GetFileUpload Обробник інформація про файл.
// GetFileUpload godoc
//
//	@Summary		Інформація про файл
//	@Description	Інформація про файл
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id файла"
//	@Success		200	{object}	response.DocGetFileUpload200
//	@Failure		400	{object}	response.DocGetFileUpload400
//	@Failure		404	{object}	response.DocGetFileUpload404
//	@Failure		500	{object}	response.DocGetFileUpload500
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
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", fileUpload))
}

// AddFileUpload Обробник створення файла.
// AddFileUpload godoc
//
//	@Summary		Створення файла
//	@Description	Створення файла
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Param          request body request.FileUploadRequest true "file upload request"
//	@Success		200	{object}	response.DocGetFileUpload200
//	@Failure		400	{object}	response.DocGetFileUpload400
//	@Failure		404	{object}	response.DocGetFileUpload404
//	@Failure		500	{object}	response.DocGetFileUpload500
//	@Router			/files [post]
func (controller *FileUploadController) AddFileUpload(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var fileRequest request.FileUploadRequest
	if err := c.BodyParser(&fileRequest); err != nil {
		logger.Log().Info(err)
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	if err := fileRequest.Validate(); err != nil {
		logger.Log().Info(strings.Join(err, ";"))
		r := response.NewResponse(fiber.StatusBadRequest, strings.Join(err, ";"), nil)
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	dto, err, code := controller.service.Create(ctx, c, fileRequest)
	if err != nil {
		r := response.NewResponse(code, err.Error(), nil)
		return c.Status(code).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}

// DeleteFileUpload Обробник видалення файла.
// DeleteFileUpload godoc
//
//	@Summary		видалення файла
//	@Description	видалення файла
//	@Tags			files
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id файла"
//	@Success		200	{object}	response.DocGetFileUpload200
//	@Failure		400	{object}	response.DocGetFileUpload400
//	@Failure		404	{object}	response.DocGetFileUpload404
//	@Failure		500	{object}	response.DocGetFileUpload500
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

	dto, err := controller.service.Delete(ctx, id)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(http.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}
