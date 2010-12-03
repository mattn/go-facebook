package graph

import (
	//"strconv"
	"os"
	"time"
)

// A Facebook Page.
// http://developers.facebook.com/docs/reference/api/page
type Page struct {
	// ID
	ID string
	// Name
	Name string
	// Profile picture
	picture string // Documented as a connection
	// Category
	Category string
	// Number of fans the page has
	FanCount float64

	// Undocumented properties but streamed
	Link            string
	Website         string
	Username        string
	Founded         *time.Time
	CompanyOverview string
	Mission         string
	Products        string

	// Connections
	feed string
	// picture  string does not exist as a connection
	tagged   string
	links    string
	photos   string
	groups   string // Not yet implemented
	albums   string
	statuses string
	videos   string
	notes    string
	posts    string
	events   string
	checkins string
}

// Gets the page's wall. Available to everyone on Facebook.
// Returns an array of Post objects.
func (p *Page) GetWall() (ps []Post, err os.Error) {
	if len(p.feed) == 0 {
		err = os.NewError("Error: Page.GetWall with ID " + p.ID + " feed URL is empty.")
	}
	return fetchPosts(p.feed)
}

// Gets the page's profile picture. Publicly available.
// Returns an HTTP 302 URL string with the location set to the picture URL.
// NOTE: This is not a connection but documented as one.
func (p *Page) GetPicture() (pic *Picture, err os.Error) {
	if len(p.picture) == 0 {
		err = os.NewError("Error: Page.GetPicture with ID " + p.ID + " picture URL is empty.")
	}
	return NewPicture(p.picture), err
}

// Gets the photos, videos, and posts in which this page has been tagged. Publicly available.
// Returns an heterogeneous array of Photo, Video or Post objects.
func (p *Page) GetTagged() (t []interface{}, err os.Error) {
	if len(p.tagged) == 0 {
		err = os.NewError("Error: Page.GetTagged with ID " + p.ID + " tagged URL is empty.")
		return
	}
	data, err := getData(p.tagged)
	if err != nil {
		return
	}
	t = make([]interface{}, len(data))
	for i, v := range data {
		tag := v.(map[string]interface{})
		tp, ok := tag["type"].(string)
		if ok {
			switch tp {
			case "status":
				t[i], err = parsePost(tag)
				if err != nil {
					return
				}
			case "photo", "link":
				t[i], err = parsePhoto(tag)
				if err != nil {
					return
				}
			case "video":
				t[i], err = parseVideo(tag)
				if err != nil {
					return
				}
			default:
				err = os.NewError("Tag has unknown type.")
				return
			}
		} else {
			err = os.NewError("Tag is corrupted.")
			return
		}
	}
	return
}

// Gets the page's posted links. Publicly available.
// Returns an array of Link objects.
func (p *Page) GetLinks() (ls []Link, err os.Error) {
	if len(p.links) == 0 {
		err = os.NewError("Error: Page.GetLinks with ID " + p.ID + " links URL is empty.")
	}
	return getLinks(p.links)
}

// Gets the photos this page has uploaded. Publicly available.
// Returns an array of Photo objects.
func (p *Page) GetPhotos() (ps []Photo, err os.Error) {
	if len(p.photos) == 0 {
		err = os.NewError("Error: Page.GetPhotos with ID " + p.ID + " photos URL is empty.")
	}
	return getPhotos(p.photos)
}

// Gets the groups this page is a member of. Available to everyone on Facebook.
// Returns an array of objects containing group id, version, name and unread fields.
// TODO: function for GetGroups

// Gets the page albums this page has created. Publicly available.
// Returns an array of Album objects.
func (p *Page) GetAlbums() (as []Album, err os.Error) {
	if len(p.albums) == 0 {
		err = os.NewError("Error: Page.GetAlbums with ID " + p.ID + " albums URL is empty.")
	}
	return getAlbums(p.albums)
}

// Gets the page's status updates. Publicly available. Requires access token.
// Returns an array of StatusMessage objects.
func (p *Page) GetStatuses() (sms []StatusMessage, err os.Error) {
	if len(p.statuses) == 0 {
		err = os.NewError("Error: Page.GetStatuses with ID " + p.ID + " statuses URL is empty.")
	}
	return getStatusMessages(p.statuses)
}

// Gets the videos this page has created. Publicly available.
// Returns an array of Video objects.
func (p *Page) GetVideos() (vs []Video, err os.Error) {
	if len(p.videos) == 0 {
		err = os.NewError("Error: Page.GetVideos with ID " + p.ID + " videos URL is empty.")
	}
	return getVideos(p.videos)
}

// Gets the page's notes. Publicly available.
// Returns an array of Note objects.
func (p *Page) GetNotes() (ns []Note, err os.Error) {
	if len(p.notes) == 0 {
		err = os.NewError("Error: Page.GetNotes with ID " + p.ID + " notes URL is empty.")
	}
	return getNotes(p.notes)
}

// Gets the page's own posts. Publicly available.
// Returns an array of Post objects.
func (p *Page) GetPosts() (feed []Post, err os.Error) {
	if len(p.posts) == 0 {
		err = os.NewError("Error: Page.GetPosts with ID " + p.ID + " posts URL is empty.")
	}
	return fetchPosts(p.posts)
}

// Gets the events this page is managing. Publicly available.
// Returns an array of Event objects.
func (p *Page) GetEvents() (es []Event, err os.Error) {
	if len(p.events) == 0 {
		err = os.NewError("Error: Page.GetEvents with ID " + p.ID + " events URL is empty.")
	}
	return getEvents(p.events)
}

// Gets Checkins made by friends of the current session user. Requires friends_checkins permissions.
// Returns an array of Checkin objects.
func (p *Page) GetCheckins() (cs []Checkin, err os.Error) {
	if len(p.checkins) == 0 {
		err = os.NewError("Error: Page.GetCheckins with ID " + p.ID + " checkins URL is empty.")
	}
	return getCheckins(p.checkins)
}

func parsePage(data map[string]interface{}) (p Page, err os.Error) {
	for key, value := range data {
		switch key {
		case "website":
			p.Website = value.(string)
		case "picture":
			p.picture = value.(string)
		case "fan_count":
			p.FanCount = value.(float64)
		case "mission":
			p.Mission = value.(string)
		case "category":
			p.Category = value.(string)
		case "name":
			p.Name = value.(string)
		case "username":
			p.Username = value.(string)
		case "link":
			p.Link = value.(string)
		case "id":
			p.ID = value.(string)
		case "products":
			p.Products = value.(string)
		case "founded":
			p.Founded, err = parseTime(value.(string))
		case "company_overview":
			p.CompanyOverview = value.(string)
		case "metadata":
			metadata := value.(map[string]interface{})
			for k, va := range metadata["connections"].(map[string]interface{}) {
				switch k {
				case "feed":
					p.feed = va.(string)
				/*
					case "picture":
						p.picture = va.(string)
				*/
				case "tagged":
					p.tagged = va.(string)
				case "links":
					p.links = va.(string)
				case "photos":
					p.photos = va.(string)
				case "groups":
					p.groups = va.(string)
				case "albums":
					p.albums = va.(string)
				case "statuses":
					p.statuses = va.(string)
				case "videos":
					p.videos = va.(string)
				case "notes":
					p.notes = va.(string)
				case "posts":
					p.notes = va.(string)
				case "events":
					p.events = va.(string)
				case "checkins":
					p.checkins = va.(string)
				}
			}
		}
	}
	return
}
