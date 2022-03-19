package core

import (
	"github.com/stretchr/testify/suite"
	events "github.com/tommzn/hdb-events-go"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

type TimestampManagerTestSuite struct {
	suite.Suite
}

func TestTimestampManagerTestSuite(t *testing.T) {
	suite.Run(t, new(TimestampManagerTestSuite))
}

func (suite *TimestampManagerTestSuite) TestCheckIsLatest() {

	mgr := NewTimestampManager()
	event := &events.ExchangeRate{Timestamp: timestamppb.Now()}
	olderEvent := &events.ExchangeRate{Timestamp: timestamppb.New(time.Now().Add(-1 * time.Minute))}
	newerEvent := &events.ExchangeRate{Timestamp: timestamppb.New(time.Now().Add(1 * time.Minute))}

	suite.True(mgr.IsLatest(event))
	mgr.Add(event)
	suite.False(mgr.IsLatest(olderEvent))
	suite.True(mgr.IsLatest(newerEvent))

	suffix := "XxX"
	suite.True(mgr.IsLatestWithSuffix(event, suffix))
	mgr.AddWithSuffix(event, suffix)
	suite.False(mgr.IsLatestWithSuffix(olderEvent, suffix))
	suite.True(mgr.IsLatestWithSuffix(newerEvent, suffix))

}
