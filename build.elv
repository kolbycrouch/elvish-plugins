use str
use path
use ./elvish-libs/pure/argparse
use ./elvish-libs/pure/list

options = [
  &plugins= 'all'
  &help= "Usage: elvish build.elv [-plugins foo,bar,baz]"
]

opts = (argparse:parse $options)

fn -plugin-paths {
  lst = []
  all [./plugins/*] | each [i]{
    lst = [(all $lst) $i]
  }
  put $lst
}

fn -all-plugins {
  lst = []
  all (-plugin-paths) | each [i]{
    lst = [(all $lst) (path:base $i)]
  }
  put $lst
}

fn -chosen-plugins {
  lst = []
  if (eq $opts[plugins] 'all') {
    put (-all-plugins)
  } else {
    put [(str:split , $opts[plugins])]
  }
}

fn -copy-plugins {
  pths = (-plugin-paths)
  all $pths | each [i]{
    cp -r $i elvish/pkg/eval/mods/
  }
}

fn -clean-plugins {
  plgs = (-all-plugins)
  all $plgs | each [i]{
    rm -rf elvish/pkg/eval/mods/$i
  }
}

fn -target-dir {
  if ?(rm -rf target 2>/dev/null) { }
  if ?(mkdir target 2>/dev/null) { }
}

fn -build-elvish {
  cd elvish/cmd/elvish
  go build -ldflags '-w -s' -o ../../../target/elvish
  cd ../../../
}

fn -build-plugins {
  chosen = (-chosen-plugins)
  cd elvish
  all $chosen | each [i]{
    if (or (eq $i 'compress') (eq $opts[plugins] 'all')) {
      go get github.com/klauspost/compress/zstd
    }
    if (or (eq $i 'glib') (eq $opts[plugins] 'all')) {
      go get github.com/gotk3/gotk3/glib
      go get github.com/gotk3/gotk3/gtk
    }
    if (or (eq $i 'gtk') (eq $opts[plugins] 'all')) {
      go get github.com/gotk3/gotk3/glib
      go get github.com/gotk3/gotk3/gtk
    }
    if (or (eq $i 'wasm') (eq $opts[plugins] 'all')) {
      go get github.com/wasmerio/wasmer-go/wasmer
    }
    go build -buildmode=plugin -ldflags '-w -s' -o ../target/$i.so pkg/eval/mods/$i/$i.go
  }
  cd ../
}

if (or (list:contains $args '-help') (list:contains $args '--help')) {
  echo $opts[help]
  exit
}
-target-dir
-clean-plugins
-copy-plugins
echo "Building elvish."
-build-elvish
echo "Done building elvish."
echo "Building Plugins: "(str:join ' ' (-chosen-plugins))
-build-plugins
echo "Done building plugins."
-clean-plugins
