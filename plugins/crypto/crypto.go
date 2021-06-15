package main

import (
	"bufio"
	"os"
	"strings"
        "io"
	"hash"
        "encoding/hex"
        "crypto/md5"
        "crypto/sha1"
        "crypto/sha256"
        "crypto/sha512"
	"src.elv.sh/pkg/eval"
)

func hashGeneric(alg string, input interface{}) string {
  val := func() io.Reader { 
    switch input := input.(type) {
      case string:
        return bufio.NewReader(strings.NewReader(string(input)))
      case *os.File:
        return bufio.NewReader(input)
    }
    return strings.NewReader("")
  }()
  crypt := func() hash.Hash {
    switch alg {
      case "md5":
        return md5.New()
      case "sha1":
        return sha1.New()
      case "sha256":
        return sha256.New()
      case "sha384":
        return sha512.New384()
      case "sha512":
        return sha512.New()
    }
    return sha256.New()
  }()
  io.Copy(crypt, val)
  return hex.EncodeToString(crypt.Sum(nil))
}

func md5Wrap(input interface{}) string {
  return hashGeneric("md5", input)
}

func sha1Wrap(input interface{}) string {
  return hashGeneric("sha1", input)
}

func sha256Wrap(input interface{}) string {
  return hashGeneric("sha256", input)
}

func sha512Wrap(input interface{}) string {
  return hashGeneric("sha512", input)
}

var Ns = eval.NsBuilder{
}.AddGoFns("crypto:", map[string]interface{}{
  "md5" : md5Wrap,
  "sha1" : sha1Wrap,
  "sha256" : sha256Wrap,
  "sha512" : sha512Wrap,
}).Ns()
