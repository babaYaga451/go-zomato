package uuid

import "github.com/google/uuid"

type RandomUUIDGenerator struct{}

func NewRandomUUIDGeneratory() *RandomUUIDGenerator {
	return &RandomUUIDGenerator{}
}

func (u *RandomUUIDGenerator) GenerateOrderID() string {
	return uuid.New().String()
}

func (u *RandomUUIDGenerator) GenerateTrackingID() string {
	return uuid.New().String()
}

func (u *RandomUUIDGenerator) GenerateAddressID() string {
	return uuid.New().String()
}
