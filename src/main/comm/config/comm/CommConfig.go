package comm

type Conf struct {
	Notify     IsOpenNotify     `yaml:"Notify"`
	LoginValid LoginValidConfig `yaml:"LoginValid"`
}

type LoginValidConfig struct {
	IsValidCaptcha bool `yaml:"IsValidCaptcha"`
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
