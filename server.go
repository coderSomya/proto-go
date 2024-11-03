package main

import (
	"context"
	"log"
	"net"

	pb "github.com/coderSomya/protogo/coffeeshop_proto"
	"google.golang.org/grpc"
)

type server struct{
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(menuRequest *pb.MenuRequest,srv grpc.ServerStreamingServer[pb.Menu]) error {
	items := []*pb.Item{
	&pb.Item{
		Id: "1",
		Name: "Black Coffee",
	},
	&pb.Item{
		Id: "2",
		Name: "Tea",
	},
   }

   for i, _ := range items{
	srv.Send(&pb.Menu{
		Items: items[0: i+1],
	})
   }

   return nil
}
func (s *server) PlaceOrder(ctx context.Context,order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "abc",
	}, nil
}
func (s *server) GetOrderStatus(ctx context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: receipt.Id,
		Status: "IN PROGRESS",
	}, nil
}

func main(){
	lis, err := net.Listen("tcp", ":9001")
	if err != nil{
		log.Fatalf("Failed to start server %v", err);
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCoffeeShopServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err!=nil{
		log.Fatalf("failed to serve :%s ", err)
	}
}