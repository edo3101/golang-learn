package main

import (
	"learn/config"
	"learn/controllers"
	"learn/routers"
	"learn/services"

	"github.com/gin-gonic/gin"
)

func main() {

	websocketService := services.NewWebsocketService()
	websocketController := controllers.NewWebSocketController(websocketService)
	// router := routers.WebSocketRouter(websocketController)
	//
	//go websocketService.HandleMessage()

	cronScheduler := config.InitCronJobs()
	defer cronScheduler.Stop()

	app := gin.Default()
	// Initialize Gin router
	r := routers.NewRouter(app, websocketController)

	r.App.Run(":8080")

	//go func() {
	//	if err := r.Run(":8080"); err != nil {
	//		panic(err)
	//	}
	//}()

	//list, err := net.Listen("tcp", ":3000")
	//if err != nil {
	//	panic(err)
	//}

	//grpcServer := grpc.NewServer()
	//proto.GrpcServerService(grpcServer, &controllers.GrpcServerController{})

	//fmt.Println("gRPC server started, Listening on port 3000")
	//if err := grpcServer.Serve(list); err != nil {
	//	panic(err)
	//}
}
