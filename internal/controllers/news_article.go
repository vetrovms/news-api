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

// NewsArticleController Контролер новин.
type NewsArticleController struct {
	service NewsService
}

// NewNewsArticleController Конструктор контролера новин.
func NewNewsArticleController(s NewsService) NewsArticleController {
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
//	@Success		200		{object}	response.DocGetNewsArticlesResponse200
//	@Failure		500		{object}	response.DocGetNewsArticlesResponse500
//	@Router			/news [get]
func (controller *NewsArticleController) GetNewsArticles(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	articles, err := controller.service.List(ctx, c)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
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
//	@Param			id		path		string	true	"uuid новини"
//	@Param			locale	query		string	false	"string enums"	Enums(en, uk)	"локаль; за замовчуванням en"
//	@Success		200		{object}	response.DocGetNewsArticleResponse200
//	@Failure		400		{object}	response.DocGetNewsArticleResponse400
//	@Failure		404		{object}	response.DocGetNewsArticleResponse400
//	@Failure		500		{object}	response.DocGetNewsArticleResponse500
//	@Router			/news/{id} [get]
func (controller *NewsArticleController) GetNewsArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id := c.Params("id")

	exists, err := controller.service.Exists(ctx, id)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	if !exists {
		r := response.NewResponse(fiber.StatusNotFound, myerrors.ResourceNotFound, nil)
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	article, err := controller.service.One(ctx, c, id)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
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
//	@Success		200		{object}	response.DocGetNewsArticleResponse200
//	@Failure		400		{object}	response.DocGetNewsArticleResponse400
//	@Failure		500		{object}	response.DocGetNewsArticleResponse500
//	@Router			/news [post]
func (controller *NewsArticleController) AddNewsArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var aRequest request.NewsArticleRequest
	if err := c.BodyParser(&aRequest); err != nil {
		logger.Log().Info(err)
		return err
	}

	if err := aRequest.Validate(); err != nil {
		logger.Log().Info(strings.Join(err, ";"))
		r := response.NewResponse(fiber.StatusBadRequest, strings.Join(err, ";"), nil)
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	dto, err := controller.service.Create(ctx, c, aRequest)
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
//	@Param			id		path		string	true	"uuid новини"
//	@Param			locale	query		string	false	"string enums"	Enums(en, uk)	"локаль; за замовчуванням en"
//	@Param          request body request.NewsArticleRequest true "news article request"
//	@Success		200		{object}	response.DocGetNewsArticleResponse200
//	@Failure		400		{object}	response.DocGetNewsArticleResponse400
//	@Failure		404		{object}	response.DocGetNewsArticleResponse404
//	@Failure		500		{object}	response.DocGetNewsArticleResponse500
//	@Router			/news/{id} [put]
func (controller *NewsArticleController) UpdateNewsArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id := c.Params("id")

	exists, err := controller.service.Exists(ctx, id)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	if !exists {
		r := response.NewResponse(fiber.StatusNotFound, myerrors.ResourceNotFound, nil)
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

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

	dto, err := controller.service.Update(ctx, c, gRequest, id)
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
//	@Param			id	path		string	true	"uuid новини"
//	@Success		200	{object}	response.DocGetNewsArticleResponse200
//	@Failure		400	{object}	response.DocGetNewsArticleResponse400
//	@Failure		404	{object}	response.DocGetNewsArticleResponse404
//	@Failure		500	{object}	response.DocGetNewsArticleResponse500
//	@Router			/news/{id}/trash [patch]
func (controller *NewsArticleController) TrashNewsArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id := c.Params("id")

	exists, err := controller.service.Exists(ctx, id)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	if !exists {
		r := response.NewResponse(fiber.StatusNotFound, myerrors.ResourceNotFound, nil)
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	dto, err := controller.service.Trash(ctx, id, request.DefaultLoc)
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
//	@Param			id	path		string	true	"uuid новини"
//	@Success		200	{object}	response.DocGetNewsArticleResponse200
//	@Failure		400	{object}	response.DocGetNewsArticleResponse400
//	@Failure		404	{object}	response.DocGetNewsArticleResponse404
//	@Failure		500	{object}	response.DocGetNewsArticleResponse500
//	@Router			/news/{id}/recover [patch]
func (controller *NewsArticleController) RecoverNewsArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id := c.Params("id")

	exists, err := controller.service.ExistsUnscoped(ctx, id)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	if !exists {
		r := response.NewResponse(fiber.StatusNotFound, myerrors.ResourceNotFound, nil)
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	dto, err := controller.service.Recover(ctx, id, request.DefaultLoc)
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
//	@Param			id	path		string	true	"uuid новини"
//	@Success		200	{object}	response.DocGetNewsArticleResponse200
//	@Failure		400	{object}	response.DocGetNewsArticleResponse400
//	@Failure		404	{object}	response.DocGetNewsArticleResponse404
//	@Failure		500	{object}	response.DocGetNewsArticleResponse500
//	@Router			/news/{id} [delete]
func (controller *NewsArticleController) DeleteNewsArticle(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id := c.Params("id")

	exists, err := controller.service.ExistsUnscoped(ctx, id)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
	if !exists {
		r := response.NewResponse(fiber.StatusNotFound, myerrors.ResourceNotFound, nil)
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	dto, err := controller.service.Delete(ctx, id, request.DefaultLoc)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(http.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}
