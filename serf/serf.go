package serf

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/KelvinWu602/node-discovery/blueprint"
	"github.com/hashicorp/serf/serf"
)

type IPv4 = blueprint.IPv4

type forusSerf struct {
	client *serf.Serf
}

func (s *forusSerf) New() error {
	// Create a Serf instance
	config := serf.DefaultConfig()
	config.MemberlistConfig.BindAddr = "0.0.0.0"
	// config.MemberlistConfig.AdvertiseAddr = contactNode.String()
	config.EventCh = make(chan serf.Event, 256)
	client, err := serf.Create(config)
	if err != nil {
		log.Printf("Failed to create Serf client: %v", err)
		return errors.New("failed to create client")
	}

	s.client = client
	return nil
}

func (s *forusSerf) JoinCluster(contactNode IPv4) error {
	// Join the Serf cluster
	contactNodeStr := fmt.Sprintf("%d.%d.%d.%d", contactNode[0], contactNode[1], contactNode[2], contactNode[3])
	_, err := s.client.Join([]string{contactNodeStr}, false)
	// _, err := s.client.Join([]string{contactNode.String()}, false)
	if err != nil {
		log.Printf("Failed to join Serf cluster: %v", err)
		return errors.New("failed to join cluster")
	}

	return nil
}

func (s *forusSerf) LeaveCluster() error {
	if s.client == nil {
		return errors.New("node discovery service aborted")
	}

	// Leave the Serf cluster
	s.client.Leave()

	// Clear client variable
	s.client = nil

	return nil
}

func (s *forusSerf) GetMembers() ([]IPv4, error) {
	if s.client == nil {
		return nil, errors.New("node discovery service aborted")
	}

	// Get the list of members in the Serf cluster
	members := s.client.Members()

	// Extract the IP addresses of the members
	ipAddresses := make([]IPv4, len(members))
	for i, member := range members {
		ipAddress, err := net.ResolveIPAddr("ip", member.Addr.String())
		if err != nil {
			log.Printf("Failed to resolve IP address: %v", err)
			return nil, errors.New("failed to get member ip")
		}
		// ipAddresses[i] = [4]byte(ipAddress.IP)
		// copy(ipAddresses[i][:], ipAddress.IP.To4())
		ipAddresses[i] = blueprint.IPv4(ipAddress.IP.To4())
	}

	return ipAddresses, nil
}
