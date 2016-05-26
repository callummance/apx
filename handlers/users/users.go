package users

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/OliWheeler/quick_apx/app/models"
	"log"
)

func Get(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	oID := bson.ObjectIdHex(c.Param("_id"))
	user := models.User{}
	err := db.C(models.CollectionUsers).FindId(oID).One(&user)
	if err != nil {
		c.Error(err)
	}
	log.Print(user.Name)
	c.JSON(200, user)
}
