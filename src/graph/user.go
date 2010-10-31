package facebook

import (
	"os"
	"time"
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
	Birthday *time.Time
	// Work history list
	Work []Workplace
	// Education history list
	Educations []Education
	// The contact email adress
	Email string
	// Link to the personal website
	Website string
	// Hometown
	Hometown Object
	// Current location
	Location Object
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

	// ##### Connections #####
	// TODO: Replace all strings with actual Connection structs
	// The News Feed. Requires read_stream permission
	Home Home
	// Wall. Requires read_stream permission to see non-public posts.
	Feed string
	// Photos, videos and posts in which the user has been tagged. Requires read_stream permission.
	Tagged string
	// Own posts. Requires read_stream permission to see non-public posts.
	Posts string
	// Profile picture
	Picture Picture
	// Friends of the user
	Friends string
	// Activities listed on the profile page
	Activities string
	// Interests listed on the profile page
	Interests string
	// Music listed on the profile page
	Music string
	// Books listed on the profile page
	Books string
	// Movies listed on the profile page
	Movies string
	// Television listed on the profile pages
	Television string
	// Pages this user has liked. Requires user_likes or friend_likes permission
	Likes string
	// Photos this user is tagged in. Requires user_photo_video_tags, friend_photo_video_tags and user_photos or friend_photos permissions
	Photos string
	// Photo albums this user has created. Requires user_photos or friend_photos permission
	Albums string
	// Videos this user has been tagged in. Requires user_videos or friend_videos permission
	Videos string
	// Groups this user is a member of. Requires user_groups or friend_groups permission
	Groups string
	// Status updates. Requires read_stream permission
	Statuses string
	// Posted links. Requires read_stream permission
	Links string
	// Notes. Requires read_stream permission
	Notes string
	// Events this user is attending. Requires user_events or friend_events permission
	Events string
	// Threads in this user's inbox. Requires read_mailbox permission
	InBox string
	// Messages in this user's outbox. Requires read_mailbox permission
	OutBox string
	// Updates in this user's inbox. Requires read_mailbox permission
	Updates string
	/* The Facebook pages owned by the current user. If the manage_pages permission has been granted,
	 * this connection also yields access_tokens that can be used to query the Graph API on behalf of the page.
	 */
	Accounts string
	// Places the current user has checked-into.
	Checkins string
	/* The user's outstanding requests for the app associated with the access token.
	 * The access token should be app secret signed and not user session signed. See more info here.
	 */
	PlatformRequests string

	// Not documented in the API but streamed
	Locale          string
	UpdatedTime     *time.Time
	FanCount        float64
	Mission         string
	Category        string
	Username        string
	Products        string
	Founded         *time.Time
	CompanyOverview string
}

func (u *User) String() string {
	return "ID: " + u.ID + "\tName: " + u.Name + "\tFirst name: " + u.FirstName +
		"\tLast name: " + u.LastName + "\tLink: " + u.Link + "\tGender: " +
		u.Gender + "\tLocale: " + u.Locale + "\tUpdated time: " + u.UpdatedTime.String() +
		"\n"
}

func FetchUser(name string) (user User, err os.Error) {
	body, err := fetchBody(name + "?metadata=1")
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
			user.Birthday, err = parseTime(value.(string))
		case "work":
			user.Work, err = parseWork(value.([]interface{}))
		case "education":
			user.Educations = parseEducations(value.([]interface{}))
		case "email":
			user.Email = value.(string)
		case "website":
			user.Website = value.(string)
		case "hometown":
			user.Hometown = parseObject(value.(map[string]interface{}))
		case "location":
			user.Location = parseObject(value.(map[string]interface{}))
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
			user.Founded, err = parseTime(value.(string))
		case "company_overview":
			user.CompanyOverview = value.(string)
		case "fan_count":
			user.FanCount = value.(float64)
		case "type":
			// TODO: Look into type

			// Parse metadata if requested
		case "metadata":
			// TODO: get and parse connections
			metadata := value.(map[string]interface{})
			for k, v := range metadata["connections"].(map[string]interface{}) {
				switch k {
				case "home":
					user.Home, err = FetchHomeByURL(v.(string)) // Pass URL
				}
			}
		default:
			debugInterface(value, key, "Person")
		}
	}
	return
}
