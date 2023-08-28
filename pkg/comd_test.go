package ftp

import (
	"fmt"
	"testing"

	"github.com/jlaffaye/ftp"
)

func TestList(t *testing.T) {

	testCases := []struct {
		m1 string
		m2 string
		m3 string

		a1 string

		w1 []*DirEntry
		w2 error
	}{
		{
			m1: "ftp.zakupki.gov.ru:21",
			m2: "free",
			m3: "free",

			a1: "/",

			w1: nil,
			w2: nil,
		},
	}

	for _, c := range testCases {

		f, _ := NewFtpConn(c.m1, c.m2, c.m3)
		defer f.DelFtpConn()

		l, lErr := f.List(c.a1)
		fmt.Println(l, lErr)

	}

}

func TestNewDirEntryType(t *testing.T) {

	testCases := []struct {
		a ftp.EntryType
		w DirEntryType
	}{
		{
			a: ftp.EntryTypeFile,
			w: DirEntryType(0),
		},
	}

	for _, c := range testCases {
		if NewDirEntryType(c.a) != c.w {
			t.Errorf(
				`
					Test failed: 	
				`,
			)
		}
	}

}
