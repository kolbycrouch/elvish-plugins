package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
        "io"
        "encoding/hex"
        "crypto/sha256"
	"src.elv.sh/pkg/eval"
	"src.elv.sh/pkg/eval/vals"
	"src.elv.sh/pkg/strutil"
)

// IntPath returns the path of the interpreter as a string.
func IntPath(fm *eval.Frame) string {
   ex, _ := os.Executable()
   return ex
}

// ScriptPath returns the path of the script as a string.
func ScriptPath(fm *eval.Frame) string {
   if len(os.Args) > 1 {
      return os.Args[1]
   } else {
      return "Can only use if called from script"
   }
}

// cmdRun takes a command name `cmd` and optional list of strings `args`
// and produces a map with the following types/values:
//
//   &exitcode= num
//   &exited= bool
//   &pid= num
//   &success= bool
//   &path= string
//   &args= list
//   &stderr= string
//   &stdin= string
//   &stdout= string
//
//   exitcode: POSIX exit code.
//   exited: bool of whether `cmd` has exited.
//   pid: pid of `cmd`.
//   success: whether `cmd` has exited successfully.
//   path: filesystem path to `cmd`.
//   args: list of args given to `cmd` (includes `cmd`).
//   stderr: stderr of `cmd`.
//   stdin: stdin of `cmd`.
//   stdout: stdout of `cmd`.
//
func cmdRun(fm *eval.Frame, cmd string, args ...string) vals.Map {
  var outs, errs strings.Builder
  ins := new(bytes.Buffer)
  infile := fm.InputFile()
  if (infile.Name() != "/dev/stdin") {
    ins.ReadFrom(infile)
  }
  proc := exec.Command(cmd, args...)
  proc.Stdin = infile
  proc.Stdout = &outs
  proc.Stderr = &errs
  proc.Start()
  proc.Wait()
  m := vals.EmptyMap
  m = m.Assoc("exitcode", int(proc.ProcessState.ExitCode()))
  m = m.Assoc("exited", proc.ProcessState.Exited())
  m = m.Assoc("pid", int(proc.ProcessState.Pid()))
  m = m.Assoc("success", proc.ProcessState.Success())
  m = m.Assoc("path", proc.Path)
  largs := vals.EmptyList
  for _, s := range proc.Args {
    largs = largs.Cons(s)
  }
  m = m.Assoc("args", largs)
  m = m.Assoc("stderr", errs.String())
  m = m.Assoc("stdin", ins.String())
  m = m.Assoc("stdout", outs.String())
  return m
}

var Ns = eval.NsBuilder{
}.AddGoFns("proc:", map[string]interface{}{
  "run" : cmdRun,
  "interp-path" : IntPath,
  "script-path" : ScriptPath,
}).Ns()
