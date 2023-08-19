package main

import (
	"github.com/knaka/s3clt"
	"os"
)

func main() {
	s3clt.RunGet(os.Args[1:])
}
