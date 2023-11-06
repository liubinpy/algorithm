package basic

// 说明: 给一个一个int32位的数字，输出其二进制的字符串
// 20 -> 00000000000000000000000000010100  最高为符号位
// -1 -> 11111111111111111111111111111111  负数等于后31位取反+1

func PrintBinaryString(num int32) string {
	var str string
	var c int32
	for i := 0; i < 32; i++ {
		c = 1 << i
		if num&c != 0 {
			str = "1" + str
		} else {
			str = "0" + str
		}
	}
	return str
}
