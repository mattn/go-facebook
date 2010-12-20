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

	// All of the comments on this post, not documented
	Comments []Comment

	// Connections
	comments string
	likes    string
}

func fetchPosts(url string) (posts []Post, err os.Error) {
	resp, err := GetResponse(url)
	if err != nil || resp.Fail {
		return
	}
	data := resp.Data
	posts = make([]Post, len(data))
	for i, val := range data {
		var post Post
		post, err = parsePost(val.(map[string]interface{}))
		posts[i] = post
	}
	return
}

// Gets all of the comments on this post. Available to everyone on Facebook.
// Returns an array of objects containing id, from, message and created_time fields.
func (p *Post) GetComments() (cs []Comment, err os.Error) {
	if p.comments == "" {
		err = os.NewError("Error: Post.GetComments: The comments URL is empty.")
	}
	return getComments(p.comments)
}

// Gets the likes on this post. Available to everyone on Facebook.
// Returns an array of objects containing the id and name fields.
func (p *Post) GetLikes() (likes []Object, err os.Error) {
	if p.likes == "" {
		err = os.NewError("Error: Post.GetLikes: The likes URL is empty.")
	}
	resp, err := GetResponse(p.likes)
	if err != nil || resp.Fail {
		return
	}
	likes = parseObjects(resp.Data)
	return
}

/*
 * Parses Post data. Returns nil for err if no error appeared.
 */
func parsePost(value map[string]interface{}) (p Post, err os.Error) {
	t, ok := value["type"].(string)
	if ok { // If ok we can do a type check
		if t != "post" {
			return p, os.NewError(value["id"].(string) + " is not a Post object it is a " + t + " object.")
		}
	}
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
		case "comments":
			data := val.(map[string]interface{})
			p.Comments, _ = parseComments(data)
		// Connections
		case "metadata":
			metadata := val.(map[string]interface{})
			for k, v := range metadata["connections"].(map[string]interface{}) {
				switch k {
				case "comments":
					p.comments = v.(string)
				case "likes":
					p.likes = v.(string)
				}
			}
		}
	}
	return
}

// ### Publishing ###
// Requires read_stream and publish_stream permission
func PublishPost(id, accessToken, message, link, picture, name, caption, description, privacy string, actions []URL) (err os.Error) {
	data := make(map[string]string)
	data["access_token"] = accessToken
	data["message"] = message
	data["link"] = link
	data["picture"] = picture
	data["name"] = name
	data["caption"] = caption
	data["description"] = description
	data["actions"] += "["
	for i, v := range actions {
		data["actions"] += "{\"name\": \"" + v.Name + "\", \"link\": \"" + v.URL + "\"}"
		if i < len(actions) {
			data["actions"] += ","
		}
	}
	data["actions"] += "]"
	data["privacy"] = privacy
	err = post(GRAPHURL+id+"/feed", data)
	return
}
