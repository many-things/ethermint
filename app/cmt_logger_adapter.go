package app

import (
	cmtlog "github.com/cometbft/cometbft/libs/log"

	"cosmossdk.io/log"
)

var _ log.Logger = (*CometLoggerAdapter)(nil)

type CometLoggerAdapter struct {
	cmtlog.Logger
}

func (cmt CometLoggerAdapter) Warn(msg string, keyVals ...any) {
	cmt.Logger.Error(msg, keyVals...)
}

func (cmt CometLoggerAdapter) With(keyVals ...any) log.Logger {
	logger := cmt.Logger.With(keyVals...)
	return CometLoggerAdapter{logger}
}

func (cmt CometLoggerAdapter) Impl() any {
	return cmt.Logger
}
