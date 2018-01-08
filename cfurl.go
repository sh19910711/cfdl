package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"strings"
)

func ExtractSourceCode(res *http.Response) string {
	doc, _ := goquery.NewDocumentFromResponse(res)
	return strings.Trim(doc.Find(".program-source").Text(), "\n")
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		println("ERR: no arguments")
		os.Exit(-1)
	}
	u := flag.Args()[0]
	res, _ := http.Get(u)
	if strings.Contains(u, "submission") {
		fmt.Print(ExtractSourceCode(res))
	} else {
		println("nothing to do :-|")
	}
}
