package facebook

import (
	"fmt"
	"strings"
)

func GetAppAccessToken(appId, appSecret string) (oauth string, err error) {
	cmd := fmt.Sprintf("https://graph.facebook.com/oauth/access_token?client_id=%s&client_secret=%s&grant_type=client_credentials", appId, appSecret)
	resp, err := Get(cmd)
	if err != nil {
		return
	}
	data := strings.Split(string(resp.Data), "=")
	if len(data) >= 2 {
		if data[0] == "access_token" {
			return data[1], nil
		}
	}
	return
}
