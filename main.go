package main

import (
	internal "github.com/Floor-Gang/suggestions/internal"
	util "github.com/Floor-Gang/utilpkg"
)

func main() {
	config := internal.GetConfig("./config.yml")
	internal.Start(config)
	util.KeepAlive()
}
