package buildingManager

import (
	"fmt"
	"strconv"
)

var buildingTable[20] *building

func Init(){
	var i int
	for i=0;i<20;i++{
		var newBuildingID="C"+strconv.Itoa(i+1)
		buildingTable[i]= new(newBuildingID)
	}
}

func ListInfo()  {
	var i int
	for i=0;i<20;i++{
		buildingTable[i].showInfo()
		fmt.Println("------------------------")
	}
}

func AddStudent(BuildingID string) bool  {
	var i int
	for i=0;i<20;i++{
		fmt.Println(buildingTable[i].buildingID)
		if(buildingTable[i].buildingID==BuildingID){
			buildingTable[i].studentCount=buildingTable[i].studentCount+1
			return true
		}
	}
	return false
}
