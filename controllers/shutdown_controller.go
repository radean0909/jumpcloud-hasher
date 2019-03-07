package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/radean0909/jumpcloud-hasher/utils/database"
	"github.com/radean0909/jumpcloud-hasher/utils/queue"
)

type ShutdownController struct {
	*BaseController
}

type ServerShutdown struct {
	Pending       int64   `json:"pending"`
	EstimatedTime float32 `json:"estimatedTime"`
}

func NewShutdownController(db *database.DB, q *queue.Queue) *ShutdownController {
	return &ShutdownController{
		BaseController: NewBaseController(db, q),
	}
}

func (sc *ShutdownController) Get(w http.ResponseWriter, r *http.Request) {

	maxTime := float32(100.0) // set max hash time to 100 ms

	if r.Method != "GET" {
		http.Error(w, "Unsupported Method Type: "+r.Method, http.StatusMethodNotAllowed)
		return
	}

	if sc.q.RunningState == false {
		http.Error(w, "Server Shutting Down", http.StatusServiceUnavailable)
		return
	}

	// Begin graceful shutdown
	sc.q.RunningState = false

	estimatedTime := sc.db.Tables["password"].Meta.Hashtime * float32(sc.db.Tables["password"].Meta.Count-sc.q.Completed)
	if sc.db.Tables["password"].Meta.Hashtime == 0 {
		estimatedTime = float32(sc.db.Tables["password"].Meta.Count-sc.q.Completed) * maxTime
	}
	// Return some stats
	stats := &ServerShutdown{
		Pending:       sc.db.Tables["password"].Meta.Count - sc.q.Completed,
		EstimatedTime: estimatedTime,
	}

	statsJSON, err := json.Marshal(stats)

	if err != nil {
		http.Error(w, "Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, string(statsJSON))

	go sc.q.Shutdown()

}
