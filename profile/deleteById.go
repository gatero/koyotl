package profile

import (
	pb "app/grpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func (rpc *RPC) DeleteById(ctx context.Context, p *pb.DeleteProfile) (*pb.Profile, error) {
	profile := &pb.Profile{Status: STATUS_DELETED}
	if e := Update(p.Id, profile); e != nil {
		return nil, grpc.Errorf(codes.Internal, e.Error())
	}

	return profile, nil
}
