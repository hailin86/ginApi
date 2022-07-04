package common

import (
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.SugaredLogger
	TraceKey = "traceId"
)

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	//logger := zap.New(core, zap.AddCaller())
	//AddCallerSkip跳过指定层级 不加这个参数 打印日志的行数永远是该文件 日志记录方法的行数
	logger := zap.New(core, zap.AddCaller(),zap.AddCallerSkip(1))
	Logger = logger.Sugar()
}

func LogDebug(c *gin.Context,args ...interface{}) {
	args = append(args, ";traceId:",c.GetString(TraceKey))
	Logger.Debug(args)
}

func LogInfo(c *gin.Context,args ...interface{})  {
	args = append(args, ";traceId:",c.GetString(TraceKey))
	Logger.Info(args)
}

func LogWarn (c *gin.Context,args ...interface{})  {
	args = append(args, ";traceId:",c.GetString(TraceKey))
	Logger.Warn(args)
}

func LogError (c *gin.Context,args ...interface{})  {
	args = append(args, ";traceId:",c.GetString(TraceKey))
	Logger.Error(args)
}


func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 格式化时间 可自定义
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 保存文件日志切割
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/app.log", // 指定日志文件目录
		MaxSize:    1,            // 文件内容大小, MB
		MaxBackups: 5,            // 保留旧文件最大个数
		MaxAge:     30,           // 保留旧文件最大天数
		Compress:   false,        // 文件是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

