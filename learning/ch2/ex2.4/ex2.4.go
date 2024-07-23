// 用来统计输入数的二进制1数目
package popcount

var pc [256]byte 

func init(){
	for i := range pc { // 直接可以将slice当参数
		pc[i] = pc[i / 2] + byte(i & 1) // byte 可以返回1的个数, pc[i] 表示数字i 二进制时1的位置个数
	}
}

func PopCount(x uint64) int{

	ans := 0
	for ; x != 0 ; x >>= 1{   // 每次右移一位
		if x & 1 == 1{
			ans ++
		}
	}
	return ans
}