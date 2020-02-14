package dormitoryManager

import "fmt"

type dormitory struct {
	dormitoryID string
	floor int
	buildingID string
	bedCount int
	usedBed int
	price int
	nextDormitory *dormitory
}

func new(newDormitoryID string,newFloor int,newBuildingID string,newBedCount int,newPrice int)*dormitory{
	var newDormitory=dormitory{dormitoryID:newDormitoryID,floor:newFloor,buildingID:newBuildingID,bedCount:newBedCount,usedBed:0,price:newPrice,nextDormitory:nil}
	return &newDormitory
}

func (this *dormitory)showInfo()  {
	fmt.Println("dormitoryID: ",this.dormitoryID)
	fmt.Println("floor: ",this.floor)
	fmt.Println("buildingID: ",this.buildingID)
	fmt.Println("bedCount: ",this.bedCount)
	fmt.Println("usedBed: ",this.usedBed)
	fmt.Println("price: ",this.price)
}
