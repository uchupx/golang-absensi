package crypt

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/uchupx/golang-absensi/internal/config"
)

type CryptService interface {
	CreateSignPSS(value string) (signatureStr string, err error)
	Verify(value string, signature string) (resp bool, err error)
	CreateJWTToken(expDuration time.Duration, content interface{}) (token *string, err error)
	VerifyJWTToken(token string) (result interface{}, err error)
}

type cryptService struct {
	conf   *config.Config
	logger *logrus.Entry
}

type Params struct {
	Conf   *config.Config
	Logger *logrus.Entry
}

func NewCryptService(params Params) CryptService {
	return &cryptService{
		conf:   params.Conf,
		logger: params.Logger,
	}
}
