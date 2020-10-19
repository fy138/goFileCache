# goFileCache
a  file cache for golang use gob encode

## Example
```go

	fc := NewFCache("./tmp/")
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
	si.Studen.Name = "Ok"
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
	
  ```
