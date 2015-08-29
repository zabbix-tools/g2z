package g2z_test

import (
	"github.com/cavaliercoder/g2z"
)

func ExampleRegisterUninitHandler() {
	panic("THIS_SHOULD_NEVER_HAPPEN")
}

func init() {
	g2z.RegisterUninitHandler(MyUninitFunc)
}

func MyUninitFunc() error {
	g2z.LogInfof("MyModule unloaded")
	return nil
}
