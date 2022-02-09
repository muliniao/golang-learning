package main

import (
	"errors"
	"fmt"
	"reflect"
	"sync/atomic"
)

type atomicValue struct {
	v atomic.Value
	t reflect.Type
}

func NewAtomicValue(example interface{}) (*atomicValue, error) {
	if example == nil {
		return nil, errors.New("atomic value: nil example")
	}
	return &atomicValue{
		t: reflect.TypeOf(example),
	}, nil
}

func (av *atomicValue) Store(v interface{}) error {
	if v == nil {
		return errors.New("atomic value: nil value")
	}
	t := reflect.TypeOf(v)
	if t != av.t {
		return fmt.Errorf("atomic value: wrong type: %s", t)
	}
	av.v.Store(v)
	return nil
}

func (av *atomicValue) Load() interface{} {
	return av.v.Load()
}

func (av *atomicValue) CompareAndSwap(old, new interface{}) bool {
	return av.v.CompareAndSwap(old, new)
}

func (av *atomicValue) Swap(new interface{}) interface{} {
	return av.v.Swap(new)
}

func (av *atomicValue) TypeOfValue() reflect.Type {
	return av.t
}
