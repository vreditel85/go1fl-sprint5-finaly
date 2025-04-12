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
	fmt.Sprintf("Имя: %s Вес: %d Рост: %1.f", p.Name, int(p.Weight), p.Height)
}
