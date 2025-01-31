package valueobject

type Address struct {
	postalCode string
	street     string
	city       string
}

func NewAddress(postalCode string, street string, city string) *Address {
	return &Address{
		postalCode: postalCode,
		street:     street,
		city:       city,
	}
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
