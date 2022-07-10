package logutil

import "github.com/sirupsen/logrus"

type NewLoggerParams struct {
	PrettyPrint bool
	ServiceName string
}

type UTCFormatter struct {
	logrus.Formatter
}

func (u UTCFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	return u.Formatter.Format(e)
}
func NewLogger(params NewLoggerParams) *logrus.Entry {
	log := logrus.New()
	log.SetFormatter(UTCFormatter{
		Formatter: &logrus.JSONFormatter{
			PrettyPrint: params.PrettyPrint,
		},
	})

	return log.WithField("service", params.ServiceName)
}
