package site

// Site represents a website's core features
type Site struct {
	IsLoggedIn bool
}

// Fetch returns a site object
func Fetch() *Site {
	return &Site{
		IsLoggedIn: false,
	}
}
