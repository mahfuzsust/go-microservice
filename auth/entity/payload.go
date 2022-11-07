package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	Id        uuid.UUID `json:id`
	TokenId   uuid.UUID `json:token_id`
	Username  string    `json:username`
	IssuedAt  time.Time `json:issuedAt`
	ExpiredAt time.Time `json:expiredAt`
}

var errExpiredToken = errors.New("Token expired")

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return errExpiredToken
	}
	return nil
}

func NewPayload(username string, tokenId uuid.UUID, duration time.Duration) *Payload {
	p := &Payload{
		Id:        uuid.New(),
		TokenId:   tokenId,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return p
}
