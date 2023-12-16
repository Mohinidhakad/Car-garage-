package Testcase

import (
	"api/models"
	"api/store"

	"github.com/DATA-DOG/go-sqlmock"
	"gofr.dev/pkg/gofr"
	"testing"
)

func TestStore_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id,CustomerName,CarName,Status FROM car").WillReturnRows(sqlmock.NewRows([]string{"id", "CustomerName", "CarName", "Status"}).AddRow(1, "John Doe", "Test Car", "Active"))
	ctx := &gofr.Context{}
	s := store.New()

	cars, err := s.Get(ctx)
	if err != nil {
		t.Errorf("Error getting cars: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}

	expectedLength := 1
	if len(cars) != expectedLength {
		t.Errorf("Expected %d cars, got %d", expectedLength, len(cars))
	}

	if cars[0].CustomerName != "John Doe" {
		t.Errorf("Expected CustomerName to be 'John Doe', got '%s'", cars[0].CustomerName)
	}
}

func TestStore_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO cars").WithArgs(sqlmock.AnyArg(), "John Doe", "Test Car", "Active").WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery("SELECT id,CustomerName,CarName,Status FROM cars").WithArgs(sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"id", "CustomerName", "CarName", "Status"}).AddRow(1, "John Doe", "Test Car", "Active"))
	ctx := &gofr.Context{}
	s := store.New()

	testCar := models.Car{
		CustomerName: "John Doe",
		CarName:      "Test Car",
		Status:       "Active",
	}

	createdCar, err := s.Create(ctx, testCar)
	if err != nil {
		t.Errorf("Error creating car: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}

	if createdCar.ID == 0 {
		t.Errorf("Expected a valid ID for the created car, got 0")
	}

	if createdCar.CustomerName != "John Doe" {
		t.Errorf("Expected CustomerName to be 'John Doe', got '%s'", createdCar.CustomerName)
	}
}

func TestStore_Update(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("UPDATE cars").WithArgs(1, "John Doe", "Test Car", "Active").WillReturnResult(sqlmock.NewResult(0, 1))

	ctx := &gofr.Context{}
	s := store.New()

	testCar := models.Car{
		ID:           1,
		CustomerName: "John Doe",
		CarName:      "Test Car",
		Status:       "Active",
	}

	updatedCar, err := s.Update(ctx, testCar)
	if err != nil {
		t.Errorf("Error updating car: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}

	if updatedCar.ID != testCar.ID {
		t.Errorf("Expected updated car ID to be %d, got %d", testCar.ID, updatedCar.ID)
	}
}
func TestStore_Delete(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM cars").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	ctx := &gofr.Context{}
	s := store.New()

	testCarID := 1

	err = s.Delete(ctx, testCarID)
	if err != nil {
		t.Errorf("Error deleting car: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
