package g2z_test

import (
	"fmt"
	"github.com/cavaliercoder/g2z"
)

func ExampleRegisterDiscoveryItem() {
	panic("THIS_SHOULD_NEVER_HAPPEN")
}

func init() {
	g2z.RegisterDiscoveryItem("go.discovery", "", Discover)
}

func Discover(request *g2z.AgentRequest) (g2z.DiscoveryData, error) {
	d := make(g2z.DiscoveryData, 5)

	for i := 0; i < 5; i++ {
		d[i] = g2z.DiscoveryItem{
			"index": fmt.Sprintf("%d", i),
		}
	}

	return d, nil
}
