package g2z_test

import (
	"github.com/cavaliercoder/g2z"
	"strings"
)

func ExampleRegisterStringItem() {
	panic("THIS_SHOULD_NEVER_HAPPEN")
}

func init() {
	g2z.RegisterStringItem("go.echo", "Hello,world", Echo)
}

func Echo(request *g2z.AgentRequest) (string, error) {
	return strings.Join(request.Params, " "), nil
}
