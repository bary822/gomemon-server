package interactor

import (
	"github.com/bary822/gomemon-server/internal/repository"
	"github.com/bary822/gomemon-server/internal/usecase"
)

type MemoEditInteractor struct {
	repo repository.MemoRepository
}

func NewMemoEditInteractor(r repository.MemoRepository) usecase.MemoEditUseCase {
	return MemoEditInteractor{
		repo: r,
	}
}

func (ir MemoEditInteractor) Handle(req usecase.MemoEditRequest) usecase.MemoEditResponse {
	id := req.ID
	content := req.Content
	memo, error := ir.repo.Edit(id, content)

	if error != nil {
		// TODO: Handle it properly
		return usecase.MemoEditResponse{}
	}

	res := usecase.MemoEditResponse{MemoID: memo.ID, Content: memo.Content}
	return res
}
