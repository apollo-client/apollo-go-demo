package main

import (
	"fmt"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/apollo-client/apollo-go"
)

var (
	app = &apollo.Application{
		Addr:    "http://81.68.181.139:8080",
		AppId:   "apollo-go",
		Secret:  "40edd9317add4489a2eaa104054b01e6",
		Cluster: "DEV",
	}
	mDef = &DConfig{}
	mPtr unsafe.Pointer
)

type DConfig struct {
	Map     map[string]string `json:"map"`
	Struct  Person            `json:"struct"`
	Strings []string          `json:"strings"`
	Ints    []int32           `json:"ints"`
	String  string            `json:"string"`
	Int     int64             `json:"int"`
	Float   float32           `json:"float"`
	Age     int64             `json:"age"`
	Ages    []int64           `json:"ages"`
}

type Person struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

// DC get dynamic config if you want
func DC() *DConfig {
	p := atomic.LoadPointer(&mPtr)
	if p != nil {
		return (*DConfig)(p)
	}
	return mDef
}

func main() {
	c, err := apollo.NewClient(app, apollo.EnableBackup(true), apollo.BackupPath("../"))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.Watch("application", mDef, &mPtr)
	if err != nil {
		fmt.Println(err)
		return
	}
	// loop and sleep for test
	for i := 0; i < 100; i++ {
		fmt.Printf("dconf:%+v", DC())
		time.Sleep(1 * time.Second)
	}
}
