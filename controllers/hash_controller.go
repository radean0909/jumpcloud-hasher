package controllers

import (
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/radean0909/jumpcloud-hasher/requests"
	"github.com/radean0909/jumpcloud-hasher/utils/database"
	"github.com/radean0909/jumpcloud-hasher/utils/queue"
)

type HashController struct {
	*BaseController
}

func NewHashController(db *database.DB, q *queue.Queue) *HashController {
	return &HashController{
		BaseController: NewBaseController(db, q),
	}
}

func (hc *HashController) Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Unsupported Method Type: "+r.Method, http.StatusMethodNotAllowed)
		return
	}

	if hc.q.RunningState == false {
		http.Error(w, "Server Shutting Down", http.StatusServiceUnavailable)
		return
	}

	// update the database immediately

	hc.db.Tables["password"].Meta.Count++
	id := hc.db.Tables["password"].Meta.Increment
	hc.db.Tables["password"].Data[id] = ""
	hc.db.Tables["password"].Meta.Increment++

	// parse the body
	var hashRequest requests.CreateHashRequest
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error processing data", http.StatusBadRequest)
		return
	}

	if r.FormValue("password") == "" {
		http.Error(w, "Invalid Password", http.StatusBadRequest)
	}

	hashRequest.Password = r.FormValue("password")

	// queue crypto operation
	job := queue.NewJob(id, hc.db.Tables["password"], hashRequest.Password)
	hc.q.Add(job)

	// Since there is a processing time, we send back the "Accepted" status with the ID and a time estimate
	w.WriteHeader(http.StatusAccepted)

	fmt.Fprintln(w, id) // A more standardized way would be to output a JSON response

}

func (hc *HashController) Get(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Unsupported Method Type: "+r.Method, http.StatusMethodNotAllowed)
		return
	}

	if hc.q.RunningState == false {
		http.Error(w, "Server Shutting Down", http.StatusServiceUnavailable)
		return
	}

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	pass, ok := hc.db.Tables["password"].Data[int64(id)]
	if !ok || pass == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, pass) // A more standardized way would be to output a JSON response

}
