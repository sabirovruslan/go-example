package hello_now

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func ShowTimeNow(host string) bool {
	t, err := ntp.Time(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting date err:%s", err)
		return false
	}
	fmt.Fprintf(os.Stdout, "Hello now %s", t)

	return true
}
