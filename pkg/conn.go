// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

package ftp

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

import (
	"errors"
	"net"
	"os"

	"github.com/MakaroffAV/tradesu-libs-cls/pkg/cls"
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

var (
	errNoConnParam = errors.New("error, conn param does not exist in env storage")
)

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //

func createConn() {

}

// ------------------------------------------------------------------------ //

func deleteConn(conn net.Conn) {

}

// ------------------------------------------------------------------------ //

// getConnParam is
// private function for
// fetching conn param for FTP server
func getConnParam(key string) (string, error) {

	if p := os.Getenv(key); p != "" {
		cls.Info(
			cls.Log{
				Msg: "fetching FTP conn param done",
				Add: map[string]string{"connParam": key},
			},
		)
		return p, nil
	} else {
		cls.Fail(
			cls.Log{
				Err: errNoConnParam,
				Msg: "fetching FTP conn param fail",
				Add: map[string]string{"connParam": key},
			},
		)
		return p, errNoConnParam
	}

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ //
