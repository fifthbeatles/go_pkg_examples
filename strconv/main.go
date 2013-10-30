package main

import (
	"fmt"
	sc "strconv"
)

func testXXXToString() {
	dst := make([]byte, 1024)
	dst = sc.AppendBool(dst, sc.CanBackquote("Hello\nWorld!"))
	dst = append(dst, ' ')
	dst = sc.AppendInt(dst, 12345, 36)
	dst = append(dst, ' ')
	dst = sc.AppendUint(dst, 54321, 36)
	dst = append(dst, ' ')
	dst = sc.AppendFloat(dst, 14.285714, 'b', 3, 64) // b for binary exponent
	dst = append(dst, ' ')
	dst = sc.AppendFloat(dst, 14.285714, 'e', 3, 64) // e and E for decimal exponent
	dst = append(dst, ' ')
	dst = sc.AppendFloat(dst, 14.285714, 'E', -1, 64)
	dst = append(dst, ' ')
	dst = sc.AppendFloat(dst, 14.285714, 'f', -1, 64) // f for no exponent
	dst = append(dst, ' ')
	dst = sc.AppendQuote(dst, "我")
	dst = append(dst, ' ')
	dst = sc.AppendQuoteRune(dst, '我')
	dst = append(dst, ' ')
	dst = sc.AppendQuoteToASCII(dst, "我")
	dst = append(dst, ' ')
	dst = sc.AppendQuoteRuneToASCII(dst, '我')
	fmt.Println(string(dst))
	fmt.Printf("%s ", sc.FormatBool(sc.IsPrint('\n')))
	fmt.Printf("%s ", sc.FormatFloat(12.345, 'f', 1, 32))
	fmt.Printf("%s ", sc.FormatInt(142857, 16))
	fmt.Printf("%s ", sc.FormatUint(714285, 16))
	fmt.Printf("%s ", sc.Itoa(1234567))
	fmt.Printf("%s ", sc.Quote("你"))
	fmt.Printf("%s ", sc.QuoteRune('你'))
	fmt.Printf("%s ", sc.QuoteToASCII("你"))
	fmt.Printf("%s ", sc.QuoteRuneToASCII('你'))
	fmt.Println()
}

func testStringToXXX() {
	if v, err := sc.Atoi("123456"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v(%T) ", v, v)
	}
	if v, err := sc.ParseBool("T"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v(%T) ", v, v)
	}
	if v, err := sc.ParseFloat("123.45", 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v(%T) ", v, v)
	}
	if v, err := sc.ParseInt("abcde", 36, 64); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v(%T) ", v, v)
	}
	fmt.Println()
	if v, err := sc.Unquote("\"Hello\""); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}

func main() {
	testXXXToString()
	testStringToXXX()
}
