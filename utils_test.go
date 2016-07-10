package utils

import (
	"fmt"
	"testing"
)

func TestGetIp(t *testing.T) {
	ips := GetAllIP()
	t.Log(ips)
}
func TestGetLocalIp(t *testing.T) {

	ip, err := GetExternalIP()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---------", ip)
}
