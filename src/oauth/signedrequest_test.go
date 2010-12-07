package oauth

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

func TestDecoding(t *testing.T) {
	for _, v := range Base64URLTests {
		t.Logf("Decoding %s\n", v.Encoded)
		_, err := DecodeBase64URL(v.Encoded)
		if err != nil {
			t.Errorf(err.String())
		}
	}
}
