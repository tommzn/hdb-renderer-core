package core

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TemplateTestSuite struct {
	suite.Suite
}

func TestTemplateTestSuite(t *testing.T) {
	suite.Run(t, new(TemplateTestSuite))
}

func (suite *TemplateTestSuite) TestRenderContent() {

	template := NewFileTemplate("fixtures/test_template_02.json")
	data := mockData{Anchor: Point{X: 100, Y: 150}}

	content, err := template.RenderWith(data)
	suite.Nil(err)
	suite.NotEqual("", content)
	assertTemplateHash(suite.Assert(), content, "3e232a9830f3baa20bfd0fd89f50103ace53049c")
}
