package main

import (
	"context"
	"github.com/sirupsen/logrus"
	pb "github.com/srahkmli/grpcTest/pb"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:9000", grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewProductShoppingClient(conn)

	logrus.Println("request save has sent")
	product := pb.Product{Name: "Product no.1", Qty: 001}
	_, err = c.SaveProduct(context.Background(), &product)
	if err != nil {
		logrus.Fatalf("Error when calling SavingProduct: %s", err)
	}

	logrus.Println("request save has sent")
	point := pb.Point{Name: "Product no.1"}
	response, err := c.GetProduct(context.Background(), &point)
	if err != nil {
		logrus.Fatalf("Error when calling SavingProduct: %s", err)
	}
	logrus.Printf("message from server is : %v and %v", response.GetName(), response.GetQty())
}
