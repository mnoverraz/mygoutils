package strings

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func GetFirstnameLastnameFromEmail(email string) (string, string) {
	firtnameDotLastname := strings.Split(email, "@")[0]
	names := strings.Split(firtnameDotLastname, ".")
	firstname := names[0]
	lastname := ""
	if len(names) == 1 {
		lastname = ""
	} else {
		lastname = names[1]
	}

	return firstname, lastname
}

func UppercaseFirstLetter(s string) string {
	// Check if the string is empty
	if s == "" {
		return ""
	}
	// Convert the first character to uppercase
	firstChar := []rune(s)[0]
	upperFirstChar := unicode.ToUpper(firstChar)

	// Replace the first character in the string with the uppercase version
	upperString := string(upperFirstChar) + s[1:]

	return upperString

}

// NoAccent returns the string without any accent.
func NoAccent(str string) (string, error) {
	var normalizer = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, err := transform.String(normalizer, str)
	if err != nil {
		return "", err
	}
	return strings.ToLower(s), err
}
