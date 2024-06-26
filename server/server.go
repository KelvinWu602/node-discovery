package server

import (
	"context"

	"github.com/KelvinWu602/node-discovery/blueprint"
	"github.com/KelvinWu602/node-discovery/protos"
)

type Server struct {
	nodeDiscovery blueprint.NodeDiscovery
	protos.UnimplementedNodeDiscoveryServer
}

func NewServer(nd blueprint.NodeDiscovery) Server {
	return Server{nodeDiscovery: nd}
}

func (s Server) JoinCluster(ctx context.Context, req *protos.JoinClusterRequest) (*protos.JoinClusterResponse, error) {
	err := s.nodeDiscovery.JoinCluster((*req).ContactNodeIP)
	return &protos.JoinClusterResponse{}, err
}

func (s Server) LeaveCluster(ctx context.Context, req *protos.LeaveClusterRequest) (*protos.LeaveClusterResponse, error) {
	err := s.nodeDiscovery.LeaveCluster()
	return &protos.LeaveClusterResponse{}, err
}

func (s Server) GetMembers(ctx context.Context, req *protos.GetMembersRequest) (*protos.GetMembersReponse, error) {
	members, err := s.nodeDiscovery.GetMembers()
	if err != nil {
		return &protos.GetMembersReponse{}, err
	}
	selfIP, err := blueprint.GetDockerHostPublicIP()
	if err != nil {
		return &protos.GetMembersReponse{}, err
	}

	for idx, elem := range members {
		if elem == selfIP {
			members = append(members[:idx], members[idx+1:]...)
			break
		}
	}

	// conversion from member in blueprint to member state in proto
	return &protos.GetMembersReponse{Member: members}, nil
}

func (s Server) mustEmbedUnimplementedNodeDiscoveryServer() {

}
