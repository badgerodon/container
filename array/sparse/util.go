package sparse

const (
	max64 uint64 = (1 << 64) - 1
	max32 uint32 = (1 << 32) - 1
	max16 uint16 = (1 << 16) - 1
	max8  uint8  = (1 << 8) - 1
)

func getArrayPosition8(bmp uint8, idx uint) int {
	mask := ^(max8 << idx)
	return popcount32(uint32(bmp & mask))
}

func getArrayPosition16(bmp uint16, idx uint) int {
	mask := ^(max16 << idx)
	return popcount32(uint32(bmp & mask))
}

func getArrayPosition32(bmp uint32, idx uint) int {
	mask := ^(max32 << idx)
	return popcount32(bmp & mask)
}

func getArrayPosition64(bmp uint64, idx uint) int {
	mask := ^(max64 << idx)
	return popcount64(bmp & mask)
}
