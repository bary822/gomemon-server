package interactor

import (
	"github.com/bary822/gomemon-server/internal/repository"
	"github.com/bary822/gomemon-server/internal/usecase"
)

type MemoGetAllInteractor struct {
	repo repository.MemoRepository
}

func NewMemoGetAllInteractor(r repository.MemoRepository) usecase.MemoGetAllUseCase {
	return MemoGetAllInteractor{
		repo: r,
	}
}

func (ir MemoGetAllInteractor) Handle(req usecase.MemoGetAllRequest) usecase.MemoGetAllResponse {
	memos, error := ir.repo.GetAll()

	if error != nil {
		// TODO: Handle it properly
		return usecase.MemoGetAllResponse{}
	}

	res := usecase.MemoGetAllResponse{Memos: memos}
	return res
}
