package utils

import (
	"errors"
	"net"
	"strings"
	"bytes"
	"strconv"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

//
func GetAllIP() []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)

	}
	size := len(addrs)
	ips := make([]string, size)
	for i, a := range addrs {

		//		println(a.Network(), a.String())
		ips[i] = strings.Split(a.String(), "/")[0]
		//		fmt.Println(strings.Split(a.String(), "/")[0])
		//		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
		//			if ipnet.IP.To4() != nil {
		//				//				os.Stdout.WriteString(ipnet.IP.String() + "\n")
		//				println(ipnet.IP.String())
		//			}
		//			println(ipnet.IP.String())
		//		}
	}
	return ips
}

func GetExternalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
//byte转16进制字符串
func ByteToHex(data []byte) string {
	buffer := new(bytes.Buffer)
	for _, b := range data {

		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}

	return buffer.String()
}

//16进制字符串转[]byte
func HexToBye(hex string) []byte {
	length := len(hex) / 2
	slice := make([]byte, length)
	rs := []rune(hex)

	for i := 0; i < length; i++ {
		s := string(rs[i*2 : i*2+2])
		value, _ := strconv.ParseInt(s, 16, 10)
		slice[i] = byte(value & 0xFF)
	}
	return slice
}