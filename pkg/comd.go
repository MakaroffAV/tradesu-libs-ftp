// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package ftp

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"io"
	"time"

	"github.com/jlaffaye/ftp"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

type DirEntry struct {
	Path string
	Name string
	Type DirEntryType
	Size uint64
	Modf time.Time
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

type DirEntryType int

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func NewDirEntryType(t ftp.EntryType) DirEntryType {
	return DirEntryType(t)
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func (fc FtpConn) List(p string) ([]*DirEntry, error) {

	e := []*DirEntry{}

	l, lErr := fc.c.List(p)
	if lErr != nil {
		return nil, lErr
	}

	for _, i := range l {
		e = append(
			e,
			&DirEntry{
				Path: p,
				Name: i.Name,
				Size: i.Size,
				Modf: i.Time,
				Type: NewDirEntryType(i.Type),
			},
		)
	}

	return e, nil

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func (fc FtpConn) Retr(p string) ([]byte, error) {

	r, rErr := fc.c.Retr(p)
	if rErr != nil {
		return nil, rErr
	}
	defer r.Close()

	b, bErr := io.ReadAll(r)
	if bErr != nil {
		return nil, bErr
	}

	return b, nil

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
