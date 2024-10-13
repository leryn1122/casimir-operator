package logger

import (
	"github.com/go-logr/logr"
	"go.uber.org/zap/zapcore"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

//var _ logr.Logger = Logger{}

type Logger struct {
	origin logr.LogSink
}

func New(origin logr.LogSink) logr.Logger {
	l := logr.New(&Logger{
		origin: origin,
	})
	return l
}

func (l *Logger) Init(info logr.RuntimeInfo) {
	l.origin.Init(info)
}

func (l *Logger) Enabled(level int) bool {
	return l.origin.Enabled(level)
}

func (l *Logger) Info(level int, msg string, keysAndValues ...any) {
	l.origin.Info(level, msg, keysAndValues...)
}

func (l *Logger) Error(err error, msg string, keysAndValues ...any) {
	l.origin.Error(err, msg, keysAndValues...)
}

func (l *Logger) WithValues(keysAndValues ...any) logr.LogSink {
	lw := *l
	lw.origin = l.origin.WithValues(keysAndValues...)
	return &lw
}

func (l *Logger) WithName(name string) logr.LogSink {
	lw := *l
	lw.origin = l.origin.WithName(name)
	return &lw
}

func InitializeLogger() {
	opts := zap.Options{
		StacktraceLevel: zapcore.PanicLevel,
	}
	zap.UseFlagOptions(&opts)
	sink := zap.New(zap.UseFlagOptions(&opts)).GetSink()
	logger := New(sink)
	klog.SetLogger(logger)
	ctrl.SetLogger(logger)
}
