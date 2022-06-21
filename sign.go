package fireblocks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

// PrivateKey is the Private key used globally in the binding.
var (
	PrivateKey = []byte(``)
	ApiKey     = ""
)

func CreateJwt(path string, body []byte) (string, error) {
	iat := time.Now()
	exp := iat.Add(1 * time.Minute)
	return signJwt(jwt.MapClaims{
		"uri":      path,
		"nonce":    uuid.NewString(),
		"iat":      iat.Unix(),
		"exp":      exp.Unix(),
		"sub":      ApiKey,
		"bodyHash": fmt.Sprintf("%x", sha256.Sum256(body)),
	})
}

func signJwt(claims jwt.MapClaims) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(PrivateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}
	return token, nil
}

func ComputeSignature(t time.Time, payload []byte, secret string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(fmt.Sprintf("%d", t.Unix())))
	mac.Write([]byte("."))
	mac.Write(payload)
	return mac.Sum(nil)
}

func encryptPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	b := h.Sum(nil)
	return "{SHA256}" + base64.StdEncoding.EncodeToString(b)
}

func encryptPassword2(password string) string {
	h := sha256.Sum256([]byte(password))
	fmt.Println()
	return "{SHA256}" + base64.StdEncoding.EncodeToString(h[:])
}
