package profile

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Create is a ethod to create user instances
// it recives a struct of Profile type
func (a *Action) Create(p *Profile) (*Profile, error) {
	// get the collection pointer
	c, _ := Collection()

	// try to Insert the profile instance
	if e := c.Insert(p); e != nil {
		// return error
		return nil, e
	}

	// return profile instance
	return p, nil
}

func (rpc *RPC) Create(ctx context.Context, p *pb.Profile) (*pb.Profile, error) {
	if e := Create(p); e != nil {
		return nil, grpc.Errorf(codes.Internal, e.Error())
	}

	return p, nil
}
