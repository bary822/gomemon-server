package interactor

import (
	"github.com/bary822/gomemon-server/internal/entity"
	"github.com/bary822/gomemon-server/internal/repository"
	"github.com/bary822/gomemon-server/internal/usecase"
	"github.com/google/uuid"
)

type MemoCreateInteractor struct {
	repo repository.MemoRepository
}

func NewMemoCreateInteractor(r repository.MemoRepository) usecase.MemoCreateUseCase {
	return MemoCreateInteractor{
		repo: r,
	}
}

func (ir MemoCreateInteractor) Handle(req usecase.MemoCreateRequest) usecase.MemoCreateResponse {
	content := req.Content
	uuid, err := uuid.NewUUID()
	if err != nil {
		return usecase.MemoCreateResponse{}
	}
	new_memo := entity.Memo{ID: uuid.String(), Content: content}

	memo, error := ir.repo.Save(new_memo)

	if error != nil {
		// TODO: Handle it properly
		return usecase.MemoCreateResponse{}
	}

	res := usecase.MemoCreateResponse{MemoID: memo.ID, Content: memo.Content}
	return res
}
