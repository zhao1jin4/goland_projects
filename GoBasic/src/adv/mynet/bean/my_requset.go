package bean

import "time"

type MyRequest struct {
	Id int
	Name string
	Weight float64
	Birthday time.Time
}

