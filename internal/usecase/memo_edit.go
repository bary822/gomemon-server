package usecase

type MemoEditRequest struct {
	ID      string
	Content string
}

type MemoEditResponse struct {
	MemoID  string
	Content string
}

type MemoEditUseCase interface {
	Handle(req MemoEditRequest) MemoEditResponse
}
