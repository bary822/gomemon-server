package interactor

import (
	"github.com/bary822/gomemon-server/internal/repository"
	"github.com/bary822/gomemon-server/internal/usecase"
)

type MemoGetByIDInteractor struct {
	repo repository.MemoRepository
}

func NewMemoGetByIDInteractor(r repository.MemoRepository) usecase.MemoGetByIDUseCase {
	return MemoGetByIDInteractor{
		repo: r,
	}
}

func (ir MemoGetByIDInteractor) Handle(req usecase.MemoGetByIDRequest) usecase.MemoGetByIDResponse {
	id := req.ID
	memo, error := ir.repo.GetByID(id)

	if error != nil {
		// TODO: Handle it properly
		return usecase.MemoGetByIDResponse{}
	}

	res := usecase.MemoGetByIDResponse{MemoID: memo.ID, Content: memo.Content}
	return res
}
