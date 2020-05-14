package car

type Service interface {
	Cars() []Car

	Car(id CardId) Car
}

type service struct {
}

func (svc *service) Cars() []Car {
	panic("implement me")
}

func (svc *service) Car(id CardId) Car {
	panic("implement me")
}
