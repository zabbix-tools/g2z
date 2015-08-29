package g2z_test

import (
	"github.com/cavaliercoder/g2z"
)

func ExampleRegisterInitHandler() {
	panic("THIS_SHOULD_NEVER_HAPPEN")
}

func init() {
	g2z.RegisterInitHandler(MyInitFunc)
}

func MyInitFunc() error {
	g2z.LogInfof("MyModule loaded")
	return nil
}
