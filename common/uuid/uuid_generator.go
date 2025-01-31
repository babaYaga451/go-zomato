package uuid

type UUIDGenerator interface {
	GenerateOrderID() string
	GenerateTrackingID() string
	GenerateAddressID() string
}
