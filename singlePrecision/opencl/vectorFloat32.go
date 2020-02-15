package opencl

type VectorFloat32 []float32

func (s VectorFloat32) Copy(data VectorFloat32) int {
	return copy(data, s)
}

func (s VectorFloat32) Length() int {
	return len(s)
}
