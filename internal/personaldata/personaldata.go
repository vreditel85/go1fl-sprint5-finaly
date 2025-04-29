package personaldata

import (
	"fmt"
)

type Personal struct {
	Name   string  //имя
	Weight float64 //вес
	Height float64 // рост
}

func (p Personal) Print() {
	//if len(p.Name) <= 0 {
	//	fmt.Sprintf("no Name: %s", p.Name)
	//}
	//if p.Weight <= 0 {
	//	fmt.Sprintf("no Weight: %2.f", p.Weight)
	//}
	//if p.Height <= 0 {
	//	fmt.Sprintf("no Height: %2.f", p.Height)
	//}
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n\n", p.Name, p.Weight,p.Height)
}
