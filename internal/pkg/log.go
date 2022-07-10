package pkg

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
	"unsafe"
)

type MyLog struct {
	name  string
	debug bool
	*zap.Logger
}

//New 返回一个日志对象
//@ name 日志名字
//@ debug ：true打印到控制台，false 打印到日志
func New(name string, debug bool) *MyLog {
	log := &MyLog{
		name:  name,
		debug: debug,
	}
	log.init()
	return log
}

// init 初始化日志
func (log *MyLog) init() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "line",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	var cores []zapcore.Core
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel
	})

	debug := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel
	})

	dPanic := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DPanicLevel
	})

	if log.debug {
		//debug 直接输出到终端中
		cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), infoLevel))
		cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), warnLevel))
		cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), errorLevel))
		cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), debug))
		cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), dPanic))
	} else {
		//使用默认日志切割
		var lumber = &Lumber{}
		infoWriter := lumber.GetDeferLumber(log.name + "_info.log")
		warnWriter := lumber.GetDeferLumber(log.name + "_warn.log")
		errorWriter := lumber.GetDeferLumber(log.name + "_error.log")
		debugWriter := lumber.GetDeferLumber(log.name + "_debug.log")
		panicWriter := lumber.GetDeferLumber(log.name + "_panic.log")
		cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(&infoWriter)), infoLevel))
		cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(&warnWriter)), warnLevel))
		cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(&errorWriter)), errorLevel))
		cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(&debugWriter)), debug))
		cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(&panicWriter)), dPanic))
	}

	// 最后创建具体的Logger
	core := zapcore.NewTee(cores...)

	caller := zap.AddCaller()
	development := zap.Development()
	logger := zap.New(core, caller, development, zap.AddStacktrace(zapcore.DPanicLevel))
	//logger.development设置为FALSE时，DPanic就不会panic
	//go.uber.org/zap/logger.go:279
	*(*bool)(unsafe.Pointer(uintptr(unsafe.Pointer(logger)) + uintptr(16))) = false
	log.Logger = logger
}

func (log *MyLog) Close() error {
	return nil
}

func (log *MyLog) MError(err error, opt ...option) {
	if err == nil {
		return
	}
	res, causeLine := log.split(err)
	register := &Option{}
	for _, o := range opt {
		o(register)
	}
	var fields []zap.Field
	fields = append(fields, register.registerOption(log)...)
	fields = append(fields, zap.String("app_name", log.name), zap.String("err_line", causeLine), zap.Any("err_stack", res))
	log.Error(err.Error(), fields...)
}

func (log *MyLog) MInfo(info string, opt ...option) {
	register := &Option{}
	for _, o := range opt {
		o(register)
	}
	var fields []zap.Field
	fields = append(fields, register.registerOption(log)...)
	fields = append(fields, zap.String("app_name", log.name))
	log.Info(info, fields...)
}

func (log *MyLog) split(err error) ([]string, string /*cause line*/) {
	str := fmt.Sprintf("%+v", err)
	tem := strings.Split(str, "\n")
	temCause := "can't get error stack "
	for i := 0; i < len(tem); i++ {
		if i == 0 {
			tem[i] = "ERR_REASON ：" + tem[i]
		}
		if i == 2 {
			temCause = tem[i]
		}
	}
	return tem, temCause
}
