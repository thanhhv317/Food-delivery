package memcache

import (
	"context"
	"fmt"
	usermodel "golang/module/user/model"
	"sync"
)

type RealStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type userCaching struct {
	store     Caching
	realStore RealStore
	once      *sync.Once
}

func NewUserCaching(store Caching, realStore RealStore) *userCaching {
	return &userCaching{
		store:     store,
		realStore: realStore,
		once:      new(sync.Once),
	}
}

func (uc *userCaching) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	userId := conditions["id"].(int)
	key := fmt.Sprintf("user-%d", userId)

	userInCache := uc.store.Read(key)

	if userInCache != nil {
		return userInCache.(*usermodel.User), nil
	}

	uc.once.Do(func() {
		user, err := uc.realStore.FindUser(ctx, conditions, moreInfo...)

		if err != nil {
			panic(err)
		}

		// Update cache
		uc.store.Write(key, user)
	})

	return uc.store.Read(key).(*usermodel.User), nil
}
