package domain

import "time"

type Condition struct {
	Type               string
	Status             ConditionStatus
	Reason             string
	Message            string
	LastTransitionTime time.Time
}

type ConditionStatus string

const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)
