package facebook

import (
	"os"
	"fmt"
	"json"
)

type TestUserAPI struct {
	appId          string
	appAccessToken string
}

func NewTestUserAPI(app_id string, appAccessToken string) *TestUserAPI {
	return &TestUserAPI{appId: app_id, appAccessToken: appAccessToken}
}

func (t *TestUserAPI) CreateTestUser(installed bool, permissions []string) (user *TestUser, err os.Error) {
	var str string
	if len(permissions) > 0 {
		for _, perm := range permissions {
			str += perm + ","
		}
		str = str[:len(str)-1] // Remove last ,
	}
	params := map[string]string{"installed": fmt.Sprintf("%t", installed), "permissions": str, "access_token": t.appAccessToken}
	// The secure url is required
	url := SECURE + GRAPH_URL + "/" + t.appId + "/accounts/test-users"
	resp, err := PostForm(url, params)
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
	url := SECURE + GRAPH_URL + "/" + t.appId + "/accounts/test-users?access_token=" + t.appAccessToken
	resp, err := Get(url)
	if err != nil {
		return
	}
	var value TestUsers
	err = json.Unmarshal(resp.Data, &value)
	tus = &value
	return
}

// TODO: TestUser Friend Requests
// TODO: TestUser deleting

type TestUsers struct {
	Data []*TestUser
}

type TestUser struct {
	ID           string
	Access_Token string
	Login_Url    string
}
