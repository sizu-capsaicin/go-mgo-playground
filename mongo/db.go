package mongo

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	mongoPath = "mongodb://localhost/"
	dbName    = "adb"
	path      = mongoPath + dbName
	cName     = "train"
)

// Insert は document を instert する関数
func Insert(data Data) {
	// MongoDB との接続
	s, _ := mgo.Dial(path)
	defer s.Close()
	db := s.DB(dbName)

	col := db.C(cName)
	if err := col.Insert(data); err != nil {
		log.Fatalln(err)
	}
}

// DropAll は指定した collection の全 document を delete する関数
func DropAll(q bson.M) {
	// MongoDB との接続
	s, _ := mgo.Dial(path)
	defer s.Close()
	db := s.DB(dbName)

	info, err := db.C(cName).RemoveAll(q)
	if err != nil {
		log.Fatalln(err)
	}
	if info != nil {
		fmt.Printf("Removed: %d, Updated: %d\n", info.Removed, info.Updated)
	}
}

// FindAll は全ての document を select する関数
func FindAll(data interface{}) {
	// MongoDB との接続
	s, _ := mgo.Dial(path)
	defer s.Close()
	db := s.DB(dbName)

	db.C(cName).Find(bson.M{}).All(data)
}
