package files

import (
	"archive/zip"
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
// Unzip unzip the src into dest
func Unzip(src, dest string) error {
	dest = filepath.Clean(dest) + string(os.PathSeparator)

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		path := filepath.Join(dest, f.Name)
		// Check for ZipSlip: https://snyk.io/research/zip-slip-vulnerability
		if !strings.HasPrefix(path, dest) {
			return fmt.Errorf("%s: illegal file path", path)
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
