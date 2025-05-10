package main

// CombineWithIdxHigh 将两个32位整数组合成一个64位整数
// idx放在高32位，appmsg_id放在低32位
func CombineWithIdxHigh(idx uint32, appmsgID uint32) uint64 {
	return (uint64(idx) << 32) | uint64(appmsgID)
}
// func main() {
// 	// 示例使用
// 	idx := uint32(1)
// 	appmsgID := uint32(2247483957)
// 	combined := CombineWithIdxHigh(idx, appmsgID)
// 	println(combined)
// }
