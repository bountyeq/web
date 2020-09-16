package sanitize

import (
	"regexp"
	"strings"
)

var (
	cleanNameRegex = regexp.MustCompile(`[^0-9A-Za-z_]+`)
)

// CleanName cleans up a name
func CleanName(name string) string {
	out := strings.Replace(name, " ", "_", -1)
	out = strings.Replace(out, "#", "", -1)
	out = strings.TrimSpace(cleanNameRegex.ReplaceAllString(out, ""))
	out = strings.Replace(out, "_", " ", -1)
	return out
}
