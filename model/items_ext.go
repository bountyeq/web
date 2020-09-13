package model

import "strings"

// Level returns an item's estamated scoring level
func (i *Items) Level() string {
	return "10"
}

// IdfileClean returns an ID file without IT prefix
func (i *Items) IdfileClean() string {
	return strings.ReplaceAll(i.Idfile, "IT", "")
}
