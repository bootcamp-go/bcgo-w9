package tasks

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetByIdHappyPath(t *testing.T) {
	// Arrange
	// Declaro el parametro que necesito
	taskId := "1"

	// Declaro la task que se supone que la bbdd deberia devolverme
	taskParam := Task{
		ID:          "1",
		Name:        "Writing",
		Description: "It's just writing",
		Completed:   true,
	}

	// Declaro un expected que el metodo del servicio deberia devolver
	taskExpected := Task{
		ID:          "1",
		Name:        "Writing",
		Description: "It's just writing",
		Completed:   true,
	}
	// Mockeo la bbdd que es la capa que no testeo
	mockBd := NewMockDatabase(taskParam, nil)
	// declaro la dependencia del repo
	repo := &DefaultRepository{mockBd}
	// declaro la dependencia del service
	serv := &DefaultService{repo}

	// Act
	// Ejecuto y guardo los que me devuelve el service
	taskObtained, err := serv.GetByID(taskId)

	// Assert
	// evaluo la igualdad entre el expected y el obtained
	assert.Equal(t, taskExpected, taskObtained)
	// Evaluo que el error venga nulo porque es el happy path
	assert.Equal(t, err, nil)
}

func TestGetByIdButIdDoesNotExist(t *testing.T) {
	// Arrange
	// Declaro el parametro que necesito
	taskId := "100000"

	errorParam := errors.New("No se encontro dicho id")

	var taskParam Task

	taskExpected := Task{}

	errorExpected := errors.New("No se encontro dicho id")

	// Mockeo la bbdd que es la capa que no testeo
	mockBd := NewMockDatabase(taskParam, errorParam)
	// declaro la dependencia del repo
	repo := &DefaultRepository{mockBd}
	// declaro la dependencia del service
	serv := &DefaultService{repo}

	// Act
	// Ejecuto y guardo los que me devuelve el service
	taskObtained, err := serv.GetByID(taskId)

	// Assert

	assert.Equal(t, taskExpected, taskObtained)

	assert.Equal(t, err, errorExpected)
}
