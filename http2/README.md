
$ curl -O https://raw.githubusercontent.com/golang/go/master/src/crypto/tls/generate_cert.go
$ go build generate_cert.go
$ ./generate_cert -host localhost
2020/03/20 02:36:19 wrote cert.pem
2020/03/20 02:36:19 wrote key.pem

