package winds

import (
	"fmt"
)

type Wind struct {
	Airport string
	Ft0     string
	Ft3000  string
	Ft6000  string
	Ft9000  string
	Ft12000 string
}

type WindsAloft struct {
	Winds map[string]Wind
}

func (wa *WindsAloft) AddWind(wind Wind) {
	if wa.Winds == nil {
		wa.Winds = make(map[string]Wind)
	}
	wa.Winds[wind.Airport] = wind
}

func (wa *WindsAloft) PrintWinds() {
	for _, wind := range wa.Winds {
		fmt.Printf("Airport: %s, 0ft: %s, 3000ft: %s, 6000ft: %s, 9000ft: %s, 12000ft: %s\n",
			wind.Airport, wind.Ft0, wind.Ft3000, wind.Ft6000, wind.Ft9000, wind.Ft12000)
	}
}
