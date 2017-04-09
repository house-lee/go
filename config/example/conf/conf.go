package conf


type serverConf struct {
	Port      uint    `soarConf:"PORT"`
	MySQLHost string  `soarConf_required:"MYSQL_HOST"`
	MySqlPort uint    `soarConf:"MYSQL_PORT"`
	Daemonize bool    `soarConf:"RUN_IN_BACKGROUND"`
	PI        float32 `soarConf:"PI"`
}

var ServerFromFile = serverConf{}
var ServerFromEnv = serverConf{}