package oauth

import (
        "strings"
        "encoding/base64"
        "os"
        "json"
	"fmt"
)

type SignedRequest struct {
        Algorithm  string
        Expires    float64
        IssuedAt   float64
        OAuthToken string
        UserID     string
}

func ParseSignedRequest(input string) (sr SignedRequest, err os.Error) {
	if len(input) == 0 {
		return sr, os.NewError("Input string is empty.")
	}
	points := strings.Count(input, ".")
	if points != 1 {
		return sr, os.NewError(fmt.Sprintf("Input string has %d '.'s, one is required.", points))
	}
        d := strings.Split(input, ".", -1)
        //signature := DecodeBase64URL(d[0])
        data, err := decodeBase64FBURL(d[1])
        if err != nil {
                return
        }
        data = data[:len(data)] // TEMP: Bypass an illegal char at the end
        var values interface{}
        if err = json.Unmarshal([]byte(data), &values); err != nil {
                return sr, os.NewError("Error: ParseSignedRequest: json.Unmarshal: " + err.String() + " in " + data)
        }
        val := values.(map[string]interface{})
	if val["algorithm"] != nil {
	  sr.Algorithm = val["algorithm"].(string)
	}
	if val["expires"] != nil {
	  sr.Expires = val["expires"].(float64)
	}
	if val["issued_at"] != nil {
	  sr.IssuedAt = val["issued_at"].(float64)
	}
	if val["oauth_token"] != nil {
	  sr.OAuthToken = val["oauth_token"].(string)
	}
	if val["user_id"] != nil {
	  sr.UserID = val["user_id"].(string)
	}
        return
        // TODO: Check SignedRequest with signature
}

func decodeBase64FBURL(s string) (string, os.Error) {
        s = s[:len(s)-2] // Facebook fix, last two chars are corrupt
        dbuf := make([]byte, base64.URLEncoding.DecodedLen(len(s)))
        _, err := base64.URLEncoding.Decode(dbuf, []byte(s))
        if err != nil {
                return "", err
        }
        // removed } is added at the end
        return string(dbuf) + "}", err
}
