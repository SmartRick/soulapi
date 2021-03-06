package conf

// mysql相关配置
type Mysql struct {
	Host       string `mapstructure:"host" json:"host" yaml:"host"`
	Port       int    `mapstructure:"port" json:"port" yaml:"port"`
	Database   string `mapstructure:"database" json:"database" yaml:"database"`
	Username   string `mapstructure:"username" json:"username" yaml:"username"`
	Password   string `mapstructure:"password" json:"password" yaml:"password"`
	InitTables bool   `mapstructure:"init_tables" json:"init_tables" yaml:"init_tables"`
}
