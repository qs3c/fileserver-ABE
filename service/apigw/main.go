package main

import (
	"fileserver_enc/service/apigw/route"
)

func main() {
	r := route.Router()
	r.Run(":8080")
}
