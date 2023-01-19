package main

import (
	"context"
	"github.com/reation/micro_service_payment_service/protoc"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	PaymentAddress = "192.168.1.104:8030"
)

func main() {
	PaymentOrder()
	getPaymentInfo()

}

func PaymentOrder() {
	conn, err := grpc.Dial(PaymentAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewPaymentClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.PaymentOrder(ctx, &protoc.PaymentOrderRequest{
		OrderNum:      "2023011623552601234567",
		OrderUserId:   1,
		PaymentUserId: 1,
		Prices:        3792.44,
		PaymentType:   5,
	})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Printf("last_id: %d", r.GetLastId())
	log.Printf("order_num: %s", r.GetOrderNum())

}

func getPaymentInfo() {
	conn, err := grpc.Dial(PaymentAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := protoc.NewPaymentClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetPaymentInfoByOrderNum(ctx, &protoc.GetPaymentInfoByOrderNumRequest{
		OrderNum: "2023011623552601234567",
	})
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	log.Printf("states: %d", r.GetStates())
	log.Println(r.PaymentInfo)
}
