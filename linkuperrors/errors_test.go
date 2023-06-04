package linkuperrors_test

import (
	"net/http"
	"testing"

	"github.com/marcbudd/linkup-service/linkuperrors"
	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	message := "Something went wrong"
	statusCode := http.StatusInternalServerError

	err := linkuperrors.New(message, statusCode)

	assert.Equal(t, message, err.Error())
	assert.Equal(t, statusCode, err.HTTPStatusCode())
}
