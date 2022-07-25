package gormlogrus

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// Level type
type Level logger.LogLevel

type mylogger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	Debug                 bool
}

func New() *mylogger {
	return &mylogger{
		SkipErrRecordNotFound: true,
	}
}

func (l *mylogger) LogMode(level logger.LogLevel) logger.Interface {
	if level > logger.Info {
		l.Debug = true
	}
	return l
}

func (l *mylogger) Info(ctx context.Context, s string, args ...interface{}) {
	logrus.WithContext(ctx).Infof(s, args...)
}

func (l *mylogger) Warn(ctx context.Context, s string, args ...interface{}) {
	logrus.WithContext(ctx).Warnf(s, args...)
}

func (l *mylogger) Error(ctx context.Context, s string, args ...interface{}) {
	logrus.WithContext(ctx).Errorf(s, args...)
}

func (l *mylogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, row := fc()
	fields := logrus.Fields{
		"line":          utils.FileWithLineNum(),
		"response_time": fmt.Sprintf("%.3f", float64(elapsed.Nanoseconds())/1e6),
		"row":           row,
		"sql":           sql,
	}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		logrus.WithContext(ctx).WithFields(fields).Error(err)
		return
	}
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		logrus.WithContext(ctx).WithFields(fields).Warn()
		return
	}
	if l.Debug {
		logrus.WithContext(ctx).WithFields(fields).Debug()
	}
}
