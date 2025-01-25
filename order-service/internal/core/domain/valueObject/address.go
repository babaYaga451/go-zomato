package valueobject

type Address struct {
	id         string
	postalCode string
	street     string
	city       string
}

func NewAddress(id string, postalCode string, street string, city string) *Address {
	return &Address{
		id:         id,
		postalCode: postalCode,
		street:     street,
		city:       city,
	}
}

func (a *Address) GetID() string {
	return a.id
}

func (a *Address) GetPostalCode() string {
	return a.postalCode
}

func (a *Address) GetStreet() string {
	return a.street
}

func (a *Address) GetCity() string {
	return a.city
}
