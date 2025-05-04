package menu

import (
	"context"

	"Menu_Service/pb"
)

type Server struct {
	pb.UnimplementedMenuServiceServer
}

func (s *Server) CalculateDiscount(ctx context.Context, req *pb.DiscountRequest) (*pb.DiscountResponse, error) {
	price := req.GetPrice()
	discount := req.GetDiscount()
	discounted := price * (1 - discount/100)

	return &pb.DiscountResponse{
		DiscountedPrice: discounted,
	}, nil
}
