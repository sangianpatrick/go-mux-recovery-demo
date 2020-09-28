package middleware

import "github.com/sirupsen/logrus"

// MyRecoveryHandlerLogger is a struct
type MyRecoveryHandlerLogger struct {
	Logger *logrus.Logger
}

// Println will log the recovered panic
func (mrhl MyRecoveryHandlerLogger) Println(recovered ...interface{}) {
	mrhl.Logger.Error(recovered...)
}
