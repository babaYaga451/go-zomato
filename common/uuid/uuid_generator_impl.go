package uuid

import "github.com/google/uuid"

type RandomUUIDGenerator struct{}

func NewRandomUUIDGeneratory() *RandomUUIDGenerator {
	return &RandomUUIDGenerator{}
}

func (u *RandomUUIDGenerator) Generate() string {
	return uuid.New().String()
}
