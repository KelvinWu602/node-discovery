package serf

import (
	"errors"
	"log"
	"net"

	"github.com/hashicorp/serf/serf"
)

type ForusSerf struct {
	agent *serf.Serf
}

func getDockerHostIP() (string, error) {
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
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil && dockerSubnet.Contains(ipNet.IP.To4()) {
			return ipNet.IP.String(), nil
		}
	}
	return "", errors.New("cannot find docker host ip")
}

func (s *ForusSerf) NewAgent() error {
	// Create a Serf instance
	config := serf.DefaultConfig()
	// This module will be run in Docker container with host driver
	// In order to make memberlist work, need to bind to the Docker host ip, i.e. not in 172.17.0.0/16
	dockerHostIP, err := getDockerHostIP()
	if err != nil {
		log.Printf("Failed to create Serf agent: %v", err)
		return errors.New("failed to create agent")
	}
	config.MemberlistConfig.BindAddr = dockerHostIP
	// config.MemberlistConfig.AdvertiseAddr = contactNode.String()
	// config.EventCh = make(chan serf.Event, 256)
	agent, err := serf.Create(config)
	if err != nil {
		log.Printf("Failed to create Serf agent: %v", err)
		return errors.New("failed to create agent")
	}

	s.agent = agent
	return nil
}

func (s ForusSerf) JoinCluster(contactNode string) error {
	if s.agent == nil {
		return errors.New("node discovery service aborted")
	}

	// Join the Serf cluster
	// contactNodeStr := fmt.Sprintf("%d.%d.%d.%d", contactNode[0], contactNode[1], contactNode[2], contactNode[3])
	// _, err := s.agent.Join([]string{contactNodeStr}, true)
	_, err := s.agent.Join([]string{contactNode}, true)
	if err != nil {
		log.Printf("Failed to join Serf cluster: %v", err)
		return errors.New("failed to join cluster")
	}

	return nil
}

func (s ForusSerf) LeaveCluster() error {
	if s.agent == nil {
		return errors.New("node discovery service aborted")
	}

	// Leave the Serf cluster
	s.agent.Leave()
	return nil
}

func (s ForusSerf) GetMembers() ([]string, error) {
	if s.agent == nil {
		return nil, errors.New("node discovery service aborted")
	}

	// Get the list of members in the Serf cluster
	members := s.agent.Members()

	// Extract the IP addresses of the members
	ipAddresses := make([]string, len(members))
	for i, member := range members {
		_, err := net.ResolveIPAddr("ip", member.Addr.String())
		if err != nil {
			log.Printf("Failed to resolve IP address: %v", err)
			return nil, errors.New("failed to get member ip")
		}
		// ipAddresses[i] = [4]byte(ipAddress.IP)
		// copy(ipAddresses[i][:], ipAddress.IP.To4())
		// ipAddresses[i] = blueprint.IPv4(ipAddress.IP.To4())
		ipAddresses[i] = member.Addr.String()
	}

	return ipAddresses, nil
}
