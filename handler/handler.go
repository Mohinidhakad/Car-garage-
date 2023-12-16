package handler

import (
	"api/models"
	"api/store"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	store store.Store
}

func New(s store.Store) handler {
	return handler{store: s}
}

type response struct {
	car []models.Car
}

func (h handler) Get(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.store.Get(ctx)
	if err != nil {
		return nil, err
	}

	return response{car: resp}, nil
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var car models.Car
	if err := ctx.Bind(&car); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	resp, err := h.store.Create(ctx, car)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	carId := ctx.Param("id")
	var car models.Car

	car.ID, _ = strconv.Atoi(carId)
	res, err := h.store.Update(ctx, car)

	if err != nil {
		return res, err

	}

	return res, nil
}
func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	carId := ctx.Param("id")
	var car models.Car

	car.ID, _ = strconv.Atoi(carId)

	err := h.store.Delete(ctx, car.ID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"message": "Deletion successful", "id": carId}, nil
}
