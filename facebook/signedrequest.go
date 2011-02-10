package facebook

import (
	"strings"
	"encoding/base64"
	"os"
	"json"
)

type SRUser struct {
	Locale  string
	Country string
	Age     *Age
}

type Age struct {
	Min float64
	Max float64
}

type SRPage struct {
	Id    string
	Liked bool
	Admin bool
}

type SignedRequest struct {
	User        *SRUser
	Algorithm   string
	Issued_at   float64
	User_id     string
	Oauth_token string
	Expires     float64
	Page        *SRPage
	Profile_id  float64
}

func ParseSignedRequest(input string) (sr *SignedRequest, err os.Error) {
	if len(input) == 0 {
		return sr, os.NewError("ParseSignedRequest: Input string is empty.")
	}
	points := strings.Count(input, ".")
	if points != 1 {
		return sr, os.NewError("Wrong input format. String must be of format <signature>.<data>")
	}
	d := strings.Split(input, ".", -1)
	//signature := DecodeBase64URL(d[0])
	data, err := base64Decode(d[1])
	if err != nil {
		return
	}
	var value SignedRequest
	if err = json.Unmarshal(data, &value); err != nil {
		return sr, os.NewError("Error: ParseSignedRequest: json.Unmarshal: " + err.String() + " in " + string(data))
	}
	sr = &value
	return
	// TODO: Check SignedRequest with signature
}

func base64Decode(str string) (dbuf []byte, err os.Error) {
	p := 4 - len(str)%4
	str = str + string(strings.Repeat("=", p))

	n := base64.URLEncoding.DecodedLen(len(str))
	dbuf = make([]byte, n)
	_, err = base64.URLEncoding.Decode(dbuf, []byte(str))
	if err != nil {
		// Try again to bypass some broken Facebook base64urls
		_, err = base64.URLEncoding.Decode(dbuf, []byte(str[:len(str)-8]))
		if err != nil {
			return
		}
		s := string(dbuf[0:n-p]) + "}"
		return []byte(s), err
	}
	return dbuf[0 : n-p], err
}
