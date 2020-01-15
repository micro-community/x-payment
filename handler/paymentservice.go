package handler

import (
	"context"
	pb "github.com/micro-community/x-payment/pb"
)

type PaymentService struct{}

func (*PaymentService) Charge(context.Context, *pb.ChargeRequest, *pb.ChargeResponse) error {
	return nil
}
