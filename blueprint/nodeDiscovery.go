// The blueprint package defines the interface for the forus Node Discovery component.
// To create a new implementation of Node Discovery, implement the NodeDiscovery interface.
package blueprint

// NodeDiscovery should be implemented by all concrete implementation.
type NodeDiscovery interface {
	// communicate with a node specified by input IP address and join the cluster of that node. In case of any errors, log it and return “failed to join cluster” error.
	JoinCluster(contactNode IPv4) error
	// notify other nodes in the cluster that you are going to leave. Log it and return “failed to leave cluster” error.
	LeaveCluster() error
	// Acquire the array of the IP addresses of all cluster members. These IP addresses are assumed to be alive.
	// Return error “node discovery service aborted” if underlying service is not running. Return error “failed to get member ip” otherwise.
	GetMembers() ([]IPv4, error)
}

type IPv4 [4]byte
