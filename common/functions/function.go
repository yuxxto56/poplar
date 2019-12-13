// 公共方法
// @kancun Team
// @Contact ly@900sui.com
package functions

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)
//md5加密
func HashMd5(str string) string{
	md5Inst := md5.New()
	md5Inst.Write([]byte(str))
	result := md5Inst.Sum([]byte(""))
	return fmt.Sprintf("%x",result)
}
//sha1加密
func HashSha1(str string) string{
	sha1Inst := sha1.New()
	_,err := sha1Inst.Write([]byte(str))
	if err != nil{
    	log.Fatal(err.Error())
	}
	result := sha1Inst.Sum([]byte(""))
	return fmt.Sprintf("%x",result)
}

//base64加密
func Base64Encode(str string) string{
	//转换成byte类型
	strB := []byte(str)
	return base64.StdEncoding.EncodeToString(strB)
}

//base64解密
func Base64Decode(str string) string{
	//转换成byte类型
	bytes,_ := base64.StdEncoding.DecodeString(str)
	return string(bytes[:])
}

//验证手机号
func CheckPhone(phone string) bool{
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}
//验证邮箱
func CheckEmail(email string) bool{
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//马赛克手机号
func MarkPhone(phone string,re ...string) string{
	if len(phone) != 11 {
		return phone
	}
	var replaceMark string
	if len(re) == 0{
		replaceMark = strings.Repeat("*",5)
	}else{
		replaceMark = strings.Repeat(string(re[0]),5)
	}
	replace := phone[3:8]
	return strings.Replace(phone,replace,replaceMark,1)
}

//使用gob编码将数据转化为byte切片
func GobEncode2Byte(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//gob编码的byte切片数据转化为其他数据
func GobDecodeByte(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

//string转json输出
func StringsToJSON(str string) string {
	var jsons bytes.Buffer
	for _, r := range str {
		rint := int(r)
		if rint < 128 {
			jsons.WriteRune(r)
		} else {
			jsons.WriteString("\\u")
			if rint < 0x100 {
				jsons.WriteString("00")
			} else if rint < 0x1000 {
				jsons.WriteString("0")
			}
			jsons.WriteString(strconv.FormatInt(int64(rint), 16))
		}
	}
	return jsons.String()
}


