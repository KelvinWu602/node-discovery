package main

import (
	"log"
	"net"

	"github.com/KelvinWu602/node-discovery/blueprint"
	"github.com/KelvinWu602/node-discovery/protos"
	"github.com/KelvinWu602/node-discovery/serf"
	"github.com/KelvinWu602/node-discovery/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// create new nodeDiscovery module (serf agent)
	forusSerf := serf.ForusSerf{}
	err := forusSerf.NewAgent()
	if err != nil {
		log.Fatal(err.Error())
	}

	// create new server with the serf agent
	server := server.NewServer(blueprint.NodeDiscovery(forusSerf))

	// register server to grpc's server (gs)
	gs := grpc.NewServer()
	protos.RegisterNodeDiscoveryServer(gs, protos.NodeDiscoveryServer(server))
	reflection.Register(gs)

	// create listener (lis), use gs to serve lis
	lis, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err.Error())
	}
	gs.Serve(lis)
}
