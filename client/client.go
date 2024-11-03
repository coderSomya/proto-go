package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/coderSomya/protogo/coffeeshop_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main(){
	conn, err:= grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Fatalf("Failed to create grpc client %s", err);
	}

	defer conn.Close()

	c := pb.NewCoffeeShopClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	menuStream, err := c.GetMenu(ctx, &pb.MenuRequest{})

	if err!=nil {
		log.Fatalf("error getting menu %s", err);
	}

	done := make(chan bool)

	var items []*pb.Item

	go func(){
		for {
			resp, err := menuStream.Recv()

			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("failed while receiving menu stream %s", err)
			}

			items = resp.Items
			log.Printf("resp received: %v", resp.Items)
		}
	}()

	<-done

	receipt, err := c.PlaceOrder(ctx, &pb.Order{Items: items})

	log.Printf("%v",receipt)

	status, err := c.GetOrderStatus(ctx, receipt)
	log.Printf("%v", status)
}