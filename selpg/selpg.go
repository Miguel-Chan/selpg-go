package selpg

import(
	"io"
	"os"
	"fmt"
	"bufio"
	"log"
)

type selpg struct {
	startPage, endPage, pageLines, totalPages int
	destination, inputFile string
	fromFeed bool
}

func NewSelpg(start, end, lines int, dest, input string, useFormFeed bool) selpg {
	return selpg{
		start,
		end,
		lines,
		end - start + 1,
		dest,
		input,
		useFormFeed,
	}
}

func (sp selpg) Run() {
	var reader io.Reader
	var writer io.Writer = os.Stdout
	if sp.inputFile != "" {
		file, err := os.Open(sp.inputFile)
		if err == nil {
			fmt.Errorf("Could not open file: %v", sp.inputFile)
			os.Exit(1)
		}
		reader = file
	} else {
		reader = os.Stdin
	}
	bufReader := bufio.NewReader(reader)
	bufWriter := bufio.NewWriter(writer)

	if sp.fromFeed {
		//"f" mode
		currentPage := 1
		for {
			ch, err := bufReader.ReadByte()
			if err != nil {
				if err == io.EOF {
					break
				} else {
					log.Fatalf("Error: %v\n", err)
					os.Exit(1)
				}
			}
			if currentPage >= sp.startPage && currentPage <= sp.endPage {
				bufWriter.WriteByte(ch)
			}
			if ch == '\f' {
				currentPage += 1
			}
		}
	} else {
		//"l" mode
		currentLine := 1
		for {
			line, err := bufReader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					break
				} else {
					log.Fatalf("Error: %v\n", err)
					os.Exit(1)
				}
			}
			if currentLine >= sp.pageLines*(sp.startPage-1)+1 &&
				currentLine <= sp.totalPages * sp.pageLines {
					bufWriter.Write(line)
			}
			currentLine += 1
		}
		totalPages := currentLine / sp.pageLines + 1
		if totalPages < sp.startPage {
			fmt.Errorf("%v: start_page (%v) greater than total pages (%v), no output written", os.Args[0], sp.startPage, totalPages)
		} else if totalPages < sp.endPage {
			fmt.Errorf("%v: end_page (%v) greater than total pages (%v), less output than expected", os.Args[0], sp.endPage, totalPages)
		}
		fmt.Errorf("%v: Done", os.Args[0])
	}
}

