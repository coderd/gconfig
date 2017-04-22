package gconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
)

const (
	ErrKeyNotFound        = "gconfig: key '%s' is not found"
	ErrNotBool            = "gconfig: key '%s' is not a bool"
	ErrNotFloat64         = "gconfig: key '%s' is not a float64"
	ErrNotString          = "gconfig: key '%s' is not a string"
	ErrNotMapStringString = "gconfig: key '%s' is not a map[string]string"
)

type ConfigFile struct {
	file string
	data map[string]interface{}
}

func LoadJsonFile(file string) (*ConfigFile, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Erase comments
	re := regexp.MustCompile("//.*?\n")
	content = re.ReplaceAll(content, []byte("\n"))

	var data map[string]interface{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	configFile := &ConfigFile{
		file: file,
		data: data,
	}

	return configFile, nil
}

func (c *ConfigFile) Get(key string) (interface{}, error) {
	if v, ok := c.data[key]; ok {
		return v, nil
	}

	return nil, fmt.Errorf(ErrKeyNotFound, key)
}

func (c *ConfigFile) Bool(key string) (bool, error) {
	v, err := c.Get(key)
	if err != nil {
		return false, err
	}

	if value, ok := v.(bool); ok {
		return value, nil
	} else {
		return false, fmt.Errorf(ErrNotBool, key)
	}
}

func (c *ConfigFile) Float64(key string) (float64, error) {
	v, err := c.Get(key)
	if err != nil {
		return 0.0, err
	}

	if value, ok := v.(float64); ok {
		return value, nil
	} else {
		return 0.0, fmt.Errorf(ErrNotFloat64, key)
	}
}

func (c *ConfigFile) String(key string) (string, error) {
	v, err := c.Get(key)
	if err != nil {
		return "", err
	}

	if value, ok := v.(string); ok {
		return value, nil
	} else {
		return "", fmt.Errorf(ErrNotString, key)
	}
}

func (c *ConfigFile) MapStringBool(key string) (map[string]bool, error) {
	v, err := c.Get(key)
	if err != nil {
		return nil, err
	}

	if value, ok := v.(map[string]interface{}); ok {
		return MapStringBool(value)
	} else {
		return nil, fmt.Errorf(ErrNotMapStringString, key)
	}
}

func (c *ConfigFile) MapStringFloat64(key string) (map[string]float64, error) {
	v, err := c.Get(key)
	if err != nil {
		return nil, err
	}

	if value, ok := v.(map[string]interface{}); ok {
		return MapStringFloat64(value)
	} else {
		return nil, fmt.Errorf(ErrNotMapStringString, key)
	}
}

func (c *ConfigFile) MapStringString(key string) (map[string]string, error) {
	v, err := c.Get(key)
	if err != nil {
		return nil, err
	}

	if value, ok := v.(map[string]interface{}); ok {
		return MapStringString(value)
	} else {
		return nil, fmt.Errorf(ErrNotMapStringString, key)
	}
}

func (c *ConfigFile) MustBool(key string) bool {
	v, err := c.Bool(key)
	if err != nil {
		panic(err)
	}

	return v
}

func (c *ConfigFile) MustFloat64(key string) float64 {
	v, err := c.Float64(key)
	if err != nil {
		panic(err)
	}

	return v
}

func (c *ConfigFile) MustString(key string) string {
	v, err := c.String(key)
	if err != nil {
		panic(err)
	}

	return v
}

func (c *ConfigFile) MustMapStringBool(key string) map[string]bool {
	v, err := c.MapStringBool(key)
	if err != nil {
		panic(err)
	}

	return v
}

func (c *ConfigFile) MustMapStringFloat64(key string) map[string]float64 {
	v, err := c.MapStringFloat64(key)
	if err != nil {
		panic(err)
	}

	return v
}

func (c *ConfigFile) MustMapStringString(key string) map[string]string {
	v, err := c.MapStringString(key)
	if err != nil {
		panic(err)
	}

	return v
}

func (c *ConfigFile) AlwaysBool(key string, defaultVal ...bool) bool {
	v, err := c.Bool(key)
	if err != nil && len(defaultVal) > 0 {
		return defaultVal[0]
	}

	return v
}

func (c *ConfigFile) AlwaysFloat64(key string, defaultVal ...float64) float64 {
	v, err := c.Float64(key)
	if err != nil && len(defaultVal) > 0 {
		return defaultVal[0]
	}

	return v
}

func (c *ConfigFile) AlwaysString(key string, defaultVal ...string) string {
	v, err := c.String(key)
	if err != nil && len(defaultVal) > 0 {
		return defaultVal[0]
	}

	return v
}

func (c *ConfigFile) AlwaysMapStringBool(key string, defaultVal ...map[string]bool) map[string]bool {
	v, err := c.MapStringBool(key)
	if err != nil && len(defaultVal) > 0 {
		return defaultVal[0]
	}

	return v
}

func (c *ConfigFile) AlwaysMapStringFloat64(key string, defaultVal ...map[string]float64) map[string]float64 {
	v, err := c.MapStringFloat64(key)
	if err != nil && len(defaultVal) > 0 {
		return defaultVal[0]
	}

	return v
}

func (c *ConfigFile) AlwaysMapStringString(key string, defaultVal ...map[string]string) map[string]string {
	v, err := c.MapStringString(key)
	if err != nil && len(defaultVal) > 0 {
		return defaultVal[0]
	}

	return v
}

// An InvalidPointerError describes an invalid argument passed to Must and Always.
type InvalidPointerError struct {
	Type reflect.Type
}

func (i *InvalidPointerError) Error() string {
	if i.Type == nil {
		return "gconfig: set (nil)"
	}

	if i.Type.Kind() != reflect.Ptr {
		return "gconfig: set(non-pointer " + i.Type.String() + ")"
	}
	return "gconfig: set(nil " + i.Type.String() + ")"
}

func set(pointer interface{}, v interface{}) error {
	pointerReflectV := reflect.ValueOf(pointer)
	pointerReflectT := reflect.TypeOf(pointer)
	if pointerReflectV.Kind() != reflect.Ptr || pointerReflectV.IsNil() {
		return &InvalidPointerError{pointerReflectT}
	}

	vReflectV := reflect.ValueOf(v)
	vReflectT := reflect.TypeOf(v)

	pointerElem := pointerReflectV.Elem()
	pointerElemT := pointerElem.Type()

	if pointerElemT != vReflectT {
		return fmt.Errorf("gconfig: set() pointer's element is a `%s`, not a `%s`", vReflectT.String(), pointerElemT.String())
	}

	pointerElem.Set(vReflectV)

	return nil
}

func (c *ConfigFile) Must(key string, value interface{}) {
	v, err := c.Get(key)
	if err != nil {
		panic(err)
	}

	if err = set(value, v); err != nil {
		panic(err)
	}
}
