package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// checks for error. Panics if an error is found.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// copies a whole directory recursively
func copy_dir(src string, dst string) error {
	return CopyDir(dst, src)
}

// reads file content at path
func read_file(path string) string {
	data, err := os.ReadFile(path)
	check(err)
	return string(data)
}

// reads a json an return a map[string]any
func read_json(path string) map[string]any {
	content := read_file(path)
	result := make(map[string]any)
	json.Unmarshal([]byte(content), &result)
	return result
}

// writes file contents
func write_file(path string, content string) error {
	err := os.WriteFile(path, []byte(content), 0644)
	check(err)
	return err
}

// parse a template file
func parse_template(name string, params any) string {
	filename := fmt.Sprintf("%s.html", name)

	tpl, err := template.ParseGlob("templates/*")
	check(err)

	buffer := new(strings.Builder)
	err = tpl.ExecuteTemplate(buffer, filename, params)

	return buffer.String()
}

// A safe html string
func safe_html(content string) template.HTML {
	return template.HTML(content)
}

// Scans recursively a directory
func walk(dir_path string, ignore []string) ([]string, []string) {

	folders := []string{}
	files := []string{}

	// Scan
	filepath.Walk(dir_path, func(path string, f os.FileInfo, err error) error {

		_continue := false

		// Loop : Ignore Files & Folders
		for _, i := range ignore {

			// If ignored path
			if strings.Index(path, i) != -1 {

				// Continue
				_continue = true
			}
		}

		if _continue == false {

			f, err = os.Stat(path)

			// If no error
			check(err)

			// File & Folder Mode
			f_mode := f.Mode()

			// Is folder
			if f_mode.IsDir() {

				// Append to Folders Array
				folders = append(folders, path)

				// Is file
			} else if f_mode.IsRegular() {

				// Append to Files Array
				files = append(files, path)
			}
		}

		return nil
	})

	return folders, files
}
