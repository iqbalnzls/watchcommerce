package logger

import (
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	XID         string
	IP          string
	ServiceName string
	Version     string
	Path        string
	Header      interface{}
	Req         interface{}
	Resp        interface{}
	Time        string
	Err         string
	ZapLog      *zap.Logger
}

type Logger interface {
	IncomingRequest(message ...string)
	SubProcessStart(message ...string) time.Time
	SubProcessEnd(start time.Time, message ...string)
	FinishedRequest(resp interface{}, message ...string)
	Info(message ...string)
	Error(message ...string)
	SetError(err error)
	SetRequest(req interface{})
}

func (l *Log) IncomingRequest(message ...string) {
	l.ZapLog.Info("Incoming Request", composeField(l, message)...)
}

func (l *Log) SubProcessStart(message ...string) time.Time {
	l.ZapLog.Info("Sub Process Start", composeField(l, message)...)
	return time.Now()
}

func (l *Log) SubProcessEnd(startTime time.Time, message ...string) {
	fields := []zap.Field{
		{Key: "processing-time", String: strconv.Itoa(int(time.Now().Sub(startTime).Milliseconds())) + "ms", Type: zapcore.StringType},
	}

	l.ZapLog.Info("Sub Process End", append(fields, composeField(l, message)...)...)
}

func (l *Log) FinishedRequest(resp interface{}, message ...string) {
	fields := []zap.Field{
		{Key: "req", Interface: l.Req, Type: zapcore.ReflectType},
		{Key: "resp", Interface: resp, Type: zapcore.ReflectType},
		{Key: "err", String: l.Err, Type: zapcore.StringType},
		{Key: "header", Interface: l.Header, Type: zapcore.ReflectType},
	}

	l.ZapLog.Info("Finished Request", append(fields, composeField(l, message)...)...)
}

func (l *Log) Info(messages ...string) {
	l.ZapLog.Info("Info", composeField(l, messages)...)
}

func (l *Log) Error(message ...string) {
	l.ZapLog.Error("Error", composeField(l, message)...)
}

func (l *Log) SetError(err error) {
	l.Err = err.Error()
}

func (l *Log) SetRequest(req interface{}) {
	l.Req = req
}

func composeField(l *Log, msg []string) []zap.Field {
	tempFields := make([]zap.Field, 0)

	for i, v := range msg {
		tempFields = append(tempFields, zap.Field{
			Key: "message_" + strconv.Itoa(i), String: v, Type: zapcore.StringType,
		})
	}

	return append(tempFields, []zap.Field{
		{Key: "xid", String: l.XID, Type: zapcore.StringType},
		{Key: "ip", String: l.IP, Type: zapcore.StringType},
		{Key: "path", String: l.Path, Type: zapcore.StringType},
		{Key: "service-name", String: l.ServiceName, Type: zapcore.StringType},
		{Key: "version", String: l.Version, Type: zapcore.StringType},
		{Key: "time", String: l.Time, Type: zapcore.StringType},
	}...)
}
