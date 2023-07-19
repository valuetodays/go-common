package utils

import (
	"testing"
)

func TestPrintIps(t *testing.T) {
	PrintIps()
}

func TestGetFirstNonLoopbackHostInfo(t *testing.T) {
	hostname, ip, err := GetFirstNonLoopbackHostInfo()
	if nil != err {
		t.Fatal(err)
	}
	t.Log("hostname=", hostname, ", ip=", ip)
}
