package profile

import (
	pb "app/grpc"
	"encoding/json"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Find method get all the user instances
// that exist into the database
func Find(query map[string]interface{}) ([]*pb.Profile, error) {
	// get the collection pointer
	c, _ := Collection()
	// declare and empty array
	var p []*pb.Profile
	// try to find all the profiles
	// and store them into the profiles slice
	if len(query) == 0 {
		query = nil
	}

	if e := c.Find(query).All(&p); e != nil {
		// if an error is ocurred then the
		// return the corresponding error
		return nil, e
	}
	return p, nil
}

func (rpc *RPC) Find(ctx context.Context, p *pb.Profile) (*pb.Profiles, error) {
	var query map[string]interface{}
	inrec, _ := json.Marshal(p)
	json.Unmarshal(inrec, &query)

	var profiles []*pb.Profile
	profiles, e := Find(query)
	if e != nil {
		return nil, grpc.Errorf(codes.Internal, e.Error())
	}

	log.Printf("length: %v", len(profiles))
	count := int32(len(profiles))

	return &pb.Profiles{
		Count:    count,
		Profiles: profiles,
	}, nil
}
