package main

import (
	"src.elv.sh/pkg/eval"
	"reflect"
)

func fieldsStruct(fm *eval.Frame, val interface{}) {
  out := fm.OutputChan()
  new := reflect.ValueOf(val).Elem()
  for i:=0;i < new.NumField();i++ {
    out <- new.Type().Field(i).Name
  }
}

var Ns = eval.NsBuilder{
}.AddGoFns("reflect:", map[string]interface{}{
  "struct-fields" : fieldsStruct,
}).Ns()