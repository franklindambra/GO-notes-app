package n

import (
	"errors"
	"time"
)

type Note struct {
	NoteTitle       string    `json:"title"`
	NoteDescription string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
}

func New(title, description string) (*Note, error) {

	if title == "" || description == "" {
		var err error = errors.New("content length must be greater than 0")
		return nil, err
	}

	return &Note{
		title,
		description,
		time.Now(),
	}, nil

}
