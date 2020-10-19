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
	header = flag.Bool(
		"header",
		false,
		"keep omit columns from first file",
	)
)

func main() {
	flag.Parse()
	if *omit <= 0 {
		*omit = 0
		*header = false
	}
	var inputList []*os.File
	fmt.Printf("%+v\n", flag.Args())
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
				var line = strings.Split(s.Text(), *omitSep)
				if i > 0 || !*header {
					text = strings.Join(line[*omit:], *omitSep)
				} else {
					text = strings.Join(line, *omitSep)
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
		println(strings.Join(lines, *sep))
	}
}
