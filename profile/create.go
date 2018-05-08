package profile

import (
	pb "app/grpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Create is a ethod to create user instances
// it recives a struct of Profile type
func Create(p *pb.Profile) error {
	// get the collection pointer
	c, _ := Collection()

	// try to Insert the profile instance
	if e := c.Insert(p); e != nil {
		return e
	}
	return nil
}

func (rpc *RPC) Create(ctx context.Context, p *pb.Profile) (*pb.Profile, error) {
	if e := Create(p); e != nil {
		return nil, grpc.Errorf(codes.Internal, e.Error())
	}

	return p, nil
}
