use flag
use net
use path
use ptr

dirptr = (flag:string "dir" (path:abs ~) "directory name")
portptr = (flag:string "port" "80" "port number")

flag:parse

net:http:handle "/" (net:http:file-server (ptr:deref $dirptr))
net:http:listen-serve ":"(ptr:deref $portptr)
