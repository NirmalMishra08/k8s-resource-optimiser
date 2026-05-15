package utils


func RoundCPU(value int) int {
	return ((value + 49) / 50) * 50
}

func RoundMemory(value int) int {
	return ((value + 127) / 128) * 128
}