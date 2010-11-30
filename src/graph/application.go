package graph

import (
	"os"
)

/*
 * An application registered on Facebook Platform.
 * The Graph API supports querying for information on existing applications.
 * To create, administer or delete applications developers must go to the Developer Application
 */
type Application struct {
	// The application ID. Publicly available.
	ID string
	// The title of the application. Publicly available.
	Name string
	// The description of the application written by the 3rd party developers. Publicly available.
	Description string
	// The category of the application. Publicly available.
	Category string
	// A link to application dashboard on Facebook. Publicly available. Contains an URL.
	Link string

	// Connections
	feed          string
	posts         string
	picture       string
	tagged        string
	links         string
	photos        string
	albums        string
	statuses      string
	videos        string
	notes         string
	events        string
	subscriptions string
	insights      string
}

// Gets the application's wall posts. Publicly available.
// Returns an array of Post objects.
func (a *Application) GetFeed() (feed []Post, err os.Error) {
	if a.feed == "" {
		err = os.NewError("Error: Application.GetFeed: The feed URL is empty.")
	}
	return fetchPosts(a.feed)
}

// Gets the applications's own posts. Publicly available.
// Returns an array of Post objects.
func (a *Application) GetPosts() (feed []Post, err os.Error) {
	if a.posts == "" {
		err = os.NewError("Error: Application.GetPosts: The posts URL is empty.")
	}
	return fetchPosts(a.posts)
}

// Gets the application's logo with maximum dimensions of 75x75 pixels suitable for embedding as the source of an image tag.
// Publicly available. Returns an HTTP 302 URL string with the location set to the picture URL.
func (a *Application) GetPicture() (pic *Picture, err os.Error) {
	if a.picture == "" {
		err = os.NewError("Error: Application.GetPicture: The picture URL is empty.")
	}
	return NewPicture(a.picture), err
}

// Gets the photos, videos, and posts in which this application has been tagged. Publicly available.
// Returns an array of Post, Photo or Video objects.
func (a *Application) GetTagged() (t []interface{}, err os.Error) {
	if a.tagged == "" {
		err = os.NewError("Error: Application.GetTagged: The tagged URL is empty.")
	}
	data, err := getData(a.tagged)
	if err != nil {
		return
	}
	t = make([]interface{}, len(data))
	for i, v := range data {
		tag := v.(map[string]interface{})
		switch tag["type"].(string) {
		case "status":
			t[i], err = parsePost(tag)
			if err != nil {
				return
			}
		case "link":
		case "photo":
			t[i], err = parsePhoto(tag)
			if err != nil {
				return
			}
		case "video":
			t[i], err = parseVideo(tag)
			if err != nil {
				return
			}
		}
	}
	return
}

// Gets the application's posted links. Publicly available.
// Returns an array of Link objects.
func (a *Application) GetLinks() (ls []Link, err os.Error) {
	if a.links == "" {
		err = os.NewError("Error: Application.GetLinks: The links URL is empty.")
	}
	return getLinks(a.links)
}

// Gets the photos this application has uploaded. Publicly available.
// Returns an array of Photo objects.
func (a *Application) GetPhotos() (ps []Photo, err os.Error) {
	if a.photos == "" {
		err = os.NewError("Error: Application.GetPhotos: The photos URL is empty.")
	}
	return getPhotos(a.photos)
}

// Gets the photo albums this page has created. Publicly available.
// Returns an array of Album objects.
func (a *Application) GetAlbums() (as []Album, err os.Error) {
	if a.albums == "" {
		err = os.NewError("Error: Application.GetAlbums: The albums URL is empty.")
	}
	return getAlbums(a.albums)
}

// Gets the application's status updates. Publicly available.
// Returns an array of StatusMessage objects.
func (a *Application) GetStatuses() (sms []StatusMessage, err os.Error) {
	if a.statuses == "" {
		err = os.NewError("Error: Application.GetStatuses: The statuses URL is empty.")
	}
	return getStatusMessages(a.statuses)
}

// Gets the videos this application has created. Publicly available.
// Returns an array of Video objects.
func (a *Application) GetVideos() (vs []Video, err os.Error) {
	if a.videos == "" {
		err = os.NewError("Error: Application.GetVideos: The videos URL is empty.")
	}
	return getVideos(a.videos)
}

// Gets the application's notes. Publicly available.
// Returns an array of Note objects.
func (a *Application) GetNotes() (ns []Note, err os.Error) {
	if a.notes == "" {
		err = os.NewError("Error: Application.GetNotes: The notes URL is empty.")
	}
	return getNotes(a.notes)
}

// Gets the events this page is managing. Publicly available.
// Returns an array of Event objects.
func (a *Application) GetEvents() (es []Event, err os.Error) {
	if a.events == "" {
		err = os.NewError("Error: Application.GetEvents: The events URL is empty.")
	}
	return getEvents(a.events)
}

// Gets all of the subscriptions this application has for real-time notifications. Requires an application access token.
// Returns an array of Subscription objects.
func (a *Application) GetSubscriptions() (s []Subscription, err os.Error) {
	if a.subscriptions == "" {
		err = os.NewError("Error: Application.GetSubscriptions: The subscriptions URL is empty.")
	}
	return getSubscriptions(a.subscriptions)
}

// Gets the usage metrics for this application. Requires an application access token.
// Returns an array of Insights objects.
func (a *Application) GetInsights() (is Insights, err os.Error) {
	if a.insights == "" {
		err = os.NewError("Error: Application.GetInsights: The insights URL is empty.")
	}
	return getInsights(a.insights)
}

func parseApplication(value map[string]interface{}) (app Application, err os.Error) {
	for key, val := range value {
		switch key {
		case "id":
			app.ID = val.(string)
		case "name":
			app.Name = val.(string)
		case "description":
			app.Description = val.(string)
		case "category":
			app.Category = val.(string)
		case "link":
			app.Link = val.(string)
			// Connections
			/*
				case "metadata":
					metadata := val.(map[string]interface{})
					for k, v := range metadata["connections"].(map[string]interface{}) {
						switch k {
						case "feed":
							app.Feed, err = GetPosts(v.(string))
						case "posts":
							app.Posts, err = GetPosts(v.(string))
						case "picture":
							app.Picture = NewPicture(v.(string))
						case "tagged":
							// TODO:
						case "links":
							// TODO
						case "events":
						}
					}
			*/
		}
	}
	return
}
