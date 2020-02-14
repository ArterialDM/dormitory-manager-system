package dormitoryManager

import (
	"fmt"
	"strconv"
)

var count int
var firstDormitory *dormitory
var lastDormitory *dormitory

func Init()  {
	firstDormitory= new("C19-634",6,"C19",4,4500)
	var i,j,k int
	var thisDormitory=firstDormitory
	for i=1;i<21;i++{
		for j=1;j<7;j++{
			for k=0;k<10;k++{
				var newDormitoryID="C"+strconv.Itoa(i)+"-"+strconv.Itoa(j*100+k)
				var newFloor=j
				var newBuildingID="C"+strconv.Itoa(i)
				var newDormitory= new(newDormitoryID,newFloor,newBuildingID,4,4500)
				thisDormitory.nextDormitory=newDormitory
				thisDormitory=newDormitory
			}
		}
	}
	lastDormitory= new("C20-305A",3,"C20",2,6500)
	thisDormitory.nextDormitory=lastDormitory
	count=1202
}

func ListInfo()  {
	var thisDormitory *dormitory
	thisDormitory=firstDormitory
	for(thisDormitory!=nil){
		thisDormitory.showInfo()
		fmt.Println("------------------------")
		thisDormitory=thisDormitory.nextDormitory
	}
}

func ListOnBuilding()  {
	var thisDormitory *dormitory
	var buildingID string
	fmt.Println("请输入查询的楼层")
	fmt.Scanln(&buildingID)
	thisDormitory=firstDormitory
	for(thisDormitory!=nil){
		if(thisDormitory.buildingID==buildingID){
			thisDormitory.showInfo()
			fmt.Println("------------------------")
		}
		thisDormitory=thisDormitory.nextDormitory
	}
	fmt.Println("查询结束")
	fmt.Println("------------------------")
}


