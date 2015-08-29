package g2z_test

import (
	"github.com/cavaliercoder/g2z"
)

func ExampleRegisterDoubleItem() {
	panic("THIS_SHOULD_NEVER_HAPPEN")
}

func init() {
	g2z.RegisterDoubleItem("go.double", "", Double)
}

func Double(request *g2z.AgentRequest) (float64, error) {
	return 1.23456, nil
}
