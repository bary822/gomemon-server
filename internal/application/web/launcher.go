package application

import (
	"log"
	"net/http"

	in_memory_repository "github.com/bary822/gomemon-server/internal/repository/in_memory"
)

type WebLauncher struct{}

func (wl WebLauncher) Launch() {
	mux := http.NewServeMux()

	create_memo_handler := CreateMemoHandler{}
	get_memo_by_id_handler := GetMemoByIDHandler{}

	// DI
	memo_repository := MemoStorage{
		repository: in_memory_repository.NewMemoInMemoryRepository(),
	}

	router := WebRouter{
		create_memo_handler:    create_memo_handler.NewCreateMemoHandler(memo_repository),
		get_memo_by_id_handler: get_memo_by_id_handler.NewGetMemoByIDHandler(memo_repository),
	}
	router.RegisterRoutes(mux)

	log.Println("Listening on :8080")
	http.ListenAndServe("localhost:8080", mux)
}
