package tencentcloud_cls_sdk_zap

import (
	"fmt"
	"math"

	"go.uber.org/zap/zapcore"
)

// clsValue convert zap Field into CLS Log Content
func clsValue(f zapcore.Field, contents map[string]string) {
	switch f.Type {
	case zapcore.UnknownType:
		contents[f.Key] = f.String
		return
	case zapcore.BoolType:
		if f.Integer == 1 {
			contents[f.Key] = "true"
		} else {
			contents[f.Key] = "false"
		}
		return
	case zapcore.Float64Type:
		contents[f.Key] = fmt.Sprintf("%f", math.Float64frombits(uint64(f.Integer)))
		return
	case zapcore.Float32Type:
		contents[f.Key] = fmt.Sprintf("%f", math.Float64frombits(uint64(f.Integer)))
		return
	case zapcore.Int64Type:
		contents[f.Key] = fmt.Sprintf("%d", f.Integer)
		return
	case zapcore.Int32Type:
		contents[f.Key] = fmt.Sprintf("%d", f.Integer)
		return
	case zapcore.Int16Type:
		contents[f.Key] = fmt.Sprintf("%d", f.Integer)
		return
	case zapcore.Int8Type:
		contents[f.Key] = fmt.Sprintf("%d", f.Integer)
		return
	case zapcore.StringType:
		contents[f.Key] = f.String
		return
	case zapcore.Uint64Type:
		contents[f.Key] = fmt.Sprintf("%d", int64(uint64(f.Integer)))
		return
	case zapcore.Uint32Type:
		contents[f.Key] = fmt.Sprintf("%d", int64(uint64(f.Integer)))
		return
	case zapcore.Uint16Type:
		contents[f.Key] = fmt.Sprintf("%d", int64(uint64(f.Integer)))
		return
	case zapcore.Uint8Type:
		contents[f.Key] = fmt.Sprintf("%d", int64(uint64(f.Integer)))
		return
	case zapcore.ErrorType:
		err := f.Interface.(error)
		if err != nil {
			contents[f.Key] = err.Error()
			return
		}
		return
	case zapcore.SkipType:
		return
	}
	// unhandled types will be treated as string
	contents[f.Key] = f.String
	return
}
