package interactor

import (
	"github.com/bary822/gomemon-server/internal/repository"
	"github.com/bary822/gomemon-server/internal/usecase"
)

type MemoDeleteInteractor struct {
	repo repository.MemoRepository
}

func NewMemoDeleteInteractor(r repository.MemoRepository) usecase.MemoDeleteUseCase {
	return MemoDeleteInteractor{
		repo: r,
	}
}

func (ir MemoDeleteInteractor) Handle(req usecase.MemoDeleteRequest) usecase.MemoDeleteResponse {
	id := req.ID
	error := ir.repo.Delete(id)

	if error != nil {
		return usecase.MemoDeleteResponse{IsSuccess: false}
	} else {
		return usecase.MemoDeleteResponse{IsSuccess: true}
	}
}
