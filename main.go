package main

import (
	"github.com/micro-community/x-payment/handler"
	"github.com/micro-community/x-payment/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	paymentservice "github.com/micro-community/x-payment/pb"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("payment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	paymentservice.RegisterPaymentServiceHandler(service.Server(), new(handler.PaymentService))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.payment", service.Server(), new(subscriber.PaymentService))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.hipstershop.payment", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
