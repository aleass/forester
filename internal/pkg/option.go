package pkg

import (
	"fmt"
	"go.uber.org/zap"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"time"
)

type Option struct {
	traceID string
	errCode int
}

func (o *Option) registerOption(log *MyLog) []zap.Field {
	var res []zap.Field
	if o.traceID != "" {
		res = append(res, zap.String("traceID", o.traceID))
	}
	if o.errCode != 0 {
		res = append(res, zap.Int("errCode", o.errCode))
	}
	return res
}

type option func(opt *Option)

//WithTraceID 打印traceID
func WithTraceID(traceID string) option {
	return func(opt *Option) {
		opt.traceID = traceID
	}
}

//WithErrCode 打印error code
func WithErrCode(errCode int) option {
	return func(opt *Option) {
		opt.errCode = errCode
	}
}

// Lumber 文件切割配置
type Lumber struct {
	Filename   string // 日志文件路径
	MaxSize    int    // 每个日志文件保存的最大尺寸 单位：M  128
	MaxBackups int    // 日志文件最多保存多少个备份 30
	MaxAge     int    // 文件最多保存多少天 7
	Compress   bool   // 是否压缩
}

//GetDeferLumber 返回默认配置
//@ param：filename文件名
//@ return: lumberjack.Logger
func (l *Lumber) GetDeferLumber(filename string) lumberjack.Logger {
	today := time.Now().Format("20060102")
	//LIVE_FRONT_SRC=/www/stock.spider.sscf.com
	filename = fmt.Sprintf("/log/%s/%s", today, filename)
	return lumberjack.Logger{
		Filename:   filename,
		MaxSize:    128,
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
	}
}

//SetLumber 返回指定配置
//@ param：Lumber配置的结构体
//@ return: lumberjack.Logger
func (l *Lumber) SetLumber(opt Lumber) lumberjack.Logger {
	return lumberjack.Logger{
		Filename:   opt.Filename,
		MaxSize:    opt.MaxSize,
		MaxBackups: opt.MaxBackups,
		MaxAge:     opt.MaxAge,
		Compress:   opt.Compress,
	}
}
