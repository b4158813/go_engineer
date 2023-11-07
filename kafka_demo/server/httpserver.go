package server

import (
	"fmt"
	"my_test/db"
	"my_test/kafka"
	"my_test/utils"

	"github.com/kataras/iris/v12"
)

func StartHttpServer() {
	app := iris.New()

	app.PartyFunc("/users", func(users iris.Party) {
		// /users/id
		users.Get("/{id:int}", func(ctx iris.Context) {
			id := ctx.Params().Get("id")
			err := kafka.ProduceMessages(ctx, db.C.Kafka.Topic1, id)
			if err != nil {
				db.L.Errorf("[%s] kafka.ProduceMessages error: %s\n", utils.GetCurrentFunctionName(), err)
				ctx.StatusCode(iris.StatusBadGateway)
				return
			}
			ctx.WriteString(fmt.Sprintf("received from id = %s, written to kafka producer done\n", id))
			ctx.StatusCode(iris.StatusOK)
		})
	})

	fmt.Println("starting http server done!")
	app.Run(iris.Addr(":8080"))
}
