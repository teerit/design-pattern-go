package main

import (
	"reflect"
	"testing"
)

func TestSayHello(t *testing.T) {
	p := NewPerson("Test", 101)
	p.SayHello()
	_ = reflect.ValueOf(p)
}
