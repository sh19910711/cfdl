package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func ExtractSourceCode(res *http.Response) string {
  doc, _ := goquery.NewDocumentFromResponse(res)
  return strings.Trim(doc.Find(".program-source").Text(), "\n")
}

func main() {
	flag.Parse()
	res, _ := http.Get(flag.Args()[0])
	fmt.Print(ExtractSourceCode(res))
}
