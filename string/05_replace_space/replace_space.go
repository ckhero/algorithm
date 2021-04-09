/**
 *@Description
 *@ClassName replace_space
 *@Date 2021/4/8 下午7:34
 *@Author ckhero
 */

package _5_replace_space

import "strings"

func replaceSpace(s string) string {
	res := strings.Builder{}
	for k, v := range s {
		if v == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteByte(s[k])
		}
	}
	return res.String()
}
