package main

import (
	"src.elv.sh/pkg/eval"
)

func mpgo (fm *eval.Frame, fn eval.Callable) {
  go func() {
    var args []interface{}
    fn.Call(fm, args, eval.NoOpts)
  }()
}

var Ns = eval.NsBuilder{}.AddGoFns("mp:", map[string]interface{}{
	"go": mpgo,
}).Ns()
