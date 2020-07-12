package kodfs_config

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"reflect"
	"strconv"
)

type KodfsConfig struct {
	HttpPort int

	HttpsPort int

	MaxCups int
}

//init a object
func NewKodsConfig() *KodfsConfig {

	k := new(KodfsConfig)

	k.SetDefaultConfig()

	return k
}

// use defalut value if config file not set
func (cfg *KodfsConfig) SetDefaultConfig() {

	cfg.HttpPort = 8880

	cfg.HttpsPort = 443

	cfg.MaxCups = 4

}

//parse config file for real set
func (cfg *KodfsConfig) ParseConfig(confPath string) {

	myfile, err := goconfig.LoadConfigFile(confPath)

	if err != nil {

		fmt.Println("读取配置文件失败[kodfs.cfg] 找不到")

	}

	sectios := myfile.GetSectionList()

	for i := 0; i < len(sectios); i++ {

		setcion := sectios[i]

		if setcion == "server" {

			keys := myfile.GetKeyList(setcion)

			v_cfg := reflect.ValueOf(cfg).Elem()

			//取出配置的项目，覆盖默认设置
			for _, key := range keys {
				f := v_cfg.FieldByName(key)
				v, _ := myfile.GetValue(setcion, key)
				setConfigValue(f, v)
			}
		}
	}

}

//根据字段的类型设置对应的值
func setConfigValue(field reflect.Value, value string) {

	switch field.Type().String() {

	case "string":
		field.SetString(value)

	case "int":
		i, _ := strconv.Atoi(value)
		field.SetInt(int64(i))

	default:
		panic("can not find field type to convert")

	}
}
