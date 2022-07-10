package crypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	jwtSvc "github.com/uchupx/golang-absensi/pkg/util/jwt"
)

func (s *cryptService) loadRsaPrivateKey() (rsaKey *rsa.PrivateKey, err error) {
	bytes, err := ioutil.ReadFile(s.conf.RSAKeyPath)
	if err != nil {
		s.logger.Errorf("[loadRsaPrivateKey] failed read file with path: %s, err: %+v", s.conf.RSAKeyPath, err)
		return
	}

	rsaKey, err = jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		s.logger.Errorf("[loadRsaPrivateKey] failed parse private key, err: %+v", err)
		return
	}

	return
}

func (s *cryptService) loadRsaPublicKey() (rsaPub *rsa.PublicKey, err error) {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("%s.pub", s.conf.RSAKeyPath))
	if err != nil {
		s.logger.Errorf("[loadRsaPublicKey] failed read file with path: %s, err: %+v", fmt.Sprintf("%s.pub", s.conf.RSAKeyPath), err)
		return
	}

	rsaPub, err = jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		s.logger.Errorf("[loadRsaPublicKey] failed parse public key, err: %+v", err)
		return
	}

	return
}

func (s *cryptService) CreateSignPSS(value string) (signatureStr string, err error) {
	msg := []byte(value)
	msgHash := sha256.New()

	_, err = msgHash.Write(msg)
	if err != nil {
		return
	}

	msgHashSum := msgHash.Sum(nil)

	privateKey, err := s.loadRsaPrivateKey()
	if err != nil {
		return
	}

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		s.logger.Errorf("[CreateSignPSS] failed to create signature, err: %+v", err)
		return
	}

	signatureStr = base64.URLEncoding.EncodeToString(signature)
	return
}

func (s *cryptService) Verify(value string, signature string) (resp bool, err error) {
	resp = false

	signatureByte, err := base64.URLEncoding.DecodeString(signature)
	if err != nil {
		return
	}

	msg := []byte(value)
	msgHash := sha256.New()

	_, err = msgHash.Write(msg)
	if err != nil {
		return
	}

	msgHashSum := msgHash.Sum(nil)

	publicKey, err := s.loadRsaPublicKey()
	if err != nil {
		return
	}

	err = rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signatureByte, nil)
	if err != nil {
		s.logger.Errorf("[Verify] failed verify token, err: %+v", err)
		return
	}

	resp = true

	return
}

func (s *cryptService) CreateJWTToken(expDuration time.Duration, content interface{}) (token *string, err error) {
	privateKey, err := s.loadRsaPrivateKey()
	if err != nil {
		return
	}

	publicKey, err := s.loadRsaPublicKey()
	if err != nil {
		return
	}

	jwtServicecryptService := jwtSvc.NewJWT(privateKey, publicKey, s.logger)

	return jwtServicecryptService.Create(expDuration, content)
}

func (s *cryptService) VerifyJWTToken(token string) (result interface{}, err error) {
	privateKey, err := s.loadRsaPrivateKey()
	if err != nil {
		return
	}

	publicKey, err := s.loadRsaPublicKey()
	if err != nil {
		return
	}

	jwtServicecryptService := jwtSvc.NewJWT(privateKey, publicKey, s.logger)

	return jwtServicecryptService.Validate(token)
}
