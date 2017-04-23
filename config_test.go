package gconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testFile = "testdata/config.json"

func TestLoadJsonFile(t *testing.T) {
	var err error

	configFile, err := LoadJsonFile(testFile)
	assert.Nil(t, err)

	var (
		b   bool
		i   float64
		f   float64
		s   string
		msb map[string]bool
		msf map[string]float64
		mss map[string]string
		sb  []bool
		sf  []float64
		ss  []string
	)

	// Test correct cases
	// ========================================================================

	// Test type-getting
	b, err = configFile.Bool("bool")
	assert.Nil(t, err)
	assert.True(t, b)

	i, err = configFile.Float64("int")
	assert.Nil(t, err)
	assert.Equal(t, 1.0, i)

	f, err = configFile.Float64("float64")
	assert.Nil(t, err)
	assert.Equal(t, 1.01, f)

	s, err = configFile.String("string")
	assert.Nil(t, err)
	assert.Equal(t, "foo", s)

	_, err = configFile.MapStringBool("map_string_bool")
	assert.Nil(t, err)

	_, err = configFile.MapStringFloat64("map_string_float64")
	assert.Nil(t, err)

	_, err = configFile.MapStringString("map_string_string")
	assert.Nil(t, err)

	_, err = configFile.SliceBool("slice_bool")
	assert.Nil(t, err)

	_, err = configFile.SliceFloat64("slice_float64")
	assert.Nil(t, err)

	_, err = configFile.SliceString("slice_string")
	assert.Nil(t, err)

	// Test must-getting
	b = configFile.MustBool("bool")
	assert.True(t, b)

	i = configFile.MustFloat64("int")
	assert.Equal(t, 1.0, i)

	f = configFile.MustFloat64("float64")
	assert.Equal(t, 1.01, f)

	s = configFile.MustString("string")
	assert.Equal(t, "foo", s)

	assert.NotPanics(t, func() {
		configFile.MustMapStringBool("map_string_bool")
	})

	assert.NotPanics(t, func() {
		configFile.MustMapStringFloat64("map_string_float64")
	})

	assert.NotPanics(t, func() {
		configFile.MustMapStringString("map_string_string")
	})

	assert.NotPanics(t, func() {
		configFile.MustSliceBool("slice_bool")
	})

	assert.NotPanics(t, func() {
		configFile.MustSliceFloat64("slice_float64")
	})

	assert.NotPanics(t, func() {
		configFile.MustSliceString("slice_string")
	})

	// Test always-getting
	b = configFile.AlwaysBool("bool")
	assert.True(t, b)

	i = configFile.AlwaysFloat64("int")
	assert.Equal(t, 1.0, i)

	f = configFile.AlwaysFloat64("float64")
	assert.Equal(t, 1.01, f)

	s = configFile.AlwaysString("string")
	assert.Equal(t, "foo", s)

	msb = configFile.AlwaysMapStringBool("map_string_bool")
	assert.True(t, len(msb) > 0)

	msf = configFile.AlwaysMapStringFloat64("map_string_float64")
	assert.True(t, len(msf) > 0)

	mss = configFile.AlwaysMapStringString("map_string_string")
	assert.True(t, len(mss) > 0)

	sb = configFile.AlwaysSliceBool("slice_bool")
	assert.True(t, len(sb) > 0)

	sf = configFile.AlwaysSliceFloat64("slice_float64")
	assert.True(t, len(sf) > 0)

	ss = configFile.AlwaysSliceString("slice_string")
	assert.True(t, len(ss) > 0)

	// Test others
	f = 0.0
	configFile.Must("float64", &f)
	assert.Equal(t, 1.01, f)

	// Test failed cases
	// ========================================================================

	// Test type-getting
	b, err = configFile.Bool("int")
	assert.NotNil(t, err)
	assert.False(t, b)

	i, err = configFile.Float64("bool")
	assert.NotNil(t, err)
	assert.Equal(t, 0.0, i)

	f, err = configFile.Float64("string")
	assert.NotNil(t, err)
	assert.Equal(t, 0.0, f)

	s, err = configFile.String("float64")
	assert.NotNil(t, err)
	assert.Equal(t, "", s)

	_, err = configFile.MapStringBool("map_string_float64")
	assert.NotNil(t, err)

	_, err = configFile.MapStringFloat64("map_string_bool")
	assert.NotNil(t, err)

	_, err = configFile.MapStringString("map_string_float64")
	assert.NotNil(t, err)

	_, err = configFile.SliceBool("slice_float64")
	assert.NotNil(t, err)

	_, err = configFile.SliceFloat64("slice_bool")
	assert.NotNil(t, err)

	_, err = configFile.SliceString("slice_float64")
	assert.NotNil(t, err)

	// Test must-getting
	assert.Panics(t, func() {
		configFile.MustBool("int")
	})

	assert.Panics(t, func() {
		configFile.MustFloat64("bool")
	})

	assert.Panics(t, func() {
		configFile.MustFloat64("string")
	})

	assert.Panics(t, func() {
		configFile.MustString("float64")
	})

	assert.Panics(t, func() {
		configFile.MustMapStringBool("map_string_float64")
	})

	assert.Panics(t, func() {
		configFile.MustMapStringFloat64("map_string_bool")
	})

	assert.Panics(t, func() {
		configFile.MustMapStringString("map_string_float64")
	})

	assert.Panics(t, func() {
		configFile.MustSliceBool("slice_float64")
	})

	assert.Panics(t, func() {
		configFile.MustSliceFloat64("slice_bool")
	})

	assert.Panics(t, func() {
		configFile.MustSliceString("slice_float64")
	})

	// Test always-getting without default values
	b = configFile.AlwaysBool("int")
	assert.False(t, b)

	i = configFile.AlwaysFloat64("bool")
	assert.Equal(t, 0.0, i)

	f = configFile.AlwaysFloat64("string")
	assert.Equal(t, 0.0, f)

	s = configFile.AlwaysString("float64")
	assert.Equal(t, "", s)

	msb = configFile.AlwaysMapStringBool("map_string_float64")
	assert.True(t, len(msb) == 0)

	msf = configFile.AlwaysMapStringFloat64("map_string_bool")
	assert.True(t, len(msf) == 0)

	mss = configFile.AlwaysMapStringString("map_string_float64")
	assert.True(t, len(mss) == 0)

	sb = configFile.AlwaysSliceBool("slice_float64")
	assert.True(t, len(sb) == 0)

	sf = configFile.AlwaysSliceFloat64("slice_bool")
	assert.True(t, len(sf) == 0)

	ss = configFile.AlwaysSliceString("slice_float64")
	assert.True(t, len(ss) == 0)

	// Test always-getting with default values
	b = configFile.AlwaysBool("int", true)
	assert.True(t, b)

	i = configFile.AlwaysFloat64("bool", 1.0)
	assert.Equal(t, 1.0, i)

	f = configFile.AlwaysFloat64("string", 1.1)
	assert.Equal(t, 1.1, f)

	s = configFile.AlwaysString("float64", "bar")
	assert.Equal(t, "bar", s)

	msb = configFile.AlwaysMapStringBool("map_string_float64", map[string]bool{"k1": true})
	assert.True(t, len(msb) > 0)

	msf = configFile.AlwaysMapStringFloat64("map_string_bool", map[string]float64{"k1": 1.0})
	assert.True(t, len(msf) > 0)

	mss = configFile.AlwaysMapStringString("map_string_float64", map[string]string{"k1": "v1"})
	assert.True(t, len(mss) > 0)

	sb = configFile.AlwaysSliceBool("slice_float64", []bool{true})
	assert.True(t, len(sb) > 0)

	sf = configFile.AlwaysSliceFloat64("slice_bool", []float64{1.0})
	assert.True(t, len(sf) > 0)

	ss = configFile.AlwaysSliceString("slice_float64", []string{"v1"})
	assert.True(t, len(ss) > 0)

	// Test others
	f = 0.0
	assert.Panics(t, func() {
		configFile.Must("string", &f)
	})

	assert.Panics(t, func() {
		f = 0.0
		configFile.Always("string", f)
	})

	assert.Panics(t, func() {
		f = 0.0
		var fp *float64 = &f
		fp = nil
		configFile.Always("string", fp)
	})

	assert.Panics(t, func() {
		f = 0.0
		configFile.Always("string", &f, 1)
	})

	f = 0.1
	configFile.Always("string", &f)
	assert.Equal(t, 0.0, f)

	f = 0.1
	configFile.Always("string", &f, 0.2)
	assert.Equal(t, 0.2, f)

}
