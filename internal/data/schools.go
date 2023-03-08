//Filename: cmd/api/schools.go

package data

import (
	"time"
)

type School struct {
	ID       int64
	Name     string
	Level    string
	Contact  string
	Phone    string
	Email    string
	Website  string
	Address  string
	Mode     []string
	CreateAt time.Time
	Version  int32
}
