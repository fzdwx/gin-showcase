package svc

import (
	"gin-showcase/config"
	"gin-showcase/mysql"
	"gin-showcase/redis"
)

type (
	ServiceContext struct {
		C     config.Config
		Mysql mysql.Conn
		Redis redis.Conn
	}
)

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		C: c,
		Mysql: mysql.Conn{
			Url: c.MysqlUrl,
		}.Open(),
		Redis: redis.Conn{
			Url: c.RedisUrl,
		}.Open(),
	}
}