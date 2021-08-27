package main

import (
	httpmod "github.com/kolbycrouch/elvish-plugins/plugins/net/http"
	"src.elv.sh/pkg/eval"
)

var Ns = eval.NsBuilder{}.AddNs("http", httpmod.Ns).Ns()
