package facebook

import (
	"strings"
	"encoding/base64"
	"os"
	"json"
	//"bytes"
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
	data = data[:len(data)] // TEMP: Bypass an illegal char at the end
	var value SignedRequest
	if err = json.Unmarshal(data, &value); err != nil {
		return sr, os.NewError("Error: ParseSignedRequest: json.Unmarshal: " + err.String() + " in " + string(data))
	}
	sr = &value
	return
	// TODO: Check SignedRequest with signature
}

func base64Decode(str string) ([]byte, os.Error) {
	p := 4 - len(str)%4
	str = str + string(strings.Repeat("=", p))

	dbuf := make([]byte, base64.URLEncoding.DecodedLen(len(str)))
	_, err := base64.URLEncoding.Decode(dbuf, []byte(str))
	if err != nil {
		return []byte{}, err
	}
	return dbuf[0 : len(dbuf)-p], err
}
