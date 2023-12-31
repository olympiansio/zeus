package types

import (
	"github.com/huweihuang/zeus/pkg/constant"
)

// Instance represents the configuration of a single Instance.
type Instance struct {
	// Standard object's metadata.
	InstanceMeta `json:"metadata"`
	// Specification of the desired behavior of a Instance.
	Spec InstanceSpec `gorm:"embedded" json:"spec"`
	// Current status of a Instance
	Status InstanceStatus `gorm:"embedded" json:"status"`
}

// InstanceMeta specified identity of Instance
type InstanceMeta struct {
	JobID     string `gorm:"column:job_id" json:"jobID,omitempty"`
	Name      string `gorm:"column:name" json:"instanceName" binding:"required"`
	Namespace string `gorm:"column:namespace" json:"namespace,omitempty" binding:"required"`
}

// InstanceSpec is the specification of  Instance.
type InstanceSpec struct {
	Image       string   `gorm:"column:image" json:"image,omitempty" binding:"required"`
	HostID      string   `gorm:"column:host_id" json:"hostID,omitempty"`
	Servers     []string `json:"servers" binding:"required,gte=1,lte=30,unique"`
	Replicas    int      `gorm:"column:replicas" json:"replicas,omitempty"`
	KeepStorage bool     `gorm:"-" json:"keepStorage,omitempty"`
}

// InstanceStatus represents the current state of Instance.
type InstanceStatus struct {
	Status      bool              `gorm:"column:status" json:"status"`
	JobState    constant.JobState `gorm:"column:job_state" json:"jobState,omitempty"`
	CreatedTime string            `gorm:"column:created_time;->" json:"createdTime,omitempty"`
	ModifyTime  string            `gorm:"column:modify_time;->" json:"modifyTime,omitempty"`
}

func (t *Instance) TableName() string {
	return "instance_tab"
}
