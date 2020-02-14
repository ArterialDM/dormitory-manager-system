package studentManager

import (
	"fmt"
)

type student struct{
	name string           //姓名
	sex bool              //true是男，false是女
	age int               //年龄
	studentID string      //学号
	phoneNumber string    //电话号码
	grade int             //年级
	college string        //所在学院
	bedID string          //床号
	nextStudent *student  //下一个节点
}

func new(newName string,newSex bool,newAge int,newStudentID string,newPhoneNumber string,newGrade int,newCollege string)*student{
	var newStudent=student{name: newName,sex: newSex,age:newAge,studentID:newStudentID,phoneNumber:newPhoneNumber,grade:newGrade,college:newCollege,nextStudent:nil,bedID:"0"}
	return &newStudent
}

func (this *student)showInfo()  {
	fmt.Println("name: ",this.name)
	if(this.sex==true){
		fmt.Println("sex: male")
	}else {
		fmt.Println("sex: female")
	}
	fmt.Println("age: ",this.age)
	fmt.Println("studentID: ",this.studentID)
	fmt.Println("phoneNumber: ",this.phoneNumber)
	fmt.Println("grade: ",this.grade)
	fmt.Println("college: ",this.college)
	fmt.Println("bedID: ",this.bedID)
}



