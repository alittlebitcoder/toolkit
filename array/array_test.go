package array

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	//将数组使用指定分隔符拼接成字符串
	//arr := []int{1, 2, 3, 4}
	//arr := []float32{1.01, 2.2, 3.3, 4.5}
	//arr := []string{"1.01", "2.2", "3.3", "4.5"}
	//actualInt := Join("", arr)
	//fmt.Println(actualInt)

	//将字符串使用指定分隔符分割成数组
	strval := "1,2,3,4,5"
	arr := Split[int](",", strval)
	fmt.Println(arr)
}
