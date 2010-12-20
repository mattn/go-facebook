package graph

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
	// An anonymous, but unique identifier for the user. Available to everyone on Facebook.
	ThirdPartyID string
	// The last time the user's profile was updated. Available to everyone on Facebook.
	LastUpdated *time.Time
	// The user's locale. Publicly available. A JSON string containing the ISO Language Code and ISO Country Code.
	Locale string
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
	UpdatedTime     *time.Time
	FanCount        float64
	Mission         string
	Category        string
	Username        string
	Products        string
	Founded         *time.Time
	CompanyOverview string

	// ##### Connections #####
	home             string
	feed             string
	tagged           string
	posts            string
	picture          string
	friends          string
	activities       string
	interests        string
	music            string
	books            string
	television       string
	likes            string
	photos           string
	albums           string
	videos           string
	groups           string
	statuses         string
	links            string
	notes            string
	events           string
	inbox            string
	outbox           string
	updates          string
	accounts         string
	checkins         string
	platformrequests string
	friendlists      string
}

func (u *User) String() string {
	return "ID: " + u.ID + "\tName: " + u.Name + "\tFirst name: " + u.FirstName +
		"\tLast name: " + u.LastName + "\tLink: " + u.Link + "\tGender: " +
		u.Gender + "\tLocale: " + u.Locale + "\tUpdated time: " + u.UpdatedTime.String() +
		"\n"
}

func parseUser(data map[string]interface{}) (user User, err os.Error) {
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
		case "third_party_id":
			user.ThirdPartyID = value.(string)
		case "last_updated":
			user.LastUpdated, err = parseTime(value.(string))
		case "locale":
			user.Locale = value.(string)

		// Not documented in the API but streamed
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
			// Ignore type here
			// Connections
		case "metadata":
			metadata := value.(map[string]interface{})
			for k, va := range metadata["connections"].(map[string]interface{}) {
				switch k {
				case "home":
					user.home = va.(string)
				case "feed":
					user.feed = va.(string)
				case "tagged":
					user.tagged = va.(string)
				case "posts":
					user.posts = va.(string)
				case "picture":
					user.picture = va.(string)
				case "friends":
					user.friends = va.(string)
				case "activities":
					user.activities = va.(string)
				case "interests":
					user.interests = va.(string)
				case "music":
					user.music = va.(string)
				case "books":
					user.books = va.(string)
				case "television":
					user.television = va.(string)
				case "likes":
					user.likes = va.(string)
				case "photos":
					user.photos = va.(string)
				case "albums":
					user.albums = va.(string)
				case "videos":
					user.videos = va.(string)
				case "groups":
					user.groups = va.(string)
				case "statuses":
					user.statuses = va.(string)
				case "links":
					user.links = va.(string)
				case "notes":
					user.notes = va.(string)
				case "events":
					user.events = va.(string)
				case "inbox":
					user.inbox = va.(string)
				case "outbox":
					user.outbox = va.(string)
				case "updates":
					user.updates = va.(string)
				case "accounts":
					user.accounts = va.(string)
				case "checkins":
					user.checkins = va.(string)
				case "platformrequests":
					user.platformrequests = va.(string)
				case "friendlists":
					user.friendlists = va.(string)
				default:
					// TODO: Print/log unsupported field
				}
			}
		}
	}
	return
}
