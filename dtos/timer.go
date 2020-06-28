package dtos

import (
	"context"
	"time"
)

type Timer struct {
	ID         string             `json:"id"`
	ModifiedAt *time.Time         `json:"modifiedAt,omitempty"`
	CreatedAt  *time.Time         `json:"createdAt,omitempty"`
	StepTime   float64            `json:"stepTime"`
	Counter    float64            `json:"counter"`
	Context    context.Context    `json:"_"`
	Cancel     context.CancelFunc `json:"_"`
}

type Data struct {
	Items []*Timer
}
