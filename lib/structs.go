package main

import "html/template"

type Post = struct {
	Title    string
	Content  template.HTML
	Date     string
	Category string
	Link     string
}

type CategoryItem = struct {
	Link  string
	Title string
}

type Category = struct {
	Title    string
	Subtitle string
	File     string
	Name     string
	Items    []CategoryItem
}
