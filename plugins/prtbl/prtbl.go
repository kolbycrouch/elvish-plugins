package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
        "io"
        "encoding/hex"
        "crypto/sha256"
	"src.elv.sh/pkg/eval"
	"src.elv.sh/pkg/eval/vals"
//	"src.elv.sh/pkg/eval/vars"
  "src.elv.sh/pkg/strutil"
)

func linesToChan(r io.Reader, ch chan<- interface{}) {
        filein := bufio.NewReader(r)
        for {
                line, err := filein.ReadString('\n')
                if line != "" {
                        ch <- strutil.ChopLineEnding(line)
                }
                if err != nil {
                        if err != io.EOF {
                                fmt.Println("error on reading:", err)
                        }
                        break
                }
        }
}

func Scan(fm *eval.Frame) string {
   scanner := bufio.NewScanner(os.Stdin)
   scanner.Scan()
   text := scanner.Text()
   return text
}

func IntPath(fm *eval.Frame) string {
   ex, _ := os.Executable()
   return ex
}

func ScriptPath(fm *eval.Frame) string {
   if len(os.Args) > 1 {
      return os.Args[1]
   } else {
      return "Can only use if called from script"
   }
}

func sha256FromFile(fm *eval.Frame, input interface{}) {
  out := fm.OutputChan()
  val := func() io.Reader { 
    switch input := input.(type) {
      case string:
        return bufio.NewReader(strings.NewReader(string(input)))
      case *os.File:
        return bufio.NewReader(input)
    }
    return strings.NewReader("")
  }()

  crypt := sha256.New()

  io.Copy(crypt, val)

  hx := hex.EncodeToString(crypt.Sum(nil))
  end := strings.NewReader(hx)
  linesToChan(end, out)

}

func execWithPid(fm *eval.Frame, cmd string, arg string) vals.Map {
  proc := exec.Command(cmd, arg)
  proc.Start()
  proc.Wait()
  m := vals.EmptyMap
  m = m.Assoc("exitcode", float64(proc.ProcessState.ExitCode()))
  m = m.Assoc("exited", proc.ProcessState.Exited())
  m = m.Assoc("pid", float64(proc.ProcessState.Pid()))
  m = m.Assoc("success", proc.ProcessState.Success())
  m = m.Assoc("path", proc.Path)
  m = m.Assoc("args", proc.Args)
  sout := fmt.Sprint(proc.Stdout)
  m = m.Assoc("stdout", sout)
  return m
}

func toRealBytes(fm *eval.Frame, inputs eval.Inputs) {
  out := fm.InputChan()
  var b []byte
  var c string
  inputs(func(v interface{}) {
    a, _ := v.([]byte)
    b = []byte(a)
    c = string(b)
    out <- c
  })
}

var Ns = eval.NsBuilder{
}.AddGoFns("prtbl:", map[string]interface{}{
  "exec-pid" : execWithPid,
  "interp-path" : IntPath,
  "script-path" : ScriptPath,
  "scan" : Scan,
  "sha256-file" : sha256FromFile,
  "to-real-bytes" : toRealBytes,
}).Ns()
