use glib
use gtk

# Append single value to TreeView's model.
fn list-append [lst x]{
  gtk:list-store-set-val $lst (gtk:list-store-append $lst) 0 $x
}

# Append several values to TreeView's model.
fn list-append-multi [lst @x]{
  all $x | each [i]{ list-append $lst $i }
}

# Handler of "changed" signal of TreeView selection.
fn selection-handler [sel lst ent]{
  rows = (gtk:tree-view-sel-selected-rows $sel $lst)
  items = []

  pth = (glib:list-data-to-treepath $rows)
  tm = (gtk:to-tree-model $lst)
  iter = (gtk:tree-model-get-iter $tm $pth)
  value = (gtk:tree-model-get-value $tm $iter 0)
  str = (glib:value-get-string $value)
  items = $str

  gtk:entry-set-text $ent (print [$items])
}

# Init Gtk.
gtk:init

# Setup Window.
win = (gtk:win-new)
gtk:win-settitle $win 'Products written in Go'
gtk:win-connect $win 'destroy' { gtk:main-quit }
#gtk:win-def-size $win 800 600

# Setup Widgets.
rootbox = (gtk:box-new (gtk:orientation-vertical) 6)
treeview = (gtk:tree-view-new)
entry = (gtk:entry-new)
liststore = (gtk:list-store-new (glib:type-string))

# Populate list.
list-append-multi $liststore 'Go' 'Docker' 'CockroachDB'

# TreeView Properties.
renderer = (gtk:cell-renderer-text-new)
column = (gtk:tree-view-col-new-with-attr Value $renderer text 0)
gtk:tree-view-append-col $treeview $column
gtk:tree-view-set-model $treeview $liststore

# TreeView Selection Properties.
selection = (gtk:tree-view-get-sel $treeview)
gtk:tree-view-sel-set-mode $selection (gtk:selection-multiple)

gtk:tree-view-sel-connect $selection changed { selection-handler $selection $liststore $entry }

# Packing.
gtk:box-pack-start $rootbox $treeview $true $true 0
gtk:box-pack-start $rootbox $entry $false $false 0
gtk:add $win $rootbox

# Start Everything.
gtk:win-showall $win
gtk:main
