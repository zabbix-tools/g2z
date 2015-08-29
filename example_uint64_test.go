package g2z_test

import (
	"github.com/cavaliercoder/g2z"
)

func ExampleRegisterUint64Item() {
	panic("THIS_SHOULD_NEVER_HAPPEN")
}

func init() {
	g2z.RegisterUint64Item("go.ping", "", Ping)
}

func Ping(request *g2z.AgentRequest) (uint64, error) {
	return 1, nil
}
