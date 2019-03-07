package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/radean0909/jumpcloud-hasher/utils/database"
	"github.com/radean0909/jumpcloud-hasher/utils/queue"
)

type StatsController struct {
	*BaseController
}

type ServerStats struct {
	Total   int64   `json:"total"`
	Average float32 `json:"average"`
	// TODO: Add a few additional metrics here
}

func NewStatsController(db *database.DB, q *queue.Queue) *StatsController {
	return &StatsController{
		BaseController: NewBaseController(db, q),
	}
}

func (sc *StatsController) Get(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Unsupported Method Type: "+r.Method, http.StatusMethodNotAllowed)
		return
	}

	if sc.q.RunningState == false {
		http.Error(w, "Server Shutting Down", http.StatusServiceUnavailable)
		return
	}

	// Get the simple stats

	stats := &ServerStats{
		Total:   sc.q.Completed,
		Average: sc.db.Tables["password"].Meta.Hashtime,
	}

	statsJSON, err := json.Marshal(stats)

	if err != nil {
		http.Error(w, "Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(statsJSON))
}
