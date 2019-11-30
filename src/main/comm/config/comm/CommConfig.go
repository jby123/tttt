package comm

type Conf struct {
	Notify IsOpenNotify `yaml:"Notify"`
	Valid  ValidConfig  `yaml:"Valid"`
}

type ValidConfig struct {
	IsLoginValidCaptcha bool `yaml:"IsLoginValidCaptcha"`
	IsOpenAuthorization bool `yaml:"IsOpenAuthorization"`
}

type IsOpenNotify struct {
	ErrorIsOpen bool              `yaml:"ErrorIsOpen"`
	EmailConfig EmailNotifyConfig `yaml:"Email"`
}

type EmailNotifyConfig struct {
	FromEmail string `yaml:"FromEmail"`
	EmailHost string `yaml:"EmailHost"`
	EmailPort int    `yaml:"EmailPort"`
	EmailPass string `yaml:"EmailPass"`
}
