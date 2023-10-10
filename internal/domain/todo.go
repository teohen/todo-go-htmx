package domain

import "github.com/google/uuid"

type Todo struct {
	Id      uuid.UUID
	Title   string
	Checked bool
}
