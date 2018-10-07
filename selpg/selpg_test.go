package selpg

import (
	"testing"
	"bytes"
)

func TestSelpg_GetPages(t *testing.T) {
	cases := []struct{
		filename string
		startPage, endPage, lines int
		formFeed bool
		want []byte
		wantErr bool
	} {
		//test Line mode
		{"../testl.txt", 1, 2, 6, false, []byte("test1\ntest1\ntest1\ntest1\ntest1\ntest1\ntest2\ntest2\ntest2\ntest2\ntest2\ntest2\n"), false},
		//test Line mode with higher end page
		{"../testl.txt", 2, 5, 6, false, []byte("test2\ntest2\ntest2\ntest2\ntest2\ntest2\ntest3\ntest3\ntest3\ntest3\ntest3\ntest3"), true},
		//test Line Mode with higher start page
		{"../testl.txt", 5, 5, 6, false, []byte(""), true},
		//test \f mode
		{"../testf.txt", 1, 1, 72, true, []byte("test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1test1\f"), false},
	}

	for index, cas := range cases {
		sp := NewSelpg(cas.startPage, cas.endPage, cas.lines, "", cas.filename, cas.formFeed)
		writer := new(bytes.Buffer)
		err := sp.GetPages(writer)
		res := make([]byte, writer.Len())
		writer.Read(res)
		if cas.wantErr != (err != nil) {
			t.Errorf("Test %v Failed with unexpected returned error", index)
		}
		flag := len(cas.want) == len(res)
		if flag {
			for ind := range cas.want {
				if res[ind] != cas.want[ind] {
					flag = false
				}
			}
		}
		if !flag {
			t.Errorf("Test %v Failed with Incorrect returned value", index)
		}
	}
}