package profile

import (
	pb "app/grpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

func Upsert(id string, upsert interface{}) (*mgo.ChangeInfo, error) {
	c, _ := Collection()

	selector := bson.M{
		"_id": bson.ObjectIdHex(id),
	}

	i, e := c.Upsert(selector, upsert)
	if e != nil {
		return nil, e
	}
	return i, nil
}

func (rpc *RPC) Upsert(ctx context.Context, p *pb.Profile) (*mgo.ChangeInfo, error) {
	i, e := Upsert(p.Id, p)
	if e != nil {
		return nil, grpc.Errorf(codes.Internal, e.Error())
	}

	return i, nil
}
