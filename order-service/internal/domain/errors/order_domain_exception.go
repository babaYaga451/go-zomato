package errors

type Exception struct {
	message string
}

func (ex *Exception) Error() string {
	return ex.message
}

func NewOrderDomainException(mssg string) *Exception {
	return &Exception{
		message: mssg,
	}
}

func NewOrderNotFoundException(mssg string) *Exception {
	return &Exception{
		message: mssg,
	}
}
