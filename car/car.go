package car

type CardId string

type Brand struct {
	name string
	*SubBrand
}

type SubBrand struct {
	name string
}

func NewBrand(name string, sub *SubBrand) *Brand {
	return &Brand{
		name:     name,
		SubBrand: sub,
	}
}

type Car struct {
	CardId
	number string
	Brand
}

type Repository interface {
	Insert(car Car) error
	Find(id CardId) Car
}

func NewCar(cardId CardId, number string, brand Brand) *Car {
	return &Car{CardId: cardId, number: number, Brand: brand}
}
