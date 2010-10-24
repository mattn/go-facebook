package facebook

import (
	"os"
)

// A user profile.
// http://developers.facebook.com/docs/reference/api/user
type User struct {
	// Identification
	ID string
	// First name
	FirstName string
	// Last name
	LastName string
	// Full name
	Name string
	// A link to the profile
	Link string
	// The Blurb that appears under the profile picture
	About string
	// Birthday
	Birthday string
	// Work history list
	Work string
	// Education history list
	Education string
	// The contact email adress
	Email string
	// Link to the personal website
	Website string
	// Hometown
	Hometown Town
	// Current location
	Location Town
	// Biography
	Bio string
	// Favorite quotes
	Quotes string
	// Gender	
	Gender string
	// Genders the user is interested in
	InterestedIn string
	// Types of relationships the user is seeking for
	MeetingFor string
	// Relationship status
	RelationshipStatus string
	// Religion
	Religion string
	// Political view
	Political string
	// Verification status
	Verified string
	// The user's significant other
	SignificantOther string
	// Timezone
	Timezone string

	// Connections
	Picture Picture

	// Not documented in the API but streamed probably Connections
	Locale          string
	UpdatedTime     string
	FanCount        float64
	Mission         string
	Category        string
	Username        string
	Products        string
	Founded         string
	CompanyOverview string
}

func (u *User) String() string {
	return "ID: " + u.ID + "\tName: " + u.Name + "\tFirst name: " + u.FirstName +
		"\tLast name: " + u.LastName + "\tLink: " + u.Link + "\tGender: " +
		u.Gender + "\tLocale: " + u.Locale + "\tUpdated time: " + u.UpdatedTime +
		"\n"
}

func FetchUserIntrospect(name string) (user User, err os.Error) {
	return FetchUser(name + "?metadata=1")
}

func FetchUser(name string) (user User, err os.Error) {
	body, err := fetchBody(name)
	if err != nil {
		return
	}
	data, err := getJsonMap(body)
	if err != nil {
		return
	}
	for key, value := range data {
		switch key {
		case "id":
			user.ID = value.(string)
		case "first_name":
			user.FirstName = value.(string)
		case "last_name":
			user.LastName = value.(string)
		case "name":
			user.Name = value.(string)
		case "link":
			user.Link = value.(string)
		case "about":
			user.About = value.(string)
		case "birthday":
			user.Birthday = value.(string)
		case "work":
			user.Work = value.(string)
		case "education":
			user.Education = value.(string)
		case "email":
			user.Email = value.(string)
		case "website":
			user.Website = value.(string)
		case "hometown":
			user.Hometown = parseTown(value.(map[string]interface{}))
		case "location":
			user.Location = parseTown(value.(map[string]interface{}))
		case "bio":
			user.Bio = value.(string)
		case "quotes":
			user.Quotes = value.(string)
		case "gender":
			user.Gender = value.(string)
		case "interested_in":
			user.InterestedIn = value.(string)
		case "meeting_for":
			user.MeetingFor = value.(string)
		case "relationship_status":
			user.RelationshipStatus = value.(string)
		case "religion":
			user.Religion = value.(string)
		case "political":
			user.Political = value.(string)
		case "verified":
			user.Verified = value.(string)
		case "significant_other":
			user.SignificantOther = value.(string)
		case "timezone":
			user.Timezone = value.(string)

		// Connections
		case "picture":
			user.Picture = NewPicture(value.(string))

		// Not documented in the API but streamed	
		case "locale":
			user.Locale = value.(string)
		case "mission":
			user.Mission = value.(string)
		case "category":
			user.Category = value.(string)
		case "username":
			user.Username = value.(string)
		case "products":
			user.Products = value.(string)
		case "founded":
			user.Founded = value.(string)
		case "company_overview":
			user.CompanyOverview = value.(string)
		case "fan_count":
			user.FanCount = value.(float64)
		case "type":
			// TODO: Look into type

			// Parse metadata if requested
		case "metadata":
			parseMetaData(value)
		default:
			debugInterface(value, key, "Person")
		}
	}
	return
}
