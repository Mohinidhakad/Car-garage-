package store

import (
	"api/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

type Store struct{}

func New() Store {
	return Store{}
}

func (c Store) Get(ctx *gofr.Context) ([]models.Car, error) {
	rows, err := ctx.DB().QueryContext(ctx, "SELECT id,CustomerName,CarName,Status FROM car")
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	car := make([]models.Car, 0)

	for rows.Next() {
		var c models.Car

		err = rows.Scan(&c.ID, &c.CustomerName, &c.CarName, &c.Status)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		car = append(car, c)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return car, nil
}

func (c Store) Create(ctx *gofr.Context, car models.Car) (models.Car, error) {
	var resp models.Car

	queryInsert := "INSERT INTO cars (id,CustomerName,CarName,Status) VALUES (?, ?, ?, ?, ?)"

	result, err := ctx.DB().ExecContext(ctx, queryInsert, car.ID, car.CustomerName, car.CarName, car.Status)

	if err != nil {
		return models.Car{}, errors.DB{Err: err}
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return models.Car{}, errors.DB{Err: err}
	}

	querySelect := "SELECT id,CustomerName,CarName,Status FROM cars WHERE id = ?"

	err = ctx.DB().QueryRowContext(ctx, querySelect, lastInsertID).
		Scan(&resp.ID, &resp.CustomerName, &resp.CarName, &resp.Status)

	if err != nil {
		return models.Car{}, errors.DB{Err: err}
	}

	return resp, nil
}

func (c Store) Update(ctx *gofr.Context, emp models.Car) (models.Car, error) {
	var resp models.Car

	//update query
	queryUpdate := "UPDATE cars (id,CustomerName,CarName,Status) VALUES (?, ?, ?, ?, ?)"

	result, err := ctx.DB().ExecContext(ctx, queryUpdate, emp.ID, emp.CustomerName, emp.CarName, emp.Status)

	if err != nil {
		return models.Car{}, errors.DB{Err: err}
	}

	return resp, nil
}

func (c Store) Delete(ctx *gofr.Context, emp models.Car) (models.Car, error) {
	var resp models.Car

	//delete query
	queryDelete := "DELETE FROM cars WHERE (id,CustomerName,CarName,Status) VALUES (?, ?, ?, ?, ?)"

	// Execute the INSERT query
	result, err := ctx.DB().ExecContext(ctx, queryDelete, emp.ID, emp.CustomerName, emp.CarName, emp.Status)

	if err != nil {
		return models.Car{}, errors.DB{Err: err}
	}
	return resp, nil
}
