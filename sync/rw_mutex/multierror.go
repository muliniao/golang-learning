package rw_mutex

import (
	"strings"
	"sync"
)

type MultiError struct {
	lock *sync.RWMutex
	errs []error
}

func New() *MultiError {
	return &MultiError{
		lock: &sync.RWMutex{},
		errs: make([]error, 0),
	}
}

func (m *MultiError) Push(err error) {
	if err == nil {
		return
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	m.errs = append(m.errs, err)
}

func (m *MultiError) Errors() []error {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.errs
}

func (m *MultiError) HasError() error {
	if len(m.errs) == 0 {
		return nil
	}
	return m
}

func (m *MultiError) Error() string {
	formattedError := make([]string, len(m.errs))
	for i, e := range m.errs {
		formattedError[i] = e.Error()
	}
	return strings.Join(formattedError, ", ")
}
