package models

import (
	"encoding/json"
	"time"
	"yulong-hids/web/models/wmongo"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

// User model definiton.
type Info struct {
	Id     bson.ObjectId       `bson:"_id"      json:"_id,omitempty"`
	Ip     string              `bson:"ip"       json:"ip"`
	System string              `bson:"system"   json:"system"`
	Type   string              `bson:"type" json:"type"`
	Data   []map[string]string `bson:"data" json:"data"`
	Uptime time.Time           `bson:"uptime"   json:"uptime,omitempty"`
	baseModel
}

func NewInfo() Info {
	mdl := Info{}
	mdl.collectionName = "info"
	return mdl
}

func (c *Info) AllValue(fieldname string, limitnum int) interface{} {
	mConn := wmongo.Conn()
	defer mConn.Close()
	var list []interface{}
	collections := mConn.DB("").C("info")

	err := collections.Find(bson.M{}).Distinct(fieldname, &list)
	if err != nil {
		beego.Error("Collections Find Distinct Error", err)
	}

	length := len(list)
	if limitnum > length {
		limitnum = length
	}

	return list[0:limitnum]
}

func (c *Info) GetInfoByIp(ip string) []Info {
	mConn := wmongo.Conn()
	defer mConn.Close()




	var cli []Info
	collections := mConn.DB("").C("info")

	//	db.getCollection('info').aggregate(
	//		[ { $match : { ip : "172.16.1.1" } } ,{$group: {_id: "$type", "time": {$first: "$uptime"}, "data": {$first: "$data"}}}, {$sort:{"uptime":-1}}]
	//);


	m := []bson.M{
		{"$match": bson.M{"ip": ip}},
		{"$group": bson.M{"_id": "$type", "time": bson.M{"$first": "$uptime"} ,  "ip": bson.M{"$first": "$ip"} ,  "system": bson.M{"$first": "$system"} ,  "type": bson.M{"$first": "$type"} ,   "data": bson.M{"$first": "$data"}    }},
		{"$sort": bson.M{"uptime": -1}}	}


	if err := collections.Pipe(m).All(&cli); err != nil {
		beego.Error("Info Find all Error", err)
	}

	//if err := collections.Find(bson.M{"ip": ip}).All(&cli); err != nil {
	//	beego.Error("Info Find all Error", err)
	//}

	return cli
}

func (c *Info) Aggregate(querylist ...bson.M) []bson.M {
	mConn := wmongo.Conn()
	defer mConn.Close()

	var res []bson.M
	collections := mConn.DB("").C("info")

	j, _ := json.Marshal(querylist)
	beego.Debug("Debug aggregate query:", string(j))

	pipe := collections.Pipe(querylist)
	err := pipe.All(&res)

	if err != nil {
		beego.Error("Collections pipe(pipe.All) aggregate all", err)
	}

	return res
}
