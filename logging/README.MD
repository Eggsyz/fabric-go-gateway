+ 使用方式
```
logging.Init(logging.Config{
    Format:     format,         //日志格式
    LogLevel:   logLevel,       //日志级别
    FilePath:   filePath,       //不配置，则打印到控制台，下列配置都可不配置
    MaxSize:    maxSize,        //每个日志文件保存的最大尺寸 单位：M
    LocalTime:  localTime,      //是否使用本地时间 默认UTC
    Compress:   compress,       //是否压缩
    MaxAge:     maxAge,         //文件最多保存多少天，默认不删除
    MaxBackups: maxBackups,     // 日志文件最多保存多少个备份，默认所有
})
log:=logging.MustGetLogger("hello")
log.Info("hello world")
```