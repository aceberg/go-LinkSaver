package main

import (
	"fmt"
)

type Link struct {
	Id 	 uint16
	Name string
	Link string
	Date string
	Tag  string
}

type Conf struct {
    DbPath   string
    GuiIP    string
    GuiPort  string
    Theme    string
}

var AllLinks  []Link
var AppConfig Conf

func main() {
	AppConfig = get_config()
	fmt.Println("AppConfig: ", AppConfig)
	db_create()
	AllLinks = db_select()
	fmt.Println("AllLinks: ", AllLinks)

	webgui() // index.go
}