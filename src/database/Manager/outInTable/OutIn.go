package outInManager

import (
	"fmt"
	"time"
)

type outIn struct {
	studentID string
	buildingID string
	in bool
	out bool
	time string
	nextOutIn *outIn
}

func newIn(newStudentID string,newBuildingID string)*outIn{
	var now=time.Now().String()
	var newOutin=outIn{studentID:newStudentID,buildingID:newBuildingID,in:true,out:false,time:now,nextOutIn:nil }
	return &newOutin
}

func newOut(newStudentID string,newBuildingID string)*outIn{
	var now=time.Now().String()
	var newOutin=outIn{studentID:newStudentID,buildingID:newBuildingID,in:false,out:true,time:now,nextOutIn:nil }
	return &newOutin
}

func (this *outIn)showInfo()  {
	if(this.out==true){
		fmt.Printf("学号为%s的学生于%s离开%s宿舍\n",this.studentID,this.time,this.buildingID)
	}
	if(this.in==true){
		fmt.Printf("学号为%s的学生于%s进入%s宿舍\n",this.studentID,this.time,this.buildingID)
	}
}
