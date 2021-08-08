package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"

	"src.elv.sh/pkg/eval"
	"src.elv.sh/pkg/eval/vals"
)

// IntPath returns the path of the interpreter as a string.
func IntPath(fm *eval.Frame) error {
	out := fm.ValueOutput()
	ex, err := os.Executable()
	if err != nil {
		return err
	}
	out.Put(ex)
	return nil
}

// ScriptPath returns the path of the script as a string.
func ScriptPath(fm *eval.Frame) error {
	out := fm.ValueOutput()
	if len(os.Args) > 1 {
		out.Put(os.Args[1])
		return nil
	}
	return errors.New("script-path not called from script.")
}

/* proc:run takes a command name `cmd` and optional list of strings `args`
and produces a map with the following types/values:

  &exitcode= num
  &exited= bool
  &pid= num
  &success= bool
  &path= string
  &args= list
  &stderr= string
  &stdin= string
  &stdout= string

  exitcode: POSIX exit code.
  exited: bool of whether `cmd` has exited.
  pid: pid of `cmd`.
  success: whether `cmd` has exited successfully.
  path: filesystem path to `cmd`.
  args: list of args given to `cmd` (includes `cmd`).
  stderr: stderr of `cmd`.
  stdin: stdin of `cmd`.
  stdout: stdout of `cmd`.

proc:run also has 2 optional arguments: `&stdout` and `&stderr`.
Both can take a string in the form of a path to use as a file to output to. */

type cmdRunOpts struct {
	Stderr string
	Stdout string
}

func (opts *cmdRunOpts) SetDefaultOptions() {}

func cmdRun(fm *eval.Frame, opts cmdRunOpts, cmd string, args ...string) error {
	out := fm.ValueOutput()
	var stdout, errs strings.Builder
	ins := new(bytes.Buffer)
	infile := fm.InputFile()
	if infile.Name() != "/dev/stdin" {
		ins.ReadFrom(infile)
	}
	proc := exec.Command(cmd, args...)
	proc.Stdin = infile
	if opts.Stderr == "" {
		proc.Stderr = &errs
	} else {
		f, _ := os.Create(opts.Stderr)
		proc.Stderr = f
	}
	if opts.Stdout == "" {
		proc.Stdout = &stdout
	} else if opts.Stdout != "" && (opts.Stdout == opts.Stderr) {
		proc.Stdout = proc.Stderr
	} else {
		f, _ := os.Create(opts.Stdout)
		proc.Stdout = f
	}
	err := proc.Start()
	if err != nil {
		return errors.New("process failed to start.")
	}
	proc.Wait()
	if err != nil {
		return errors.New("process failed to wait.")
	}
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
	m = m.Assoc("stdout", stdout.String())
	out.Put(m)
	return nil
}

var Ns = eval.NsBuilder{}.AddGoFns("proc:", map[string]interface{}{
	"run":         cmdRun,
	"interp-path": IntPath,
	"script-path": ScriptPath,
}).Ns()
