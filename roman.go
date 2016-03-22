package roman

import (
	"bytes"
	// "fmt"
	// "strconv"
)

var romanNumbers = []struct {
	name  string
	value int
}{
	{"I", 1},
	{"V", 5},
	{"X", 10},
	{"L", 50},
	{"C", 100},
	{"D", 500},
	{"M", 1000},
}

type Number struct {
	Num int
}

func getSubRoman(times int, indentation int) string {
	// fmt.Printf("number:" + strconv.Itoa(number) + "\n")
	if times <= 0 {
		return ""
	}

	number := times * indentation
	pos := 0
	for i := 0; i < len(romanNumbers); i++ {
		if indentation == romanNumbers[i].value {
			pos = i
			break
		}
	}

	var buffer bytes.Buffer
	if pos == len(romanNumbers)-1 {
		for i := 0; i < times; i++ {
			buffer.WriteString(romanNumbers[pos].name)
		}

		return buffer.String()
	}

	first := romanNumbers[pos]
	mid := romanNumbers[pos+1]
	last := romanNumbers[pos+2]

	if number == mid.value {
		buffer.WriteString(mid.name)
	} else if number < mid.value {
		if number == (mid.value - indentation) {
			buffer.WriteString(first.name)
			buffer.WriteString(mid.name)
		} else {
			for i := 0; i < times; i++ {
				buffer.WriteString(first.name)
			}
		}
	} else if number < last.value {
		if number == (last.value - indentation) {
			buffer.WriteString(first.name)
			buffer.WriteString(last.name)
		} else {
			buffer.WriteString(mid.name)
			for i := 0; i < times-mid.value; i++ {
				buffer.WriteString(first.name)
			}
		}
	}

	return buffer.String()
}

func NewNumber(num int) Number {
	return Number{num}
}

func (n *Number) ToArabic() int {
	return n.Num
}

func (n *Number) ToRoman() string {
	var buffer bytes.Buffer
	num := n.Num
	indenter := 1

	for (num / (indenter * 10)) != 0 {
		indenter *= 10
	}

	for num > 0 {
		times := num / indenter

		str := getSubRoman(times, indenter)
		buffer.WriteString(str)

		num = num % indenter
		indenter /= 10
	}

	return buffer.String()
}
