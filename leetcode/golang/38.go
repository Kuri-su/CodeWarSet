package golang

import (
	"bytes"
	"fmt"
	"strconv"
)

func CountAndSay(n int) string {
	res := "1"

	var (
		temp   = ""
		number = 0
	)

	for i := 2; i <= n; i++ {
		lastStr := string(res[0])

		for k, v := range res {
			str := string(v)
			if str == lastStr {
				number++
			} else {
				temp += fmt.Sprintf("%d", number) + lastStr
				number = 1
				lastStr = str
			}

			if k == len(res)-1 {
				temp += fmt.Sprintf("%d", number) + lastStr
			}
		}
		lastStr = ""
		number = 0
		res = temp
		temp = ""
	}

	return res
}

func CountAndSay2(n int) string {
	if n == 1 {
		return "1"
	}
	lastStr := CountAndSay2(n - 1)
	res := bytes.NewBufferString("")
	curNum := lastStr[0]
	curCoefficient := 1
	lastStrLen := len(lastStr)
	for i := 1; i < lastStrLen; i++ {
		if lastStr[i] == curNum {
			curCoefficient += 1
		} else {
			res.WriteString(strconv.Itoa(curCoefficient))
			res.WriteString(string(curNum))
			curNum = lastStr[i]
			curCoefficient = 1
		}
	}
	res.WriteString(strconv.Itoa(curCoefficient))
	res.WriteString(string(curNum))
	return res.String()
}
