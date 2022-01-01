package core

import (
	"errors"
	"time"

	"github.com/golang/protobuf/proto"
	core "github.com/tommzn/hdb-core"
	events "github.com/tommzn/hdb-events-go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type dataSourceMock struct {
	returnNil, returnError bool
}

type rendererMock struct {
	template   Template
	dataSource DataSource
	anchor     Point
}

type mockData struct {
	Anchor Point
	Event  interface{}
}

func newDataSourceMock(returnNil, returnError bool) DataSource {
	return &dataSourceMock{returnNil: returnNil, returnError: returnError}
}

func newRendererMock(template Template, dataSource DataSource, anchor Point) Renderer {
	return &rendererMock{template: template, dataSource: dataSource, anchor: anchor}
}

func (mock *dataSourceMock) Latest(datasource core.DataSource) (proto.Message, error) {

	if mock.returnNil {
		return nil, nil
	}
	if mock.returnError {
		return nil, errors.New("Error occured.")
	}
	return &events.ExchangeRate{
		FromCurrency: "USD",
		ToCurrency:   "EUR",
		Rate:         1.23445,
		Timestamp:    timestamppb.New(time.Now()),
	}, nil
}

func (mock *dataSourceMock) All(datasource core.DataSource) ([]proto.Message, error) {

	if mock.returnNil {
		return nil, nil
	}
	if mock.returnError {
		return nil, errors.New("Error occured.")
	}
	return []proto.Message{
		&events.ExchangeRate{
			FromCurrency: "USD",
			ToCurrency:   "EUR",
			Rate:         0.8765,
			Timestamp:    timestamppb.New(time.Now()),
		},
		&events.ExchangeRate{
			FromCurrency: "EUR",
			ToCurrency:   "USD",
			Rate:         1.23445,
			Timestamp:    timestamppb.New(time.Now()),
		}}, nil
}

func (mock *dataSourceMock) Observe(filter []core.DataSource) <-chan proto.Message {

	eventChan := make(chan proto.Message, 3)
	eventChan <- &events.ExchangeRate{
		FromCurrency: "USD",
		ToCurrency:   "EUR",
		Rate:         0.8765,
		Timestamp:    timestamppb.New(time.Now()),
	}
	eventChan <- &events.ExchangeRate{
		FromCurrency: "EUR",
		ToCurrency:   "USD",
		Rate:         1.23445,
		Timestamp:    timestamppb.New(time.Now()),
	}
	return eventChan
}

func (renderer *rendererMock) Size() Size {
	return Size{Height: 0, Width: 0}
}

func (renderer *rendererMock) Content() (string, error) {

	event, err := renderer.dataSource.Latest(core.DATASOURCE_EXCHANGERATE)
	if err != nil {
		return "", err
	}
	if event == nil {
		return "", errors.New("Missing event data.")
	}
	data := mockData{
		Anchor: renderer.anchor,
		Event:  event,
	}
	return renderer.template.RenderWith(data)
}
