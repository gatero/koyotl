package profile

import (
	pb "app/grpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Find method get all the user instances
// that exist into the database
func Find(query *pb.Profile, p []*pb.Profile) error {
	// get the collection pointer
	c, _ := Collection()

	if e := c.Find(query).All(p); e != nil {
		// if an error is ocurred then the
		// return the corresponding error
		return e
	}

	return nil
}

func (rpc *RPC) Find(ctx context.Context, p *pb.Profile) ([]*pb.Profile, error) {
	var profiles []*pb.Profile

	if e := Find(p, profiles); e != nil {
		return nil, grpc.Errorf(codes.Internal, e.Error())
	}

	return profiles, nil
}
