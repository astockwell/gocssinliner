package cssInliner

import (
	"bytes"
	"strings"
	"testing"
)

func Test_Inliner_Inline(t *testing.T) {
	src := strings.NewReader(`<html>
  <head>
    <style>.blue { color: blue; }</style>
  </head>
  <body>
    <h1 class="blue" style="display: block;">Hello!</h1>
  </body>
</html>`)

	i := NewInliner()

	var result bytes.Buffer

	err := i.Inline(&result, src)
	if err != nil {
		t.Error(err)
	}

	// fmt.Println(result.String())

	want := `<html>
  <head>
    <style>.blue { color: blue; }</style>
  </head>
  <body>
    <h1 class="blue" style="display:block;color:blue;" >Hello!</h1>
  </body>
</html>`

	if result.String() != want {
		t.Errorf("Got %v, Want %v", result.String(), want)
	}
}
