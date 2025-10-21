package mqtt

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
}

func NewHandler(logger *zap.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func RegisterHandlers(client mqtt.Client, handler *Handler) {
	client.Subscribe("site/+/node/+/telemetry", 0, handler.handleNodeTelemetry)
	client.Subscribe("site/+/coord/+/telemetry", 0, handler.handleCoordTelemetry)
	client.Subscribe("site/+/coord/+/mmwave", 0, handler.handleCoordMMWave)
}

func (h *Handler) defaultMessageHandler(client mqtt.Client, msg mqtt.Message) {
	log.Default()
	h.logger.Info("Received message: %s from topic: %s\n", zap.ByteString("payload", msg.Payload()), zap.String("topic", msg.Topic()))
}

func (h *Handler) connectHandler(client mqtt.Client) {
	h.logger.Info("Connected")
}

func (h *Handler) connectLostHandler(client mqtt.Client, err error) {
	h.logger.Error("Connect lost: %v", zap.Error(err))
}

func (h *Handler) handleCoordTelemetry(client mqtt.Client, msg mqtt.Message) {
	h.logger.Info("Received coord telemetry")
}

func (h *Handler) handleNodeTelemetry(client mqtt.Client, msg mqtt.Message) {
	h.logger.Info("Received node telemetry")
}

func (h *Handler) handleCoordMMWave(client mqtt.Client, msg mqtt.Message) {
	h.logger.Info("Received coord mmwave")
}
