package serf

import (
	"errors"
	"log"
	"net"

	"github.com/KelvinWu602/node-discovery/blueprint"
	"github.com/hashicorp/serf/serf"
)

type ForusSerf struct {
	agent *serf.Serf
}

func (s *ForusSerf) NewAgent() error {
	// Create a Serf instance
	config := serf.DefaultConfig()
	// This module will be run in Docker container with host driver
	// In order to make memberlist work, need to bind to the Docker host private ip, i.e. not in 172.17.0.0/16
	dockerHostPrivateIP, err := blueprint.GetDockerHostPrivateIP()
	if err != nil {
		log.Printf("Failed to create Serf agent: %v", err)
		return errors.New("failed to create agent")
	}
	config.MemberlistConfig.BindAddr = dockerHostPrivateIP
	// Advertise the Docker host public ip
	dockerHostPublicIP, err := blueprint.GetDockerHostPublicIP()
	if err != nil {
		log.Printf("Failed to create Serf agent: %v", err)
		return errors.New("failed to create agent")
	}
	config.MemberlistConfig.AdvertiseAddr = dockerHostPublicIP
	// config.EventCh = make(chan serf.Event, 256)
	agent, err := serf.Create(config)
	if err != nil {
		log.Printf("Failed to create Serf agent: %v", err)
		return errors.New("failed to create agent")
	}

	s.agent = agent
	return nil
}

func (s *ForusSerf) JoinCluster(contactNode string) error {
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

func (s *ForusSerf) LeaveCluster() error {
	if s.agent == nil {
		return errors.New("node discovery service aborted")
	}

	// Leave the Serf cluster
	if err := s.agent.Leave(); err != nil {
		log.Println("Failed to leave cluster", err)
		return err
	}
	// Restore Serf agent
	if err := s.agent.Shutdown(); err != nil {
		log.Println("Failed to shutdown serf agent", err)
		return err
	}
	<-s.agent.ShutdownCh()
	if err := s.NewAgent(); err != nil {
		log.Println("Failed to restore serf agent, cannot JoinCluster again. Please stop this service.", err)
		return err
	}
	return nil
}

func (s *ForusSerf) GetMembers() ([]string, error) {
	if s.agent == nil {
		return nil, errors.New("node discovery service aborted")
	}

	// Get the list of members in the Serf cluster
	members := s.agent.Members()

	// Extract the IP addresses of the members
	memberInfo := make([]string, 0)
	for _, member := range members {
		_, err := net.ResolveIPAddr("ip", member.Addr.String())
		if err != nil {
			log.Printf("Failed to resolve IP address: %v", err)
			return nil, errors.New("failed to get member ip")
		}

		if member.Status == serf.StatusAlive {
			memberInfo = append(memberInfo, member.Addr.String())
		}
	}

	return memberInfo, nil
}
