package profile

import (
	pb "app/grpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

//RH_Delete method removes a profile instance using the uid pased through
//URL params
func DeleteById(id string) {
}

func (rpc *RPC) DeleteById(ctx context.Context, p *pb.DeleteProfile) (*pb.Profile, error) {
	return &pb.Profile{}, grpc.Errorf(codes.Internal, "hola")
}
