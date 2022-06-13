package mysql

// Pool 连接池配置
type Pool struct {
	Min uint32 `json:"min"`
	Max uint32 `json:"max"`
}

// Single MysqlSingle  配置
type Single struct {
	Host     string `json:"host"`
	Port     uint32 `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// Conig Mysql配置
type Conig struct {
	Model  string `json:"model"`
	Single Single `json:"single"`
	Pool   Pool   `json:"pool"`
}
