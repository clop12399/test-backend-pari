package handlers

import (
	"strconv"
	"test-backend-pari/items"
	"test-backend-pari/items/entities"

	"github.com/gofiber/fiber/v2"
)

type ItemHandler struct {
	uc items.Usecase
}

func NewItemHandler(uc items.Usecase) items.Handler {
	return &ItemHandler{
		uc: uc,
	}
}

// Delete implements Handler.
func (h *ItemHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	err = h.uc.Delete(id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

// Index implements Handler.
func (u *ItemHandler) Index(c *fiber.Ctx) error {
	items, err := u.uc.GetAll()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(items)
}

// Show implements Handler.
func (h *ItemHandler) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	items, err := h.uc.GetById(id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(items)
}

// Store implements Handler.
func (h *ItemHandler) Store(c *fiber.Ctx) error {
	var request entities.CreateItemRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	err = request.Validate()
	if err != nil {
		return err
	}

	err = h.uc.Create(request)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

// Update implements Handler.
func (h *ItemHandler) Update(c *fiber.Ctx) error {
	var request entities.UpdateItemRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	err = request.Validate()
	if err != nil {
		return err
	}

	err = h.uc.Update(request, id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
