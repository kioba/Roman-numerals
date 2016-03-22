package roman

import (
	"flag"
	"os"
	"testing"
)

// testing
func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestConvertFromArabToRoman(t *testing.T) {
	cases := []struct {
		from     int
		expected string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{20, "XX"},
		{30, "XXX"},
		{55, "LV"},
		{99, "XCIX"},
		{148, "CXLVIII"},
		{2016, "MMXVI"},
	}

	for _, test := range cases {
		number := Number{test.from}
		got := number.ToRoman()
		if got != test.expected {
			t.Errorf("[ToRoman(%d) == %q] Expected: %q", test.from, got, test.expected)
		}
	}
}
