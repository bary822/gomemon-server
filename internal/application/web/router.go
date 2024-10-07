package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/bary822/gomemon-server/internal/controller"
	"github.com/bary822/gomemon-server/internal/interactor"
	repository "github.com/bary822/gomemon-server/internal/repository"
)

type WebRouter struct {
	create_memo_handler    *CreateMemoHandler
	get_memo_by_id_handler *GetMemoByIDHandler
	get_all_memos_handler  *GetAllMemosHandler
	delete_memo_handler    *DeleteMemoHandler
	edit_memo_handler      *EditMemoHandler
}

// interfaceとしてのMemoRepositoryと名前がカブるのでMemoStorageと命名しておく
// 実態はMemoRepositoryだけをメンバに持つ単純な構造体
type MemoStorage struct {
	// 具象クラス(MemoInMemoryRepository)ではなく、抽象クラス(MemoRepository interface)に依存させているのがポイント
	repository repository.MemoRepository
}

type CreateMemoHandler struct {
	controller controller.MemoController
}

type GetMemoByIDHandler struct {
	controller controller.MemoController
}

type GetAllMemosHandler struct {
	controller controller.MemoController
}

type DeleteMemoHandler struct {
	controller controller.MemoController
}

type EditMemoHandler struct {
	controller controller.MemoController
}

func (h *CreateMemoHandler) NewCreateMemoHandler(s MemoStorage) *CreateMemoHandler {
	// DI
	usecase := interactor.NewMemoCreateInteractor(s.repository)
	controller := controller.NewCreateMemoController(usecase)

	return &CreateMemoHandler{
		controller: controller,
	}
}

func (h *GetMemoByIDHandler) NewGetMemoByIDHandler(s MemoStorage) *GetMemoByIDHandler {
	// DI
	usecase := interactor.NewMemoGetByIDInteractor(s.repository)
	controller := controller.NewGetMemoByIDController(usecase)

	return &GetMemoByIDHandler{
		controller: controller,
	}
}

func (h *GetAllMemosHandler) NewGetAllMemosHandler(s MemoStorage) *GetAllMemosHandler {
	// DI
	usecase := interactor.NewMemoGetAllInteractor(s.repository)
	controller := controller.NewGetAllMemosController(usecase)

	return &GetAllMemosHandler{
		controller: controller,
	}
}

func (h *DeleteMemoHandler) NewDeleteMemoHandler(s MemoStorage) *DeleteMemoHandler {
	// DI
	usecase := interactor.NewMemoDeleteInteractor(s.repository)
	controller := controller.NewDeleteMemoController(usecase)

	return &DeleteMemoHandler{
		controller: controller,
	}
}

func (h *EditMemoHandler) NewEditMemoHandler(s MemoStorage) *EditMemoHandler {
	// DI
	usecase := interactor.NewMemoEditInteractor(s.repository)
	controller := controller.NewEditMemoController(usecase)

	return &EditMemoHandler{
		controller: controller,
	}
}

func (h *CreateMemoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if jsonBody, err := parseJSON(*r); !errors.Is(err, nil) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Print(err)
		fmt.Fprintln(w, "{}")
	} else {
		memoContent := jsonBody["content"].(string)
		memo_res := h.controller.CreateMemo(memoContent)

		if json, err := json.Marshal(memo_res); !errors.Is(err, nil) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Print(err)
			fmt.Fprintln(w, "{}")
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, string(json))
		}
	}
}

func (h *GetMemoByIDHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id := strings.TrimPrefix(r.URL.Path, "/memos/")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "{}")
	} else {
		memo_res := h.controller.GetMemoByID(id)
		if memo_res.MemoID == "" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "{}")
		}

		if json, err := json.Marshal(memo_res); !errors.Is(err, nil) {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "{}")
		} else {
			fmt.Fprintln(w, string(json))
		}
	}
}

func (h *GetAllMemosHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	memos_res := h.controller.GetAllMemos()
	if json, err := json.Marshal(memos_res); !errors.Is(err, nil) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "{}")
	} else {
		fmt.Fprintln(w, string(json))
	}
}

func (h *DeleteMemoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id := strings.TrimPrefix(r.URL.Path, "/memos/")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "{}")
	} else {
		memo_res := h.controller.DeleteMemo(id)
		if memo_res.IsSuccess {
			w.WriteHeader(http.StatusNoContent)
			fmt.Fprintln(w, "{}")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "{}")
		}
	}
}

func (h *EditMemoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id := strings.TrimPrefix(r.URL.Path, "/memos/")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "{}")
		return
	}

	jsonBody, err := parseJSON(*r)
	if !errors.Is(err, nil) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Print(err)
		fmt.Fprintln(w, "{}")
		return
	}

	if jsonBody["content"] == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Print("Body must contain entry 'content'")
		fmt.Fprintln(w, "{}")
		return
	}

	memoContent := jsonBody["content"].(string)
	memo_res := h.controller.EditMemo(id, memoContent)

	if memo_res.MemoID == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "{}")
	}

	if json, err := json.Marshal(memo_res); !errors.Is(err, nil) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "{}")
	} else {
		fmt.Fprintln(w, string(json))
	}
}

func HandleCORSPreflight(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wl := NewWhiteList()
		if requester_origin := r.Header.Get("Origin"); wl.IsAllowedOrigin(requester_origin) {
			w.Header().Set("Access-Control-Allow-Origin", requester_origin)
		}
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
		}

		next.ServeHTTP(w, r)
	}
}

func (wr *WebRouter) HandleMemos(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request " + r.Method + ": " + r.URL.Path)

	switch r.Method {
	case http.MethodPost:
		wr.create_memo_handler.Handle(w, r)
	case http.MethodGet:
		wr.get_all_memos_handler.Handle(w, r)
	}
}

func (wr *WebRouter) HandleMemo(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request " + r.Method + ": " + r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		wr.get_memo_by_id_handler.Handle(w, r)
	case http.MethodDelete:
		wr.delete_memo_handler.Handle(w, r)
	case http.MethodPut:
		wr.edit_memo_handler.Handle(w, r)
	}
}

func (wr *WebRouter) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/memos", HandleCORSPreflight(wr.HandleMemos))
	mux.HandleFunc("/memos/", HandleCORSPreflight(wr.HandleMemo))
}

func parseJSON(r http.Request) (map[string]any, error) {
	if r.Header.Get("Content-Type") != "application/json" {
		err := errors.New("Content-Type must be 'application/json'.")
		return nil, err
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		err := errors.New("Content-Length must be present.")
		return nil, err
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		err := errors.New("Something went wrong while reading request body")
		return nil, err
	}

	//parse json
	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		err := errors.New("Something went wrong while parsing body")
		return nil, err
	}

	return jsonBody, nil
}
