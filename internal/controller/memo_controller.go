package controller

import "github.com/bary822/gomemon-server/internal/usecase"

type MemoController struct {
	memoCreateUsecase  usecase.MemoCreateUseCase
	memoGetByIDUsecase usecase.MemoGetByIDUseCase
	memoGetAllUseCase  usecase.MemoGetAllUseCase
	memoDeleteUseCase  usecase.MemoDeleteUseCase
	memoEditUseCase    usecase.MemoEditUseCase
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

func NewGetAllMemosController(uc_get_all_memo usecase.MemoGetAllUseCase) MemoController {
	return MemoController{
		memoGetAllUseCase: uc_get_all_memo,
	}
}

func NewDeleteMemoController(uc_delete_memo usecase.MemoDeleteUseCase) MemoController {
	return MemoController{
		memoDeleteUseCase: uc_delete_memo,
	}
}

func NewEditMemoController(uc_edit_memo usecase.MemoEditUseCase) MemoController {
	return MemoController{
		memoEditUseCase: uc_edit_memo,
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

func (c MemoController) GetAllMemos() usecase.MemoGetAllResponse {
	req := usecase.MemoGetAllRequest{}
	return c.memoGetAllUseCase.Handle(req)
}

func (c MemoController) DeleteMemo(id string) usecase.MemoDeleteResponse {
	req := usecase.MemoDeleteRequest{ID: id}
	return c.memoDeleteUseCase.Handle(req)
}

func (c MemoController) EditMemo(id string, content string) usecase.MemoEditResponse {
	req := usecase.MemoEditRequest{ID: id, Content: content}
	return c.memoEditUseCase.Handle(req)
}
