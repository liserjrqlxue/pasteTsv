package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
)

func main() {
	flag.Parse()
	if *omit <= 0 {
		*omit = 0
	}
	var inputList []*os.File
	for _, file := range flag.Args() {
		var f = osUtil.Open(file)
		inputList = append(inputList, f)
	}
	defer func() {
		for _, f := range inputList {
			simpleUtil.CheckErr(f.Close())
		}
	}()
	var scannerList []*bufio.Scanner
	for _, f := range inputList {
		var s = bufio.NewScanner(f)
		scannerList = append(scannerList, s)
	}

	var done = false
	for !done {
		var lines []string
		var n = 0
		for i, s := range scannerList {
			var text = ""
			if s.Scan() {
				text = s.Text()
				if *omit > 0 && i > 0 {
					text = strings.Join(strings.Split(s.Text(), *omitSep)[*omit:], *omitSep)
				}
			} else {
				n++
			}
			if n == len(scannerList) {
				done = true
				break
			}
			lines = append(lines, text)
		}
		fmt.Println(strings.Join(lines, *sep))
	}
}
