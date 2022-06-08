package helper

import (
	"errors"
	"strings"
)

type ErrorCustom struct {
	Messages []string
}

func (e *ErrorCustom) Append(err error) *ErrorCustom {
	if err != nil {
		e.Messages = append(e.Messages, err.Error())
	}
	return e
}

func (e *ErrorCustom) Message() error {
	if len(e.Messages) > 0 {
		msg := strings.Join(e.Messages, "; ")
		return errors.New(msg)
	}
	return nil
}
