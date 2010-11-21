package oauth

import (
	"strings"
	"encoding/base64"
	"os"
	"json"
)

type SignedRequest struct {
	Algorithm  string
	Expires    float64
	IssuedAt   float64
	OAuthToken string
	UserID     string
}

func ParseSignedRequest(input string) (sr SignedRequest, err os.Error) {
	d := strings.Split(input, ".", -1)
	//signature := DecodeBase64URL(d[0])
	data := DecodeBase64URL(d[1])

	var values interface{}
	if err = json.Unmarshal([]byte(data), &values); err != nil {
		return
	}
	val := values.(map[string]interface{})
	sr.Algorithm = val["algorithm"].(string)
	sr.Expires = val["expires"].(float64)
	sr.IssuedAt = val["issued_at"].(float64)
	sr.OAuthToken = val["oauth_token"].(string)
	sr.UserID = val["user_id"].(string)
	return
	// TODO: Check SignedRequest with signature
}

func DecodeBase64URL(s string) string {
	s += "=" // Facebook fix, facebook sends corrupted base64_url
	dbuf := make([]byte, base64.URLEncoding.DecodedLen(len(s)))
	_, err := base64.URLEncoding.Decode(dbuf, []byte(s))
	if err != nil {
		println(err.String())
	}
	return string(dbuf)
}
