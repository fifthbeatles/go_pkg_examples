package main

import (
	"fmt"
	"os"
	ss "strings"
)

func testSearch() {
	s := "abcdefg一二三四五cdefghi四五六七八hijklmn"
	fmt.Println(ss.Contains(s, "g一"))
	fmt.Println(ss.ContainsAny(s, "12345"))
	fmt.Println(ss.ContainsRune(s, '五'))
	fmt.Println(ss.Count(s, "四五"))
	fmt.Println(ss.EqualFold("Hello", "hello"))
	fmt.Println(ss.HasPrefix(s, "abc"))
	fmt.Println(ss.HasSuffix(s, "lmn"))
	fmt.Println(ss.Index(s, "hi"), ss.Index(s, "12345"), ss.LastIndex(s, "hi"), ss.IndexRune(s, '五'))
	fmt.Println(ss.IndexAny(s, "note"), ss.LastIndexAny(s, "note"))
	fmt.Println(ss.IndexFunc(s, func(r rune) bool { return r > 'g' && r < 'n' }))
	fmt.Println(ss.LastIndexFunc(s, func(r rune) bool { return r > 'd' && r < 'g' }))
}

func testSplit() {
	s := "  abc  def ab cd efg bcd ef "
	fmt.Println(ss.Join(ss.Fields(s), ","))
	fmt.Println(ss.FieldsFunc(s, func(r rune) bool { return r == 'c' }))
	fmt.Println(ss.Split(s, "ef"), ss.SplitN(s, "ef", 2), ss.SplitN(s, "ef", -2), ss.SplitN(s, "ef", 0))
	fmt.Println(ss.SplitAfter(s, "ef"), ss.SplitAfterN(s, "ef", 2), ss.SplitAfterN(s, "ef", -2), ss.SplitAfterN(s, "ef", 0))
	fmt.Println(ss.Repeat("hello ", 3))
}

func testConvert() {
	s := "abcdefgbcdfghefbc"
	fmt.Println(ss.Map(func(r rune) rune { return r + 2 }, s))
	fmt.Println(ss.Replace(s, "bc", "23", -1), ss.Replace(s, "bc", "23", 2))
	fmt.Println(ss.Title("my dream"))
	fmt.Println(ss.ToLower("WWW"), ss.ToUpper("www"))
}

func testTrim() {
	s := "abbaabcccdefefabaabbbdefaab"
	fmt.Println(ss.Trim(s, "ab"), ss.TrimLeft(s, "ab"), ss.TrimRight(s, "ab"))
	fmt.Println(ss.TrimSpace(" \t\n Hello World \n\t\r\n"))
	fmt.Println(ss.TrimPrefix(s, "ab"), ss.TrimSuffix(s, "ab"))
	fmt.Println(ss.TrimFunc(s, func(r rune) bool { return r < 'e' }))
}

func testReplacer() {
	r := ss.NewReplacer("a", "A", "c", "C")
	s := "ababcbcdabcdbcdcda"
	fmt.Println(r.Replace(s))
	r.WriteString(os.Stdout, s)
}

func main() {
	//testSearch()
	//testSplit()
	//testConvert()
	//testTrim()
	testReplacer()
}
