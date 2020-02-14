package Manager

import (
	allocManager "database/Manager/allocTable"
	bedManager "database/Manager/bedTable"
	buildingManager "database/Manager/buildingTable"
	dormitoryManager "database/Manager/dormitoryTable"
	outInManager "database/Manager/outInTable"
	studentManager "database/Manager/studentTable"
	"fmt"
)

func Init()  {
	allocManager.Init()
	bedManager.Init()
	buildingManager.Init()
	dormitoryManager.Init()
	outInManager.Init()
	studentManager.Init()
}

func DormitoryAlloc()  {
	bedManager.GetBed()
}

func StudentHandel()  {
	var a int
	for(true){
		fmt.Println("1. 学生入学")
		fmt.Println("2. 学生入住")
		fmt.Println("3. 学生维护")
		fmt.Println("4. 学生查询")
		fmt.Scanln(&a)
		if(a>4||a<1){
			fmt.Println("非法输入请重新输入")
			fmt.Println("------------------------")
			continue
		}else {
			break
		}
	}
	if(a==1){
		studentManager.AddStudent()
	}
	if(a==2){
		studentManager.StudentAddBed()
	}
	if(a==3){
		studentManager.UpdataStudent()
	}
	if(a==4){
		studentManager.SearchStudent()
	}
}

func InformationSearch()  {
	var a int
	for(true){
		fmt.Println("1. 按公寓楼号查询住宿信息")
		fmt.Println("2. 按学号查询住宿信息")
		fmt.Println("3. 显示所有学生")
		fmt.Println("4. 显示所有宿舍楼")
		fmt.Println("5. 显示所有宿舍")
		fmt.Println("6. 显示所有床位")
		fmt.Println("7. 显示所有分配策略")
		fmt.Println("8. 显示进出表")
		fmt.Scanln(&a)
		if(a>8||a<1){
			fmt.Println("非法输入请重新输入")
			fmt.Println("------------------------")
			continue
		}else {
			break
		}
	}
	if(a==1){
		bedManager.ListOnBuilding()
	}
	if(a==2){
		studentManager.SearchStudent()
	}
	if(a==3){
		studentManager.ListInfo()
	}
	if(a==4){
		buildingManager.ListInfo()
	}
	if(a==5){
		dormitoryManager.ListInfo()
	}
	if(a==6){
		bedManager.ListInfo()
	}
	if(a==7){
		allocManager.ListInfo()
	}
	if(a==8){
		outInManager.ListInfo()
	}
}

func OutInTable()  {
	var a int
	for(true){
		fmt.Println("1. 学生进入")
		fmt.Println("2. 学生离开")
		fmt.Scanln(&a)
		if(a>4||a<1){
			fmt.Println("非法输入请重新输入")
			fmt.Println("------------------------")
			continue
		}else {
			break
		}
	}
	outInManager.AddOutIn(a)
	fmt.Println("------------------------")
	outInManager.ListInfo()
}



