package configs

type Database struct {
	Host        string `yaml:"host`
	Port        int    `yaml:"port"`
	Accountname string `yaml:"accountname"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
}
