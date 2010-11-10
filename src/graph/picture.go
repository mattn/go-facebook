package facebook

// Profile picture
type Picture struct {
	// Link to the profile picture
	URL string
}

func NewPicture(url string) (pic *Picture) {
	pic = new(Picture)
	pic.URL = url
	return
}
