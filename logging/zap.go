package logging

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"
)

func NewZapLogger(core zapcore.Core, options ...zap.Option) *zap.Logger {
	return zap.New(
		core,
		append([]zap.Option{
			zap.AddCaller(),
			zap.AddStacktrace(zapcore.ErrorLevel),
		}, options...)...,
	)
}

// NewGRPCLogger creates a grpc.Logger that delegates to a zap.Logger.
func NewGRPCLogger(l *zap.Logger) *zapgrpc.Logger {
	l = l.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(3),
	)
	return zapgrpc.NewLogger(l, zapgrpc.WithDebug())
}

// NewLogger creates a logger that delegates to the zap.SugaredLogger.
func NewLogger(l *zap.Logger, options ...zap.Option) *Logger {
	return &Logger{
		s: l.WithOptions(append(options, zap.AddCallerSkip(1))...).Sugar(),
	}
}

type Logger struct{ s *zap.SugaredLogger }

func (f *Logger) DPanic(args ...interface{})                    { f.s.DPanicf(formatArgs(args)) }
func (f *Logger) DPanicf(template string, args ...interface{})  { f.s.DPanicf(template, args...) }
func (f *Logger) DPanicw(msg string, kvPairs ...interface{})    { f.s.DPanicw(msg, kvPairs...) }
func (f *Logger) Debug(args ...interface{})                     { f.s.Debugf(formatArgs(args)) }
func (f *Logger) Debugf(template string, args ...interface{})   { f.s.Debugf(template, args...) }
func (f *Logger) Debugw(msg string, kvPairs ...interface{})     { f.s.Debugw(msg, kvPairs...) }
func (f *Logger) Error(args ...interface{})                     { f.s.Errorf(formatArgs(args)) }
func (f *Logger) Errorf(template string, args ...interface{})   { f.s.Errorf(template, args...) }
func (f *Logger) Errorw(msg string, kvPairs ...interface{})     { f.s.Errorw(msg, kvPairs...) }
func (f *Logger) Fatal(args ...interface{})                     { f.s.Fatalf(formatArgs(args)) }
func (f *Logger) Fatalf(template string, args ...interface{})   { f.s.Fatalf(template, args...) }
func (f *Logger) Fatalw(msg string, kvPairs ...interface{})     { f.s.Fatalw(msg, kvPairs...) }
func (f *Logger) Info(args ...interface{})                      { f.s.Infof(formatArgs(args)) }
func (f *Logger) Infof(template string, args ...interface{})    { f.s.Infof(template, args...) }
func (f *Logger) Infow(msg string, kvPairs ...interface{})      { f.s.Infow(msg, kvPairs...) }
func (f *Logger) Panic(args ...interface{})                     { f.s.Panicf(formatArgs(args)) }
func (f *Logger) Panicf(template string, args ...interface{})   { f.s.Panicf(template, args...) }
func (f *Logger) Panicw(msg string, kvPairs ...interface{})     { f.s.Panicw(msg, kvPairs...) }
func (f *Logger) Warn(args ...interface{})                      { f.s.Warnf(formatArgs(args)) }
func (f *Logger) Warnf(template string, args ...interface{})    { f.s.Warnf(template, args...) }
func (f *Logger) Warnw(msg string, kvPairs ...interface{})      { f.s.Warnw(msg, kvPairs...) }
func (f *Logger) Warning(args ...interface{})                   { f.s.Warnf(formatArgs(args)) }
func (f *Logger) Warningf(template string, args ...interface{}) { f.s.Warnf(template, args...) }

// for backwards compatibility
func (f *Logger) Critical(args ...interface{})                   { f.s.Errorf(formatArgs(args)) }
func (f *Logger) Criticalf(template string, args ...interface{}) { f.s.Errorf(template, args...) }
func (f *Logger) Notice(args ...interface{})                     { f.s.Infof(formatArgs(args)) }
func (f *Logger) Noticef(template string, args ...interface{})   { f.s.Infof(template, args...) }

func (f *Logger) Named(name string) *Logger { return &Logger{s: f.s.Named(name)} }
func (f *Logger) Sync() error               { return f.s.Sync() }
func (f *Logger) Zap() *zap.Logger          { return f.s.Desugar() }

func (f *Logger) IsEnabledFor(level zapcore.Level) bool {
	return f.s.Desugar().Core().Enabled(level)
}

func (f *Logger) With(args ...interface{}) *Logger {
	return &Logger{s: f.s.With(args...)}
}

func (f *Logger) WithOptions(opts ...zap.Option) *Logger {
	l := f.s.Desugar().WithOptions(opts...)
	return &Logger{s: l.Sugar()}
}

func formatArgs(args []interface{}) string { return strings.TrimSuffix(fmt.Sprintln(args...), "\n") }
