package car

type Service interface {
	Cars() []Car

	Car(id CardId) Car

	Save(car Car) error
}

type service struct {
	Repository
}

func (svc *service) Save(car Car) error {
	svc.Repository.Insert(car)
}

func (svc *service) Cars() []Car {
	panic("implement me")
}

func (svc *service) Car(id CardId) Car {
	return svc.Find(id)
}
