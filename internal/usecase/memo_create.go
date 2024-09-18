package usecase

type MemoCreateRequest struct {
	Content string
}

type MemoCreateResponse struct {
	MemoID  string
	Content string
}

type MemoCreateUseCase interface {
	Handle(req MemoCreateRequest) MemoCreateResponse
}
