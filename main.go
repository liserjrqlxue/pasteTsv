package main

import (
	"bufio"
	"flag"
	"fmt"
	"strings"

	"github.com/liserjrqlxue/goUtil/osUtil"
	"github.com/liserjrqlxue/goUtil/simpleUtil"
)

var (
	sep = flag.String(
		"sep",
		"\t",
		"sep to paste",
	)
	omit = flag.Int(
		"omit",
		0,
		"omit columns with index < -omit ( index start from 0 )",
	)
	omitSep = flag.String(
		"omitSep",
		"\t",
		"sep to split columns",
	)
	chanLen = flag.Int(
		"size",
		1000,
		"chan buffer size",
	)
)

type bChan struct {
	ch   chan string
	flag chan int
}

func main() {
	flag.Parse()
	if *omit <= 0 {
		*omit = 0
	}
	var chanList []bChan
	for i, file := range flag.Args() {
		var ch = bChan{
			ch:   make(chan string, *chanLen),
			flag: make(chan int, *chanLen),
		}
		chanList = append(chanList, ch)
		var f = osUtil.Open(file)
		go func(flag bool) {
			defer simpleUtil.DeferClose(f)
			var s = bufio.NewScanner(f)
			var text = ""
			for s.Scan() {
				text = s.Text()
				if *omit > 0 && flag {
					text = strings.Join(strings.Split(s.Text(), *omitSep)[*omit:], *omitSep)
				}
				ch.ch <- text
				ch.flag <- 1
			}
			simpleUtil.CheckErr(s.Err())
			for {
				ch.ch <- ""
				ch.flag <- 0
			}
		}(i > 0)
	}

	for {
		var lines []string
		var n = 0
		for _, ch := range chanList {
			lines = append(lines, <-ch.ch)
			n += <-ch.flag
		}
		if n == 0 {
			break
		}
		fmt.Println(strings.Join(lines, *sep))
	}
}
