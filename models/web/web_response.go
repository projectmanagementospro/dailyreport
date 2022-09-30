package web

import "time"

type WebResponse struct {
	Code     int           `json:"code"`
	Status   string        `json:"status"`
	Errors   interface{}   `json:"errors"`
	Data     interface{}   `json:"data"`
	Duration time.Duration `json:"duration"`
}
