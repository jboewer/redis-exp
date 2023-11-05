package main

import (
	"flag"
	"github.com/jboewer/redis/pkg/redis"
	"log"
	"net"
)

var FlagHost = flag.String("host", "0.0.0.0", "Host to listen to")
var FlagPort = flag.String("port", "6379", "Port to listen to")

func main() {
	flag.Parse()

	server := redis.Server{}

	log.Fatal(server.ListenAndServe(net.JoinHostPort(*FlagHost, *FlagPort)))
}
