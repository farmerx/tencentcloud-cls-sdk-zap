package tencentcloud_cls_sdk_zap

import (
	clssdk "github.com/tencentcloud/tencentcloud-cls-sdk-go"
	"go.uber.org/zap/zapcore"
)

// NewCLSCore creates new OpenTelemetry Core to export logs in CLS format
func NewCLSCore(topic string, callback clssdk.CallBack, client *clssdk.AsyncProducerClient, opts ...Option) zapcore.Core {
	c := &clsCore{
		logger: &CLSLogger{client, topic, callback},
		level:  zapcore.InfoLevel,
	}
	for _, apply := range opts {
		apply(c)
	}

	return c
}

// Option is a function that applies an option to an OpenTelemetry Core
type Option func(c *clsCore)

// WithLevel sets the minimum level for the OpenTelemetry Core log to be exported
func WithLevel(level zapcore.Level) Option {
	return Option(func(c *clsCore) {
		c.level = level
	})
}
