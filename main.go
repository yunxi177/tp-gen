package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

//Conf 配置
type Conf struct {
	Path           string `yaml:"path"`
	AppName        string `yaml:"application_name"`
	Module         string `yaml:"module"`
	ModelName      string `yaml:"model_name"`
	ControllerName string `yaml:"controller_name"`
	ServiceName    string `yaml:"service_name"`
	ValidateName   string `yaml:"validate_name"`
	GenPath        string
	CDATA          CCond
}

// CCond Controller 生成条件
type CCond struct {
	GenAdd  bool
	GenUp   bool
	GenList bool
	GenInfo bool
	GenDel  bool
}

// 项目路径
var path string

func main() {
	conf := Conf{}
	table := flag.String("t", "", "表名")
	genParam := flag.String("g", "aduli", "要生成的方法名")
	module := flag.String("m", "", "要生成的数据所属模块,不指定则用配置文件中配置的模块")
	origin := flag.String("o", "", "表的来源")

	flag.Parse()

	for _, value := range *genParam {
		genFlag := string(value)
		switch genFlag {
		case "a":
			conf.CDATA.GenAdd = true
		case "d":
			conf.CDATA.GenDel = true
		case "u":
			conf.CDATA.GenUp = true
		case "l":
			conf.CDATA.GenList = true
		case "i":
			conf.CDATA.GenInfo = true
		}
	}

	if *table == "" {
		fmt.Println("表名不能为空")
		return
	}
	// 获取配置文件信息
	data, err := ioutil.ReadFile("./config/conf.yaml")
	if err != nil {
		fmt.Println("获取配置文件失败")
		return
	}

	//解析配置文件
	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		fmt.Println("解析配置文件失败", err)
		return
	}

	if *module != "" {
		conf.Module = *module
	}
	if conf.Path == "" {
		conf.Path = "./"
	}

	conf.GenPath = conf.Path + "/" + conf.AppName + "/" + conf.Module

	fileName := cameCase(*table)
	fmt.Println("name:", *genParam)
	fmt.Println("table:", fileName)
	// 解析模型
	modelName := fileName
	if *origin != "" {
		*origin = strings.ToUpper(*origin)
		modelName = *origin + fileName
		paseModel(modelName, conf, "OriginModel.tpl", *table, *origin)
	} else {
		paseModel(fileName, conf, "model.tpl", *table, *origin)
	}

	// 解析service
	parseServices(fileName, conf, modelName, *origin)
	// 解析controller
	parseController(fileName, conf, *origin)
	//解析 validate
	parseValidate(fileName, conf)
}

func parseValidate(fileName string, cfg Conf) {
	tplContent := parseTpl("validate.tpl", map[string]interface{}{"fileName": fileName, "cfg": cfg})

	mPath := cfg.GenPath + "/" + cfg.ValidateName + "/" + fileName + "Validate.php"
	writeFile(mPath, tplContent)
}
func parseController(fileName string, cfg Conf, origin string) {
	tplContent := parseTpl("controller.tpl", map[string]interface{}{"fileName": fileName, "origin": origin, "cfg": cfg})

	mPath := cfg.GenPath + "/" + cfg.ControllerName + "/" + fileName + ".php"
	writeFile(mPath, tplContent)
}

func parseServices(fileName string, cfg Conf, modelName string, origin string) {
	tplContent := parseTpl("service.tpl", map[string]interface{}{"fileName": fileName, "cfg": cfg, "modelName": modelName, "origin": origin})

	mPath := cfg.GenPath + "/" + cfg.ServiceName + "/"
	if origin != "" {
		mPath = cfg.GenPath + "/" + cfg.ServiceName + "/" + origin + "/"
	}
	mkAllDir(mPath)
	writeFile(mPath+fileName+"Service.php", tplContent)
}

// paseModel 解析模型
func paseModel(fileName string, cfg Conf, tplName string, tableName string, origin string) {
	tplContent := parseTpl(tplName, map[string]interface{}{"fileName": fileName, "cfg": cfg, "tableName": tableName, "origin": origin})

	mPath := cfg.GenPath + "/" + cfg.ModelName + "/" + fileName + ".php"
	writeFile(mPath, tplContent)
}

// writeFile 文件写入
func writeFile(path string, content []byte) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("文件创建失败", err)
		return
	}
	defer file.Close()
	file.Write(content)
}

// parseTpl 解析模板
func parseTpl(tplName string, data map[string]interface{}) []byte {
	newbytes := bytes.NewBufferString("")

	t := template.Must(template.ParseFiles("./template/" + tplName))
	t.Execute(newbytes, data)
	tplContent, err := ioutil.ReadAll(newbytes)
	if err != nil {
		fmt.Printf("%v", err)
		return []byte{}
	}

	return tplContent
}
func cameCase(str string) string {
	// 是否有表前缀, 设置了就先去除表前缀
	var text string
	for _, p := range strings.Split(str, "_") {
		// 字段首字母大写的同时, 是否要把其他字母转换为小写
		if p == "id" {
			text += "ID"
			continue
		}
		switch len(p) {
		case 0:
		case 1:
			text += strings.ToUpper(p[0:1])
		default:
			text += strings.ToUpper(p[0:1]) + p[1:]
		}
	}
	return text
}

// mkAllDir 递归创建目录
func mkAllDir(path string) {
	res := IsDir(path)
	if res {
		return
	}
	err := os.MkdirAll(path, 0777)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		return
	}

}

//IsDir 判断文件是否存在
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
