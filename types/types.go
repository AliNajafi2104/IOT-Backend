package types

import "time"

type Status_mode string

const (
	Operational Status_mode = "operational"
	Pairing     Status_mode = "pairing"
	Ota         Status_mode = "ota"
	Error       Status_mode = "error"
)

type Node struct {
	Id             string      `bson:"_id,omitempty" json:"id"`
	Light_id       string      `bson:"light_id" json:"light_id"`
	Status_mode    Status_mode `bson:"status_mode" json:"status_mode"`
	Avg_r          int         `bson:"avg_r" json:"avg_r"`
	Temp_c         float32     `bson:"temp_c" json:"temp_c"`
	Vbat_mv        int         `bson:"vbat_mv" json:"vbat_mv"`
	Fw_version     string      `bson:"fw_version" json:"fw_version"`
	Last_seen      time.Time   `bson:"last_seen" json:"last_seen"`
	Site_id        string      `bson:"site_id" json:"site_id"`
	Coordinator_id string      `bson:"coordinator_id" json:"coordinator_id"`
}

type Coordinator struct {
	Id                string    `bson:"_id,omitempty" json:"id"`
	Site_id           string    `bson:"site_id" json:"site_id"`
	Fw_version        string    `bson:"fw_version" json:"fw_version"`
	Nodes_online      int       `bson:"nodes_online" json:"nodes_online"`
	Wifi_rssi         int       `bson:"wifi_rssi" json:"wifi_rssi"`
	Mmwave_event_rate float32   `bson:"mmwave_event_rate" json:"mmwave_event_rate"`
	Last_seen         time.Time `bson:"last_seen" json:"last_seen"`
}

type Site struct {
	Id       string `bson:"_id,omitempty" json:"id"`
	Name     string `bson:"name" json:"name"`
	Location string `bson:"location" json:"location"`
	Config   string `bson:"config" json:"config"`
}

type Telemetry struct {
	Ts        time.Time `bson:"ts" json:"ts"`
	Device_id string    `bson:"device_id" json:"device_id"`
	Payload   string    `bson:"payload" json:"payload"`
}

type Target_type string

const (
	NodeTarget        Target_type = "node"
	CoordinatorTarget Target_type = "coordinator"
)

type Status string

const (
	OTAStatusPending    Status = "pending"
	OTAStatusInProgress Status = "in_progress"
	OTAStatusSuccess    Status = "success"
	OTAStatusRollback   Status = "rollback"
	OTAStatusFailed     Status = "failed"
)

type OTAJob struct {
	Id             string      `bson:"_id,omitempty" json:"id"`
	Target_type    Target_type `bson:"target_type" json:"target_type"`
	Target_version string      `bson:"target_version" json:"target_version"`
	Status         Status      `bson:"status" json:"status"`
	Created_at     time.Time   `bson:"created_at" json:"created_at"`
	Updated_at     time.Time   `bson:"updated_at" json:"updated_at"`
}
