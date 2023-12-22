package main

import (
	"github.com/bongochat/utils/v1/date"
	"github.com/bongochat/utils/v1/logger"
	"github.com/bongochat/utils/v1/resterrors"
)

func errTest() (string, resterrors.RestError) {
	if 1+1 == 1 {
		return "", resterrors.NewNotFoundError("This is not found error")
	}
	return "No error", nil
}

func main() {
	logger.Info("This is info log")
	msg, err := errTest()
	if err != nil {
		logger.Error("There was an error", err)
	}
	logger.Info(msg)
	logger.Info(date.GetCurrentDate())
}
