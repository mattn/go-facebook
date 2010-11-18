package graph

const (
	PRIVACY_EVERYONE           = "EVERYONE"
	PRIVACY_CUSTOM             = "CUSTOM"
	PRIVACY_SELF               = "SELF"
	PRIVACY_ALL_FRIENDS        = "ALL_FRIENDS"
	PRIVACY_SOME_FRIENDS       = "SOME_FRIENDS"
	PRIVACY_NO_FRIENDS         = "NO_FRIENDS"
	PRIVACY_NETWORKS_FRIENDS   = "NETWORKS_FRIENDS"
	PRIVACY_FRIENDS_OF_FRIENDS = "FRIENDS_OF_FRIENDS"
)


/* 
 * An object that defines the privacy setting for a post, video, or album.
 * Note: This privacy setting only applies to posts to the current or specified user's own Wall; Facebook ignores this setting for targeted Wall posts (when the user is writing on the Wall of a friend, Page, event, group connected to the user). Consistent with behavior on Facebook, all targeted posts are viewable by anyone who can see the target's Wall.
 * Privacy Policy: Any non-default privacy setting must be intentionally chosen by the user. You may not set a custom privacy setting unless the user has proactively specified that they want this non-default setting.
 */
type Privacy struct {
	// The privacy value for the object, specify one of EVERYONE, CUSTOM, ALL_FRIENDS, NETWORKS_FRIENDS, FRIENDS_OF_FRIENDS.
	Value string
	// For CUSTOM settings, this indicates which users can see the object. Can be one of EVERYONE, NETWORKS_FRIENDS (when the object can be seen by networks and friends), FRIENDS_OF_FRIENDS, ALL_FRIENDS, SOME_FRIENDS, SELF, or NO_FRIENDS (when the object can be seen by a network only).
	Friends string
	// For CUSTOM settings, specify a comma-separated list of network IDs that can see the object, or 1 for all of a user's networks.
	Networks string
	// When friends is set to SOME_FRIENDS, specify a comma-separated list of user IDs and friend list IDs that 'can' see the post.
	Allow string
	// When friends is set to SOME_FRIENDS, specify a comma-separated list of user IDs and friend list IDs that 'cannot' see the post. 
	Deny string

	Description string // Not documented in the API
}

func parsePrivacy(value map[string]interface{}) (privacy Privacy) {
	privacy.Value = value["value"].(string)
	privacy.Friends = value["friends"].(string)
	privacy.Networks = value["networks"].(string)
	privacy.Allow = value["allow"].(string)
	privacy.Deny = value["deny"].(string)
	privacy.Description = value["description"].(string)
	return
}
