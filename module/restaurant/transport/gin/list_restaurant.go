package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"golang/common"
	"golang/component/appctx"
	bizrestaurant "golang/module/restaurant/business"
	restaurantmodel "golang/module/restaurant/model"
	reporestaurant "golang/module/restaurant/repository"
	restaurantstorage "golang/module/restaurant/storage"
	grpcclientrestaurant "golang/module/restaurant/storage/grpcclient"
	"golang/proto"
	"google.golang.org/grpc"
	"net/http"
)

type fakeListStore struct {
}

func (fakeListStore) ListDataWithCondition(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	return []restaurantmodel.Restaurant{}, nil
}

func ListRestaurants(appContext appctx.AppContext, grpcClientConn *grpc.ClientConn) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var result []restaurantmodel.Restaurant

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Process()

		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appContext.GetMaiDBConnection())
		likedStore := grpcclientrestaurant.NewGRPCClient(proto.NewRestaurantLikeServiceClient(grpcClientConn))
		//likedStore := restaurantlikestorage.NewSQLStore(appContext.GetMaiDBConnection())
		repo := reporestaurant.NewListRestaurantRepo(store, likedStore)
		biz := bizrestaurant.NewListRestaurantBiz(repo)
		result, err := biz.ListDataWithCondition(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
