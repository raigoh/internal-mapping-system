package data

type Station struct {
	Name        string
	X, Y        int
	Connections []*Station
}
