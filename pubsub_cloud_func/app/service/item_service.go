package service

import (
	"errors"

	"github.com/chazool/go_medium_samples/pubsub_cloud_func/pubsubcloudfunc/app/dto"
	"github.com/chazool/go_medium_samples/pubsub_cloud_func/pubsubcloudfunc/app/repository"
)

type ItemService interface {
	New(item dto.Item) error
}

type ItemServiceImpl struct {
	ItemRepo repository.ItemRepo
}

func NewItemService() ItemService {
	return ItemServiceImpl{
		ItemRepo: repository.NewItemRepo(),
	}
}

func (srv ItemServiceImpl) New(i dto.Item) error {
	item := srv.ItemRepo.Get(i.ID)
	if item != nil {
		return errors.New("item already exists")
	}

	srv.ItemRepo.New(i)
	return nil
}
