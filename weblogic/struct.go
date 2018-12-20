/***********************************************************************
# File Name: struct.go
# Author: TiDao
# mail: songzhaofeng411@gmail.com
# Created Time: 2018-12-19 12:10:14
*********************************************************************/
package weblogic

type Data struct {
	Env     Env
	Install Install
	Domain  Domain
}

//all string of path must be write full path
type Env struct {
	JavaHome  string `json:"JavaHome"`
	UserName  string `json:"UserName"`
	Password  string `json:"Password"`
	JreHome   string `json:"JreHome"`
	ClassPath string `json:"ClassPath"`
	Path      string `json:"Path"` //the path which find executed file
	EnvFile   string `json:"EnvFile"`
	WlsHome   string `json:"WlsHome"`
}

type Install struct {
	PackagePath string
	InstallTpl  string
	InstallRsp  string
	Locfile     string
	OracleHome  string
}

type Domain struct {
	Tpl           string
	Rsp           string
	Path          string
	ServerMode    string
	AdminAddr     string
	AdminName     string
	AdminPassword string
	HttpPort      string
	IsSSL         string
	SSLPort       string
}
