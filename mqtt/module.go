package mqtt

import "go.uber.org/fx"

var Module = fx.Module("mqtt",
	fx.Provide(
		NewHandler,
		NewMQTTClient,
	),
	fx.Invoke(
		RegisterHandlers,
	),
)
