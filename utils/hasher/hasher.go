package hasher

import (
	"crypto/sha512"
	"encoding/base64"
	"time"

	"github.com/radean0909/jumpcloud-hasher/utils/queue"
)

func Process(j *queue.Job) *queue.Job {
	// Here is where we do our processing
	// Sleep 5 seconds
	time.Sleep(5 * time.Second)

	// Now hash the password
	startTime := time.Now()
	hasher := sha512.New()
	hasher.Write([]byte(j.In))
	j.Out = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	endTime := time.Now()
	j.ProcessTime = endTime.Sub(startTime).Nanoseconds() / 1000.0
	return j
}
