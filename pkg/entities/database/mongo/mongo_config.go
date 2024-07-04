package mongo

type DatabaseConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	DBName    string `yaml:"dbname"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	UserTable string `yaml:"usertable"`
	JobTable  string `yaml:"jobtable"`
}


