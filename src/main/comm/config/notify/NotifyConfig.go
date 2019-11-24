package notify

type Conf struct {
	Notify      IsOpenNotify      `yaml:"Notify"`
	EmailConfig EmailNotifyConfig `yaml:"Email"`
}

type IsOpenNotify struct {
	ErrorIsOpen bool `yaml:"ErrorIsOpen"`
}

type EmailNotifyConfig struct {
	FromEmail string `yaml:"FromEmail"`
	EmailHost string `yaml:"EmailHost"`
	EmailPort int    `yaml:"EmailPort"`
	EmailPass string `yaml:"EmailPass"`
}
