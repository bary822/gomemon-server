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

func (repo *MemoMySqlRepository) GetByID(id string) (*entity.Memo, error) {
	// TODO: implement
	return &entity.Memo{}, nil
}

func (repo *MemoMySqlRepository) GetAll() ([]*entity.Memo, error) {
	// TODO: implement
	return []*entity.Memo{}, nil
}

func (repo *MemoMySqlRepository) Delete(id string) error {
	// TODO: implement
	return nil
}

func (repo *MemoMySqlRepository) Edit(id string, content string) (*entity.Memo, error) {
	// TODO: implement
	return &entity.Memo{}, nil
}
