package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	TIMEFORMAT = "2006-01-02 15-04-05.000"
)
const (
	//logmod 默认是 1 文件  2 stdout 4 其他
	FILE_MODE = 1 << iota
	STDOUT_MODE
	OTHER_MODE

	//time
	TIME_KEY        = "time"
	LEVEL_KEY       = "level"
	NAME_KEY        = "logger"
	CALLER_KEY      = "line"
	MESSAGE_KEY     = "data"
	STACK_TRACE_KEY = "stacktrace"
	SERVER_NAME     = "server"
)

// format
func currentTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	t = t.Local()
	enc.AppendString(t.Format(TIMEFORMAT))
}

type LogFileConfig struct {
	Filename   string `toml:"filename"`
	MaxSize    int    `toml:"maxsize"` //MB
	MaxAge     int    `toml:"maxage"`
	MaxBackups int    `toml:"maxbackups"`
	LocalTime  bool   `toml:"localtime"`
	Compress   bool   `toml:"compress"`
}

type LogKeyConfig struct {
	LogName       string `toml:"logname"`
	Timekey       string `toml:"timekey"`
	LevelKey      string `toml:"levelkey"`
	NameKey       string `toml:"namekey"`
	CallerKey     string `toml:"callerkey"`
	MessageKey    string `toml:"messagekey"`
	StacktraceKey string `toml:"stacktracekey"`
	ServerName    string `toml:serverName"`

	Level  string `toml:"level"`
	IsDev  bool   `toml:"isdev"`
	LogMod int8   `toml:"logmod"`
}

type LogConfig struct {
	LogFileConfig
	LogKeyConfig
}

type Logger struct {
	zlogs     *zap.Logger
	logConfig *LogConfig
}

func NewLogger(lfg *LogConfig) *Logger {

	logger := new(Logger)
	logger.logConfig = lfg
	logger.setDefaultConfig()
	logger.initLogger()

	return logger
}

func (l *Logger) initLogger() {

	lj := l.getLumberjackLog()

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.UnmarshalText([]byte(l.logConfig.Level))

	zws := make([]zapcore.WriteSyncer, 0)
	zws = append(zws, zapcore.AddSync(lj))

	if l.logConfig.LogMod&STDOUT_MODE > 0 {
		zws = append(zws, zapcore.AddSync(os.Stdout))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(l.getZapEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zws...),
		atomicLevel,
	)

	zapOptions := make([]zap.Option, 0)

	//get aller
	caller := zap.AddCaller()
	zapOptions = append(zapOptions, caller)

	development := zap.Development()
	zapOptions = append(zapOptions, development)

	filed := zap.Fields(zap.String(SERVER_NAME, l.logConfig.ServerName))

	zapOptions = append(zapOptions, filed)
	l.zlogs = zap.New(core, zapOptions...)
}
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	writerSlice := []zap.Field{zap.Field(zap.String("level", "debug"))}
	writerSlice = append(writerSlice, fields...)
	l.zlogs.Debug(msg, writerSlice...)
}
func (l *Logger) Info(msg string, fields ...zap.Field) {
	writerSlice := []zap.Field{zap.Field(zap.String("level", "debug"))}
	writerSlice = append(writerSlice, fields...)
	l.zlogs.Info(msg, writerSlice...)
}
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	writerSlice := []zap.Field{zap.Field(zap.String("level", "debug"))}
	writerSlice = append(writerSlice, fields...)
	l.zlogs.Warn(msg, writerSlice...)
}
func (l *Logger) Error(msg string, fields ...zap.Field) {
	writerSlice := []zap.Field{zap.Field(zap.String("level", "debug"))}
	writerSlice = append(writerSlice, fields...)
	l.zlogs.Error(msg, writerSlice...)
}

func (l *Logger) setDefaultConfig() {
	if len(l.logConfig.Timekey) == 0 {
		l.logConfig.Timekey = TIME_KEY
	}

	if len(l.logConfig.LevelKey) == 0 {
		l.logConfig.LevelKey = LEVEL_KEY
	}

	if len(l.logConfig.NameKey) == 0 {
		l.logConfig.NameKey = NAME_KEY
	}
	if len(l.logConfig.CallerKey) == 0 {
		l.logConfig.CallerKey = CALLER_KEY
	}
	if len(l.logConfig.MessageKey) == 0 {
		l.logConfig.MessageKey = MESSAGE_KEY
	}
	if len(l.logConfig.StacktraceKey) == 0 {
		l.logConfig.StacktraceKey = STACK_TRACE_KEY
	}

	if l.logConfig.MaxSize == 0 {
		l.logConfig.MaxSize = 500
	}

	if l.logConfig.MaxBackups == 0 {
		l.logConfig.MaxBackups = 30
	}

	if l.logConfig.MaxAge == 0 {
		l.logConfig.MaxAge = 30
	}
}

// make zconfig
func (l *Logger) getZapEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        l.logConfig.Timekey,
		LevelKey:       l.logConfig.LevelKey,
		NameKey:        l.logConfig.NameKey,
		CallerKey:      l.logConfig.CallerKey,
		MessageKey:     l.logConfig.MessageKey,
		StacktraceKey:  l.logConfig.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     currentTimeEncoder,             // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
}

// get lumber jacklog
func (l *Logger) getLumberjackLog() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   l.logConfig.Filename,   // 日志文件路径
		MaxSize:    l.logConfig.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: l.logConfig.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     l.logConfig.MaxAge,     // 文件最多保存多少天
		Compress:   l.logConfig.Compress,   //压缩
	}
}
