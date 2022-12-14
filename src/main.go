package main

import (
	"embed"
)

type Link struct {
	Id   uint16
	Name string
	Link string
	Date string
	Tag  string
	Tags []string
}

type Conf struct {
	DbPath  string
	GuiIP   string
	GuiPort string
	Theme   string
}

var AllLinks []Link
var AppConfig Conf

//go:embed templates/*
var TemplHTML embed.FS

func main() {
	AppConfig = get_config()

	db_create()            // Create DB if not existed
	AllLinks = db_select() // Select all links from DB

	webgui() // web-index.go
}
