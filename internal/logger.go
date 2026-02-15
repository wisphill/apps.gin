package internal

import "go.uber.org/zap"

var Log *zap.SugaredLogger

func Init() error {
	l, err := zap.NewProduction()
	if err != nil {
		return err
	}
	Log = l.Sugar()
	return nil
}
