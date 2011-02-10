package facebook

import (
	"testing"
)

// Tests

type Base64URLTest struct {
	Encoded string
	Decoded string
}

var Base64URLTests = []Base64URLTest{
	{"eyJhbGdvcml0aG0iOiJITUFDLVNIQTI1NiIsIjAiOiJwYXlsb2FkIn0", "{\"algorithm\":\"HMAC-SHA256\",\"0\":\"payload\"}"},
}

type SignedRequestTest struct {
	Data string
	Sr   SignedRequest
}

var SignedRequestTests = []SignedRequestTest{
	{"vlXgu64BQGFSQrY0ZcJBZASMvYvTHu9GQ0YM9rjPSso.eyJhbGdvcml0aG0iOiJITUFDLVNIQTI1NiIsIjAiOiJwYXlsb2FkIn0", SignedRequest{nil, "HMAC-SHA256", 0, "", "", 0, nil, 0}},
}

func TestDecoding(t *testing.T) {
	for _, v := range Base64URLTests {
		dec, err := base64Decode(v.Encoded)
		if err != nil {
			t.Error(err.String())
		}
		if string(dec) != v.Decoded {
			t.Errorf("%s is not expected value %s", dec, v.Decoded)
		}
	}
}

func TestSignedRequest(t *testing.T) {
	for _, test := range SignedRequestTests {
		sr, err := ParseSignedRequest(test.Data)
		if err != nil {
			t.Error(err.String())
		}
		if sr.Algorithm != test.Sr.Algorithm {
			t.Errorf("%s is not expected value %s", sr.Algorithm, test.Sr.Algorithm)
		}
	}
}
