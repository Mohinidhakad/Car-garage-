package Testcase

import (
	"api/handler"
	"api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gofr.dev/pkg/gofr"
	"testing"
)

type MockCarStore struct {
	mock.Mock
}

func (m *MockCarStore) Create(ctx *gofr.Context, car *models.Car) (*models.Car, error) {
	args := m.Called(ctx, car)
	return args.Get(0).(*models.Car), args.Error(1)
}

func (m *MockCarStore) Update(ctx *gofr.Context, car *models.Car) (*models.Car, error) {
	args := m.Called(ctx, car)
	return args.Get(0).(*models.Car), args.Error(1)
}

func (m *MockCarStore) Delete(ctx *gofr.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateHandler(t *testing.T) {
	mockStore := &MockCarStore{}
	h := handler.New(mockStore)

	mockStore.On("Create", mock.Anything, mock.AnythingOfType("*model.car")).Return(&models.Car{
		CarName: "KIA", CustomerName: "JOhn", Status: "dispatched"}, nil)

	mockContext := new(gofr.Context)

	result, err := h.Create(mockContext)

	expected := &models.Car{
		CarName: "KIA", CustomerName: "JOhn", Status: "dispatched"}
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	mockStore.AssertExpectations(t)
}

func TestUpdateHandler(t *testing.T) {
	mockStore := &MockCarStore{}
	h := handler.New(MockCarStore{})

	mockStore.On("Update", mock.Anything, mock.AnythingOfType("*model.car")).Return(&models.Car{
		CarName: "KIA", CustomerName: "JOhn", Status: "dispatched"}, nil)

	mockContext := new(gofr.Context)

	result, err := h.Update(mockContext)

	expected := &models.Car{
		CarName: "KIA", CustomerName: "JOhn", Status: "dispatched"}
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	mockStore.AssertExpectations(t)
}

func TestDeleteHandler(t *testing.T) {
	mockStore := &MockCarStore{}

	h := handler.New(MockCarStore{})

	mockStore.On("Delete", mock.Anything, 1).Return(nil)

	mockContext := new(gofr.Context)

	result, err := h.Delete(mockContext)

	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"message": "Deletion successful", "id": 1}, result)
	mockStore.AssertExpectations(t)
}
