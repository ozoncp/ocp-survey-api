package models

import "fmt"

type Survey struct {
	Id     uint64
	UserId uint64
	Link   string
}

func (s Survey) String() string {
	return fmt.Sprintf("Id: %d, UserId: %d, Link: %s", s.Id, s.UserId, s.Link)
}
