package usecase

import "github.com/bary822/gomemon-server/internal/entity"

type MemoGetAllRequest struct{}

type MemoGetAllResponse struct {
	Memos []*entity.Memo
}

type MemoGetAllUseCase interface {
	Handle(req MemoGetAllRequest) MemoGetAllResponse
}
