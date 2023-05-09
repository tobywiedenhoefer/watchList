package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tobywiedenhoefer/watchList/database/models"
	"io"
	"net"
	"net/http"

	"github.com/tobywiedenhoefer/watchList/database"
	"github.com/tobywiedenhoefer/watchList/database/tables"
)

const servePort = 3333

func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func routeWatchList(w http.ResponseWriter, r *http.Request) {
	db, err := database.Database("watchlist")
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, fmt.Sprintf("database.Database.error: %s", err))
		return
	}
	var wlr models.WatchListRow
	switch r.Method {
	case "GET":
		watchListRows, err := watchlist.GetAll(db)
		if err != nil {
			// set error headers
			io.WriteString(w, fmt.Sprintf("routeWatchList.GET.error: %s", err))
			return
		} else {
			setCORSHeaders(w)
			jsonBody, _ := json.Marshal(&watchListRows)
			w.Write(jsonBody)
		}
	case "POST":
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&wlr); err != nil {
			// set error headers
			io.WriteString(w, fmt.Sprintf("routeWatchList.POST.error: %s", err))
			return
		}
		addedID, err := watchlist.PostOne(db, wlr)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("routeWatchList.POST.error: %s", err))
			return
		}
		wlr.ID = addedID
		setCORSHeaders(w)
		w.WriteHeader(http.StatusCreated)
		jsonBody, _ := json.Marshal([]models.WatchListRow{wlr})
		w.Write(jsonBody)
	case "PUT":
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&wlr); err != nil {
			// set error headers
			io.WriteString(w, fmt.Sprintf("routeWatchList.PUT.error: %s", err))
			return
		}
		rowsAffected, err := watchlist.PutOne(db, wlr)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("routeWatchList.PUT.error: %s", err))
			return
		} else if rowsAffected == 0 {
			io.WriteString(w, fmt.Sprintf("routeWatchList.PUT.ZeroRowsAffected: %v", err))
			return
		}
		setCORSHeaders(w)
		data := map[string]interface{}{
			"rowsAffected": rowsAffected,
		}
		jsonBody, _ := json.Marshal(data)
		w.WriteHeader(http.StatusNoContent)
		w.Write(jsonBody)
	case "DELETE":
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&wlr); err != nil {
			io.WriteString(w, fmt.Sprintf("routeWatchList.PUT.error: %v", err))
			return
		}
		rowsDeleted, err := watchlist.Delete(db, &wlr.ID)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("routeWatchList.DELETE.error: %v", err))
			return
		}
		setCORSHeaders(w)
		w.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
			"rowsDeleted": rowsDeleted,
		}
		jsonBody, _ := json.Marshal(data)
		w.Write(jsonBody)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/watchlist", routeWatchList)
	ctx := context.Background()
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", servePort),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, servePort, l.Addr().String())
			return ctx
		},
	}
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("error running http server: %s \n", err)
		}
	}
}
