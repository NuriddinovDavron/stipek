package main

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 45.2, "cake": 36.3},
		tip:   0,
	}
	return b
}
