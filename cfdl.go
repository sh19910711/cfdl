package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func main() {
	flag.Parse()
	res, _ := http.Get(flag.Args()[0])
	doc, _ := goquery.NewDocumentFromResponse(res)
	fmt.Print(strings.Trim(doc.Find(".program-source").Text(), "\n"))
}
