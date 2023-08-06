package main

import (
	"gotasks/11-lession/server/pkg/cache"
	netsrv "gotasks/11-lession/server/pkg/netsrv"
	"gotasks/11-lession/server/pkg/webapp"
)

func main() {
	go netsrv.Listen()
	webapp.Start(":12345", &cache.FileCache{})
}
