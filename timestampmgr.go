package core

import (
	"reflect"
	"time"

	"github.com/golang/protobuf/proto"
	events "github.com/tommzn/hdb-events-go"
)

// NewTimestampManager returns an empty timestamp manager.
func NewTimestampManager() TimestampManager {
	return &EventTimestampManager{
		timestamps: make(map[eventTypeName]time.Time),
	}
}

// IsLatest will returns true if there's no similar event in local storage
// or if local timestamp is older than in passed event.
func (mgr *EventTimestampManager) IsLatest(event proto.Message) bool {

	typeName := getTypeName(event)
	eventTimestamo := eventTimestamo(event)
	if currentTimestamo, ok := mgr.timestamps[typeName]; ok {
		return currentTimestamo.Before(eventTimestamo)
	}
	return true
}

// Add timestamp of passed event to local storage for later checks.
func (mgr *EventTimestampManager) Add(event proto.Message) {
	typeName := getTypeName(event)
	mgr.timestamps[typeName] = eventTimestamo(event)
}

// GetTypeName will return name of a type, given by reflect.TypeOf().String.
// In case of indoor climate events type name will contain datasource suffix as well.
func getTypeName(event proto.Message) eventTypeName {

	typeName := reflect.TypeOf(event).String()
	if IndoorClimateData, ok := event.(*events.IndoorClimate); ok {
		typeName = typeName + IndoorClimateData.Type.String()
	}
	return eventTypeName(typeName)
}

// EventTimestamo returns timestamp from passed event.
// If an event doesn't contain a timestamp value, NOW is returned.
func eventTimestamo(event proto.Message) time.Time {

	switch v := event.(type) {
	case *events.IndoorClimate:
		return v.Timestamp.AsTime()
	case *events.WeatherData:
		return v.Current.Timestamp.AsTime()
	case *events.ExchangeRate:
		return v.Timestamp.AsTime()
	default:
		return time.Now()
	}
}
