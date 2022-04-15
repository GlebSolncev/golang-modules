package database

import (
	"fmt"
)

type (
	StructMethod struct {
		Method DBMethods
	}

	database struct {
		//ID   int
		Data []byte
	}
)

var (
	db  = database{}
	seq = 1
)

func (m *StructMethod) Get() []byte {
	//var result = []byte{} //map[int]byte{}
	fmt.Println(">>>>>", db.Data)
	//for id, item := range db {
	//	fmt.Print("Get >>> ", string(item.Data), id)
	//	json.Unmarshal(item.Data, &resule)
	//	//result[id] = item.Data
	//}
	//data, err := json.Marshal(db.Data)
	//check(err)

	return db.Data
}
func (m *StructMethod) Save(data []byte) bool {

	fmt.Println("DATA >>", data, string(data))
	db.Data = data
	//fmt.Println(string(db))
	//db = db.Data&database{
	//	Data: data,
	//}

	// fmt.Println(dt)
	//if err := c.Bind(u); err != nil {
	//	return err
	//}
	//db[dt.ID] = dt
	//seq++
	return true
}
