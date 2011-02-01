package facebook

import (
	"json"
	"os"
)

// A user profile.
// http://developers.facebook.com/docs/reference/api/user
type User struct {
	// Identification
	ID string
	// First name
	First_Name string
	// Last name
	Last_Name string
	// Full name
	Name string
	// A link to the profile
	Link string
	// The Blurb that appears under the profile picture
	About string
	// Birthday
	Birthday string
	// Work history list
	Work []*Workplace
	// Education history list
	//Educations []Education
	// The contact email adress
	Email string
	// Link to the personal website
	Website string
	// Hometown
	Hometown *Object
	// Current location
	Location *Object
	// Biography
	Bio string
	// Favorite quotes
	Quotes string
	// Gender
	Gender string
	// Genders the user is interested in
	Interested_In string
	// Types of relationships the user is seeking for
	Meeting_For string
	// Relationship status
	Relationship_Status string
	// Religion
	Religion string
	// Political view
	Political string
	// Verification status
	Verified string
	// The user's significant other
	Significant_Other string
	// Timezone
	Timezone string
	// An anonymous, but unique identifier for the user. Available to everyone on Facebook.
	Third_Party_ID string
	// The last time the user's profile was updated. Available to everyone on Facebook.
	Updated_Time string
	// The user's locale. Publicly available. A JSON string containing the ISO Language Code and ISO Country Code.
	Locale string
}

func GetUser(id string) (user *User, err os.Error) {
	resp, err := Call(id, map[string]string{})
	if err != nil {
		return
	}
	var value User
	err = json.Unmarshal(resp.Data, &value)
	user = &value
	return
}
