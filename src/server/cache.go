package main

import "errors"

type Cache struct {
	C map[string]interface{}
}

func NewCache() *Cache {
	cache := new(Cache)
	cache.C = make(map[string]interface{})
	return cache
}

func (c *Cache) SetValue(key string, value interface{}) {
	c.C[key] = value
}

func (c *Cache) GetValue(key string) (interface{}, error) {
	if value, ok := c.C[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("No Entry")
	}
}

func (c *Cache) GetStringValue(key string) (string, error) {
	if value, ok := c.C[key]; ok {
		switch str := value.(type) {
		case string:
			return str, nil
		}
		return "", errors.New("Not a string value")
	} else {
		return "", errors.New("No Entry")
	}
}

func (c *Cache) GetIntValue(key string) (int64, error) {
	if value, ok := c.C[key]; ok {
		switch i := value.(type) {
		case int:
			return int64(i), nil
		case int64:
			return i, nil
		case int32:
			return int64(i), nil
		}
		return 0, errors.New("Not a int value")
	} else {
		return 0, errors.New("No Entry")
	}
}
