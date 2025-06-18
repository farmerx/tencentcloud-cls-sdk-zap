package tencentcloud_cls_sdk_zap

import (
	clssdk "github.com/tencentcloud/tencentcloud-cls-sdk-go"
	"go.uber.org/zap/zapcore"
)

type CLSLogger struct {
	Client   *clssdk.AsyncProducerClient
	Topic    string
	Callback clssdk.CallBack
}

// This class provide interface for TencentCLoud CLS logger
type clsCore struct {
	logger *CLSLogger
	fields []zapcore.Field
	level  zapcore.Level
}

// Enabled ...
func (c *clsCore) Enabled(level zapcore.Level) bool {
	return c.level.Enabled(level)
}

func (c *clsCore) With(f []zapcore.Field) zapcore.Core {
	fields := c.fields
	fields = append(fields, f...)
	return &clsCore{
		logger: c.logger,
		fields: fields,
		level:  c.level,
	}
}

// Check CLS zap extension method to check if logger is enabled
func (c *clsCore) Check(entry zapcore.Entry, checked *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(entry.Level) {
		return checked.AddCore(entry, c)
	}
	return checked
}

// Write CLS zap extension method to write
func (c *clsCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	contents := make(map[string]string, 0)
	// add common zap log fields as attributes
	for _, s := range c.fields {
		clsValue(s, contents)
	}
	// add zap log fields as attributes
	for _, s := range fields {
		clsValue(s, contents)
	}
	if entry.Level > zapcore.InfoLevel {
		callerString := entry.Caller.String()
		if len(callerString) > 0 {
			contents["Caller"] = callerString
		}

		if len(entry.Stack) > 0 {
			contents["StackTrace"] = entry.Stack
		}
	}
	contents["Level"] = entry.Level.String()
	contents["Message"] = entry.Message
	contents["LoggerName"] = entry.LoggerName
	log := clssdk.NewCLSLog(entry.Time.UnixMilli(), contents)
	if err := c.logger.Client.SendLog(c.logger.Topic, log, c.logger.Callback); err != nil {
		return err
	}
	return nil
}

// Sync CLS zap extension method to sync
func (c *clsCore) Sync() error {
	return nil
}
