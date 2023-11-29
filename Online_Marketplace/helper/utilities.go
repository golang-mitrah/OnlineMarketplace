package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"onlinemarketplace/model"
	"time"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

var SecretKey = "mjnhbfutyebdycs"

type Response struct {
	Data      interface{} `json:"Data"`
	StartTime time.Time   `json:"-"`
	EndTime   time.Time   `json:"-"`
	Message   string      `json:"Message"`
	Duration  string      `json:"Duration"`
}

func NewResponse() Response {
	r := Response{
		StartTime: time.Now(),
	}
	return r
}

func (r Response) Response(data interface{}) Response {
	r.Data = data
	r.Message = "Success"
	r.Duration = time.Now().Sub(r.StartTime).String()
	r.EndTime = time.Now()
	return r
}

func (r Response) ErrorResponse(err error) Response {
	r.Data = struct{}{}
	r.EndTime = time.Now()
	r.Message = err.Error()
	r.Duration = time.Now().Sub(r.StartTime).String()
	return r
}

func Encryption(stringToEncrypt, mySecrect string) (string, error) {

	block, err := aes.NewCipher([]byte(mySecrect))
	if err != nil {
		return "", fmt.Errorf("Error : %s", err)

	}

	plainText := []byte(stringToEncrypt)

	bytes := make([]byte, aes.BlockSize)
	cfb := cipher.NewCFBEncrypter(block, bytes)

	cipherText := make([]byte, len(plainText))

	cfb.XORKeyStream(cipherText, plainText)

	return Encode(cipherText), nil
}

func Encode(cipherText []byte) string {
	Encoding := base64.StdEncoding.EncodeToString(cipherText)
	return Encoding
}

func GenerateToken(user *model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("Invalid JWT Token")
	}
	claims["username"] = user.UserName
	claims["password"] = user.Password
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", errors.New("Some technical error")
	}
	return tokenString, nil
}

func LogInfo(payload interface{}, res Response) {
	payloadJsn, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	responseJsn, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	log.WithFields(log.Fields{"Request": payloadJsn, "Response": responseJsn}).Info()
}
