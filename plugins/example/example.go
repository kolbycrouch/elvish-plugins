package main

import (
	"src.elv.sh/pkg/eval"
)

func greeting() string {
  return "Greetings from plugin!"
}

var Ns = eval.NsBuilder{}.AddGoFns("example:", map[string]interface{}{
	"greeting": greeting,
}).Ns()
