package core

import (
	"errors"
	"time"

	events "github.com/tommzn/hdb-events-go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type eventSourceMock struct {
	returnNil, returnError bool
}

func newEventSourceMock(returnNil, returnError bool) DataSource {
	return &eventSourceMock{returnNil: returnNil, returnError: returnError}
}

func (mock *eventSourceMock) Get() (interface{}, error) {

	if mock.returnNil {
		return nil, nil
	}
	if mock.returnError {
		return nil, errors.New("Error occured.")
	}
	return events.ExchangeRate{
		FromCurrency: "USD",
		ToCurrency:   "EUR",
		Rate:         1.23445,
		Timestamp:    timestamppb.New(time.Now()),
	}, nil
}
