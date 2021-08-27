package main

import (
	"fmt"
	"reflect"

	"github.com/dgraph-io/ristretto"
	"src.elv.sh/pkg/diag"
	"src.elv.sh/pkg/eval"
	"src.elv.sh/pkg/eval/errs"
	"src.elv.sh/pkg/parse"
)

var memoCache = createCache()

func createCache() *ristretto.Cache {
	rcache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}
	return rcache
}

func memoize(fm *eval.Frame, fn eval.Callable, args ...interface{}) error {
	clos := reflect.ValueOf(fn).Elem()
	out := fm.ValueOutput()
	var argstring string
	for _, i := range args {
		argstring = argstring + fmt.Sprintf("%v", i)
	}
	if !clos.FieldByName("Op").IsValid() {
		return errs.BadValue{
			What:  `Argument 0 of "memoize"`,
			Valid: "callable", Actual: "builtin",
		}
	}
	rang := clos.FieldByName("Op").Interface().(diag.Ranger).Range()
	code := fmt.Sprintf("%v", clos.FieldByName("SrcMeta").
		Interface().(parse.Source).Code[rang.From:rang.To])

	val, found := memoCache.Get(code + argstring)
	// We didn't find the memo. Call the function and send its return over `out`.
	if !found {
		// Wrap our function for use by fm.CaptureOutput().
		fnc := func(nfm *eval.Frame) error {
			err := fn.Call(nfm, args, eval.NoOpts)
			return err
		}
		caps, err := fm.CaptureOutput(fnc)
		if err != nil {
			return err
		}
		memoCache.Set(code+argstring, caps, 1)
		memoCache.Wait()
		for _, i := range caps {
			out.Put(i)
		}
		// We did find the memo. Send the values over `out`.
	} else {
		for _, i := range val.([]interface{}) {
			out.Put(i)
		}
	}
	return nil
}

var Ns = eval.NsBuilder{}.AddGoFns("memo:", map[string]interface{}{
	"memoize": memoize,
}).Ns()
