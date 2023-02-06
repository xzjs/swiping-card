package model

type DB struct {
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	DBName string `yaml:"dbname"`
}

type AES struct {
	Key string `yaml:"key"`
}

type Cookie struct {
	Name   string `yaml:"name"`
	Domain string `yaml:"domain"`
}

type Wechat struct {
	Appid  string `yaml:"appid"`
	Secret string `yaml:"secret"`
}

type Conf struct {
	DB     DB     `yaml:"db"`
	AES    AES    `yaml:"aes"`
	Cookie Cookie `yaml:"cookie"`
	Wechat Wechat `yaml:"wechat"`
	QiNiu  QiNiu  `yaml:"qiniu"`
}

type QiNiu struct {
	AK     string `yaml:"ak"`
	SK     string `yaml:"sk"`
	Bucket string `yaml:"bucket"`
	Domain string `yaml:"domain"`
}
