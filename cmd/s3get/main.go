package main

import (
	"github.com/knaka/s3clt"
	"os"
)

func main() {
	s3clt.Run(s3clt.CommandGet, os.Args[1:])
}
