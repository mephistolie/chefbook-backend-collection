package entity

import "github.com/google/uuid"

type Collection struct {
	Id        uuid.UUID
	OwnerId   uuid.UUID
	Coauthors []uuid.UUID
	Name      string
	Emoji     *string
}
