package controllers

import (
	"context"
	"net/http"
	myerrors "news/internal/errors"
	"news/internal/logger"
	"news/internal/models"
	"news/internal/request"
	"news/internal/response"
	"news/internal/services"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type NewsArticleController struct {
	service *services.NewsArticleService
}

// NewNewsArticleController Конструктор контролера новин.
func NewNewsArticleController(s *services.NewsArticleService) NewsArticleController {
	return NewsArticleController{
		service: s,
	}
}

// GetNewsArticles Обробник список новин.
func (controller *NewsArticleController) GetNewsArticles(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	loc := c.Query("loc")
	if !request.LocInWhiteList(loc) {
		loc = request.DefaultLoc
	}

	articles, err := controller.service.List(ctx, c.Queries(), loc)
	if err != nil {
		r := response.NewResponse(fiber.StatusServiceUnavailable, err.Error(), nil)
		return c.Status(fiber.StatusServiceUnavailable).JSON(r)
	}
	return c.JSON(response.NewResponse(fiber.StatusOK, "", articles))
}

// GetNewsArticle Обробник інформація про новину.
func (controller *NewsArticleController) GetNewsArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	loc := c.Query("loc")
	if !request.LocInWhiteList(loc) {
		loc = request.DefaultLoc
	}

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

	article, err := controller.service.One(ctx, id, loc)
	if err != nil {
		r := response.NewResponse(fiber.StatusServiceUnavailable, err.Error(), nil)
		return c.Status(fiber.StatusServiceUnavailable).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", article))
}

// AddNewsArticle Обробник створення нової новини.
func (controller *NewsArticleController) AddNewsArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newsArticleDTO models.NewsArticleDTO
	var aRequest request.NewsArticleRequest

	loc := c.Query("loc")
	if !request.LocInWhiteList(loc) {
		loc = request.DefaultLoc
	}

	if err := c.BodyParser(&aRequest); err != nil {
		logger.Log().Info(err)
		return err
	}

	if err := aRequest.Validate(); err != nil {
		logger.Log().Info(strings.Join(err, ";"))
		r := response.NewResponse(fiber.StatusBadRequest, strings.Join(err, ";"), nil)
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	aRequest.Fill(&newsArticleDTO)
	dto, err := controller.service.Create(ctx, newsArticleDTO, loc)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}

// UpdateNewsArticle Обробник оновлення новини.
func (controller *NewsArticleController) UpdateNewsArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	loc := c.Query("loc")
	if !request.LocInWhiteList(loc) {
		loc = request.DefaultLoc
	}

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

	var newsArticleDTO models.NewsArticleDTO
	var gRequest request.NewsArticleRequest

	if err := c.BodyParser(&gRequest); err != nil {
		logger.Log().Info(err)
		return err
	}
	if err := gRequest.Validate(); err != nil {
		logger.Log().Info(strings.Join(err, ";"))
		r := response.NewResponse(fiber.StatusBadRequest, strings.Join(err, ";"), nil)
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	gRequest.Fill(&newsArticleDTO)
	newsArticleDTO.ID = id
	dto, err := controller.service.Update(ctx, newsArticleDTO, loc)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}

// TrashNewsArticle Обробник м'яке видалення новини.
func (controller *NewsArticleController) TrashNewsArticle(c *fiber.Ctx) error {
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

	var dto = &models.NewsArticleDTO{ID: id}
	// @todo розділити методи репозиторія для пошуку запису з перекладами та без
	dto, err = controller.service.Trash(ctx, dto, request.DefaultLoc)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(http.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}

// RecoverNewsArticle Обробник відновлення новини після м'якого видалення.
func (controller *NewsArticleController) RecoverNewsArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		r := response.NewResponse(fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	exists, err := controller.service.ExistsUnscoped(ctx, id)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	if !exists {
		r := response.NewResponse(fiber.StatusNotFound, myerrors.ResourceNotFound, nil)
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	var dto = &models.NewsArticleDTO{ID: id}
	// @todo розділити методи репозиторія для пошуку запису з перекладами та без
	dto, err = controller.service.Recover(ctx, dto, request.DefaultLoc)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(http.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}

// DeleteNewsArticle Обробник остаточного видалення новини.
func (controller *NewsArticleController) DeleteNewsArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id, err := c.ParamsInt("id")
	if err != nil {
		r := response.NewResponse(fiber.StatusBadRequest, err.Error(), nil)
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	exists, err := controller.service.ExistsUnscoped(ctx, id)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	if !exists {
		r := response.NewResponse(fiber.StatusNotFound, myerrors.ResourceNotFound, nil)
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	var dto = &models.NewsArticleDTO{ID: id}
	// @todo розділити методи репозиторія для пошуку запису з перекладами та без
	dto, err = controller.service.Delete(ctx, dto, request.DefaultLoc)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(http.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}
