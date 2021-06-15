package main

import (
	"bytes"
	"os"
	"src.elv.sh/pkg/eval"
	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

func newMod(store *wasmer.Store, file *os.File) *wasmer.Module {
  buf := new(bytes.Buffer)
  buf.ReadFrom(file)
  module, _ := wasmer.NewModule(store, buf.Bytes())
  return module
}

func newInstanceWrap(mod *wasmer.Module, obj *wasmer.ImportObject) *wasmer.Instance {
  inst, _ := wasmer.NewInstance(mod, obj)
  return inst
}

func getFuncWrap(inst *wasmer.Instance, name string) wasmer.NativeFunction {
  fun, _ := inst.Exports.GetFunction(name)
  return fun
}

func callWrap(name wasmer.NativeFunction, args ...interface{}) interface{} {
  res, _ := name(args)
  return res
}

var Ns = eval.NsBuilder{}.AddGoFns("wasm:", map[string]interface{}{
	"call": callWrap,
	"get-function": getFuncWrap,
	"new-engine": wasmer.NewEngine,
	"new-module": newMod,
	"new-store": wasmer.NewStore,
	"new-import-object": wasmer.NewImportObject,
	"new-instance": newInstanceWrap,
}).Ns()
