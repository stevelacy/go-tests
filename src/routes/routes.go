package routes

import "os"
import "net/http"
import "fmt"

import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"

type Config struct {
  Url string
}

type post struct {
  Id bson.ObjectId `bson:"_id"`
  Content string   `bson:"content"`
}

var config = Config{Url: "mongodb://localhost:27017/go-lang-test"}

var session, err = mgo.Dial(config.Url)




func Index (w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Index"))
}

func Home (w http.ResponseWriter, r *http.Request) {

  if err != nil {
    fmt.Print("Error: can not connect to mongo", err)
    os.Exit(1)
  }

  comments := session.DB("go-lang-test").C("comments")
  comment := post{Id: bson.NewObjectId(), Content: "Test comment"}
  err = comments.Insert(comment)
  if err != nil {
    fmt.Print("Error: inserting document", err)
    os.Exit(1)
  }

  fmt.Print(err, session)
  w.Write([]byte("Home page"))
}
