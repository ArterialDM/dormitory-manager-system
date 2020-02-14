package outInManager

import "fmt"

var count int
var firstOutIn *outIn
var lastOutIn *outIn

func Init()  {
	firstOutIn= newIn("201710097041","C19")
	lastOutIn=firstOutIn
	count=1
}

func AddOutIn(oi int)  {
	var newStudentID string
	var newBuildingID string
	fmt.Println("请输入学生学号")
	fmt.Scanln(&newStudentID)
	fmt.Printf("学生学号为: %s\n",newStudentID)
	fmt.Println("请输入宿舍楼")
	fmt.Scanln(&newBuildingID)
	fmt.Printf("宿舍楼为: %s\n",newBuildingID)
	if(oi==1){
		var newIn= newIn(newStudentID,newBuildingID)
		lastOutIn.nextOutIn=newIn
		lastOutIn=newIn
		newIn.showInfo()
	}else {
		var newOut= newOut(newStudentID,newBuildingID)
		lastOutIn.nextOutIn=newOut
		lastOutIn=newOut
		newOut.showInfo()
	}
	count=count+1
}

func ListInfo(){
	var thisOutIn *outIn
	thisOutIn=firstOutIn
	if(thisOutIn==nil){
		fmt.Println("进出登记表还没初始化")
		return
	}
	for(true){
		thisOutIn.showInfo()
		fmt.Println("------------------------")
		if(thisOutIn.nextOutIn==nil){
			break
		}
		thisOutIn=thisOutIn.nextOutIn
	}
}

