package strings

import "strings"

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
