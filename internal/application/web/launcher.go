package application

import (
	"log"
	"net/http"

	file_repository "github.com/bary822/gomemon-server/internal/repository/file"
)

type WebLauncher struct{}

func (wl WebLauncher) Launch() {
	mux := http.NewServeMux()

	create_memo_handler := CreateMemoHandler{}
	get_memo_by_id_handler := GetMemoByIDHandler{}
	get_all_memos_handler := GetAllMemosHandler{}

	// DI
	memo_repository := MemoStorage{
		repository: file_repository.NewMemoFileRepository(),
	}

	router := WebRouter{
		create_memo_handler:    create_memo_handler.NewCreateMemoHandler(memo_repository),
		get_memo_by_id_handler: get_memo_by_id_handler.NewGetMemoByIDHandler(memo_repository),
		get_all_memos_handler:  get_all_memos_handler.NewGetAllMemosHandler(memo_repository),
	}
	router.RegisterRoutes(mux)

	log.Println("Listening on :8080")
	http.ListenAndServe("localhost:8080", mux)
}
