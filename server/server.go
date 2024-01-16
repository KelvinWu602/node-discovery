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
	memberInfo, err := s.nodeDiscovery.GetMembers()
	if err != nil {
		return &protos.GetMembersReponse{}, err
	}

	// conversion from member in blueprint to member state in proto
	members := make([]*protos.MemberState, len(memberInfo))
	for idx, info := range memberInfo {
		members[idx] = &protos.MemberState{MembersIP: info.MemberIP, Status: info.Status}
	}
	return &protos.GetMembersReponse{Member: members}, nil
}

func (s Server) mustEmbedUnimplementedNodeDiscoveryServer() {

}
