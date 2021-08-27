package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"strings"

	"src.elv.sh/pkg/eval"
)

func hashGeneric(alg string, input string) string {
	reader := strings.NewReader(input)
	crypt := func() hash.Hash {
		switch alg {
		case "md5":
			return md5.New()
		case "sha1":
			return sha1.New()
		case "sha256":
			return sha256.New()
		case "sha512":
			return sha512.New()
		}
		return sha256.New()
	}()
	io.Copy(crypt, reader)
	return hex.EncodeToString(crypt.Sum(nil))
}

func md5Wrap(input string) string {
	return hashGeneric("md5", input)
}

func sha1Wrap(input string) string {
	return hashGeneric("sha1", input)
}

func sha256Wrap(input string) string {
	return hashGeneric("sha256", input)
}

func sha512Wrap(input string) string {
	return hashGeneric("sha512", input)
}

var Ns = eval.NsBuilder{}.AddGoFns("crypto:", map[string]interface{}{
	"md5":    md5Wrap,
	"sha1":   sha1Wrap,
	"sha256": sha256Wrap,
	"sha512": sha512Wrap,
}).Ns()
