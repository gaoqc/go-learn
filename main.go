package main

import (
	"fmt"
	"html/template"
	"os"
)

type ClzInfo struct {
	ClzName  string
	InstName string
	Desc     string
}

func main() {
	path := "/Users/gaoqc/Documents/codes/pingan/doc/tmp/tpl"
	tpl := template.Must(template.ParseFiles(path+"/controller/TplController.java", path+"/controller/TplESAController.java", path+"/service/TplService.java", path+"/service/TplServiceImpl.java"))
	clzInfo := ClzInfo{"Test", "test", "just for test"}
	// tpl.ExecuteTemplate(os.Stdout, "tplESAController", clzInfo)
	// tpl.ExecuteTemplate(os.Stdout, "tplService", clzInfo)
	// tpl.ExecuteTemplate(os.Stdout, "tplServiceImpl", clzInfo)

	destPath := "/Users/gaoqc/Documents/codes/pingan/doc/tmp/tpl/dest/"

	// tplMap := make(map[string]string)
	// tplMap["tplController"] = destPath + clzInfo.ClzName + "Controller.java"
	// tplMap["tplESAController"] = destPath + clzInfo.ClzName + "Controller.java"
	// tplMap["tplSerivce"] = destPath + clzInfo.ClzName + "Service.java"
	// tplMap["tplServiceImpl"] = destPath + clzInfo.ClzName + "ServiceImpl.java"
	// for k, v := range tplMap {
	// 	tpl.ExecuteTemplate(getFile(destPath+clzInfo.ClzName+"Controller.java"), "tplController", clzInfo)
	// 	// fmt.Println("k=%s ,v=%s", k, v)
	// }
	tpl.ExecuteTemplate(getFile(destPath+clzInfo.ClzName+"Controller.java"), "tplController", clzInfo)
	tpl.ExecuteTemplate(getFile(destPath+clzInfo.ClzName+"ESAController.java"), "tplESAController", clzInfo)
	tpl.ExecuteTemplate(getFile(destPath+clzInfo.ClzName+"Service.java"), "tplService", clzInfo)
	tpl.ExecuteTemplate(getFile(destPath+clzInfo.ClzName+"ServiceImpl.java"), "tplServiceImpl", clzInfo)
}
func getFile(filepath string) *os.File {
	f, err := os.Create(filepath)
	if err != nil {
		fmt.Println("err:%v", err)
		return nil
	}
	return f

}
