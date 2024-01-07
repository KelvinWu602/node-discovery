package blueprint

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
)

func GetDockerHostPrivateIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	dockerSubnet := net.IPNet{
		IP:   net.ParseIP("172.17.0.0"),
		Mask: net.CIDRMask(16, 32),
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil && !dockerSubnet.Contains(ipNet.IP.To4()) {
			return ipNet.IP.String(), nil
		}
	}
	return "", errors.New("cannot find docker host ip")
}

func GetDockerHostPublicIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	parsedIP := net.ParseIP(string(ip))
	if parsedIP == nil {
		return "", errors.New("ipify returning string not ip addr")
	}
	return string(ip), nil
}
