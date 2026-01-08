package domain

import (
	"time"
)

type Instance struct {
	Name   string
	Spec   InstanceSpec
	Status InstanceStatus
}

type InstanceSpec struct {
	MachineType  string
	Region       string
	DiskSizeGB   int
	DesiredState InstanceDesiredState
}

type InstanceDesiredState string

const (
	DesiredRunning InstanceDesiredState = "RUNNING"
	DesiredStopped InstanceDesiredState = "STOPPED"
	DesiredDeleted InstanceDesiredState = "DELETED"
)

type InstanceStatus struct {
	State          InstanceState
	ProviderID     string
	Conditions     []Condition
	LastError      string
	LastUpdateTime time.Time
}

type InstanceState string

const (
	StateNone         InstanceState = "NONE"
	StateProvisioning InstanceState = "PROVISIONING"
	StateRunning      InstanceState = "RUNNING"
	StateStopping     InstanceState = "STOPPING"
	StateStopped      InstanceState = "STOPPED"
	StateDeleting     InstanceState = "DELETING"
	StateError        InstanceState = "ERROR"
)
