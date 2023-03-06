package restaurantlikestorage

import (
	"golang.org/x/net/context"
	"golang/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type gRPCServer struct {
	db *gorm.DB
	proto.UnimplementedRestaurantLikeServiceServer
}

func NewGRPCServer(db *gorm.DB) *gRPCServer {
	return &gRPCServer{db: db}
}

func (s *gRPCServer) GetRestaurantLikeStat(ctx context.Context, request *proto.RestaurantLikeStatRequest) (*proto.RestaurantLikeStatResponse, error) {
	storage := NewSQLStore(s.db)

	ids := make([]int, len(request.ResIds))

	for i := range ids {
		ids[i] = int(request.ResIds[i])
	}

	result, err := storage.GetRestaurantLikes(ctx, ids)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "method GetRestaurantLikeStat has something error %s", err.Error())
	}

	mapRs := make(map[int32]int32)

	for k, v := range result {
		mapRs[int32(k)] = int32(v)
	}

	return &proto.RestaurantLikeStatResponse{Result: mapRs}, nil
}
