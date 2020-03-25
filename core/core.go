package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)


type ICore interface {
	GenerateFile(file string, fType string)
}

type Core struct {
	Name    string
	Format  map[string]string
	FileDir string
}

func (c Core) GenerateFile(file string, fType string) {
	fType = strings.ToLower(fType)
	for key, val := range c.Format {
		file = strings.ReplaceAll(file, key, val)
	}
	fileDir := fType
	if fileDir == "model" {
		fileDir = "models"
	}
	if c.FileDir != "" {
		fileDir = filepath.Join(c.FileDir, fileDir)
	}else{
		fileDir = filepath.Join("output", fileDir)
	}
	fileName := c.Name + ToUpper(fType) + ".go"
	WriteFile(fileDir, fileName, file, 0755)
}

//写入文件
func WriteFile(fileDir string, fileName string, file string, mode os.FileMode) error {
	_, err := os.Stat(fileDir)
	if err != nil {
		err = os.Mkdir(fileDir, mode)
		if err != nil {
			log.Fatalln(err.Error() + ": " + fileDir)
		}
	}
	fn := filepath.Join(fileDir, fileName)
	err = ioutil.WriteFile(fn, []byte(file), mode)
	if err != nil {
		log.Fatalln(err.Error() + ": " + fn)
	}
	fmt.Println("success create :" + fn)
	return err
}

// 字符首字母大写
func ToUpper(str string) string {
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

func UcFirst(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

func LcFirst(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}
