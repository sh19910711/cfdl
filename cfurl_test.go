package main_test

import (
	main "github.com/sh19910711/cfurl"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestExtractSourceCode(t *testing.T) {
	input := `
<pre class="program-source">
Hello World
</pre>
  `
	res := &http.Response{
		Request: &http.Request{},
		Body:    ioutil.NopCloser(strings.NewReader(input)),
	}

	sc := main.ExtractSourceCode(res)
	if sc != "Hello World" {
		t.Fail()
	}
}

func TestExtractProblemInput(t *testing.T) {
	input := `
<div class="sample-test"><div class="input">
  <pre>4<br />42<br /></pre>
</div></div>
  `
	res := &http.Response{
		Request: &http.Request{},
		Body:    ioutil.NopCloser(strings.NewReader(input)),
	}

	in := main.ExtractProblemInput(res)
	if in != "4\n42" {
		t.Fail()
	}
}
