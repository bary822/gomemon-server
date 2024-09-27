package repository

import "github.com/bary822/gomemon-server/internal/entity"

type MemoRepository interface {
	Save(memo entity.Memo) (*entity.Memo, error)
	GetByID(id string) (*entity.Memo, error)
	GetAll() ([]*entity.Memo, error)
	Delete(id string) error
}
