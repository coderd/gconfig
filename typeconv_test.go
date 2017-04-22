package gconfig

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapStringBool(t *testing.T) {
	var err error

	// Test correct cases
	v := map[string]interface{}{
		"k1": true,
		"k2": false,
	}
	_, err = MapStringBool(v)
	assert.Nil(t, err)

	// Test failed cases
	v1 := map[string]interface{}{
		"k1": 1,
		"k2": 2,
	}
	_, err = MapStringBool(v1)
	assert.NotNil(t, err)
}

func TestMapStringInt(t *testing.T) {
	var err error

	// Test correct cases
	v := map[string]interface{}{
		"k1": 1,
		"k2": 2,
	}
	_, err = MapStringInt(v)
	assert.Nil(t, err)

	// Test failed cases
	v1 := map[string]interface{}{
		"k1": 1,
		"k2": 2.0,
	}
	_, err = MapStringInt(v1)
	assert.NotNil(t, err)
}

func TestMapStringFloat64(t *testing.T) {
	var err error

	// Test correct cases
	v := map[string]interface{}{
		"k1": 1.0,
		"k2": 2.1,
	}
	_, err = MapStringFloat64(v)
	assert.Nil(t, err)

	// Test failed cases
	v1 := map[string]interface{}{
		"k1": 1,
		"k2": 2,
	}
	_, err = MapStringFloat64(v1)
	assert.NotNil(t, err)
}

func TestMapStringString(t *testing.T) {
	var err error

	// Test correct cases
	v := map[string]interface{}{
		"k1": "v1",
		"k2": "v2",
	}
	_, err = MapStringString(v)
	assert.Nil(t, err)

	// Test failed cases
	v1 := map[string]interface{}{
		"k1": 1,
		"k2": 2,
	}
	_, err = MapStringString(v1)
	assert.NotNil(t, err)
}

func TestSliceBool(t *testing.T) {
	var err error

	// Test correct cases
	v := []interface{}{
		true,
		false,
	}
	_, err = sliceBool(v)
	assert.Nil(t, err)

	// Test failed cases
	v1 := []interface{}{
		1,
		2,
	}
	_, err = sliceBool(v1)
	assert.NotNil(t, err)
}

func TestSliceFloat64(t *testing.T) {
	var err error

	// Test correct cases
	v := []interface{}{
		1.0,
		2.1,
	}
	_, err = sliceFloat64(v)
	assert.Nil(t, err)

	// Test failed cases
	v1 := []interface{}{
		1,
		2,
	}
	_, err = sliceFloat64(v1)
	assert.NotNil(t, err)
}

func TestSliceString(t *testing.T) {
	var err error

	// Test correct cases
	v := []interface{}{
		"v1",
		"v2",
	}
	_, err = sliceString(v)
	assert.Nil(t, err)

	// Test failed cases
	v1 := []interface{}{
		1,
		2,
	}
	_, err = sliceString(v1)
	assert.NotNil(t, err)
}
