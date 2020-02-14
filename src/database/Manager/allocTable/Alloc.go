package allocManager

import "fmt"

type alloc struct {
	college string           //所在学院
	buildingID string
	nextAlloc *alloc
}

func new(newCollege string,newBuildingID string)*alloc{

	var newAlloc=alloc{college:newCollege,buildingID:newBuildingID}
	return &newAlloc
}

func (this *alloc)showInfo()  {
	fmt.Println("college: ",this.college,"  分配到  ","buildingID: ",this.buildingID)
}
