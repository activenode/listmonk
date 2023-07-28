package main

import (
	"fmt"
	"reflect"

	"github.com/zcalusic/sysinfo" // to replace syscall info
)

type Employee struct {
	Name string
	Age  int
	Job  string
}

func getFieldString(e *Employee, field string) string {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func getFieldInteger(e *Employee, field string) int {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

func main() {
	var (
		si sysinfo.SysInfo
		// TODO: Maybe change this because of https://stackoverflow.com/questions/29415909/cannot-get-uname-by-golang
	)

	si.GetSysInfo()

	// data, err := json.MarshalIndent(&si, "", "  ")
	// if err != nil {
	// 	fmt.Printf("WARNING: error getting system info: %v", err)
	// }

	// r := reflect.ValueOf(&si.OS)
	// f := reflect.Indirect(r).FieldByName("Release")
	fmt.Println("##################")
	fmt.Println(si.OS.Name)
	fmt.Println(si.OS.Release)
	fmt.Println("##################")
	// fmt.Println(os.Hostname())
	// fmt.Println(runtime.GOOS)

	e := Employee{"Adam", 36, "CEO"}
	fmt.Println(getFieldString(&e, "Name"))

	fmt.Println(getFieldInteger(&e, "Age"))

	fmt.Println(getFieldString(&e, "Job"))

}
