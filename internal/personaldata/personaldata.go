package personaldata

import (
	"fmt"
)
type Personal struct {
	Name string //имя 
	Weight float64 //вес 
	Height float64 // рост
}

func (p Personal) Print() {
	fmt.Sprintf("Имя: %s\nВес: %1.f\nРост: %1.f\n", p.Name, p.Weight, p.Height)
}
