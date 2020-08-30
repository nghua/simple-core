package setting

import "simple-core/utils/toml"

var (
	Mode           string
	Port           string
	ReadTimeout    int64
	WriteTimeout   int64
	MaxHeaderBytes int64
	DbUser         string
	DbPassword     string
	DbHost         string
	DbName         string
	SecureKey      string
)

func init() {
	config := toml.Load("conf/server.toml")

	Mode = config.GetString("development.mode")
	Port = config.GetString("server.port")
	ReadTimeout = config.GetInt64("server.read_timeout")
	WriteTimeout = config.GetInt64("server.write_timeout")
	MaxHeaderBytes = config.GetInt64("server.max_header_bytes")
	SecureKey = config.GetString("server.secure_key")
	DbUser = config.GetString("database.db_user")
	DbPassword = config.GetString("database.db_password")
	DbHost = config.GetString("database.db_host")
	DbName = config.GetString("database.db_name")
}
