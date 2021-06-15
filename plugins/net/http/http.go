package http

import (
	"context"
	"net/http"

	"src.elv.sh/pkg/eval"
)

func fserver(path string) http.Handler {
  return http.FileServer(http.Dir(path))
}

func listenserve(fm *eval.Frame, port string) {
  var srv http.Server
  srv.Addr = port

  go func() {
    done := fm.Interrupts()
    <-done
    srv.Shutdown(context.Background())
  }()
  srv.ListenAndServe()
}

var Ns = eval.NsBuilder{}.AddGoFns("net:http:", map[string]interface{}{
	"file-server": fserver,
	"handle": http.Handle,
	"listen-serve": listenserve,
}).Ns()
