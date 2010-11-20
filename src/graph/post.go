package graph

import (
	"os"
	"time"
)

/*
 * An individual entry in a profile's feed.
 * The read_stream extended permission is required to access any information in a profile's feed that is not shared with everyone.
 */
type Post struct {
	// The post ID
	ID string
	// An object containing the ID and name of the user who posted the message
	From Object
	// A list of the profiles mentioned or targeted in this post
	To []Object
	// The message
	Message string
	// If available, a link to the picture included with this post
	Picture *Picture
	// The link attached to this post
	Link string
	// The name of the link
	Name string
	// The caption of the link (appears beneath the link name)
	Caption string
	// A description of the link (appears beneath the link caption)
	Description string
	// If available, the source of the stream attachment to add to this post (for example, a flash file or image)
	Source string
	// A link to an icon representing the type of this post
	Icon string
	// A string indicating which application was used to create this post
	Attribution string
	/* A list of available actions on the post (including commenting, liking, and an optional app-specified action),
	 * encoded as objects with keys for the 'name' and 'link'.
	 */
	Actions []URL
	/* 
	 * An object that defines the privacy setting for a post, video, or album. It contains the following fields:
	 * value(string) : The privacy value for the object, specify one of EVERYONE, CUSTOM, ALL_FRIENDS, NETWORKS_FRIENDS, FRIENDS_OF_FRIENDS.
	 * friends(string) : For CUSTOM settings, this indicates which users can see the object. Can be one of EVERYONE, NETWORKS_FRIENDS (when the object can be seen by networks and friends), FRIENDS_OF_FRIENDS, ALL_FRIENDS, SOME_FRIENDS, SELF, or NO_FRIENDS (when the object can be seen by a network only).
	 * networks(string) : For CUSTOM settings, specify a comma-separated list of network IDs that can see the object, or 1 for all of a user's networks.
	 * allow(string) : When friends is set to SOME_FRIENDS, specify a comma-separated list of user IDs and friend list IDs that 'can' see the post.
	 * deny(string) : When friends is set to SOME_FRIENDS, specify a comma-separated list of user IDs and friend list IDs that 'cannot' see the post. 
	 * 
	 * Note: This privacy setting only applies to posts to the current or specified user's own Wall; Facebook ignores this setting for targeted Wall posts (when the user is writing on the Wall of a friend, Page, event, group connected to the user). Consistent with behavior on Facebook, all targeted posts are viewable by anyone who can see the target's Wall.
	 * Privacy Policy: Any non-default privacy setting must be intentionally chosen by the user. You may not set a custom privacy setting unless the user has proactively specified that they want this non-default setting.
	 */
	Privacy Object //?
	// The number of likes on this post
	Likes float64
	// The time the post was initially published
	CreatedTime *time.Time
	// The time of the last comment on this post
	UpdatedTime *time.Time

	// Connections
	// All of the comments on this post (this is no real connection, data is passed with the post)
	Comments []Comment
}

func fetchPosts(url string) (posts []Post, err os.Error) {
	d, err := getObjByURL(url)
	if err != nil {
		return
	}
	for key, value := range d {
		switch key {
		case "data":
			data := value.([]interface{})
			posts = make([]Post, len(data))
			for i, val := range data {
				var post Post
				post, err = parsePost(val.(map[string]interface{}))
				posts[i] = post
			}
		case "paging":
		}
	}
	return
}

/*
 * Parses Post data. Returns nil for err if no error appeared.
 */
func parsePost(value map[string]interface{}) (p Post, err os.Error) {
	for key, val := range value {
		switch key {
		case "id":
			p.ID = val.(string)
		case "from":
			p.From = parseObject(val.(map[string]interface{}))
		case "to":
			data := val.(map[string]interface{})
			p.To = parseObjects(data["data"].([]interface{}))
		case "message":
			p.Message = val.(string)
		case "picture":
			p.Picture = NewPicture(val.(string))
		case "link":
			p.Link = val.(string)
		case "name":
			p.Name = val.(string)
		case "caption":
			p.Caption = val.(string)
		case "description":
			p.Description = val.(string)
		case "source":
			p.Source = val.(string)
		case "icon":
			p.Icon = val.(string)
		case "attribution":
			p.Attribution = val.(string)
		case "actions":
			p.Actions = parseURLs(val.([]interface{}))
		case "privacy":
			// TODO: Privacy				
		case "likes":
			p.Likes = val.(float64)
		case "created_time":
			p.CreatedTime, err = parseTime(val.(string))
		case "updated_time":
			p.UpdatedTime, err = parseTime(val.(string))
		// Connections
		case "comments":
			data := val.(map[string]interface{})
			p.Comments, _ = parseComments(data)
		}
	}
	return
}
