package main

import (
	"fmt"
	"goAdmin/src/main/utils"
	"testing"
)

func TestMD5(t *testing.T) {

	fmt.Println(utils.EncodeMD5("123456"))

	fmt.Println(Capitalize("createTime"))

	str := "hellOWorlD" //返回str is all lower char
	b := make([]string, len(str))
	for i, _ := range str {
		s := str[i]
		if 'A' <= s && s <= 'Z' {
			s = s - 'A' + 'a'
		}
		b[i] = string(s)
	}
	fmt.Println(str)      //返回hellOWorlD
	fmt.Printf("%s\n", b) //返回helloworld
}

//
//func ParseUpperTo(str string)  {
//
//	public static String upperCharToUnderLine(String param) {
//		        Pattern  p=Pattern.compile("[A-Z]");
//		        if(param==null ||param.equals("")){
//			            return "";
//			        }
//		        StringBuilder builder=new StringBuilder(param);
//		        Matcher mc=p.matcher(param);
//		        int i=0;
//		        while (mc.find()) {
//			            System.out.println(builder.toString());
//			            System.out.println("mc.start():" + mc.start() + ", i: " + i);
//			            System.out.println("mc.end():" + mc.start() + ", i: " + i);
//			            builder.replace(mc.start()+i, mc.end()+i, "_"+mc.group().toLowerCase());
//			            i++;
//			        }
//
//		        if('_' == builder.charAt(0)){
//			            builder.deleteCharAt(0);
//			        }
//		        return builder.toString();
//		    }
//
//}
// Capitalize 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
