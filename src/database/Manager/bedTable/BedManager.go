package bedManager

import (
	allocManager "database/Manager/allocTable"
	"fmt"
	"strconv"
)

var count int
var firstBed *bed
var lastBed *bed

func Init()  {
	firstBed= new("C19-634-2",6,"C20-634","C19",4500)
	var i,j,k,l int
	var thisBed=firstBed
	for i=1;i<21;i++{
		var newBuildingID="C"+strconv.Itoa(i)
		for j=1;j<7;j++{
			for k=0;k<10;k++{
				var newDormitoryID=newBuildingID+"-"+strconv.Itoa(j*100+k)
				for l=1;l<5;l++{
					var newBedID=newDormitoryID+"-"+strconv.Itoa(l)
					var newBed= new(newBedID,j,newDormitoryID,newBuildingID,4500)
					thisBed.nextBed=newBed
					thisBed=newBed
				}
			}
		}
	}
	lastBed= new("C20-305A-5",3,"C20-305A","C20",6500)
	thisBed.nextBed=lastBed
	count=4802
}

func ListInfo()  {
	var thisBed *bed
	thisBed=firstBed
	for(thisBed!=nil){
		thisBed.showInfo()
		fmt.Println("------------------------")
		thisBed=thisBed.nextBed
	}
}

func AddBed()  {
	var newBed *bed
	var bedID string
	var floor int
	var dormitoryID string
	var buildingID string
	var price int
	fmt.Println("请输入床号")
	fmt.Scanln(&bedID)
	fmt.Println("确认床号：",bedID)
	fmt.Println("请输入楼层")
	fmt.Scanln(&floor)
	fmt.Println("确认楼层：",floor)
	fmt.Println("请输入房间号")
	fmt.Scanln(&dormitoryID)
	fmt.Println("确认房间号：",dormitoryID)
	fmt.Println("请输入楼层号")
	fmt.Scanln(&buildingID)
	fmt.Println("确认楼层号：",buildingID)
	fmt.Println("请输入价格")
	fmt.Scanln(&price)
	fmt.Println("确认价格：",price)
	newBed= new(bedID,floor,dormitoryID,buildingID,price)
	lastBed.nextBed=newBed
	lastBed=newBed
	count=count+1
}

func ListOnBuilding()  {
	var thisBed *bed
	var buildingID string
	fmt.Println("请输入查询的楼层")
	fmt.Scanln(&buildingID)
	thisBed=firstBed
	for(thisBed!=nil){
		if(thisBed.buildingID==buildingID){
			thisBed.showInfo()
			fmt.Println("------------------------")
		}
		thisBed=thisBed.nextBed
	}
	fmt.Println("查询结束")
	fmt.Println("------------------------")
}

func UpdataUseBaseBedID(bedID string,studentID string)bool  {
	var thisBed *bed
	thisBed=firstBed
	for(thisBed!=nil){
		if(thisBed.bedID==bedID){
			thisBed.uesd=true
			thisBed.studentID=studentID
			return true
		}
		thisBed=thisBed.nextBed
	}
	return false
}

func GetBed()  {
	var college string
	var buildingID string
	fmt.Println("请输入年级学院")
	fmt.Scanln(&college)
	fmt.Println("------------------------")
	fmt.Println("确认年级学院：",college)
	fmt.Println("------------------------")
	buildingID=allocManager.SearchBuildingID(college)
	if(buildingID!=""){
		var thisBed=firstBed
		for(thisBed!=nil){
			if(thisBed.buildingID==buildingID&&thisBed.studentID==""){
				fmt.Println("你分配到的床号： ",thisBed.bedID)
				fmt.Println("------------------------")
				return
			}
			thisBed=thisBed.nextBed
		}
		fmt.Println("没有床位了")
		fmt.Println("------------------------")
	} else{
		fmt.Println("没有设置此年级学院的分配策略")
		fmt.Println("------------------------")
	}
}
