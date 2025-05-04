package menu

import (
	"context"

	"github.com/CallenCaracy/ByteBites/services/Menu_Service/pb"
)

type MenuServiceServer struct {
	pb.UnimplementedMenuServiceServer
}

func (s *MenuServiceServer) CalculateDiscount(ctx context.Context, req *pb.DiscountRequest) (*pb.DiscountResponse, error) {
	price := req.GetPrice()
	discount := req.GetDiscount()
	discounted := price * (1 - discount/100)

	return &pb.DiscountResponse{
		DiscountedPrice: discounted,
	}, nil
}
