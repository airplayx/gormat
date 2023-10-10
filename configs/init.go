/*Package configs ...
@Time : 2019/12/25 16:05
@Software: GoLand
@File : init
@Author : Bingo <airplayx@gmail.com>
*/
package configs

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"os"
)

var (
	//JSON the default conf
	JSON = []byte(`{"const":{"theme":"light","scale":"1.0","language":"en"},"language":{"black":"黑色","light":"白色","default":"默认","others":"其它","effective after restart":"重启后生效","base":"基本","screen":"屏幕","language":"语言","rules":"规则","mapping":"映射","special":"特殊","sql to struct":"sql 转 struct","json to struct":"json 转 struct","edit link":"编辑连接","add link":"添加连接","bad link parameters":"错误的连接参数","Please input %s to delete the current group record":"请输入 %s 确认删除当前组记录","action":"操作","save ok":"保存成功","confirm":"确定","cancel":"取消","Please confirm to delete the current: %s ?":"确认删除当前 %s 连接记录吗?","no":"否","yes":"是","info":"信息","connection test":"测试连接","testing":"测试中","connection successful":"连接成功","the database can not be empty":"数据库名称不能为空","duplicate links":"重复的连接记录","engine":"引擎","host":"主机地址","port":"端口","user":"用户名","password":"密码","database":"数据库","table":"表结构","build only":"仅生成","build and include omitempty":"生成并包含omitempty","carriage return":"回车换行","auto save files":"保存文件","save factory func":"生成工厂函数","output folder":"输出文件夹","tags":"标签","exclusion table":"排除表","always execute":"始终执行","type conversion":"类型转换","the fields starting with is_ are automatically converted to bool":"is_开头字段自动转为bool","specified fields transformation":"指定字段转换","bool":"布尔值"},"sql2struct":{"target_dir":"./models","auto_save":false,"tags":["gorm","json"],"json_omitempty":false,"exclude_tables":[],"try_complete":true,"tinyint2bool":true,"sourceMap":[],"reflect":"{\"tinyint\":\"int8\",\"smallint\":\"int16\",\"int\":\"int32\",\"bigint\":\"int64\",\"float\":\"float64\",\"double\":\"float64\",\"decimal\":\"float64\",\"char\":\"string\",\"varchar\":\"string\",\"text\":\"string\",\"mediumtext\":\"string\",\"longtext\":\"string\",\"time\":\"time.Time\",\"date\":\"time.Time\",\"datetime\":\"time.Time\",\"timestamp\":\"time.Time\",\"enum\":\"string\",\"set\":\"string\",\"blob\":\"string\"}","special":"{\"id\":\"uint\"}"}}`)
	//CustomFile the default conf file
	CustomFile = "./config.json"
)

func init() {
	_, err := os.Stat(CustomFile)
	if err == nil {
		JSON, _ = ioutil.ReadFile(CustomFile)
	}
}

//Text the func translation
func Text(key string, replace ...interface{}) string {
	l, _ := jsonparser.GetString(JSON, "const", "language")
	if str, err := jsonparser.GetString(JSON, "language", key); l != "en" && err == nil {
		key = str
	}
	return fmt.Sprintf(key, replace...)
}
