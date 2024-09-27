package repository

import (
	"errors"

	"github.com/bary822/gomemon-server/internal/entity"
)

type MemoInMemoryRepository struct {
	memos []*entity.Memo
}

func NewMemoInMemoryRepository() *MemoInMemoryRepository {
	return &MemoInMemoryRepository{
		memos: make([]*entity.Memo, 0),
	}
}

func (repo *MemoInMemoryRepository) Save(memo entity.Memo) (*entity.Memo, error) {
	repo.memos = append(repo.memos, &memo)

	return &memo, nil
}

func (repo *MemoInMemoryRepository) GetByID(id string) (*entity.Memo, error) {
	for _, memo := range repo.memos {
		if memo.ID == id {
			return memo, nil
		}
	}

	return nil, errors.New("Memo not found")
}

func (repo *MemoInMemoryRepository) GetAll() ([]*entity.Memo, error) {
	return repo.memos, nil
}

func (repo *MemoInMemoryRepository) Delete(id string) error {
	var del_index int

	for i, memo := range repo.memos {
		if memo.ID == id {
			del_index = i
		}
	}

	repo.memos = append(repo.memos[:del_index], repo.memos[del_index+1:]...)

	return nil
}
