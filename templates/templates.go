package templates

import (
	"fmt"
	"html/template"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/mnoverraz/mygoutils/strings"
)

// New create an HTML template and adds:
//
//   - sprig (http://masterminds.github.io/sprig/)
//   - specials functions https://github.com/mnoverraz/mygoutils/blob/main/templates/templates.go#L18
func New(name string) *template.Template {
	return template.New(name).Funcs(sprig.FuncMap()).Funcs(PersoFuncMap())
}

func PersoFuncMap() template.FuncMap {

	return template.FuncMap{
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
		// implemented from strings
		"noAccent":             strings.NoAccent,
		"upperCaseFirstLetter": strings.UppercaseFirstLetter,
		"isFirstLetterVowel":   isFirstLetterVowel,
		"lowercaseFirstLetter": strings.LowercaseFirstLetter,
		"humanDuration":        HumanDuration,
	}
}

func isFirstLetterVowel(word string) bool {

	firstLetterFromString := []rune(word)[0]

	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, value := range vowels {
		if value == firstLetterFromString {
			return true
		}
	}
	return false
}

func HumanDuration(d time.Duration) string {
	days := d / (24 * time.Hour)
	d -= days * 24 * time.Hour

	hours := d / time.Hour
	d -= hours * time.Hour

	minutes := d / time.Minute
	d -= minutes * time.Minute

	seconds := d / time.Second

	if days > 0 {
		return fmt.Sprintf("%dj %dh %dm", days, hours, minutes)
	}

	if hours > 0 {
		return fmt.Sprintf("%dh %dm", hours, minutes)
	}

	if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, seconds)
	}

	return fmt.Sprintf("%ds", seconds)
}
