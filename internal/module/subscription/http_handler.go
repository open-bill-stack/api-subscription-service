package subscription

import (
	"api-subscription-service/internal/module/subscription/structure"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

type HttpHandle struct {
	log     *zap.Logger
	service *Service
}

func NewHttpHandler(p Params) (HttpResult, error) {
	return HttpResult{
		Router: &HttpHandle{
			log:     p.Log,
			service: p.Service,
		},
	}, nil
}

func (h *HttpHandle) Register(app *fiber.App) {
	groupTariff := app.Group("/subscription")
	groupTariff.Post("", h.CreateSubscription)
	groupTariff.Get("", h.ListSubscriptions)
	groupTariff.Get("/:id", h.GetSubscription)
	groupTariff.Put("/:id", h.UpdateSubscription)
	groupTariff.Delete("/:id", h.DeleteSubscription)
}

func (h *HttpHandle) CreateSubscription(c *fiber.Ctx) error {
	var req structure.CreateSubscription
	var validate *validator.Validate
	validate = validator.New(validator.WithRequiredStructEnabled())

	// Парсимо JSON з тіла запиту
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Перевіряємо валідацію структури
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var userID, errParseUserID = uuid.Parse(req.UserID)
	if errParseUserID != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}
	var tariffID, errParseTariffID = uuid.Parse(req.TariffID)
	if errParseTariffID != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	tariff := Subscription{
		UserID: pgtype.UUID{
			Bytes: userID,
			Valid: true,
		},
		TariffID: pgtype.UUID{
			Bytes: tariffID,
			Valid: true,
		},
	}
	createSubscription, err := h.service.Create(c.Context(), &tariff)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(structure.ResponseSubscription{
		ID:       createSubscription.ID.String(),
		UserID:   createSubscription.UserID.String(),
		TariffID: createSubscription.TariffID.String(),
	})
}
func (h *HttpHandle) GetSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	var subscriptionID, errParse = uuid.Parse(id)
	if errParse != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}
	subscription, err := h.service.GetByID(c.Context(), subscriptionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if subscription == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{})
	}
	return c.Status(fiber.StatusOK).JSON(structure.ResponseSubscription{
		ID:       subscription.ID.String(),
		UserID:   subscription.UserID.String(),
		TariffID: subscription.TariffID.String(),
	})
}
func (h *HttpHandle) ListSubscriptions(c *fiber.Ctx) error {
	subscriptions, err := h.service.List(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var structSubscriptions = make([]structure.ResponseSubscription, 0)
	for _, item := range subscriptions {
		structSubscriptions = append(structSubscriptions, structure.ResponseSubscription{
			ID:       item.ID.String(),
			UserID:   item.UserID.String(),
			TariffID: item.TariffID.String(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(structSubscriptions)
}
func (h *HttpHandle) UpdateSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	var subscriptionID, errParse = uuid.Parse(id)
	if errParse != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}
	var req structure.CreateSubscription
	var validate *validator.Validate
	validate = validator.New(validator.WithRequiredStructEnabled())

	// Парсимо JSON з тіла запиту
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Перевіряємо валідацію структури
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var userID, errParseUserID = uuid.Parse(req.UserID)
	if errParseUserID != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}
	var tariffID, errParseTariffID = uuid.Parse(req.TariffID)
	if errParseTariffID != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	subscription := Subscription{
		ID: pgtype.UUID{
			Bytes: subscriptionID,
			Valid: true,
		},
		UserID: pgtype.UUID{
			Bytes: userID,
			Valid: true,
		},
		TariffID: pgtype.UUID{
			Bytes: tariffID,
			Valid: true,
		},
	}
	createSubscription, err := h.service.UpdateByID(c.Context(), &subscription)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(structure.ResponseSubscription{
		ID:       createSubscription.ID.String(),
		UserID:   createSubscription.UserID.String(),
		TariffID: createSubscription.TariffID.String(),
	})
}
func (h *HttpHandle) DeleteSubscription(c *fiber.Ctx) error {
	id := c.Params("id")
	var subscriptionID, errParse = uuid.Parse(id)
	if errParse != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}
	status, err := h.service.DeleteByID(c.Context(), subscriptionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if !status {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Subscription not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Subscription deleted successfully",
	})
}
