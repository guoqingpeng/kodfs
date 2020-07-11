package kodfs_config

type KodfsConfig struct {
	httpPort int

	httpsPort int

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

	cfg.httpPort = 8880

	cfg.httpsPort = 4443

	cfg.MaxCups = 4

}

//parse config file for real
func (cfg *KodfsConfig) ParseConfig(confPath string) {

}

//获取http端口
func (cfg *KodfsConfig) GetHttpPort() int {

	return cfg.httpPort
}