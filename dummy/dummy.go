package main

import (
	"github.com/cavaliercoder/g2z"
)

func main() {
	g2z.Log(g2z.LogLevelInformation, "This is dummy.main()")
}

func init() {
	g2z.Log(g2z.LogLevelInformation, "This is dummy.init()")
}
