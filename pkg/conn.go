// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package ftp

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"github.com/jlaffaye/ftp"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

type FtpConn struct {
	c *ftp.ServerConn
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func (fc FtpConn) DelFtpConn() error {
	return fc.c.Quit()
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func NewFtpConn(a string, u string, p string) (*FtpConn, error) {

	c, cErr := ftp.Dial(a)
	if cErr != nil {
		return nil, cErr
	}

	lErr := c.Login(u, p)
	if lErr != nil {
		return nil, lErr
	}

	return &FtpConn{c: c}, nil

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
