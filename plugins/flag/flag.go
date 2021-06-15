package main

import (
	"flag"
	"os"
	"src.elv.sh/pkg/eval"
)

func parseWrap() {
  flag.CommandLine.Parse(os.Args[2:])
}

var Ns = eval.NsBuilder{}.AddGoFns("flag:", map[string]interface{}{
	"args": flag.Args,
	"string": flag.String,
	"int": flag.Int,
	"bool": flag.Bool,
	"parse": parseWrap,
	"parsed": flag.Parsed,
}).Ns()
