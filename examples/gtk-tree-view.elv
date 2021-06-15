use glib
use gtk

cols = [0 1 2]

# Add a column to the tree view.
fn create-column [title id]{
  renderer = (gtk:cell-renderer-text-new)
  put (gtk:tree-view-col-new-with-attr $title $renderer "text" $id)
}

fn setup-tree-view {
  tv = (gtk:tree-view-new)

  gtk:tree-view-append-col $tv (create-column "Version" $cols[0])
  gtk:tree-view-append-col $tv (create-column "Feature" $cols[1])
  gtk:tree-view-append-col $tv (create-column "BOOGY" $cols[2])

  ls = (gtk:list-store-new (glib:type-string))

  gtk:tree-view-set-model $tv $ls

  put $tv
  put $ls
}

fn add-row [lst map]{
  iter = (gtk:list-store-append $lst)
  gtk:list-store-set $lst $iter $cols $map
}

fn setup-window [title]{
  win = (gtk:win-new)
  gtk:win-settitle $win $title
  gtk:win-connect $win "destroy" { gtk:main-quit }
  gtk:win-def-size $win 600 300
  put $win
}

# Init Gtk.
gtk:init

# Setup Window.
win = (setup-window "Go Feature Timeline")

tv ls = (setup-tree-view)

gtk:add $win $tv

add-row $ls ["r57" "Gofix command added for rewriting code for new APIs" "hello" "ors"]

# Start Everything.
gtk:win-showall $win
gtk:main
