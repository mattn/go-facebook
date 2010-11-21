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
	feed    string
	posts   string
	picture string
	// The photos, videos, and posts in which this application has been tagged. Publicly available. An array of Post, Photo or Video objects
	//Tagged TODO
	links  string
	photos string
	albums string
	/*
		// The application's status updates. Publicly available.
		// Statuses []Status // TODO
		// The videos this application has created. Publicly available.
		// Videos []Video // TODO
		// The application's notes. Publicly available.
		// Notes []Note // TODO
		// The events this page is managing. Publicly available.
		Events []Event
		// All of the subscriptions this application has for real-time notifications. Requires an application access token.
		// Subscriptions []Subscription // TODO
		// Usage metrics for this application. Requires an application access token.
		// Insights []Insight // TODO
	*/
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
