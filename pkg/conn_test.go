// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package ftp

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"testing"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func TestDelFtpConn(t *testing.T) {

	testCases := []struct {
		a string
		u string
		p string

		n string
		r func(e error) bool
		w bool
	}{
		{
			a: "ftp.zakupki.gov.ru:21",
			u: "free",
			p: "free",

			n: `
			deleting instance of ftp 
			client with correct connection params`,
			r: func(e error) bool { return e == nil },
			w: true,
		},
	}

	for _, c := range testCases {
		f, _ := NewFtpConn(c.a, c.u, c.p)
		defer func() {
			if err := f.DelFtpConn(); c.w != c.r(err) {
				t.Errorf(
					`
						Test failed:	%s
										want %t, got %t
					`,
					c.n,
					c.w,
					c.r(err),
				)
			}
		}()
	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func TestNewFtpConn(t *testing.T) {

	testCases := []struct {
		a string
		u string
		p string

		n string
		r func(e error) bool
		w bool
	}{
		{
			a: "ftp.zakupki.gov.ru:21",
			u: "free",
			p: "free",

			n: `
			creating instance of ftp 
			client with correct connection params`,
			r: func(e error) bool { return e == nil },
			w: true,
		},
		{
			a: "ftp.zakupki.gov.ru",
			u: "free",
			p: "free",

			n: `
			creating instance of ftp
			client with incorrect connection param: address`,
			r: func(e error) bool { return e != nil },
			w: true,
		},
		{
			a: "ftp.zakupki.gov.ru:21",
			u: "free1",
			p: "free",

			n: `
			creating instance of ftp
			client with incorrect connection param: username`,
			r: func(e error) bool { return e != nil },
			w: true,
		},
	}

	for _, c := range testCases {
		_, fErr := NewFtpConn(c.a, c.u, c.p)
		if c.w != c.r(fErr) {
			t.Errorf(
				`
					Test failed:	%s
									want %t, got %t
				`,
				c.n,
				c.w,
				c.r(fErr),
			)
		}
	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
