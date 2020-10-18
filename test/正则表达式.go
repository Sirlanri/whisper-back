package main

import "regexp"

func main2() {
	source := "mail@ri-cocn"
	reg1, err := regexp.Compile(`\S+@\S+\.`)
	if err != nil {
		println(err.Error())
	}
	result := reg1.MatchString(source)
	println("结果为 ", result)
}
