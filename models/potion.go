package models

type Potion struct {
	ID          int64
	Name        string
	Description string
	Price       float64
	Properties  []string
}

type CartItem struct {
	Potion   *Potion
	Quantity int
}

func (ci *CartItem) Subtotal() float64 {
	return ci.Potion.Price * float64(ci.Quantity)
}
