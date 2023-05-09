package test

import (
	"fmt"
	"testing"
)

func Nil(t *testing.T, obj any) {
	t.Helper()
	isNil(obj).true(t)
}

func NotNil(t *testing.T, obj any) {
	t.Helper()
	isNil(obj).false(t)
}

func NoError(t *testing.T, err error) {
	t.Helper()
	isNil(err).true(t)
}

func Error(t *testing.T, err error) {
	t.Helper()
	isNil(err).false(t)
}

func isNil(a any) res {
	switch a := a.(type) {
	case nil:
		return res{true, "nil"}
	default:
		return res{false, fmt.Sprint(a)}
	}
}