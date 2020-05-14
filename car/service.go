package car

type Service interface {
	Cars() []Car

	Car(id CardId) Car
}
