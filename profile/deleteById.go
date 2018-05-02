package profile

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RH_Delete method removes a profile instance using the uid pased through
// URL params
func RH_DeleteById(c *gin.Context) {
	var update map[string]interface{}
	update["status"] = STATUS_DELETED

	i, e := Upsert(c.Param("id"), update)
	if e != nil {
		c.JSON(http.StatusInternalServerError, e)
	}

	c.JSON(http.StatusOK, i)
}

//func (rpc *RPC) FindById(ctx context.Context, p *pb.Profile) (*pb.Profile, error) {
//return nil, grpc.Errorf(codes.Internal, "hola")
//}

//func (rpc *RPC) DeleteById(ctx context.Context, p *pb.Profile) (*pb.Profile, error) {
//return nil, grpc.Errorf(codes.Internal, "hola")
//}
