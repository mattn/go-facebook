package facebook

import (
	"os"
	"fmt"
	"json"
)

type TestUserAPI struct {
	appId string
}

func NewTestUserAPI(app_id string) *TestUserAPI {
	return &TestUserAPI{appId: app_id}
}

func (t *TestUserAPI) CreateTestUser(installed bool, permissions []string, appAccessToken string) (user *TestUser, err os.Error) {
	var str string
	if len(permissions) > 0 {
		for _, perm := range permissions {
			str += perm + ","
		}
		str = str[:len(str)-1] // Remove last ,
	}
	params := map[string]string{"installed": fmt.Sprintf("%t", installed), "permissions": str, "access_token": appAccessToken}
	// The secure url is required
	url := SECURE + GRAPH_URL + "/" + t.appId + "/accounts/test-users"
	resp, err := Post(url, params)
	if err != nil {
		return
	}
	data := string(resp.Data)
	if data == "Error code: 2900 (Too many test accounts)" {
		return nil, os.NewError(data)
	}
	var value TestUser
	err = json.Unmarshal(resp.Data, &value)
	user = &value
	return
}

func (t *TestUserAPI) GetTestUsers() (tus *TestUsers, err os.Error) {
	url := SECURE + GRAPH_URL + "/" + t.appId + "/accounts/test-users"
	resp, err := Get(url)
	if err != nil {
		return
	}
	var value TestUsers
	err = json.Unmarshal(resp.Data, &value)
	tus = &value
	return
}

type TestUsers struct {
	Data []*TestUser
}

type TestUser struct {
	ID           string
	Access_Token string
	Login_Url    string
}
