package main

import (
	"src.elv.sh/pkg/eval"
	httpmod "src.elv.sh/pkg/eval/mods/net/http"
)

var Ns = eval.NsBuilder{}.AddNs("http", httpmod.Ns).Ns()
