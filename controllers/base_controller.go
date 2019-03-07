package controllers

import (
	"github.com/radean0909/jumpcloud-hasher/utils/database"
	"github.com/radean0909/jumpcloud-hasher/utils/queue"
)

type BaseController struct {
	db *database.DB
	q  *queue.Queue
}

func NewBaseController(db *database.DB, q *queue.Queue) *BaseController {
	return &BaseController{
		db: db,
		q:  q,
	}
}
