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
