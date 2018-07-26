package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

// But html.Parse parses pretty much anything, so this test is useless.
func TestPrettyOutputCanBeParsed(t *testing.T) {
	input := `
<html>
<body>
	<p class="something" id="short"><span class="special">hi</span></p><br/>
</body>
</html>
`
	want := `<html>
  <head/>
  </head>
  <body>
    <p class=something id=short>
      <span class=special>
        hi
      </span>
    </p>
    <br/>
    </br>
  </body>
</html>
`

	out = bytes.Buffer{}
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Error(err)
	}
	forEachNode(doc, start, end)
	result := out.String()
	if !reflect.DeepEqual(result, want) {
		t.Errorf("got=%#v, want=%#v\n", result, want)
	}
}
