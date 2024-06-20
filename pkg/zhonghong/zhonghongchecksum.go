package zhonghongchecksum


func Checksum(data []byte) int {
	sum := 0
	for _, b := range data {
		sum = sum + int(b)
	}
	return sum % 256
}