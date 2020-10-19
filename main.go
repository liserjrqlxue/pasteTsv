package main

import (
	"bufio"
	"flag"
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
	for _, file := range os.Args {
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
	for {
		var lines []string
		for i, s := range scannerList {
			var n = 0
			var text = ""
			if s.Scan() {
				var line = strings.Split(s.Text(), *omitSep)
				if i > 0 || !*header {
					text = strings.Join(line[*omit:], *omitSep)
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
		if done {
			break
		}
	}
}
