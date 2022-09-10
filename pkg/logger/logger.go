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
	LOG_EXT    = ".log"
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

	Level  string `toml:"level"`
	IsDev  bool   `toml:"isdev"`
	LogMod int8   `toml:"logmod"`
}

type LogConfig struct {
	LogFileConfig
	LogKeyConfig
	isMultiFile bool
}

func WithMultiFile(lfg *LogConfig) *LogConfig {
	lfg.isMultiFile = true
	return lfg
}
func (lfg *LogConfig) IsMultiFile() bool {
	return lfg.isMultiFile
}

type Logger struct {
	zlogs     *zap.Logger
	logConfig *LogConfig
}

func NewLogger(lfg *LogConfig) *Logger {
	logger := new(Logger)
	logger.logConfig = lfg
	logger.setDefaultConfig()

	if lfg.IsMultiFile() {
		logger.initLoggerMulti()
	} else {
		logger.initLogger()
	}
	return logger
}

func (l *Logger) initLoggerMulti() {
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.UnmarshalText([]byte(l.logConfig.Level))
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	cores := [...]zapcore.Core{
		zapcore.NewCore(zapcore.NewJSONEncoder(l.getZapEncoderConfig()),
			l.getWriteSyncer(l.logConfig.Filename+".debug"+LOG_EXT), debugPriority),

		zapcore.NewCore(zapcore.NewJSONEncoder(l.getZapEncoderConfig()),
			l.getWriteSyncer(l.logConfig.Filename+".info"+LOG_EXT), infoPriority),

		zapcore.NewCore(zapcore.NewJSONEncoder(l.getZapEncoderConfig()),
			l.getWriteSyncer(l.logConfig.Filename+".warn"+LOG_EXT), warnPriority),

		zapcore.NewCore(zapcore.NewJSONEncoder(l.getZapEncoderConfig()),
			l.getWriteSyncer(l.logConfig.Filename+".error"+LOG_EXT), errorPriority),
	}
	zapOptions := make([]zap.Option, 0)
	zapOptions = append(zapOptions, zap.AddCaller())
	zapOptions = append(zapOptions, zap.AddCallerSkip(1))
	zapOptions = append(zapOptions, zap.Development())
	l.zlogs = zap.New(zapcore.NewTee(cores[:]...), zapOptions...)
}

func (l *Logger) initLogger() {
	zws := make([]zapcore.WriteSyncer, 0)
	zws = append(zws, zapcore.AddSync(l.getLumberjackLog(l.logConfig.Filename+LOG_EXT)))
	if l.logConfig.LogMod&STDOUT_MODE > 0 {
		zws = append(zws, zapcore.AddSync(os.Stdout))
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.UnmarshalText([]byte(l.logConfig.Level))

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(l.getZapEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zws...),
		atomicLevel,
	)

	zapOptions := make([]zap.Option, 0)
	zapOptions = append(zapOptions, zap.AddCaller())
	zapOptions = append(zapOptions, zap.AddCallerSkip(1))
	zapOptions = append(zapOptions, zap.Development())
	l.zlogs = zap.New(core, zapOptions...)
}

func (l *Logger) GetZlog() *zap.Logger {
	return l.zlogs
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
		LineEnding:     zapcore.DefaultLineEnding,      //最后结尾
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     currentTimeEncoder,             // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, // 时间编码器
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,        //函数编码
	}
}

// get lumber jacklog
func (l *Logger) getLumberjackLog(fileName string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   fileName,               // 日志文件路径
		MaxSize:    l.logConfig.MaxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: l.logConfig.MaxBackups, // 日志文件最多保存多少个备份
		MaxAge:     l.logConfig.MaxAge,     // 文件最多保存多少天
		Compress:   l.logConfig.Compress,   //压缩
	}
}

func (l *Logger) getWriteSyncer(fileName string) zapcore.WriteSyncer {
	lumberjackObj := l.getLumberjackLog(fileName)
	if l.logConfig.LogMod&STDOUT_MODE > 0 {
		return zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(lumberjackObj))
	}
	return zapcore.AddSync(lumberjackObj)
}
