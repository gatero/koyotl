package profile

import (
	pb "app/grpc"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Find method get all the user instances
// that exist into the database
func Find(query interface{}, options map[string]interface{}) ([]*pb.Profile, error) {
	// get the collection pointer
	c, e := Collection()
	if e != nil {
		return nil, e
	}
	// declare and empty array
	var p []*pb.Profile

	offset := int(options["page"].(int32) * options["offset"].(int32))
	limit := int(options["limit"].(int32))

	if e := c.Find(query.(*pb.Profile)).Sort("name").Skip(offset).Limit(limit).All(&p); e != nil {
		// if an error is ocurred then the
		// return the corresponding error
		return nil, e
	}

	return p, nil
}

func (rpc *RPC) Find(ctx context.Context, p *pb.FindProfile) (*pb.Profiles, error) {
	options := map[string]interface{}{
		"offset": int32(0),
		"page":   int32(1),
		"limit":  int32(10),
	}

	if p.Offset > 0 {
		options["offset"] = p.Offset
	}

	if p.Page > 0 {
		options["page"] = p.Page
	}

	if p.Limit > 0 {
		options["limit"] = p.Limit
	}

	var profiles []*pb.Profile

	profiles, e := Find(p.Query, options)
	if e != nil {
		return nil, grpc.Errorf(codes.Internal, e.Error())
	}

	count := int32(len(profiles))

	log.Printf("options: %v", options)

	return &pb.Profiles{
		Count:    count,
		Profiles: profiles,
	}, nil
}
