package facebook

// A subscription to an application to get real-time updates for an Graph object type.
// For more details, see the Real-time Overview.
type Subscription struct {
	// The object type to subscribe to.Available to everyone in Facebook by default. Contains code or permissions.
	Object string
	// The list of fields for the object type. Available to everyone in Facebook by default. Contains a comma-seperated list of field names.
	// Fields TODO
	// An endpoint on your domain which can handle the real-time notifications. Available to everyone in Facebook by default. Contains a valid URL.
	CallbackURL string
	// Whether or not the subscription is active or not. Available to everyone in Facebook by default.
	Active bool
}
