package valueObject

type Money struct {
	amount float64
}

func NewMoney(value float64) Money {
	return Money{
		amount: value,
	}
}

func (m Money) Add(o Money) Money {
	return Money{
		amount: m.amount + o.amount,
	}
}

func (m Money) IsGreaterThanZero() bool {
	return m.amount > float64(0)
}

func (m Money) Equals(o Money) bool {
	return m.amount == o.amount
}

func (m Money) GetAmount() float64 {
	return m.amount
}

func (m Money) Multiply(multiplier int) Money {
	return Money{
		amount: m.amount * float64(multiplier),
	}
}
