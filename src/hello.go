package main

import (
	"database/Manager"
	"fmt"
)

func main(){

	Manager.Init()
	for(true){
		var a int
		fmt.Println("1. 寝室分配")
		fmt.Println("2. 学生管理")
		fmt.Println("3. 信息查询")
		fmt.Println("4. 出入登记")
		fmt.Scanln(&a)
		fmt.Println("------------------------")
		if(a>4||a<1){
			fmt.Println("非法输入请重新输入")
			fmt.Println("------------------------")
			continue
		}
		if(a==1){
			Manager.DormitoryAlloc()
			continue
		}
		if(a==2){
			Manager.StudentHandel()
			continue
		}
		if(a==3){
			Manager.InformationSearch()
			continue
		}
		if(a==4){
			Manager.OutInTable()
			continue
		}
	}
