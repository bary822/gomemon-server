package usecase

type MemoGetByIDRequest struct {
	ID string
}

type MemoGetByIDResponse struct {
	MemoID  string
	Content string
}

type MemoGetByIDUseCase interface {
	Handle(req MemoGetByIDRequest) MemoGetByIDResponse
}
