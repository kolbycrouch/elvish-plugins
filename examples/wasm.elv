use file
use wasm

engine = (wasm:new-engine)
store = (wasm:new-store $engine)
module = (wasm:new-module $store (file:open "./examples/example.wasm"))
importobj = (wasm:new-import-object)
instance = (wasm:new-instance $module $importobj)
sum = (wasm:get-function $instance "sum")
wasm:call $sum 3 2
