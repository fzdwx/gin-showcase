package redis

import "fmt"

type (
	Conn struct {
		Url string
	}
)

func (c Conn) Open() Conn {
	fmt.Println("open redis connection")
	return c
}

func (c Conn) Close() {
	fmt.Println("close redis connection")
}

func (c Conn) Get(key string) string {
	return "ping"
}