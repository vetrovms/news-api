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

// NewsArticleController Контролер новин.
type NewsArticleController struct {
	service services.INewsService
}

// NewNewsArticleController Конструктор контролера новин.
func NewNewsArticleController(s services.INewsService) NewsArticleController {
	return NewsArticleController{
		service: s,
	}
}

// GetNewsArticles Обробник список новин.
// GetNewsArticles godoc
//
//	@Summary		Список новин
//	@Description	Отримати список новин
//	@Tags			news
//	@Accept			json
//	@Produce		json
//	@Param			locale	query		string	false	"string enums"	Enums(en, uk)	"локаль; за замовчуванням en"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		404		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/news [get]
func (controller *NewsArticleController) GetNewsArticles(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	locale := c.Query("locale")
	if !request.LocInWhiteList(locale) {
		locale = request.DefaultLoc
	}

	articles, err := controller.service.List(ctx, c.Queries(), locale)
	if err != nil {
		r := response.NewResponse(fiber.StatusServiceUnavailable, err.Error(), nil)
		return c.Status(fiber.StatusServiceUnavailable).JSON(r)
	}
	return c.JSON(response.NewResponse(fiber.StatusOK, "", articles))
}

// GetNewsArticle Обробник інформація про новину.
// GetNewsArticle godoc
//
//	@Summary		Інформація про новину
//	@Description	Інформація про новину
//	@Tags			news
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int		true	"id новини"
//	@Param			locale	query		string	false	"string enums"	Enums(en, uk)	"локаль; за замовчуванням en"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		404		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/news/{id} [get]
func (controller *NewsArticleController) GetNewsArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	locale := c.Query("locale")
	if !request.LocInWhiteList(locale) {
		locale = request.DefaultLoc
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

	article, err := controller.service.One(ctx, id, locale)
	if err != nil {
		r := response.NewResponse(fiber.StatusServiceUnavailable, err.Error(), nil)
		return c.Status(fiber.StatusServiceUnavailable).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", article))
}

// AddNewsArticle Обробник створення нової статті.
// AddNewsArticle godoc
//
//	@Summary		Створення новини
//	@Description	Створення новини
//	@Tags			news
//	@Accept			json
//	@Produce		json
//	@Param			locale	query		string	false	"string enums"	Enums(en, uk)	"локаль; за замовчуванням en"
//	@Param          request body request.NewsArticleRequest true "news article request"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		404		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/news [post]
func (controller *NewsArticleController) AddNewsArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newsArticleDTO models.NewsArticleDTO
	var aRequest request.NewsArticleRequest

	locale := c.Query("locale")
	if !request.LocInWhiteList(locale) {
		locale = request.DefaultLoc
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
	dto, err := controller.service.Create(ctx, newsArticleDTO, locale)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}

// UpdateNewsArticle Обробник оновлення новини.
// UpdateNewsArticle godoc
//
//	@Summary		Оновлення новини
//	@Description	Оновлення новини
//	@Tags			news
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int		true	"id новини"
//	@Param			locale	query		string	false	"string enums"	Enums(en, uk)	"локаль; за замовчуванням en"
//	@Param          request body request.NewsArticleRequest true "news article request"
//	@Success		200		{object}	response.Response
//	@Failure		400		{object}	response.Response
//	@Failure		404		{object}	response.Response
//	@Failure		500		{object}	response.Response
//	@Router			/news/{id} [put]
func (controller *NewsArticleController) UpdateNewsArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	locale := c.Query("locale")
	if !request.LocInWhiteList(locale) {
		locale = request.DefaultLoc
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
	dto, err := controller.service.Update(ctx, newsArticleDTO, locale)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}

// TrashNewsArticle Обробник м'яке видалення новини.
// TrashNewsArticle godoc
//
//	@Summary		м'яке видалення новини
//	@Description	м'яке видалення новини
//	@Tags			news
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id новини"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/news/{id}/trash [patch]
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
	//	@todo	розділити методи репозиторія для пошуку запису з перекладами та без
	dto, err = controller.service.Trash(ctx, dto, request.DefaultLoc)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(http.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}

// RecoverNewsArticle Обробник відновлення новини після м'якого видалення.
// RecoverNewsArticle godoc
//
//	@Summary		відновлення новини після м'якого видалення
//	@Description	відновлення новини після м'якого видалення
//	@Tags			news
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id новини"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/news/{id}/recover [patch]
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
	//	@todo	розділити методи репозиторія для пошуку запису з перекладами та без
	dto, err = controller.service.Recover(ctx, dto, request.DefaultLoc)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(http.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}

// DeleteNewsArticle Обробник остаточного видалення новини.
// DeleteNewsArticle godoc
//
//	@Summary		остаточне видалення новини
//	@Description	остаточне видалення новини
//	@Tags			news
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"id новини"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		404	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/news/{id} [delete]
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
	//	@todo	розділити методи репозиторія для пошуку запису з перекладами та без
	dto, err = controller.service.Delete(ctx, dto, request.DefaultLoc)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(http.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}
