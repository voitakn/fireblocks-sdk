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
	PrivateKey = ""
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
		//crypto.createHash("sha256").update(JSON.stringify(bodyJson || "")).digest().toString("hex")
	})
}

func signJwt(tokenData jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, tokenData)
	tokenString, err := token.SignedString(PrivateKey)
	if err != nil {
		fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
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

func main() {
	fmt.Println(encryptPassword("abcd1234"))
	fmt.Println(encryptPassword2(`{"main": "test"}`))
}
