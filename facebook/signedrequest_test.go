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
	{"eyJhbGdvcml0aG0iOiJITUFDLVNIQTI1NiIsImlzc3VlZF9hdCI6MTI5NzM3OTAyNiwidXNlciI6eyJjb3VudHJ5IjoiZGUiLCJsb2NhbGUiOiJlbl9VUyIsImFnZSI6eyJtaW4iOjIxfX19",
		"{\"algorithm\":\"HMAC-SHA256\",\"issued_at\":1297379026,\"user\":{\"country\":\"de\",\"locale\":\"en_US\",\"age\":{\"min\":21}}}"},
}

type SignedRequestTest struct {
	Data string
	Sr   SignedRequest
}

var SignedRequestTests = []SignedRequestTest{
	{"vlXgu64BQGFSQrY0ZcJBZASMvYvTHu9GQ0YM9rjPSso.eyJhbGdvcml0aG0iOiJITUFDLVNIQTI1NiIsIjAiOiJwYXlsb2FkIn0", SignedRequest{nil, "HMAC-SHA256", 0, "", "", 0, nil, 0}},
	{"D2Jp6a-RkTnPG0lDUl5PT1Ob1FK9vQTH9srKrgqdcK4.eyJhbGdvcml0aG0iOiJITUFDLVNIQTI1NiIsImlzc3VlZF9hdCI6MTI5NzM3OTAyNiwidXNlciI6eyJjb3VudHJ5IjoiZGUiLCJsb2NhbGUiOiJlbl9VUyIsImFnZSI6eyJtaW4iOjIxfX19",
		SignedRequest{&SRUser{"de", "en_US", &Age{Min: 21}}, "HMAC-SHA256", 1297379026, "", "", 0, nil, 0},
	},
}

func TestDecoding(t *testing.T) {
	for _, v := range Base64URLTests {
		dec, err := base64Decode(v.Encoded)
		if err != nil {
			t.Fatalf("Error with %s : %s", v.Encoded, err.String())
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
			t.Fatalf(err.String())
		}
		if sr.Algorithm != test.Sr.Algorithm {
			t.Errorf("%s is not expected value %s", sr.Algorithm, test.Sr.Algorithm)
		}
	}
}
