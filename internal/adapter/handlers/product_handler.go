package handlers

import (

	"sistemaLogistica/internal/controllers"
	"sistemaLogistica/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	handler controllers.ProductUsecase
}

func (p ProductHandler) HandlerPost(ctx *fiber.Ctx) error {
	var info entity.Product
	err := ctx.BodyParser(&info)
	if err != nil {
		return err
	}

	product, err := p.handler.Create(info)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(product)
}

func (p ProductHandler) HandlerGet(ctx *fiber.Ctx) error {
	var info entity.Product
	err := ctx.BodyParser(&info)
	if err != nil {
		return err
	}

	product, err := p.handler.Read(info)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(product)
}

func (p ProductHandler) HandlerPut(ctx *fiber.Ctx) error {
	var info entity.Product
	err := ctx.BodyParser(&info)
	if err != nil {
		return err
	}

	product, err := p.handler.Update(info)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(product)
}

func (p ProductHandler) HandlerDelete(ctx *fiber.Ctx) error {
	var info entity.Product
	err := ctx.BodyParser(&info)
	if err != nil {
		return err
	}

	err = p.handler.Delete(info)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(info)
}

func NewProductHandler(h controllers.ProductUsecase) ProductHandler {
	return ProductHandler{handler: h}
}
