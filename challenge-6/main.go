package main

import "challenge/routers"

func main() {
	var port = ":4000"

	routers.StartSever().Run(port)
}
