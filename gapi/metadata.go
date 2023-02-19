package gapi

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"log"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	grpcUserAgentHeader        = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("Extracting metadata md :%+v", md)

		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		if userAgents := md.Get(grpcUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		if clientIP := md.Get(xForwardedForHeader); len(clientIP) > 0 {
			mtdt.ClientIP = clientIP[0]
		}

		if p, ok := peer.FromContext(ctx); ok {
			mtdt.ClientIP = p.Addr.String()
		}
	}
	return mtdt
}
