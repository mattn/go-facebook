package facebook

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
	Hometown string
	// Current location
	Location string
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
