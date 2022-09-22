package logx

import (
	"context"
	"os"
	"time"

	"github.com/panyaxbo/libs/contextx"
	"github.com/sirupsen/logrus"
)

type Severity string
type LogType string

const (
	severityKey = "severity"

	SeverityInfo      Severity = "INFO"
	SeverityDebug     Severity = "DEBUG"
	SeverityWarn      Severity = "WARNING"
	SeverityError     Severity = "ERROR"
	SeverityEmergency Severity = "EMERGENCY"

	logTypeKey = "log_type"

	LogTypeInBound      LogType = "INBOUND"
	LogTypeReqInBound   LogType = "REQ_INBOUND"
	LogTypeRespInBound  LogType = "RESP_INBOUND"
	LogTypeOutBound     LogType = "OUTBOUND"
	LogTypeReqOutBound  LogType = "REQ_OUTBOUND"
	LogTypeRespOutBound LogType = "RESP_OUTBOUND"
	LogTypeWorkFlow     LogType = "WORKFLOW"
	LogTypeProcessing   LogType = "PROCESSING"

	workFlowKey = "work_flow"
	refIdKey    = "ref_id"

	timestampFormat = time.RFC3339Nano
)

var (
	LimitMSG = 10000
)

func StandardLogger() *logrus.Logger {
	return logrus.StandardLogger()
}

func Init(level, env string) {
	logrus.SetOutput(os.Stdout)

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(lvl)
	}

	if env == "local" || env == "dev" {
		logrus.SetFormatter(&logrus.TextFormatter{})
		return
	}

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: timestampFormat,
	})
}

func LimitMSGByte(b []byte) string {
	if LimitMSG < len(b) {
		return string(b[:LimitMSG]) + "..."
	}
	return string(b)
}

func LimitMSGString(s string) string {
	if LimitMSG < len(s) {
		return s[:LimitMSG] + "..."
	}
	return s
}

func WithContext(ctx context.Context) logrus.FieldLogger {
	return logrus.WithField("id", contextx.GetID(ctx))
}

func WithSeverityInfo(ctx context.Context) logrus.FieldLogger {
	return logrus.WithFields(logrus.Fields{
		"id":        contextx.GetID(ctx),
		severityKey: SeverityInfo,
	})
}

func WithSeverityDebug(ctx context.Context) logrus.FieldLogger {
	return logrus.WithFields(logrus.Fields{
		"id":        contextx.GetID(ctx),
		severityKey: SeverityDebug,
	})
}

func WithSeverityWarn(ctx context.Context) logrus.FieldLogger {
	return logrus.WithFields(logrus.Fields{
		"id":        contextx.GetID(ctx),
		severityKey: SeverityWarn,
	})
}

func WithSeverityError(ctx context.Context) logrus.FieldLogger {
	return logrus.WithFields(logrus.Fields{
		"id":        contextx.GetID(ctx),
		severityKey: SeverityError,
	})
}

func WithSeverityEmergency(ctx context.Context) logrus.FieldLogger {
	return logrus.WithFields(logrus.Fields{
		"id":        contextx.GetID(ctx),
		severityKey: SeverityEmergency,
	})
}

func WithSeverityInfoWithLogTypeWorkFlow(ctx context.Context, logType LogType, workFlow, refId string) logrus.FieldLogger {
	return logrus.WithFields(logrus.Fields{
		"id":        contextx.GetID(ctx),
		severityKey: SeverityInfo,
		logTypeKey:  logType,
		workFlowKey: workFlow,
		refIdKey:    refId,
	})
}

func WithSeverityDebugWithLogTypeWorkFlow(ctx context.Context, logType LogType, workFlow, refId string) logrus.FieldLogger {
	return logrus.WithFields(logrus.Fields{
		"id":        contextx.GetID(ctx),
		severityKey: SeverityDebug,
		logTypeKey:  logType,
		workFlowKey: workFlow,
		refIdKey:    refId,
	})
}

