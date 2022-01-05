package core

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type MockTestSuite struct {
	suite.Suite
}

func TestMockTestSuite(t *testing.T) {
	suite.Run(t, new(MockTestSuite))
}

func (suite *MockTestSuite) TestRenderContent() {

	template := NewFileTemplate("fixtures/test_template_02.json")
	datasource := newDataSourceMock(false, false)
	anchor := Point{X: 100, Y: 200}
	renderer := newRendererMock(template, datasource, anchor)

	content, err := renderer.Content()
	suite.Nil(err)
	suite.NotEqual("", content)
	assertTemplateHash(suite.Assert(), content, "d76b69b0f79044e8f3df3222b1a127bf57bbd472")
}

func (suite *MockTestSuite) TestRenderContentWithError() {

	template := NewFileTemplate("fixtures/test_template_02.json")
	anchor := Point{X: 100, Y: 200}
	renderer := newRendererMock(template, newDataSourceMock(true, false), anchor)

	content, err := renderer.Content()
	suite.NotNil(err)
	suite.Equal("", content)

	renderer2 := newRendererMock(template, newDataSourceMock(false, true), anchor)

	content2, err2 := renderer2.Content()
	suite.NotNil(err2)
	suite.Equal("", content2)
}
