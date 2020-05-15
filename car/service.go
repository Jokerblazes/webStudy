package car

type Service interface {
	Cars() []Car

	Car(id CardId) Car

	Save(car Car) error
}

type service struct {
}

func (svc *service) Save(car Car) error {
	panic("implement me")
}

func (svc *service) Cars() []Car {
	panic("implement me")
}

func (svc *service) Car(id CardId) Car {
	panic("implement me")
}