func WithSeverityWarnWithLogTypeWorkFlow(ctx context.Context, logType LogType, workFlow, refId string) logrus.FieldLogger {
	return logrus.WithFields(logrus.Fields{
		"id":        contextx.GetID(ctx),
		severityKey: SeverityWarn,
		logTypeKey:  logType,
		workFlowKey: workFlow,
		refIdKey:    refId,
	})
}

func WithSeverityErrorWithLogTypeWorkFlow(ctx context.Context, logType LogType, workFlow, refId string) logrus.FieldLogger {
	return logrus.WithFields(logrus.Fields{
		"id":        contextx.GetID(ctx),
		severityKey: SeverityError,
		logTypeKey:  logType,
		workFlowKey: workFlow,
		refIdKey:    refId,
	})
}

func WithSeverityEmergencyWithLogTypeWorkFlow(ctx context.Context, logType LogType, workFlow, refId string) logrus.FieldLogger {
	return logrus.WithFields(logrus.Fields{
		"id":        contextx.GetID(ctx),
		severityKey: SeverityEmergency,
		logTypeKey:  logType,
		workFlowKey: workFlow,
		refIdKey:    refId,
	})
}

func WithField(key string, value interface{}) *logrus.Entry {
	return logrus.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logrus.WithFields(fields)
}

func Debugf(format string, args ...interface{}) {
	logrus.WithField(severityKey, SeverityDebug).Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logrus.WithField(severityKey, SeverityInfo).Infof(format, args...)
}

func Printf(format string, args ...interface{}) {
	logrus.WithField(severityKey, SeverityInfo).Printf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logrus.WithField(severityKey, SeverityWarn).Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	logrus.WithField(severityKey, SeverityWarn).Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.WithField(severityKey, SeverityError).Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.WithField(severityKey, SeverityEmergency).Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	logrus.WithField(severityKey, SeverityEmergency).Panicf(format, args...)
}

func Debug(args ...interface{}) {
	logrus.WithField(severityKey, SeverityDebug).Debug(args...)
}

func Info(args ...interface{}) {
	logrus.WithField(severityKey, SeverityInfo).Info(args...)
}

func Print(args ...interface{}) {
	logrus.WithField(severityKey, SeverityInfo).Print(args...)
}

func Warn(args ...interface{}) {
	logrus.WithField(severityKey, SeverityWarn).Warn(args...)
}

func Warning(args ...interface{}) {
	logrus.WithField(severityKey, SeverityWarn).Warning(args...)
}

func Error(args ...interface{}) {
	logrus.WithField(severityKey, SeverityError).Error(args...)
}

func Fatal(args ...interface{}) {
	logrus.WithField(severityKey, SeverityEmergency).Fatal(args...)
}

func Panic(args ...interface{}) {
	logrus.WithField(severityKey, SeverityEmergency).Panic(args...)
}

func Debugln(args ...interface{}) {
	logrus.WithField(severityKey, SeverityDebug).Debugln(args...)
}

func Infoln(args ...interface{}) {
	logrus.WithField(severityKey, SeverityInfo).Infoln(args...)
}

func Println(args ...interface{}) {
	logrus.WithField(severityKey, SeverityInfo).Println(args...)
}

func Warnln(args ...interface{}) {
	logrus.WithField(severityKey, SeverityWarn).Warnln(args...)
}

func Warningln(args ...interface{}) {
	logrus.WithField(severityKey, SeverityWarn).Warningln(args...)
}

func Errorln(args ...interface{}) {
	logrus.WithField(severityKey, SeverityError).Errorln(args...)
}

func Fatalln(args ...interface{}) {
	logrus.WithField(severityKey, SeverityEmergency).Fatalln(args...)
}

func Panicln(args ...interface{}) {
	logrus.WithField(severityKey, SeverityEmergency).Panicln(args...)
}
