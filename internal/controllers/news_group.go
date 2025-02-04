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

// NewsGroupController Контролер груп новин.
type NewsGroupController struct {
	service GroupsService
}

// NewNewsGroupController Конструктор контролера груп новин.
func NewNewsGroupController(s GroupsService) NewsGroupController {
	return NewsGroupController{
		service: s,
	}
}

// GetNewsGroups Обробник список груп новин.
// GetNewsGroups godoc
//
//	@Summary		список груп новин
//	@Description	список груп новин
//	@Tags			groups
//	@Accept			json
//	@Produce		json
//	@Param			locale	query		string	false	"string enums"	Enums(en, uk)	"локаль; за замовчуванням en"
//	@Success		200		{object}	response.DocGetNewsGroupsResponse200
//	@Failure		400		{object}	response.DocGetNewsGroupResponse400
//	@Failure		404		{object}	response.DocGetNewsGroupResponse400
//	@Failure		500		{object}	response.DocGetNewsGroupsResponse500
//	@Router			/groups [get]
func (controller *NewsGroupController) GetNewsGroups(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	groups, err := controller.service.List(ctx, c)
	if err != nil {
		r := response.NewResponse(fiber.StatusServiceUnavailable, err.Error(), nil)
		return c.Status(fiber.StatusServiceUnavailable).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", groups))
}

// GetNewsGroup Обробник інформація про групу новин.
// GetNewsGroup godoc
//
//	@Summary		Інформація про групу новин
//	@Description	Інформація про групу новин
//	@Tags			groups
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"uuid групи новин"
//	@Param			locale	query		string	false	"string enums"	Enums(en, uk)	"локаль; за замовчуванням en"
//	@Success		200		{object}	response.DocGetNewsGroupResponse200
//	@Failure		400		{object}	response.DocGetNewsGroupResponse400
//	@Failure		404		{object}	response.DocGetNewsGroupResponse404
//	@Failure		500		{object}	response.DocGetNewsGroupResponse500
//	@Router			/groups/{id} [get]
func (controller *NewsGroupController) GetNewsGroup(c *fiber.Ctx) error {
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

	group, err := controller.service.One(ctx, c, id)
	if err != nil {
		r := response.NewResponse(fiber.StatusServiceUnavailable, err.Error(), nil)
		return c.Status(fiber.StatusServiceUnavailable).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", group))
}

// AddNewsGroup Обробник створення нової групи новин.
// AddNewsGroup godoc
//
//	@Summary		Створення групи новин
//	@Description	Створення групи новин
//	@Tags			groups
//	@Accept			json
//	@Produce		json
//	@Param			locale	query		string	false	"string enums"	Enums(en, uk)	"локаль; за замовчуванням en"
//	@Success		200		{object}	response.DocGetNewsGroupResponse200
//	@Param          request body request.NewsGroupRequest true "news group request"
//	@Failure		400		{object}	response.DocGetNewsGroupResponse400
//	@Failure		404		{object}	response.DocGetNewsGroupResponse404
//	@Failure		500		{object}	response.DocGetNewsGroupResponse500
//	@Router			/groups [post]
func (controller *NewsGroupController) AddNewsGroup(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRequest request.NewsGroupRequest

	if err := c.BodyParser(&gRequest); err != nil {
		logger.Log().Info(err)
		return err
	}
	if err := gRequest.Validate(); err != nil {
		logger.Log().Info(strings.Join(err, ";"))
		r := response.NewResponse(fiber.StatusBadRequest, strings.Join(err, ";"), nil)
		return c.Status(fiber.StatusBadRequest).JSON(r)
	}

	dto, err := controller.service.Create(ctx, c, gRequest)
	if err != nil {
		r := response.NewResponse(fiber.StatusInternalServerError, err.Error(), nil)
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	return c.JSON(response.NewResponse(fiber.StatusOK, "", dto))
}

// UpdateNewsGroup Обробник оновлення групи новин.
// UpdateNewsGroup godoc
//
//	@Summary		Оновлення групи новин
//	@Description	Оновлення групи новин
//	@Tags			groups
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"uuid групи новин"
//	@Param			locale	query		string	false	"string enums"	Enums(en, uk)	"локаль; за замовчуванням en"
//	@Param          request body request.NewsGroupRequest true "news group request"
//	@Success		200		{object}	response.DocGetNewsGroupResponse200
//	@Failure		400		{object}	response.DocGetNewsGroupResponse400
//	@Failure		404		{object}	response.DocGetNewsGroupResponse404
//	@Failure		500		{object}	response.DocGetNewsGroupResponse500
//	@Router			/groups/{id} [put]
func (controller *NewsGroupController) UpdateNewsGroup(c *fiber.Ctx) error {
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

	var gRequest request.NewsGroupRequest

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

// TrashNewsGroup Обробник м'яке видалення групи новин.
// TrashNewsGroup godoc
//
//	@Summary		м'яке видалення групи новин
//	@Description	м'яке видалення групи новин
//	@Tags			groups
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"uuid групи новин"
//	@Success		200	{object}	response.DocGetNewsGroupResponse200
//	@Failure		400	{object}	response.DocGetNewsGroupResponse400
//	@Failure		404	{object}	response.DocGetNewsGroupResponse404
//	@Failure		500	{object}	response.DocGetNewsGroupResponse500
//	@Router			/groups/{id}/trash [patch]
func (controller *NewsGroupController) TrashNewsGroup(c *fiber.Ctx) error {
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

	return c.JSON(dto)
}

// RecoverNewsGroup Обробник відновлення групи новин після м'якого видалення.
// RecoverNewsGroup godoc
//
//	@Summary		відновлення групи новин після м'якого видалення
//	@Description	відновлення групи новин після м'якого видалення
//	@Tags			groups
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"uuid групи новин"
//	@Success		200	{object}	response.DocGetNewsGroupResponse200
//	@Failure		400	{object}	response.DocGetNewsGroupResponse400
//	@Failure		404	{object}	response.DocGetNewsGroupResponse404
//	@Failure		500	{object}	response.DocGetNewsGroupResponse500
//	@Router			/groups/{id}/recover [patch]
func (controller *NewsGroupController) RecoverNewsGroup(c *fiber.Ctx) error {
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

	return c.JSON(dto)
}

// DeleteNewsGroup Обробник остаточного видалення групи новин.
// DeleteNewsGroup godoc
//
//	@Summary		остаточне видалення групи новин
//	@Description	остаточне видалення групи новин
//	@Tags			groups
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"uuid групи новин"
//	@Success		200	{object}	response.DocGetNewsGroupResponse200
//	@Failure		400	{object}	response.DocGetNewsGroupResponse400
//	@Failure		404	{object}	response.DocGetNewsGroupResponse404
//	@Failure		500	{object}	response.DocGetNewsGroupResponse500
//	@Router			/groups/{id} [delete]
func (controller *NewsGroupController) DeleteNewsGroup(c *fiber.Ctx) error {
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

	return c.JSON(dto)
}
