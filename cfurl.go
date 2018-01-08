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

func ExtractProblemStatement(res *http.Response) string {
	doc, _ := goquery.NewDocumentFromResponse(res)
	return strings.Trim(doc.Find(".problem-statement").Text(), "\n")
}

func ExtractProblemInput(res *http.Response) string {
	doc, _ := goquery.NewDocumentFromResponse(res)
	return strings.Join(doc.Find(".sample-test .input pre").Map(
		func(_ int, s *goquery.Selection) string {
			h, _ := s.Html()
			return strings.Trim(strings.Replace(h, "<br/>", "\n", -1), "\n")
		}), "\n")
}

func ExtractProblemOutput(res *http.Response) string {
	doc, _ := goquery.NewDocumentFromResponse(res)
	return strings.Join(doc.Find(".sample-test .output pre").Map(
		func(_ int, s *goquery.Selection) string {
			h, _ := s.Html()
			return strings.Trim(strings.Replace(h, "<br/>", "\n", -1), "\n")
		}), "\n")
}

func main() {
	input := flag.Bool("input", false, "")
	output := flag.Bool("output", false, "")
	flag.Parse()
	if len(flag.Args()) == 0 {
		println("ERR: no arguments")
		os.Exit(-1)
	}

	u := flag.Args()[0]
	res, _ := http.Get(u)

	if strings.Contains(u, "submission") {
		fmt.Print(ExtractSourceCode(res))
	} else if strings.Contains(u, "/problem/") {
		if *input {
			fmt.Print(ExtractProblemInput(res))
		} else if *output {
			fmt.Print(ExtractProblemOutput(res))
		} else {
			fmt.Print(ExtractProblemStatement(res))
		}
	} else {
		println("nothing to do :-|")
	}
}
