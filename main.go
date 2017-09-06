package main

import (
	"crypto/tls"
	"encoding/pem"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("You need to specify a target with the form host or host:port")
		os.Exit(1)
	}

	target := os.Args[1]

	if !strings.Contains(target, ":") {
		target = target + ":443"
	}

	tlsConn, err := tls.Dial("tcp", target, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		panic("failed to connect: " + err.Error())
	}

	certs := tlsConn.ConnectionState().PeerCertificates

	rootCert := certs[len(certs)-1]
	block := pem.Block{
		Type:  "CERTIFICATE",
		Bytes: rootCert.Raw,
	}
	err = pem.Encode(os.Stdout, &block)
	if err != nil {
		panic(err)
	}
	tlsConn.Close()
}
