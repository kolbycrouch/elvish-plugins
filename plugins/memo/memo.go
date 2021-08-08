package main

import (
	"fmt"
	"reflect"

	"github.com/dgraph-io/ristretto"
	"src.elv.sh/pkg/eval"
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
	// TODO: Possibly change to something more sophisticated than concating ...interface{}.
	// this is possibly very inefficient.
	var argstring string
	for _, i := range args {
		argstring = argstring + fmt.Sprintf("%v", i)
	}
	ns := fmt.Sprintf("%p", clos.FieldByName("Captured").
		Interface().(*eval.Ns))

	val, found := memoCache.Get(ns + argstring)
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
		memoCache.Set(ns+argstring, caps, 1)
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
