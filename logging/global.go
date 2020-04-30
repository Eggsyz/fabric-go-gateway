package logging

const (
	defaultFormat     = "%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}"
	defaultLevel      = "info"
	defaultMaxSize    = 100
	defaultMaxAge     = 7
	defaultMaxBackups = 100
)

var Global *Logging

func init() {
	logging, err := New(Config{})
	if err != nil {
		panic(err)
	}
	Global = logging
}

func Init(config Config) {
	err := Global.Apply(config)
	if err != nil {
		panic(err)
	}
}
func MustGetLogger(loggerName string) *Logger {
	return Global.Logger(loggerName)
}
