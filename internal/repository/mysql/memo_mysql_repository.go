package repository

import "github.com/bary822/gomemon-server/internal/entity"

type MemoMySqlRepository struct{}

func NewMemoMySqlRepository() *MemoMySqlRepository {
	return &MemoMySqlRepository{}
}

func (repo *MemoMySqlRepository) Save(memo entity.Memo) (*entity.Memo, error) {
	// TODO: implement
	return &entity.Memo{}, nil
}

func (repo *MemoMySqlRepository) GetByID(content string) (*entity.Memo, error) {
	// TODO: implement
	return &entity.Memo{}, nil
}

func (repo *MemoMySqlRepository) GetAll() ([]*entity.Memo, error) {
	// TODO: implement
	return []*entity.Memo{}, nil
}

func (repo *MemoMySqlRepository) Delete() error {
	// TODO: implement
	return nil
}
