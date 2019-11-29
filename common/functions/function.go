// 公共方法
// @kancun Team
// @Contact ly@900sui.com
package functions

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
)
//md5加密
func HashMd5(str string) string{
	md5Inst := md5.New()
	md5Inst.Write([]byte(str))
	result := Md5Inst.Sum([]byte(""))
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