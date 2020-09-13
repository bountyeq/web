package site

import "github.com/bountyeq/web/config"

// Site represents a website's core features
type Site struct {
	IsLoggedIn bool
	Title      string
}

// Fetch returns a site object
func Fetch() *Site {
	return &Site{
		IsLoggedIn: false,
		Title:      config.Title(),
	}
}
