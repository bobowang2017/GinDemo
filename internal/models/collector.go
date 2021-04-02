package models

type Collector struct {
	BaseModel
	Name                 string `gorm:"size:45;not null" json:"name" `
	Description          string `gorm:"size:300;" json:"description" `
	PlanId               *int   `gorm:"not null;" json:"planId"`
	ProjectId            *int   `gorm:"not null" json:"projectId"`
	ProjectCode          string `gorm:"size:45;" json:"projectCode"`
	EnvId                *int   `gorm:"not null;" json:"envId"`
	EnvName              string `gorm:"size:45;" json:"envName"`
	Cluster              string `gorm:"size:45;" json:"cluster"`
	TokenUrl             string `gorm:"size:200;" json:"tokenUrl"`
	AppId                *int   `json:"appId"`
	TokenRefreshInterval *int   `json:"tokenRefreshInterval"`
	MetricUrl            string `gorm:"size:200;" json:"metricUrl"`
	AlertUrl             string `gorm:"size:200;" json:"alertUrl"`
	HeartBeatUrl         string `gorm:"size:200;" json:"heartBeatUrl"`
	HeartBeatInterval    *int   `json:"heartBeatInterval"`
	Enabled              *int   `gorm:"default:0" json:"enabled"`
	ExecCnt              *int   `json:"execCnt"`
	AlertCnt             *int   `json:"alertCnt"`
}
