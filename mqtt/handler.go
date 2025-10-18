package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func (h *handler) handleCoordTelemetry(client mqtt.Client, msg mqtt.Message) {
	h.logger.Info("Received coord telemetry")
}

func (h *handler) handleNodeTelemetry(client mqtt.Client, msg mqtt.Message) {
	h.logger.Info("Received node telemetry")
}

func (h *handler) handleCoordMMWave(client mqtt.Client, msg mqtt.Message) {
	h.logger.Info("Received coord mmwave")
}
