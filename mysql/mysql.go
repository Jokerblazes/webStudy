package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"webStudy/car"
)

type repository struct {
	engine *xorm.Engine
}

func NewRepository() *repository {
	engine, err := xorm.NewEngine("mysql", "root:@/test?charset=utf8")
	if err != nil {
		panic("connect db failed!")
	}
	return &repository{engine: engine}
}

func (r repository) Insert(car car.Car) error {
	_, err := r.engine.InsertOne(car)
	return err
}

func (r repository) Find(id car.CardId) car.Car {
	car := car.Car{CardId: id}
	_, _ = r.engine.Get(&car)
	return car
}
