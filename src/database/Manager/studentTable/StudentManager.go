package studentManager

import (
	bedManager "database/Manager/bedTable"
	buildingManager "database/Manager/buildingTable"
	"fmt"
	"strings"
)

var count int
var firstStudent *student
var lastStudent *student

func Init(){
	firstStudent= new("王栋铭",true,20,"201710097041","13760657316",3,"17计算机工程学院")
	lastStudent=firstStudent
	count=1
}

func AddStudent()  {
	var newStudent *student
	var name string           //姓名
	var sex bool              //true是男，false是女
	var age int               //年龄
	var studentID string      //学号
	var phoneNumber string    //电话号码
	var grade int             //年级
	var college string        //所在学院
	fmt.Println("请输入学生姓名")
	fmt.Scanln(&name)
	fmt.Println("确认你的学生姓名：",name)
	fmt.Println("请输入学生性别，0是女，1是男")
	fmt.Scanln(&sex)
	if(sex==false){
		fmt.Println("确认你的性别：female")
	}else {
		fmt.Println("确认你的性别：male")
	}
	fmt.Println("请输入学生年龄")
	fmt.Scanln(&age)
	fmt.Println("确认你的学生年龄：",age)
	fmt.Println("请输入学生学号")
	fmt.Scanln(&studentID)
	fmt.Println("确认你的学生学号：",studentID)
	fmt.Println("请输入学生电话号码")
	fmt.Scanln(&phoneNumber)
	fmt.Println("确认你的学生电话号码：",phoneNumber)
	fmt.Println("请输入学生年级")
	fmt.Scanln(&grade)
	fmt.Println("确认你的学生年级：",grade)
	fmt.Println("请输入学生所在学院")
	fmt.Scanln(&college)
	fmt.Println("确认你的学生所在学院：",college)
	newStudent= new(name,sex,age,studentID,phoneNumber,grade,college)
	fmt.Println("------------------------")
	fmt.Println("您新增了一条记录")
	newStudent.showInfo()
	fmt.Println("------------------------")
	lastStudent.nextStudent=newStudent
	lastStudent=newStudent
	count=count+1
}

func SearchStudent()  {
	var studentID string
	var thisStudent *student
	thisStudent=firstStudent
	fmt.Println("请输入需要查找学生学号")
	fmt.Scanln(&studentID)
	fmt.Println("确认学生学号：",studentID)
	fmt.Println("------------------------")
	for(thisStudent!=nil){
		if(thisStudent.studentID==studentID){
			thisStudent.showInfo()
			fmt.Println("------------------------")
			return
		}
		thisStudent=thisStudent.nextStudent
	}
	fmt.Println("找不到此学生 ")
	fmt.Println("------------------------")
}

func UpdataStudent(){
	var studentID string
	var thisStudent *student
	thisStudent=firstStudent
	fmt.Println("请输入需要更新学生学号")
	fmt.Scanln(&studentID)
	fmt.Println("确认学生学号：",studentID)
	fmt.Println("------------------------")
	for(thisStudent!=nil){
		if(thisStudent.studentID==studentID){
			thisStudent.showInfo()
			fmt.Println("------------------------")
			fmt.Println("开始修改信息")
			fmt.Println("------------------------")
			var name string           //姓名
			var sex bool              //true是男，false是女
			var age int               //年龄
			var phoneNumber string    //电话号码
			var grade int             //年级
			var college int           //所在学院
			fmt.Println("请输入学生姓名")
			fmt.Scanln(&name)
			fmt.Println("确认你的学生姓名：",name)
			thisStudent.name=name
			fmt.Println("请输入学生性别，0是女，1是男")
			fmt.Scanln(&sex)
			if(sex==false){
				fmt.Println("确认你的性别：female")
				thisStudent.sex=false
			}else {
				fmt.Println("确认你的性别：male")
				thisStudent.sex=true
			}
			fmt.Println("请输入学生年龄")
			fmt.Scanln(&age)
			fmt.Println("确认你的学生年龄：",age)
			thisStudent.age=age
			fmt.Println("请输入学生电话号码")
			fmt.Scanln(&phoneNumber)
			fmt.Println("确认你的学生电话号码：",phoneNumber)
			thisStudent.phoneNumber=phoneNumber
			fmt.Println("请输入学生年级")
			fmt.Scanln(&grade)
			fmt.Println("确认你的学生年级：",grade)
			thisStudent.grade=grade
			fmt.Println("请输入学生所在学院")
			fmt.Scanln(&college)
			fmt.Println("确认你的学生所在学院：",college)
			thisStudent.grade=grade
			fmt.Println("------------------------")
			fmt.Println("确认学生更新后的信息")
			fmt.Println("------------------------")
			thisStudent.showInfo()
			return
		}
		thisStudent=thisStudent.nextStudent
	}
	fmt.Println("找不到此学生 ")
	fmt.Println("------------------------")
}

func StudentAddBed(){
	var studentID string
	var thisStudent *student
	thisStudent=firstStudent
	fmt.Println("请输入需要分配床位的学生学号")
	fmt.Scanln(&studentID)
	fmt.Println("确认学生学号：",studentID)
	fmt.Println("------------------------")
	for(thisStudent!=nil){
		if(thisStudent.studentID==studentID){
			var bedID string
			thisStudent.showInfo()
			fmt.Println("------------------------")
			fmt.Println("开始分配床号")
			fmt.Println("------------------------")
			fmt.Println("请输入分配给这位学生的床号")
			fmt.Scanln(&bedID)
			thisStudent.bedID=bedID
			if(bedManager.UpdataUseBaseBedID(bedID,studentID)==false){
				fmt.Println("没有这张床")
				fmt.Println("------------------------")
				return
			}
			var buildingID=strings.Split(bedID,"-")[0]
			if(buildingManager.AddStudent(buildingID)==true){
				fmt.Println("分配成功")
				fmt.Println("------------------------")
				return
			}else{
				fmt.Println("宿舍楼计数出错")
				fmt.Println("------------------------")
			}
			return
		}
		thisStudent=thisStudent.nextStudent
	}
	fmt.Println("找不到此学生")
	fmt.Println("------------------------")
}

func DeleteStudent()  {
	var studentID string
	var thisStudent *student
	thisStudent=firstStudent
	fmt.Println("请输入学生学号")
	fmt.Scanln(&studentID)
	fmt.Println("确认你的学生学号：",studentID)
	fmt.Println("------------------------")
	if(thisStudent==nil){
		fmt.Println("学生表还没初始化")
		return
	}
	if(thisStudent.studentID==studentID){
		fmt.Println("不能删除第一个学生")
		return
	}
	for(true){
		if(thisStudent.studentID==studentID){
			var priorStudent *student
			priorStudent=firstStudent
			for(true){
				if(priorStudent.nextStudent==thisStudent){
					priorStudent.nextStudent=thisStudent.nextStudent
					count=count-1
					return
				}
			}
		}
		if(thisStudent.nextStudent==nil){
			fmt.Printf("找不到学号为%s的学生\n",studentID)
			break
		}
		thisStudent=thisStudent.nextStudent
	}
}

func GetCount() int {
	return count
}

func ListInfo(){
	var thisStudent *student
	thisStudent=firstStudent
	if(thisStudent==nil){
		fmt.Println("学生表还没初始化")
		return
	}
	for(true){
		thisStudent.showInfo()
		fmt.Println("------------------------")
		if(thisStudent.nextStudent==nil){
			break
		}
		thisStudent=thisStudent.nextStudent
	}
}
