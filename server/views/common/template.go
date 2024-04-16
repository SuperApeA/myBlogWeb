package common

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"myBlogWeb/config"
)

var htmlTemplate *HTMLTemplate

type TemplateBlog struct {
	*template.Template
}

type HTMLTemplate struct {
	// TODO： 这里要优化，以TemplateBlog为基类，各自继承
	Category   *TemplateBlog
	Custom     *TemplateBlog
	Detail     *TemplateBlog
	Home       *TemplateBlog
	Index      *TemplateBlog
	Login      *TemplateBlog
	Pigeonhole *TemplateBlog
	Writing    *TemplateBlog
}

// 前端代码中使用的函数
func isODD(num int) bool {
	return num%2 == 0
}

func getNextName(nameList []string, index int) string {
	return nameList[index+1]
}

func getDate(layout string) string {
	return time.Now().Format(layout)
}

func getDateDay(layout string) int {
	return time.Now().Day()
}

// getTemplateBlog 初始化templateBlog，加载前端公共文件
func getTemplateBlog(templateName string) *TemplateBlog {
	t := template.New(templateName + ".html")

	path := config.AppLocalPath
	t.Funcs(template.FuncMap{"isODD": isODD, "getNextName": getNextName, "date": getDate, "dateDay": getDateDay})
	homePath := filepath.Join(path, "/viewsrc/template/home.html")
	// 加载公用部分
	t, err := t.ParseGlob(filepath.Join(path, "/viewsrc/template/layout/*.html"))
	if err != nil {
		log.Println("解析前端html文件报错", err)
	}
	t, err = t.ParseFiles(homePath)
	if err != nil {
		log.Println("解析前端html文件报错", err)
	}
	// 加载对应入参的前端文件
	templatePath := filepath.Join(path, "/viewsrc/template/"+templateName+".html")
	t, err = t.ParseFiles(templatePath)
	if err != nil {
		log.Println("解析前端html文件报错", err)
	}

	return &TemplateBlog{
		Template: t,
	}
}

// InitHTMLTemplate 初始化HTMLTemplate
func InitHTMLTemplate(htmlTemplate *HTMLTemplate) {
	// 通过反射获取成员名称并根据成员名称进行相应初始化
	htmlTemplateType := reflect.TypeOf(htmlTemplate).Elem()
	htmlTemplateValue := reflect.ValueOf(htmlTemplate).Elem()
	for i := 0; i < htmlTemplateType.NumField(); i++ {
		// 获取每个字段的元数据
		field := htmlTemplateType.Field(i)
		// 获取字段的名称
		fieldName := field.Name
		// 使用 FieldByName 获取指定字段的值
		fieldValue := htmlTemplateValue.FieldByName(fieldName)
		// 检查字段是否存在且可被设置
		if fieldValue.IsValid() && fieldValue.CanSet() {
			// 根据名称初始化TemplateBlog
			newValue := getTemplateBlog(strings.ToLower(fieldName))
			fieldValue.Set(reflect.ValueOf(newValue))
			//log.Println(fmt.Sprintf("%s赋值结果：%s", fieldName, htmlTemplateValue.FieldByName(fieldName)))
		} else {
			log.Println(fmt.Sprintf("%s无法赋值", fieldName))
		}
	}
	//log.Println("htmlTemplate初始化后结果：", htmlTemplate)
}

func InitHTMLTemplateCtl() {
	if htmlTemplate == nil {
		htmlTemplate = &HTMLTemplate{
			Category:   &TemplateBlog{},
			Custom:     &TemplateBlog{},
			Detail:     &TemplateBlog{},
			Home:       &TemplateBlog{},
			Index:      &TemplateBlog{},
			Login:      &TemplateBlog{},
			Pigeonhole: &TemplateBlog{},
			Writing:    &TemplateBlog{},
		}
		InitHTMLTemplate(htmlTemplate)
	}
}

func GetHTMLTemplateCtl() *HTMLTemplate {
	if htmlTemplate == nil {
		InitHTMLTemplateCtl()
	}
	return htmlTemplate
}

// WriteError templateBlog执行解析
func (t *TemplateBlog) WriteError(wr io.Writer, err error) {
	if err != nil {
		// 直接返回错误
		_, _ = wr.Write([]byte(err.Error()))
	}
}

// WriteData templateBlog执行解析
func (t *TemplateBlog) WriteData(wr io.Writer, data any) error {
	if err := t.Template.Execute(wr, data); err != nil {
		log.Println("前端文件加载解析执行报错", err)
		return err
	}
	return nil
}
