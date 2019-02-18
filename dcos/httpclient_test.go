package dcos

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultTransportBaseNil(t *testing.T) {
	c := &DefaultTransport{
		Base: nil,
	}
	rt := c.base()

	assert.ObjectsAreEqual(http.DefaultTransport, rt)
}

func TestDefaultTransportBase(t *testing.T) {
	c := &DefaultTransport{
		Base: &DefaultTransport{Base: nil},
	}
	rt := c.base()

	assert.IsType(t, &DefaultTransport{}, rt)
}
