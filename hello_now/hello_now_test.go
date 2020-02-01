package hello_now

import "testing"

func TestShowTimeNow(t *testing.T) {
	host := "0.beevik-ntp.pool.ntp.org"
	if done := ShowTimeNow(host); done != true {
		t.Fatalf("bad show timr %s", host)
	}

	hostFail := "test"
	if done := ShowTimeNow(hostFail); done == true {
		t.Fatalf("bad show timr %s", hostFail)
	}
}
