package parse

import "testing"

func TestRawTextTrim(t *testing.T) {
	type test struct{ input, output string }
	var tests = []test{
		{"", ""},
		{" ", ""},
		{"  ", ""},
		{"\n", ""},
		{"\n\n  \r\n\t ", ""},
		{"a", "a"},
		{"a ", "a"},
		{" a", "a"},
		{"a\n", "a"},
		{"\na", "a"},
		{"a \n ", "a"},
		{" \n a", "a"},
		{"a\nb", "a b"},
		{" <a> b ", "<a> b"},
	}
	for _, test := range tests {
		var actual = string(rawtext(test.input, true, true))
		if test.output != actual {
			t.Errorf("input: %q, expected %q, got %q", test.input, test.output, actual)
		}
	}
}

func TestRawTextNoTrim(t *testing.T) {
	type test struct{ input, output string }
	var tests = []test{
		{"", ""},
		{" ", " "},
		{"  ", "  "},
		{"\n", ""},
		{"\n\n  \r\n\t ", ""},
		{"a", "a"},
		{"a ", "a "},
		{" a", " a"},
		{"a\n", "a"},
		{"\na", "a"},
		{"a \n ", "a"},
		{" \n a", "a"},
		{"a\nb", "a b"},
		{"\n\t a \nb\n\n", "a b"},
		{"a / b", "a / b"},
		{"a \t /\nb", "a \t / b"},
		{"<br>  \n", "<br>"},
		{"<a>", "<a>"},
		{" <a>\n\t", " <a>"},
		{"<a> \n\t b \r\n\t <c>", "<a>b<c>"},
		{"a <b> \n\t<c> \n d\nd", "a <b><c>d d"},
		{"a <br>\n\t b \n\n\t \n\t c", "a <br>b c"},
		{"\u2222", "\u2222"},
		{" \u2222", " \u2222"},
		{"\u2222 ", "\u2222 "},
		{" \n \u2222", "\u2222"},
		{"\u2222 \n ", "\u2222"},
		{" \n\t\u2222 \n\t\r ", "\u2222"},
		{"\u2222 <\uEEEE> \n\t<\u9EC4> \n \u607A\n\u607A", "\u2222 <\uEEEE><\u9EC4>\u607A \u607A"},
	}

	for _, test := range tests {
		var actual = string(rawtext(test.input, false, false))
		if test.output != actual {
			t.Errorf("input: %q, expected %q, got %q", test.input, test.output, actual)
		}
	}
}
