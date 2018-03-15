package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var sess *mgo.Session

type SampleDoc struct {
	Id        *bson.ObjectId `bson:"_id,omitempty"`
	Field_one string
	Field_two string
}

var mongoURL = ""

func init() {
	var err error

	f, err := os.Open("/run/secrets/mongo-url")
	if err != nil {
		panic("You must attach a secret with key 'mongo-url'")
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	mongoURL = string(data)

	sess, err = mgo.Dial(mongoURL)

	if err != nil {
		panic(err)
	}
}

// Handle a serverless request
func Handle(req []byte) []byte {
	doc := SampleDoc{}

	err := json.Unmarshal(req, &doc)

	if err != nil {
		return []byte(err.Error())
	}

	err = sess.DB("").C("ofdocs").Insert(doc)
	if err != nil {
		return []byte(err.Error())
	}

	count, err := sess.DB("").C("ofdocs").Count()
	if err != nil {
		return []byte(err.Error())
	}

	return []byte(fmt.Sprintf("%d", count))
}
