package signature

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
)

func GenerateSignature(apiSecret, requestMethod, fullURLPath, requestData string, expireTimestamp int64) string {
	key := []byte(apiSecret)
	u, err := url.Parse(fullURLPath)
	if err != nil {
		fmt.Println(fullURLPath, err)
	}
	u.RawQuery = u.Query().Encode()
	message := requestMethod + u.String() + strconv.FormatInt(expireTimestamp, 10) + requestData
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
