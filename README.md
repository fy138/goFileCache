# goFileCache
a  file cache for golang use gob encode

## Example
```go

package main

import (
	"log"
	"time"

	fc "github.com/fy138/goFileCache"
)

type MyInfo struct {
	Name string
	Age  int
}
type School struct {
	SchoolName string
	Studen     struct {
		Name string
		Age  int
	}
}

func main() {
	fc := fc.NewFileCache("./tmp/")

	//key 1
	k1 := "myinfo"
	d := &MyInfo{Name: "fy", Age: 100}
	err := fc.SetCache(k1, d)
	if err != nil {
		log.Print(err)
	}
	mydata := &MyInfo{}
	err = fc.GetCache(k1, 7200, mydata)
	if err != nil {
		log.Print(err)
	}
	log.Printf("%s=>%d", mydata.Name, mydata.Age)

	//key2
	k2 := "schoolinfo"
	si := &School{SchoolName: "myschool"}
	si.Studen.Name = "test"
	si.Studen.Age = 10
	err = fc.SetCache(k2, si)
	if err != nil {
		log.Print(err)
	}
	scinfo := &School{}
	err = fc.GetCache(k2, 7200, scinfo)
	if err != nil {
		log.Print(err)
	}
	log.Printf("%s,%s=>%d", scinfo.SchoolName, si.Studen.Name, si.Studen.Age)
	//expired
	time.Sleep(time.Second * 5)
	err = fc.GetCache(k2, 4, scinfo)
	if err != nil {
		log.Print(err)
	}
}

	
  ```
