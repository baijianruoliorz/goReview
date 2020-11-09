package split

import "strings"

/*
*  @author liqiqiorz
*  @data 2020/11/8 22:09
 */
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	for i > -i {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
