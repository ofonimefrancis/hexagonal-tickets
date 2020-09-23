package main

import "flag"

func main() {
	var server bool
	var dbType, dbURL, redisPassword string
	flag.StringVar(&dbType, "database", "redis", "database type[redis, psql]")
	flag.BoolVar(&server, "server", false, "run in server mode")
}
