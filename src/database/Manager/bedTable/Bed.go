package bedManager

import "fmt"

type bed struct {
	bedID string
	floor int
	dormitoryID string
	buildingID string
	uesd bool
	studentID string
	price int
	nextBed *bed
}

func new(newBedID string,newFloor int,newDormitoryID string,newBuildingID string,newPrice int)*bed{
	var newBed=bed{bedID:newBedID,floor:newFloor,dormitoryID:newDormitoryID,buildingID:newBuildingID,uesd:false,price:newPrice}
	return &newBed
}

func (this *bed)showInfo()  {
	fmt.Println("bedID: ",this.bedID)
	fmt.Println("floor: ",this.floor)
	fmt.Println("dormitoryID: ",this.dormitoryID)
	fmt.Println("buildingID: ",this.buildingID)
	fmt.Println("used: ",this.uesd)
	fmt.Println("studentID: ",this.studentID)
	fmt.Println("price: ",this.price)
}
