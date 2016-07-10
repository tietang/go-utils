package utils

import (
	"errors"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func PanicError(err error) {
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

func GetRunningDir() string {
	file, _ := exec.LookPath(os.Args[0])
	runningDir, _ := filepath.Abs(filepath.Dir(file))
	return runningDir
}
