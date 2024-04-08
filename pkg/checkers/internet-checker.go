package checkers

import (
	"net"
	"time"
)

func CheckInternet() {
	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", "8.8.8.8:53", timeout)

	if err != nil {
		panic(err)
	}

	conn.Close()
}
