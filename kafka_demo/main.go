package main

import (
	"my_test/cache"
	"my_test/db"
	"my_test/kafka"
	"my_test/server"
)

func main() {
	db.InitConfig()
	db.InitLog()
	cache.InitGCache()
	kafka.InitKafkaCfg()
	go server.StartConsumerServer()
	server.StartHttpServer()
}
