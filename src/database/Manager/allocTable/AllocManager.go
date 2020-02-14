package allocManager

import "fmt"

var count int
var firstAlloc *alloc
var lastAlloc *alloc

func Init()  {
	firstAlloc= new("17计算机工程学院","C19")
	var secondAlloc= new("16管理学院","C9")
	var thirdAlloc= new("16计算机工程学院","C18")
	lastAlloc= new("17国际商学院","C13")
	firstAlloc.nextAlloc=secondAlloc
	secondAlloc.nextAlloc=thirdAlloc
	thirdAlloc.nextAlloc=lastAlloc
	count=4
}

func ListInfo()  {
	var thisAlloc *alloc
	thisAlloc=firstAlloc
	for(thisAlloc!=nil){
		thisAlloc.showInfo()
		fmt.Println("------------------------")
		thisAlloc=thisAlloc.nextAlloc
	}
}

func SearchBuildingID(college string) string {
	var thisAlloc=firstAlloc;
	for(thisAlloc!=nil){
		if(thisAlloc.college==college){
			return thisAlloc.buildingID
		}
		thisAlloc=thisAlloc.nextAlloc
	}
	return ""
}
