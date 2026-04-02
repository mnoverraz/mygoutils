package strings

import (
	"strings"
	"unicode"
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

func (s String) Test() {}

func Coucou() {}

func Coucou2() {}

func Coucou3() {}

type String struct {
}
