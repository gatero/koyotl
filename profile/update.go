package profile

import (
	pb "app/grpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	bson "gopkg.in/mgo.v2/bson"
)

func Upsert(id string, upsert *pb.Profile) error {
	c, _ := Collection()

	selector := bson.M{
		"_id": bson.ObjectIdHex(id),
	}

	_, e := c.Upsert(selector, &upsert)
	if e != nil {
		return e
	}
	return nil
}

func (rpc *RPC) Upsert(ctx context.Context, p *pb.Profile) (*pb.Profile, error) {
	e := Upsert(p.Id, p)
	if e != nil {
		return nil, grpc.Errorf(codes.Internal, e.Error())
	}

	return p, nil
}
