package dto

type CollectorDto struct {
	Name                 string `binding:"required,min=6,max=12"`
	Description          string
	PlanId               int `binding:"required"`
	ProjectId            int `binding:"required"`
	ProjectCode          string
	EnvId                int `binding:"required"`
	EnvName              string
	Cluster              string `binding:"required"`
	TokenUrl             string
	AppId                int
	TokenRefreshInterval int
	MetricUrl            string
	AlertUrl             string
	HeartBeatUrl         string
	HeartBeatInterval    int
	Enabled              int
	ExecCnt              int
	AlertCnt             int
}
