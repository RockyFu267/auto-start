package controller

//EnvConf yaml配置的key不能用 _ 否则会信息丢失
type EnvConf struct {
	Interval int64         `json:"interval"`
	Service  []ServiceInfo `json:"service"`
}

type ServiceInfo struct {
	Port      string `json:"port"`
	Startpath string `json:"startpath"`
}
