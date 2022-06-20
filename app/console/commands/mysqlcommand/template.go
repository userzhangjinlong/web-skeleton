package mysqlcommand

import (
	"bytes"
	"fmt"
	"linkr-frame/app/utils"
	"os"
	"text/template"
)

const databaseTpl = `package {{.DatabaseName}}

import "linkr-frame/app/model"

type {{.TableName | index}} struct {
 {{range.Column}} {{$length := len .Comment}} {{if gt $length 0}}
	//{{.Comment}}  {{else}} // {{.Name}} {{ end }}
	{{$typeLen := len .Type}} {{ if gt $typeLen 0}}{{.Name | index}} {{.Type}} {{.Tag}} {{else}} {{.Name}} {{end}}
{{end}}
	//继承父类model
	model.Model
}

func (model *{{.TableName | index}}) TableName() string {
	return "{{.TableName}}"
}`

type DatabaseTemplate struct {
	databaseTpl string
}

type DatabaseColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName    string
	DatabaseName string
	Column       []*DatabaseColumn
}

//NewDatabaseTemplate 实例化表名称
func NewDatabaseTemplate() *DatabaseTemplate {
	return &DatabaseTemplate{
		databaseTpl: databaseTpl,
	}
}

//AssemblyColumns 格式化获取到的字段输出模板内容
func (db *DatabaseTemplate) AssemblyColumns(tbColumns []*TableColumn) []*DatabaseColumn {
	tplColumns := make([]*DatabaseColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		if !utils.InArray(column.ColumnName, []string{"createAt", "updateAt"}) {
			tag := fmt.Sprintf("`"+"gorm:"+"\"column:%s\""+" "+"json:"+"\"%s\""+"`", column.ColumnName, column.ColumnName)
			tplColumns = append(tplColumns, &DatabaseColumn{
				Name:    column.ColumnName,
				Type:    DBTypeToStructType[column.DataType],
				Tag:     tag,
				Comment: column.ColumnComment,
			})
		}

	}

	return tplColumns
}

//Generate 处理渲染模板内容
func (db *DatabaseTemplate) Generate(dataBaseName string, tableName string, tplColumns []*DatabaseColumn) error {
	tpl := template.Must(template.New("mysqlcommand").Funcs(template.FuncMap{
		"index": utils.UnderscoreToUpperCamelCase,
	}).Parse(db.databaseTpl))

	tplDB := StructTemplateDB{
		TableName:    tableName,
		DatabaseName: utils.ToLower(dataBaseName),
		Column:       tplColumns,
	}

	buf := new(bytes.Buffer)
	err := tpl.Execute(buf, tplDB)

	if err != nil {
		return err
	}
	putStringInFile(dataBaseName, tableName, buf.String())

	return nil
}

//putInFile 将写入buff池内的内容写入指定文件
func putStringInFile(dataBaseName string, fileName string, bufString string) error {
	utils.CreateIfNotExistDir("app/model/" + utils.UnderscoreToUpperCamelCase(dataBaseName))
	file, err := os.OpenFile("app/model/"+dataBaseName+"/"+utils.ToLower(fileName)+".go", os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		fmt.Println("open file err", err)
		return err
	}

	fmt.Fprintf(file, bufString)
	file.Close()

	return nil
}
