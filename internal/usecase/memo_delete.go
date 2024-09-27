package usecase

type MemoDeleteRequest struct {
	ID string
}

type MemoDeleteResponse struct {
	IsSuccess bool
}

type MemoDeleteUseCase interface {
	Handle(req MemoDeleteRequest) MemoDeleteResponse
}
