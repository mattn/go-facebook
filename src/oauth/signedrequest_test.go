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
	{"eyJhbGdvcml0aG0iOiJITUFDLVNIQTI1NiIsImV4cGlyZXMiOjEyOTI1NDQwMDAsImlzc3VlZF9hdCI6MTI5MjUzODkwNywib2F1dGhfdG9rZW4iOiIxMDE0Nzc4ODY1OTIwNjN8Mi51MEpKd0ZOZHhldHptQXRFNEZfbl9RX18uMzYwMC4xMjkyNTQ0MDAwLTEwMDAwMTE3NzU5MDAyMXx0NzhubmZvX29Cb2p6S1FMd09mejd2R0toMUUiLCJ1c2VyIjp7ImxvY2FsZSI6ImVuX1VTIiwiY291bnRyeSI6ImRlIn0sInVzZXJfaWQiOiIxMDAwMDExNzc1OTAwMjEifQ", "{\"algorithm\":\"HMAC-SHA256\",\"expires\":1292544000,\"issued_at\":1292538907,\"oauth_token\":\"101477886592063|2.u0JJwFNdxetzmAtE4F_n_Q__.3600.1292544000-100001177590021|t78nnfo_oBojzKQLwOfz7vGKh1E\",\"user\":{\"locale\":\"en_US\",\"country\":\"de\"},\"user_id\":\"100001177590021\"}"},
}
func TestDecoding(t *testing.T) {
	for _, v := range Base64URLTests {
		t.Logf("Decoding %s\n", v.Encoded)
		_, err := decodeBase64FBURL(v.Encoded)
		if err != nil {
			t.Errorf(err.String())
		}
	}
}
