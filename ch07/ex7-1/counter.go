package counter

import (
	"bufio"
	"bytes"
	"fmt"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	r := bytes.NewBuffer(p)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		*w++
	}
	return int(*w), nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	r := bytes.NewBuffer(p)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		*l++
	}
	return int(*l), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	var w WordCounter
	w.Write([]byte("one two three four five"))
	w.Write([]byte("six seven eight nine ten"))
	fmt.Println(w)

	var l LineCounter
	l.Write([]byte("one\ntwo\nthree\nfour\nfive\n"))
	fmt.Println(l)
}
