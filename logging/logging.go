package logging

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"regexp"
	"sync"
)

/**
 * FilePath 日志文件路径
 * Format 日志格式
 * LogLevel 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 */
type Config struct {
	Format     string
	LogLevel   string
	FilePath   string
	MaxSize    int
	LocalTime  bool
	Compress   bool
	MaxAge     int
	MaxBackups int
}

type Logging struct {
	mutex        sync.RWMutex
	levelEnabler zapcore.LevelEnabler
	encoder      zapcore.Encoder
	writer       zapcore.WriteSyncer
}

// 生成默认logging,并根据配置文件更新logging
func New(c Config) (*Logging, error) {
	s := &Logging{}
	err := s.Apply(c)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// 根据配置文件更新logging
func (l *Logging) Apply(c Config) error {
	l.mutex.Lock()
	initialize(&c)
	// 则保存到指定文件
	if c.FilePath != "" {
		if isValidLoggerFilePath(c.FilePath) {
			l.writer = zapcore.AddSync(&lumberjack.Logger{
				Filename:   c.FilePath,
				MaxSize:    c.MaxSize,
				LocalTime:  c.LocalTime,
				Compress:   c.Compress,
				MaxAge:     c.MaxAge,
				MaxBackups: c.MaxBackups,
			})
			l.encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		} else {
			fmt.Println("the logger file path is invalided")
			return fmt.Errorf("the logger file path is invalided")
		}
	} else { // 控制台打印
		l.writer = zapcore.AddSync(os.Stdout)
		formats, err := ParseFormat(c.Format)
		if err != nil {
			return err
		}
		multiFormatter := NewMultiFormatter()
		multiFormatter.SetFormatters(formats)
		l.encoder = NewFormatEncoder(multiFormatter)
	}
	l.levelEnabler = zap.NewAtomicLevelAt(getLoggerLevel(c.LogLevel))
	l.mutex.Unlock()
	return nil
}

func (l *Logging) ZapLogger(name string) *zap.Logger {
	l.mutex.Lock()
	if !isValidLoggerName(name) {
		panic(fmt.Sprintf("invalid logger name: %s", name))
	}
	core := zapcore.NewCore(l.encoder, l.writer, l.levelEnabler)
	l.mutex.Unlock()
	return NewZapLogger(core).Named(name)
}

func (l *Logging) Logger(name string) *Logger {
	zl := l.ZapLogger(name)
	return NewLogger(zl)
}

var loggerNameRegexp = regexp.MustCompile(`^\w+(\.\w+)*$`)
var loggerFilePathRegexp = regexp.MustCompile(`^[a-zA-Z]:(\\\w+)+|(/\w+)+|\.(/\w+)+$`)

func isValidLoggerName(loggerName string) bool {
	return loggerNameRegexp.MatchString(loggerName)
}

func isValidLoggerFilePath(loggerFilePath string) bool {
	return loggerFilePathRegexp.MatchString(loggerFilePath)
}

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func initialize(c *Config) {
	if c.Format == "" {
		c.Format = defaultFormat
	}
	if c.LogLevel == "" {
		c.LogLevel = defaultLevel
	}
	if c.MaxAge == 0 {
		c.MaxAge = defaultMaxAge
	}
	if c.MaxBackups == 0 {
		c.MaxBackups = defaultMaxBackups
	}
	if c.MaxSize == 0 {
		c.MaxSize = defaultMaxSize
	}
}
