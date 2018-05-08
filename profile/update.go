package profile

import (
	pb "app/grpc"
	"encoding/json"
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

	var profile map[string]interface{}
	inrec, _ := json.Marshal(p)
	json.Unmarshal(inrec, &profile)

	var result map[string]interface{}
	if e := c.Find(selector).One(&result); e != nil {
		return e
	}

	for k, v := range profile {
		result[k] = v
	}

	updated, e := FillStruct(result, &pb.Profile{})
	if e != nil {
		return e
	}

	if e := c.Update(selector, &updated); e != nil {
		return e
	}

	return nil
}

func FillStruct(input map[string]interface{}, output interface{}) (interface{}, error) {
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
	return structValue.Interface(), nil
}

func (rpc *RPC) Update(ctx context.Context, p *pb.UpdateProfile) (*pb.Profile, error) {
	if e := Update(p.Id, p.Profile); e != nil {
		return nil, grpc.Errorf(codes.Internal, e.Error())
	}

	return p.Profile, nil
}
