package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func main() {
	m1 := make(map[string]string, 0)
	m2 := make(map[string]string, 0)
	for i := 0; i < 100; i++ {
		iStr := fmt.Sprintf("%d", i)
		m1[iStr] = iStr
		m2[iStr] = iStr
	}
	//log.Println(m1)
	//log.Println(m2)
	fmt.Println(CmIsEq(m1, m2))

}
func Md5Str(str string) string {
	w := md5.New()
	_, _ = io.WriteString(w, str)
	return fmt.Sprintf("%x", w.Sum(nil))
}

//是否相等。就是判断md5
func CmIsEq(cm1 map[string]string, cm2 map[string]string) bool {
	return Md5Data(cm1) == Md5Data(cm2)
}

//把 map 变成md5 string
func Md5Data(data map[string]string) string {
	str := strings.Builder{}
	for k, v := range data {
		str.WriteString(k)
		str.WriteString(v)
	}
	fmt.Println(str.String())
	return Md5Str(str.String())
}
