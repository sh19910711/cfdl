package cfdl_test

import (
  "testing"
  "net/http"
  "io/ioutil"
  "strings"
  "github.com/sh19910711/cfdl"
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

  sc := cfdl.ExtractSourceCode(res)
  if sc != "Hello World" {
    t.Fail()
  }
}
