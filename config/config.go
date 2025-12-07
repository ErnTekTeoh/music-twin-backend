package config

var liveConfiguration = &Configuration{
	//// when deploying need to put 0.0.0.0 instead of localhost to work
	HostURL:   "0.0.0.0:8080",
	FeHostURL: "0.0.0.0:3000",
	//DbConnection: DBConnection{
	//	Username: "admin",
	//	Password: "chattrpublicdb9707!",
	//	DBHost:   "chattr-sg.cru0a6m4ihlu.ap-southeast-1.rds.amazonaws.com",
	//	Port:     "3306",
	//	DBName:   "chattr_sg",
	//},
	//HashSecretKey: "5jkj2g19aa84mafs",
	//SMTPServer: SMTPServer{
	//	Host:     "smtp.gmail.com",
	//	Username: "chattr.sg@gmail.com",
	//	Password: "jhke xrfq xkab rcbs",
	//	Port:     587,
	//},
}

var devConfiguration = &Configuration{
	HostURL:   "0.0.0.0:8080",
	FeHostURL: "0.0.0.0:3000",
	DbConnection: DBConnection{
		Username: "root",
		Password: "password",
		DBHost:   "localhost",
		Port:     "3306",
		DBName:   "music_twin_sg",
	},
	HashSecretKey: "35jvx19gfa864nms",
	SMTPServer: SMTPServer{
		Host:     "smtp.gmail.com",
		Username: "chattr.sg@gmail.com",
		Password: "jhke xrfq xkab rcbs",
		Port:     587,
	},
}

type Configuration struct {
	HostURL            string       `json:"host_url"`
	FeHostURL          string       `json:"fe_host_url"`
	DbConnection       DBConnection `json:"db_connection"`
	HashSecretKey      string       `json:"hash_secret_key"`
	ImageHashSecretKey string       `json:"image_hash_secret_key"`
	SMTPServer         SMTPServer   `json:"smtp_server"`
}

type SMTPServer struct {
	Host     string
	Username string
	Password string
	Port     int
}

type DBConnection struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DBHost   string `json:"db_host"`
	Port     string `json:"port"`
	DBName   string `json:"db_name"`
}

func GetHostURL() string {
	configuration := devConfiguration
	if GetEnv() == LIVE {
		configuration = liveConfiguration
	}
	return configuration.HostURL
}

func GetFeHostURL() string {
	configuration := devConfiguration
	if GetEnv() == LIVE {
		configuration = liveConfiguration
	}
	return configuration.FeHostURL
}

func GetDBConnection() DBConnection {
	configuration := devConfiguration
	if GetEnv() == LIVE {
		configuration = liveConfiguration
	}
	return configuration.DbConnection
}

// NOTE: key needs to be 16 char
func GetHashSecretKey() string {
	configuration := devConfiguration
	if GetEnv() == LIVE {
		configuration = liveConfiguration
	}
	return configuration.HashSecretKey
}

func GetImageHashSecretKey() string {
	configuration := devConfiguration
	if GetEnv() == LIVE {
		configuration = liveConfiguration
	}
	return configuration.ImageHashSecretKey
}

func GetSMTPServerSetting() SMTPServer {
	configuration := devConfiguration
	if GetEnv() == LIVE {
		configuration = liveConfiguration
	}
	return configuration.SMTPServer
}
