package main

import (
	"bufio"
	"cleanImport/util"
	"io"
	"os"
	"path/filepath"
	"strings"
	"log"
)

var templateDir = "./"

func main() {
	filepath.Walk(templateDir, walkFunc)
	util.ShowSuccess()
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if !info.IsDir() && strings.HasSuffix(info.Name(), "java") {
		unUseImpClass := handle(path, info)
		if unUseImpClass != nil && len(unUseImpClass) > 0 {
			delUnUserImport(path, unUseImpClass)
		}
	}
	return nil
}

func handle(path string, info os.FileInfo) map[int]string {
	if f, err := os.Open(path); err == nil {

		defer f.Close()

		buf := bufio.NewReader(f)
		//所有当前文件导入的的类
		impClass := make(map[string]int, 15)
		//无用的导入类
		unUseImpClass := make(map[int]string, 15)
		lineNum := 0

		for {
			lineNum++
			line, err := buf.ReadString('\n')

			//判断读取结束和打开失败的情况
			if err != nil {
				if err == io.EOF {
					log.Println(info.Name(), "处理完成")
					break
				}
				log.Println(info.Name(), "打开失败")
				return unUseImpClass
			}

			//删除文件前后空格
			line = strings.TrimSpace(line)
			//如果当前行长度为0 或者已经被注释则不处理
			if len(line) <= 1 || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "*") {
				continue
			}

			//如果当前行是import开头，则为引入的类
			if strings.HasPrefix(line, "import ") {
				classNameStartIndex := strings.LastIndex(line, ".") + 1
				className := util.Substr(line, classNameStartIndex, len(line) - classNameStartIndex - 1)
				if _, ok := impClass[className]; ok {
					//如果已经存在则直接判断为当前导入的类为重复导入 需要删除
					unUseImpClass[lineNum] = className
				} else {
					//添加到所有需要导入的map中
					impClass[className] = lineNum
				}

			} else {
				for key, _ := range impClass {
					// 循环所有导入的类，查看当前行中是否包含该类，如果包含则删除。最后剩余的就是无用的导入
					if strings.Contains(line, key) {
						delete(impClass, key)
					}
				}
			}

		}

		for key, value := range impClass {
			// 把剩余的导入加入待删除map
			unUseImpClass[value] = key
		}
		return unUseImpClass

	}
	return nil
}

/**
删除文件中无用的import
*/
func delUnUserImport(path string, unUseImpClass map[int]string) {
	desPath := path + ".bak"

	destFile, deserr := os.OpenFile(desPath, os.O_WRONLY | os.O_CREATE, 0666)
	if deserr != nil {
		return
	}
	defer destFile.Close()

	f, err := os.Open(path)
	if err != nil {
		return
	}

	defer f.Close()
	buf := bufio.NewReader(f)
	lineNum := 0
	for {
		lineNum++
		line, err := buf.ReadString('\n')

		//判断读取结束和打开失败的情况
		if err != nil {
			if err == io.EOF {
				log.Println("处理完成")
				break
			}
			log.Println("打开失败")
			return
		}

		if _, ok := unUseImpClass[lineNum]; ok {
			continue
		}
		destFile.WriteString(line)
	}

	f.Close()
	destFile.Close()

	println(strings.Replace(desPath, "\\", "/", -1), strings.Replace(path, "\\", "/", -1))
	os.Rename(strings.Replace(desPath, "\\", "/", -1), strings.Replace(path, "\\", "/", -1))

}


