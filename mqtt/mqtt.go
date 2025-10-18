package mqtt

import (
	"context"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("mqtt",
	fx.Provide(NewHandler),
	fx.Invoke(NewMQTTClient),
)

func NewHandler(logger *zap.Logger) *handler {
	return &handler{
		logger: logger,
	}
}

type handler struct {
	logger *zap.Logger
}

func (h *handler) defaultMessageHandler(client mqtt.Client, msg mqtt.Message) {
	h.logger.Info("Received message: %s from topic: %s\n", zap.ByteString("payload", msg.Payload()), zap.String("topic", msg.Topic()))
}

func (h *handler) connectHandler(client mqtt.Client) {
	h.logger.Info("Connected")
}

func (h *handler) connectLostHandler(client mqtt.Client, err error) {
	h.logger.Error("Connect lost: %v", zap.Error(err))
}

func NewMQTTClient(lc fx.Lifecycle, handler *handler) {
	broker := "tcp://localhost:1883"
	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetUsername("user1")
	opts.SetPassword("user1")
	opts.SetDefaultPublishHandler(handler.defaultMessageHandler)
	opts.OnConnect = handler.connectHandler
	opts.OnConnectionLost = handler.connectLostHandler
	client := mqtt.NewClient(opts)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			token := client.Connect()
			if !token.WaitTimeout(10 * time.Second) {
				return fmt.Errorf("mqtt connect timeout")
			}
			if err := token.Error(); err != nil {
				return err
			}

			RegisterSubscriptions(client, handler)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			client.Disconnect(250)
			return nil
		},
	})
}

func RegisterSubscriptions(client mqtt.Client, handler *handler) {
	client.Subscribe("site/+/node/+/telemetry", 0, handler.handleNodeTelemetry)
	client.Subscribe("site/+/coord/+/telemetry", 0, handler.handleCoordTelemetry)
	client.Subscribe("site/+/coord/+/mmwave", 0, handler.handleCoordMMWave)
}
