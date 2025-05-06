package model

type Lab []LabElement

type LabElement struct {
	// seconds
	Start float64
	// seconds
	End float64
	// phoneme index
	Index int
}
