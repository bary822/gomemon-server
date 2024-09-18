package controller

import "github.com/bary822/gomemon-server/internal/usecase"

type MemoController struct {
	memoCreateUsecase  usecase.MemoCreateUseCase
	memoGetByIDUsecase usecase.MemoGetByIDUseCase
}

func NewCreateMemoController(uc_memo_create usecase.MemoCreateUseCase) MemoController {
	return MemoController{
		memoCreateUsecase: uc_memo_create,
	}
}

func NewGetMemoByIDController(uc_get_memo_by_id usecase.MemoGetByIDUseCase) MemoController {
	return MemoController{
		memoGetByIDUsecase: uc_get_memo_by_id,
	}
}

func (c MemoController) CreateMemo(content string) usecase.MemoCreateResponse {
	req := usecase.MemoCreateRequest{Content: content}
	return c.memoCreateUsecase.Handle(req)
}

func (c MemoController) GetMemoByID(id string) usecase.MemoGetByIDResponse {
	req := usecase.MemoGetByIDRequest{ID: id}
	return c.memoGetByIDUsecase.Handle(req)
}
