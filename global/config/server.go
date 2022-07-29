package config

type ServerConfig struct {
	S  Run     `mapstructure:"run"`
	SW Swagger `mapstructure:"swagger"`
	L  Log     `mapstructure:"log"`
	T  Task    `mapstructure:"task"`
	M  []Mysql `mapstructure:"mysql"`
	R  []Redis `mapstructure:"redis"`
}

type Run struct {
	Name          string `mapstructure:"name"`
	Env           string `mapstructure:"env"`
	RunMode       string `mapstructure:"runMode"`
	SystemVersion string `mapstructure:"systemVersion"`
	Ip            string `mapstructure:"ip"`
	Port          string `mapstructure:"port"`
}

type Swagger struct {
	Title    string `mapstructure:"title"`
	Desc     string `mapstructure:"desc"`
	Host     string `mapstructure:"host"`
	BasePath string `mapstructure:"base_path"`
	Version  string `mapstructure:"version"`
}

type Log struct {
	AppName    string `mapstructure:"appName"`
	LogFileDir string `mapstructure:"logFileDir"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
	Dev        bool   `mapstructure:"dev"`
}

type Mysql struct {
	Name      string `mapstructure:"name"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Ip        string `mapstructure:"ip"`
	Port      int    `mapstructure:"port"`
	DB        string `mapstructure:"db"`
	Parameter string `mapstructure:"parameter"`
}

type Redis struct {
	Name     string `mapstructure:"name"`
	Ip       string `mapstructure:"ip"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Task struct {
	Task1 `mapstructure:"task1"`
	Task2 `mapstructure:"task2"`
	Task3 `mapstructure:"task3"`
	Task4 `mapstructure:"task4"`
	Task5 `mapstructure:"task5"`
}

type Task1 struct {
	Interval int64 `mapstructure:"interval"`
}

type Task2 struct {
	Interval int64 `mapstructure:"interval"`
}

type Task3 struct {
	Interval int64 `mapstructure:"interval"`
}

type Task4 struct {
	Interval int64 `mapstructure:"interval"`
}

type Task5 struct {
	Interval int64 `mapstructure:"interval"`
}
