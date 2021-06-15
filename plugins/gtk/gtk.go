package main

import (
	"fmt"
	"strconv"
	"src.elv.sh/pkg/eval"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/glib"
	"src.elv.sh/pkg/eval/vals"
)

func gtkInit(fm *eval.Frame) {
  gtk.Init(nil)
}

func windowNew(fm *eval.Frame) *gtk.Window {
  win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
  return win
}

func windowShowAll(fm *eval.Frame, win *gtk.Window) {
  win.ShowAll()
}

func gtkMain(fm *eval.Frame) {
  gtk.Main()
}

func mainQuit(fm *eval.Frame) {
  gtk.MainQuit()
}

func windowConnect(fm *eval.Frame, win *gtk.Window, s string, call eval.Callable) {
  win.Connect(s, func() {
    call.Call(fm, []interface{}{}, nil)
  })
}

func windowSetTitle(fm *eval.Frame, win *gtk.Window, title string) {
  win.SetTitle(title)
}

func labelNew(fm *eval.Frame, lbl string) gtk.IWidget {
  l, _ := gtk.LabelNew(lbl)
  return l
}

func add(fm *eval.Frame, win *gtk.Window, widg gtk.IWidget) {
  win.Add(widg)
}

func boxPackStart(fm *eval.Frame, box *gtk.Box, widg gtk.IWidget, ta, tb bool, num float64) {
  n := uint(num)
  box.PackStart(widg, ta, tb, n)
}

func windowDefaultSize(fm *eval.Frame, win *gtk.Window, width int, height int) {
  win.SetDefaultSize(width, height)
}

func textViewNew(fm *eval.Frame) gtk.IWidget {
  tv, _ := gtk.TextViewNew()
  return tv
} 

func orientationVertical(fm *eval.Frame) gtk.Orientation {
  return gtk.ORIENTATION_VERTICAL
}

func orientationHorizontal(fm *eval.Frame) gtk.Orientation {
  return gtk.ORIENTATION_HORIZONTAL
}

func boxNew(fm *eval.Frame, orient gtk.Orientation, num int) gtk.IWidget {
  box, _ := gtk.BoxNew(orient, num)
  return box
}

func treeViewNew(fm *eval.Frame) gtk.IWidget {
  tv, _ := gtk.TreeViewNew()
  return tv
}

func entryNew(fm *eval.Frame) *gtk.Entry {
  ent, _ := gtk.EntryNew()
  return ent
}

func listStoreNew(fm *eval.Frame, typ glib.Type) *gtk.ListStore {
  tv, _ := gtk.ListStoreNew(typ)
  return tv
}

func cellRendererTextNew(fm *eval.Frame) *gtk.CellRendererText {
  rend, _ := gtk.CellRendererTextNew()
  return rend
}

func treeViewColumnNewWithAttribute(fm *eval.Frame, title string, rend *gtk.CellRendererText, attr string, col int) *gtk.TreeViewColumn {
  column, _ := gtk.TreeViewColumnNewWithAttribute(title, rend, attr, col)
  return column
}

func treeViewAppendColumn(fm *eval.Frame, tv *gtk.TreeView, col *gtk.TreeViewColumn) {
  tv.AppendColumn(col)
}

func treeViewSetModel(fm *eval.Frame, tv *gtk.TreeView, lst *gtk.ListStore) {
  tv.SetModel(lst)
}

func treeViewGetSelection(fm *eval.Frame, tv *gtk.TreeView) *gtk.TreeSelection {
  sel, _ := tv.GetSelection()
  return sel
}

func selectionMultiple(fm *eval.Frame) gtk.SelectionMode {
  return gtk.SELECTION_MULTIPLE
}

func treeViewSelectionSetMode(fm *eval.Frame, sel *gtk.TreeSelection, mode gtk.SelectionMode) {
  sel.SetMode(mode)
}

func treeViewSelectionConnect(fm *eval.Frame, sel *gtk.TreeSelection, s string, call eval.Callable) {
  sel.Connect(s, func() {
    call.Call(fm, []interface{}{}, nil)
  })
}

func entrySetText(fm *eval.Frame, entry *gtk.Entry, s string) {
  entry.SetText(s)
}

func listStoreAppend(fm *eval.Frame, lst *gtk.ListStore) *gtk.TreeIter {
  return lst.Append()
}

func listStoreSet(fm *eval.Frame, lst *gtk.ListStore, iter *gtk.TreeIter, col vals.List, val vals.List) {
  var c []int
  var i []interface{}
  for x := 0;x < col.Len();x++ {
    t, _ := col.Index(x)
    u := t.(string)
    v, _ := strconv.ParseInt(u, 10, 32)
    c = append(c, int(v))
  }
  for y := 0;y < val.Len();y++ {
    t, _ := val.Index(y)
    i = append(i, t)
  }
  fmt.Println(c)
  fmt.Println(i)
  lst.Set(iter, c, i)
}

func listStoreSetValue(fm *eval.Frame, lst *gtk.ListStore, iter *gtk.TreeIter, col int, val interface{}) {
  lst.SetValue(iter, col, val)
}

func treeViewSelectionSelectedRows(fm *eval.Frame, sel *gtk.TreeSelection, lst *gtk.ListStore) *glib.List {
  return sel.GetSelectedRows(lst)
}

func toTreeModel(fm *eval.Frame, tp *gtk.ListStore) *gtk.TreeModel {
  return tp.ToTreeModel()
}

func treeModelGetIter(fm *eval.Frame, tm *gtk.TreeModel, path *gtk.TreePath) *gtk.TreeIter {
  gi, _ := tm.GetIter(path)
  return gi
}

func treeModelGetValue(fm *eval.Frame, tm *gtk.TreeModel, itr *gtk.TreeIter, col int) *glib.Value {
  gv, _ := tm.GetValue(itr, col)
  return gv
}

var Ns = eval.NsBuilder{
}.AddGoFns("gtk:", map[string]interface{}{
  "add" : add,
  "box-new" : boxNew,
  "box-pack-start" : boxPackStart,
  "cell-renderer-text-new" : cellRendererTextNew,
  "entry-new" : entryNew,
  "entry-set-text" : entrySetText,
  "init" : gtkInit,
  "label-new" : labelNew,
  "list-store-append" : listStoreAppend,
  "list-store-new" : listStoreNew,
  "list-store-set" : listStoreSet,
  "list-store-set-val" : listStoreSetValue,
  "main" : gtkMain,
  "main-quit" : mainQuit,
  "orientation-horizontal" : orientationHorizontal,
  "orientation-vertical" : orientationVertical,
  "selection-multiple" : selectionMultiple,
  "text-view-new" : textViewNew,
  "to-tree-model" : toTreeModel,
  "tree-model-get-iter" : treeModelGetIter,
  "tree-model-get-value" : treeModelGetValue,
  "tree-view-append-col" : treeViewAppendColumn,
  "tree-view-col-new-with-attr" : treeViewColumnNewWithAttribute,
  "tree-view-get-sel" : treeViewGetSelection,
  "tree-view-new" : treeViewNew,
  "tree-view-sel-connect" : treeViewSelectionConnect,
  "tree-view-sel-selected-rows" : treeViewSelectionSelectedRows,
  "tree-view-sel-set-mode" : treeViewSelectionSetMode,
  "tree-view-set-model" : treeViewSetModel,
  "win-def-size" : windowDefaultSize,
  "win-connect" : windowConnect,
  "win-new" : windowNew,
  "win-settitle" : windowSetTitle,
  "win-showall" : windowShowAll,
}).Ns()
