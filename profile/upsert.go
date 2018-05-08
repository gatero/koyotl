package profile

import (
	pb "app/grpc"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	bson "gopkg.in/mgo.v2/bson"
)

func Update(id string, p *pb.Profile) error {
	c, _ := Collection()
	log.Printf("\n\n UPDATE: %v \n\n", p)

	selector := bson.M{
		"_id": bson.ObjectIdHex(id),
	}

	var result map[string]interface{}
	if e := c.Find(selector).One(&result); e != nil {
		return e
	}
	profile, e := FillProfile(result, p)
	if e != nil {
		return e
	}
	log.Printf("PROFILE: %v\n\n", profile)

	if e := c.Update(selector, &p); e != nil {
		return e
	}

	return nil
}

func FillProfile(input map[string]interface{}, output *pb.Profile) (pb.Profile, error) {
	structValue := reflect.ValueOf(output).Elem()
	for k, v := range input {
		if strings.Title(k) != "_id" {
			structField := structValue.FieldByName(strings.Title(k))
			if !structField.IsValid() {
				return pb.Profile{}, fmt.Errorf("No such field: %s in obj", k)
			}

			if !structField.CanSet() {
				return pb.Profile{}, fmt.Errorf("Cannot set %s field value", k)
			}

			structFieldType := structField.Type()
			val := reflect.ValueOf(v)
			if structFieldType != val.Type() {
				return pb.Profile{}, errors.New("Provided value type didn't match obj field type")
			}

			structField.Set(val)
		}
	}
	log.Printf("structValue: %v\n\n", reflect.Indirect(structValue))
	profile := reflect.Indirect(structValue)
	log.Printf("PROFILE: %T\n\n", profile)
	return pb.Profile{}, nil
}

func (rpc *RPC) Update(ctx context.Context, p *pb.UpdateProfile) (*pb.Profile, error) {
	if e := Update(p.Id, p.Profile); e != nil {
		return nil, grpc.Errorf(codes.Internal, e.Error())
	}

	return p.Profile, nil
}
