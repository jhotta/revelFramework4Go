package controllers

import ("github.com/robfig/revel"
		"fmt"
		)


func init() {
	revel.OnAppStart(Init)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
//	revel.InterceptMethod(Application.AddUser, revel.BEFORE)
//	revel.InterceptMethod(Hotels.checkUser, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
	fmt.Printf("done init func in init.go file\n")
}
