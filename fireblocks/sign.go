package fireblocks

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"log"
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

func SignatureVerify(signature string, body []byte, key []byte) (bool, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return false, fmt.Errorf("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	bodySig, err := rsa.EncryptOAEP(
		sha512.New(),
		rand.Reader,
		pub,
		body,
		nil)
	if err != nil {
		return false, fmt.Errorf("error rsa encrypt signature")
	}
	log.Println("SignatureVerify = signature", signature)
	body64 := Base64Enc(bodySig)
	log.Println("SignatureVerify ==== body64", body64)
	if signature == body64 {
		return true, nil
	}
	return false, fmt.Errorf("signature was not validated")
}

func RsaEncrypt(origData []byte, key []byte) ([]byte, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, fmt.Errorf("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(nil, pub, origData)
}

func Base64Enc(b1 []byte) string {
	s1 := base64.StdEncoding.EncodeToString(b1)
	s2 := ""
	var LEN int = 76
	for len(s1) > 76 {
		s2 = s2 + s1[:LEN] + "\n"
		s1 = s1[LEN:]
	}
	s2 = s2 + s1
	return s2
}
