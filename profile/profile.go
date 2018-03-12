package profile

import (
	"app/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

type Model struct{}

type Profile struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	FirstName string
	LastName  string
	Birthday  string
}

func (m *Model) Create(c *gin.Context) {
	mongo := db.Mongo{}
	session, _ := mongo.Open()
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	p := session.DB("Admin").C("profiles")

	profile := &Profile{
		bson.NewObjectId(),
		"hola mundo",
		"hola mundo",
		"hola mundo",
	}

	err := p.Insert(profile)
	if err != nil {
		panic(err)
	}

	c.String(http.StatusOK, "hello")
}

func (m *Model) Find(c *gin.Context) {
	mongo := db.Mongo{}
	session, _ := mongo.Open()
	defer session.Close()

	fmt.Printf("%s", session)
	c.String(http.StatusOK, "hello")
}

func (m *Model) Update(c *gin.Context) {
	fmt.Println("profile test create method")
}

func (m *Model) Delete(c *gin.Context) {
	fmt.Println("profile test create method")
}
