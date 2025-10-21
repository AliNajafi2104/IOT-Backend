package mqtt

import (
	"context"
	"fmt"
	"time"

	"github.com/IOT-Backend/config"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/fx"
)

func NewMQTTClient(lc fx.Lifecycle, handler *Handler, cfg *config.Config) mqtt.Client {
	broker := cfg.MQTT.Broker
	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetUsername(cfg.MQTT.Username)
	opts.SetPassword(cfg.MQTT.Password)
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
			return nil
		},
		OnStop: func(ctx context.Context) error {
			client.Disconnect(250)
			return nil
		},
	})
	return client
}
