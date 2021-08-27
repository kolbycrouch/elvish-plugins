package main

import (
	"src.elv.sh/pkg/eval"
	"src.elv.sh/pkg/eval/errs"
	"src.elv.sh/pkg/eval/vals"
)

func mapFn(fm *eval.Frame, fn eval.Callable, cont interface{}) error {
    out := fm.ValueOutput()
    switch t := cont.(type) {
        case vals.List:
            newcont := vals.EmptyList
            for iter := t.Iterator(); iter.HasElem(); iter.Next() {
                elem :=  iter.Elem()
                slice := make([]interface{}, 1)
                slice[0] = elem
                fnc := func(nfm *eval.Frame) error {
			              err := fn.Call(nfm, slice, eval.NoOpts)
			              return err
		            }
		            caps, err := fm.CaptureOutput(fnc)
		            if err != nil {
			              return err
		            }
		            if len(caps) > 0 {
		                newcont = newcont.Cons(caps[0])
                }
            }
            out.Put(newcont)
            return nil
        case vals.Map:
            newcont := vals.EmptyMap
            for iter := t.Iterator(); iter.HasElem(); iter.Next() {
                key, val :=  iter.Elem()
                slice := make([]interface{}, 1)
                slice[0] = val
                fnc := func(nfm *eval.Frame) error {
			              err := fn.Call(nfm, slice, eval.NoOpts)
			              return err
		            }
		            caps, err := fm.CaptureOutput(fnc)
		            if err != nil {
			              return err
		            }
		            if len(caps) > 0 {
  	                newcont = newcont.Assoc(key, caps[0])
                }
            }
            out.Put(newcont)
            return nil
        default:
            return errs.BadValue{
                What: `argument 1 of "map"`,
                Valid: "List or Map", Actual: vals.Kind(t)}
    }
    return nil
}

var Ns = eval.NsBuilder{}.AddGoFns("fun:", map[string]interface{}{
	"map": mapFn,
}).Ns()
