package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/srahkmli/grpcTest/controller"
	"github.com/srahkmli/grpcTest/pb"
	"google.golang.org/grpc"
	"net"
)

func ListenAndServe() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		logrus.Fatalf("failed to listen : %v", err)
	}
	grpcServer := grpc.NewServer()

	grpcTest.RegisterProductShoppingServer(grpcServer, &controller.ProductController{})
	logrus.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve : %v", err)
	}
}
