package grpcclientrestaurant

import (
	"context"
	"golang/common"
	"golang/proto"
)

type grpcClient struct {
	client proto.RestaurantLikeServiceClient
}

func NewGRPCClient(client proto.RestaurantLikeServiceClient) *grpcClient {
	return &grpcClient{client: client}
}

func (c *grpcClient) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	newIds := make([]int32, len(ids))

	for i := range newIds {
		newIds[i] = int32(ids[i])
	}

	res, err := c.client.GetRestaurantLikeStat(ctx, &proto.RestaurantLikeStatRequest{ResIds: newIds})

	if err != nil {
		return nil, common.ErrDB(err)
	}

	result := make(map[int]int)

	for k, v := range res.Result {
		result[int(k)] = int(v)
	}

	return result, nil
}
