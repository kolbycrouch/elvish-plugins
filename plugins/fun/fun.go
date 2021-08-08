package main

import (
	"errors"
	"reflect"
	"src.elv.sh/pkg/eval"
	"src.elv.sh/pkg/eval/errs"
	"src.elv.sh/pkg/eval/vals"
)

func mapFn(fm *eval.Frame, fn eval.Callable, cont interface{}) error {
    var newcont interface{}
    out := fm.OutputChan()
    clos := reflect.ValueOf(fn).Elem()
    err := errors.New("HI")
    switch t := cont.(type) {
        case vals.List:
            newcont = vals.EmptyList
        case vals.Map:
            newcont = vals.EmptyMap
        default:
            return errs.BadValue{
                What: `argument 1 of "map"`,
                Valid: "List or Map", Actual: vals.Kind(t)}
    }


    _ = newcont
    return err
}

var Ns = eval.NsBuilder{}.AddGoFns("fun:", map[string]interface{}{
	"map": mapFn,
}).Ns()
