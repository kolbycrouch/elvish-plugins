package main

import (
	"src.elv.sh/pkg/eval"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/glib"
)

func typeString(fm *eval.Frame) glib.Type {
  return glib.TYPE_STRING
}

func listData(fm *eval.Frame, lst *glib.List) interface{} {
  return lst.Data()
}

func listLength(fm *eval.Frame, lst *glib.List) uint {
  return lst.Length()
}

func listFirst(fm *eval.Frame, lst *glib.List) *glib.List {
  return lst.First()
}

func listNext(fm *eval.Frame, lst *glib.List) *glib.List {
  return lst.Next()
}

func listDataToTreePath(fm *eval.Frame, lst *glib.List) *gtk.TreePath {
  return lst.Data().(*gtk.TreePath)
}

func valueGetString(fm *eval.Frame, val *glib.Value) {
	out := fm.OutputChan()

	s, _ := val.GetString()
	out <- s
}

var Ns = eval.NsBuilder{
}.AddGoFns("glib:", map[string]interface{}{
  "list-data" : listData,
  "list-data-to-treepath" : listDataToTreePath,
  "list-first" : listFirst,
  "list-length" : listLength,
  "list-next" : listNext,
  "type-string" : typeString,
  "value-get-string" : valueGetString,
}).Ns()
