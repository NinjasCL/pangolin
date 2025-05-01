package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const build_dir string = "_build"

func make_build_dir() {
	os.RemoveAll(build_dir)
	err := os.Mkdir(build_dir, 0755)
	check(err)
}

func copy_static_files() {
	err := copy_dir("public", build_dir)
	check(err)
}

func print_status_message(message string) {
	fmt.Printf("\n%s 🛠️", message)
}

func print_welcome_message() {
	fmt.Println("Pangolin Generator v1.0 🦔")
	fmt.Println("by ninjas.cl 🥷")
}

func print_done_message() {
	fmt.Printf("\n\nDone✨\n")
}

func compile_page(name string, content string, outpath string) string {
	dir := filepath.Join(outpath, fmt.Sprintf("%s.html", name))
	write_file(dir, content)
	return content
}

func compile_section(name string) string {
	print_status_message(name)
	content := parse_template(name, "")
	return compile_page(name, content, build_dir)
}

func read_post(path string) Post {
	json := read_json(filepath.Join(path, "page.json"))

	title := fmt.Sprintf("%s", json["title"])
	date := fmt.Sprintf("%s", json["date"])
	category := fmt.Sprintf("%s", json["category"])

	content := safe_html(read_file(filepath.Join(path, "page.html")))
	uri := filepath.Base(path)

	return Post{
		Title:    title,
		Content:  content,
		Date:     date,
		Category: category,
		Link:     uri,
	}
}

func compile_posts() {

	_, files := walk("data", []string{})

	for _, file := range files {

		if strings.HasSuffix(file, "page.json") {
			post := read_post(strings.TrimSuffix(file, "page.json"))
			page := parse_template("post", post)
			compile_page(post.Link, page, build_dir)
			print_status_message(post.Link)
		}
	}
}

func compile_categories() {
	path := filepath.Join("data", "categories.json")
	json := read_json(path)

	var items []any = json["items"].([]any)

	for _, item_json := range items {

		var item map[string]any = item_json.(map[string]any)

		title := fmt.Sprintf("%s", item["title"])
		subtitle := fmt.Sprintf("%s", item["subtitle"])
		category_name := fmt.Sprintf("%s", item["category"])
		category_file := fmt.Sprintf("%s", item["file"])

		var category_items []CategoryItem

		// Check for posts with the category
		_, files := walk("data", []string{})

		for _, file := range files {
			if strings.HasSuffix(file, "page.json") {
				post := read_post(strings.TrimSuffix(file, "page.json"))
				if post.Category == category_name {

					category_item := CategoryItem{
						Title: post.Title,
						Link:  fmt.Sprintf("%s.html", post.Link),
					}

					category_items = append(category_items, category_item)
				}
			}
		}

		category := Category{
			Title:    title,
			Subtitle: subtitle,
			Name:     category_name,
			File:     category_file,
			Items:    category_items,
		}

		page := parse_template("category", category)
		compile_page(category.File, page, build_dir)
		print_status_message(category.File)
	}
}

func compile_sections() {
	json := read_json(filepath.Join("data", "pages.json"))
	var items []any = json["items"].([]any)

	for _, page := range items {
		compile_section(fmt.Sprintf("%s", page))
	}
}

func main() {
	print_welcome_message()

	// Make build directory
	make_build_dir()

	// Copy static files
	copy_static_files()

	compile_posts()

	compile_categories()

	compile_sections()

	print_done_message()
}
