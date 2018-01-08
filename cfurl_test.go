package main_test

import (
  "testing"
  "net/http"
  "io/ioutil"
  "strings"
  main "github.com/sh19910711/cfurl"
)

func TestExtractSourceCode(t *testing.T) {
  input := `
<pre class="program-source">
Hello World
</pre>
  `
  res := &http.Response{
    Request: &http.Request{},
    Body: ioutil.NopCloser(strings.NewReader(input)),
  }

  sc := main.ExtractSourceCode(res)
  if sc != "Hello World" {
    t.Fail()
  }
}
