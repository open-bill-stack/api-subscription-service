package subscription

import (
	"api-subscription-service/internal/module/subscription/structure"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/open-bill-stack/kinkajou"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ParamsRun struct {
	fx.In

	Log         *zap.Logger
	AQMPChannel *amqp.Channel
}

type EventHandle struct {
	log     *zap.Logger
	service *Service
}

func NewEventHandler(p Params) (EventResult, error) {
	return EventResult{
		Router: &EventHandle{
			log:     p.Log,
			service: p.Service,
		},
	}, nil
}

func (h *EventHandle) Register(app *kinkajou.App) {
	app.OnEvent("user.deleted", h.UserDeleted)
	app.OnEvent("tariff.deleted", h.TariffDeleted)
}

func (h *EventHandle) UserDeleted(message *amqp.Delivery) {
	var event structure.EventUser
	err := json.Unmarshal(message.Body, &event)
	if err != nil {
		_ = message.Reject(true)
	}

	var userID, errParse = uuid.Parse(event.UUID)
	if errParse != nil {
		_ = message.Reject(true)
	}

	if _, err := h.service.DeleteByUserID(context.Background(), userID); err != nil {
		_ = message.Reject(true)
	}
	_ = message.Ack(false)

}

func (h *EventHandle) TariffDeleted(message *amqp.Delivery) {
	var event structure.EventUser
	err := json.Unmarshal(message.Body, &event)
	if err != nil {
		_ = message.Reject(true)
	}

	var tariffID, errParse = uuid.Parse(event.UUID)
	if errParse != nil {
		_ = message.Reject(true)
	}

	if _, err := h.service.DeleteByTariffID(context.Background(), tariffID); err != nil {
		_ = message.Reject(true)
	}
	_ = message.Ack(false)

}
