package buildingManager

import (
	"fmt"
)

type building struct {
	buildingID string
	studentCount int
}

func new(newBuildingID string)*building{
	var newBuilding=building{buildingID:newBuildingID,studentCount:0}
	return &newBuilding
}

func (this *building)showInfo()  {
	fmt.Printf("宿舍楼: %s    入住: %d人",this.buildingID,this.studentCount)
	fmt.Println()
}
