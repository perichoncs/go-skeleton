package domain

import "time"

type Player struct {
	Name         string    `json:"name,omitempty"`
	Age          int       `json:"age,omitempty"`
	CreationTime time.Time `json:"-"`
}
