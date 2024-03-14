package ports

import (
	server "distributed_log/internal/common/server/http"
	"distributed_log/internal/event"
	"distributed_log/internal/log"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"time"
)

type Routes struct {
	handler *event.Handler
}

// func NewRoutes(r *chi.Mux, ser IndexInterface) Routes {
func NewRoutes(r *chi.Mux, ser *event.Handler) Routes {
	ru := Routes{handler: ser}
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Route("/index", func(r chi.Router) {

		r.Post("/new", ru.CreateIndex)   // POST /index
		r.Get("/create", ru.CreateIndex) // DEBUG
		r.Get("/index", ru.GetIndex)     // GET /index/search

	})
	return ru
}

func (ru *Routes) CreateIndex(w http.ResponseWriter, r *http.Request) {

	i := log.Index{}
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		slog.Error("(ru *Routes) CreateIndex NewDecoder:", err)
		WriteResponse(w, server.Response{Status: "Err", StatusCode: http.StatusBadRequest, Message: http.StatusText(http.StatusBadRequest), RequestTime: time.Now()})

		return
	}

	newEvent := event.Event{Type: string(event.EventInsert), Data: i}

	off, err := ru.handler.EventBus.Publish(&newEvent)
	fmt.Println(off)
	if err != nil {
		slog.Error("(ru *Routes) CreateIndex Publish:", err)
		WriteResponse(w, server.Response{Status: "Err", StatusCode: http.StatusInternalServerError, Message: http.StatusText(http.StatusInternalServerError), RequestTime: time.Now()})
	}

	WriteResponse(w, server.Response{Status: "OK", StatusCode: http.StatusCreated, Message: http.StatusText(http.StatusCreated), RequestTime: time.Now()})
}

func (ru *Routes) GetIndex(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	//index, ok := ctx.Value("index").(*index.Index)
	//if !ok {
	//	http.Error(w, http.StatusText(422), 422)
	//	return
	//}
	//w.write([]byte(fmt.Sprintf("index id:%s", index.Id)))
	fmt.Println("Log(w http.ResponseWriter, r *http.Request) ")
	//ru.ser.Get(&index.Index{Id: "___x7777"})
}

func WriteResponse(w http.ResponseWriter, response server.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	err := enc.Encode(response)
	if err != nil {
		slog.Error("WriteResponse", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
