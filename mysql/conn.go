package mysql

import "fmt"

type (
	Conn struct {
		Url string
	}
)

func (c Conn) Open() Conn {
	fmt.Println("open mysql connection")
	return c
}

func (c Conn) Close() {
	fmt.Println("close mysql connection")
}