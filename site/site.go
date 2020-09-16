package site

import "github.com/bountyeq/web/config"

// Site represents a website's core features
type Site struct {
	IsLoggedIn bool
	Title      string
	Page       *Page
}

// Page represents pagination
type Page struct {
	Start           string
	StartDisplay    string
	Next            string
	NextDisplay     string
	Previous        string
	PreviousDisplay string
	CurrentDisplay  string
}

// Fetch returns a site object
func Fetch() *Site {
	return &Site{
		IsLoggedIn: false,
		Title:      config.Title(),
		Page:       &Page{},
	}
}

// Generate will generate pages
func (p *Page) Generate(start string, startDisplay string, previous string, previousDisplay string, currentDisplay string, next string, nextDisplay string) {
	p.Start = start
	p.StartDisplay = startDisplay
	p.Previous = previous
	p.PreviousDisplay = previousDisplay
	p.CurrentDisplay = currentDisplay
	p.Next = next
	p.NextDisplay = nextDisplay
}
