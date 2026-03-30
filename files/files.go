package files

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mnoverraz/mygoutils/system"
)

// WriteFile take a pathFile in arg and create the
// directories if not exists and write the content in buf
// in the file
func WriteFile(buf []byte, pathFile string) error {
	path := filepath.Dir(pathFile)
	filename := filepath.Base(pathFile)

	if err := os.MkdirAll(path, 0700); err != nil {
		return err
	}

	file := filepath.Join(path, filename)
	return os.WriteFile(file, buf, 0600)
}

func MarkdownToPDF(markdownFilePath string) error {

	// confirm that pandoc exist on the device
	if system.CommandExists("pandoc") == false {
		return fmt.Errorf("pandoc is not on the system. Try to install it with\n  - brew install pandoc")
	}
	if system.CommandExists("pdftex") == false {
		return fmt.Errorf("pdftex is not on the system. Try to install it with\n  - brew install mactex")
	}

	filename := filepath.Base(markdownFilePath)
	directory := filepath.Dir(markdownFilePath)
	filenameWithoutExt := strings.TrimSuffix(filename, filepath.Ext(filename))

	outputPdfFilename := filenameWithoutExt + ".pdf"

	cmd := exec.Command("pandoc", markdownFilePath, "-o", filepath.Join(directory, outputPdfFilename), "-f", "markdown-implicit_figures", "--resource-path", directory)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
