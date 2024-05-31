package repository

import (
	"cmp"

	"slices"

	"github.com/chazool/go_medium_samples/pubsub_cloud_func/pubsubcloudfunc/app/dto"
)

var items = make([]dto.Item, 0, 100)

type ItemRepo interface {
	New(i dto.Item)
	Get(id uint) *dto.Item
}

type ItemRepoImpl struct {
}

func NewItemRepo() ItemRepo {
	return ItemRepoImpl{}
}

func (repo ItemRepoImpl) Get(id uint) *dto.Item {
	i, _ := slices.BinarySearchFunc(items, dto.Item{ID: id}, func(e, t dto.Item) int {
		return cmp.Compare(e.ID, t.ID)
	})
	return &items[i]
}

func (repo ItemRepoImpl) New(i dto.Item) {
	items = append(items, i)
}
