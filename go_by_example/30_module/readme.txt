steps:
- copy 30_module/ to GOPATH/simplemodule

imported as "simplemodule"
used as "simplemaths"

building package:
$ go build

installing package:
$ go install

----------------------------------------------

troubleshooting:
error: "package XXX is not in GOROOT" when building a Go project
solution: 
$ go env -w GO111MODULE=off
