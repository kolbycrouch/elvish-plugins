package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
        "io"
	"src.elv.sh/pkg/eval"
	"github.com/klauspost/compress/zstd"
)

func compressGeneric(alg string, input interface{}) io.Writer {
  val := func() io.Reader {
    switch input := input.(type) {
      case string:
        return bufio.NewReader(strings.NewReader(string(input)))
      case *os.File:
        return bufio.NewReader(input)
    }
    return strings.NewReader("")
  }()
  var b *bytes.Buffer
  enc := func() io.Writer {
    switch alg {
      case "zstd":
        e, _ := zstd.NewWriter(b)
        return e
    }
    e, _ := zstd.NewWriter(b)
    return e
  }()
  _, _ = io.Copy(enc, val)
  return enc
}

func zstdWrap(input interface{}) io.Writer {
  return compressGeneric("zstd", input)
}

func zstdCompress(fm *eval.Frame, input interface{}) {
  val := func() io.Reader {
    switch input := input.(type) {
      case string:
        return bufio.NewReader(strings.NewReader(string(input)))
      case *os.File:
        return bufio.NewReader(input)
    }
    return strings.NewReader("")
  }()
  buf := new(bytes.Buffer)
  enc, _ := zstd.NewWriter(buf)  
  _, _ = io.Copy(enc, val)
  enc.Close()
  out := fm.OutputFile()
  out.Write(buf.Bytes())
}

var Ns = eval.NsBuilder{
}.AddGoFns("compress:", map[string]interface{}{
  "zstd" : zstdCompress,
}).Ns()
