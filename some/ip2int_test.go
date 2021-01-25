package some

import (
	"testing"
)

func TestIp2int_Ip2int(t *testing.T) {
	ip := "192.168.0.1"
	ipint := Ip2int(ip)
	ipint1 := Ip2int1(ip)
	if ipint != 3232235521 {
		t.Fail()
	}
	if ipint1 != 3232235521 {
		 t.Fail()
	}
}