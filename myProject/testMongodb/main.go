package main




import ( 
	"gopkg.in/mgo.v2" 
	"gopkg.in/mgo.v2/bson" 
	)
	session, err := mgo.Dial("127.0.0.1:27017")