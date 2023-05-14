package generator

import (
	"testing"
)

func TestReplaceUnescapeChar(t *testing.T) {
	test := `#$%&\_{}^~`
	actual := replaceUnescapedChar(test)
	expected := `{\#}{\$}{\%}{\&}\textbackslash{\_}{\{}{\}}\textasciicircum\textasciitilde`
	if actual != expected {
		t.Errorf("failed when replacing escape character, got: %s, expected: %s", actual, expected)
	}
}
