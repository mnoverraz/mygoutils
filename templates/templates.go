package templates

import (
	"html/template"

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
