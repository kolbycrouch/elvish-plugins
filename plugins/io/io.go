package main

import (
	"bufio"
	"os"
	"src.elv.sh/pkg/eval"
)

func Scan(fm *eval.Frame) string {
   scanner := bufio.NewScanner(os.Stdin)
   scanner.Scan()
   text := scanner.Text()
   return text
}

var Ns = eval.NsBuilder{
}.AddGoFns("io:", map[string]interface{}{
  "scan" : Scan,
}).Ns()
