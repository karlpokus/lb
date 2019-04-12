# lb
A load balancer in go in ~60 LOC. Supports both tcp-, and http proxying.

The implementation is stupefiedly simple and thus has its shortcomings - reading http from a net.Conn with keep-alive enabled will never end in an EOF - so unless you want resource leakage you better use a short hard timeout. That timeout will not work for tcp obviously. Round robin is the default schedule.

# usage
```bash
# start lb on port 12000
$ go run lb.go 12000
# start two http servers
$ go run httpserver.go 3300
$ go run httpserver.go 3301
# run curl against a http path and lb port
$ curl localhost:12000/foo
```

# license
MIT
