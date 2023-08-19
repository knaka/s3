package main

import (
	"github.com/knaka/s3clt"
	"os"
)

func main() {
	//var err error
	command := s3clt.CommandUnknown
	args := os.Args[1:]
	switch os.Args[0] {
	case "s3get":
		command = s3clt.CommandGet
	case "s3put":
		command = s3clt.CommandPut
	default:
		if len(os.Args) < 2 {
			panic("Too few arguments")
		}
		switch os.Args[1] {
		case "get":
			command = s3clt.CommandGet
			args = args[1:]
		case "put":
			command = s3clt.CommandPut
			args = args[1:]
		default:
			panic("Unknown command " + os.Args[1])
		}
	}
	s3clt.Run(command, args)
}
