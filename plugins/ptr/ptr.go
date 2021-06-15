package main

import (
	"reflect"
	"src.elv.sh/pkg/eval"
)

func deref(name interface{}) interface{} {
  return reflect.ValueOf(name).Elem().Interface()
}

var Ns = eval.NsBuilder{}.AddGoFns("ptr:", map[string]interface{}{
	"deref": deref,
}).Ns()
