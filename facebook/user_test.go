package facebook

import (
	"testing"
	"fmt"
)

func TestGetUser(t *testing.T) {
	for _, id := range users {
		user, err := GetUser(id)
		if err != nil {
			t.Errorf("%s\n", err)
		}
		if user == nil {
			t.Errorf("User is nil of object: %s\n", user)
		}
		if len(user.ID) == 0 {
			t.Errorf("User.ID is empty of object: %s\n", user)
		}
	}
}
