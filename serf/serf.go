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

func (s *ForusSerf) NewAgent() error {
	// Create a Serf instance
	config := serf.DefaultConfig()
	config.MemberlistConfig.BindAddr = "0.0.0.0"
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
