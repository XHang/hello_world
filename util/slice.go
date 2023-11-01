package util

type slice struct{}

var Slice slice

func (s *slice) CreateSliceWithRandomValues(len int) []int {
	array := make([]int, 0, len)
	for i := 0; i < len; i++ {
		array = append(array, Random.GenerateNumber())
	}
	return array
}
