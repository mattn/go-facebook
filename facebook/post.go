package facebook

import "encoding/json"

/*
* An individual entry in a profile's feed.
* The read_stream extended permission is required to access any information in a profile's feed that is not shared with everyone.
 */
type Post struct {
	// The number of likes.
	//Likes string
	// The post ID
	ID string
	// An object containing the ID and name of the user who posted the message
	From *Object
	// A list of the profiles mentioned or targeted in this post
	To *ObjectCollection
	// The message
	Message string
	// If available, a link to the picture included with this post
	Picture string
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
	//Actions []URL //TODO: name, link
	// TODO: Privacy
	// The time the post was initially published
	Created_Time string
	// The time of the last comment on this post
	Updated_Time string
	// TODO: Targeting 

	// Metadata contains Connections
	*Metadata
}

func GetPost(id string) (post *Post, err error) {
	resp, err := Call(id, RequestMetadata)
	if err != nil {
		return
	}
	var value Post
	err = json.Unmarshal(resp.Data, &value)
	post = &value
	return
}
