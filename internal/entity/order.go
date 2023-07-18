package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.Validate()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("ID é requerida")
	}
	if o.Price <= 0 {
		return errors.New("Preço tem que ser maior que 0")
	}
	if o.Tax <= 0 {
		return errors.New("A Taxa tem que ser maior que 0")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Price
	err:= o.Validate()
	if err != nil{
		return err
	}
	return nil
}
