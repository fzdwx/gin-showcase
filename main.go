package main

import (
	"fmt"
	"gin-showcase/config"
	"gin-showcase/svc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	c := config.Config{
		MysqlUrl: "mysql://root:root@localhost:3306/test",
		RedisUrl: "redis://localhost:6379",
	}
	e := gin.Default()

	context := svc.NewServiceContext(c)
	defer context.Mysql.Close()
	defer context.Redis.Close()

	e.GET(Ping(context))
}

func Ping(svContext *svc.ServiceContext) (string, gin.HandlerFunc) {
	return "ping", func(c *gin.Context) {
		c.String(http.StatusOK, svContext.Redis.Get("ping"))
	}
}
