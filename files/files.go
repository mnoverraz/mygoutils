package files

import (
	"fmt"
	"os"
)

func WriteFile(content []byte, filename string) int {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer f.Close()

	l, err := f.Write(content)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return 0
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return l
}
