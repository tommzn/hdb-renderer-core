package core

import (
	"crypto/sha1"
	"encoding/hex"
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

	template := NewEinkTemplate("fixtures/test_template_01.json", newEventSourceMock(false, false))
	template.SetAnchor(Point{X: 100, Y: 200})

	canvas, err := template.Render()
	suite.Nil(err)
	suite.NotEqual("", canvas)
	suite.assertTemplateHash(canvas, "95d70c7d72bbc2e227014571839e88adc0766259")

	suite.Equal(Size{Height: 0, Width: 0}, template.GetSize())
}

func (suite *TemplateTestSuite) TestRenderWithErrors() {

	template1 := NewEinkTemplate("fixtures/test_template_01.json", newEventSourceMock(true, false))
	canvas1, err1 := template1.Render()
	suite.NotNil(err1)
	suite.Equal("", canvas1)

	template2 := NewEinkTemplate("fixtures/test_template_01.json", newEventSourceMock(false, true))
	canvas2, err2 := template2.Render()
	suite.NotNil(err2)
	suite.Equal("", canvas2)
}

func (suite *TemplateTestSuite) assertTemplateHash(template string, expectedHash string) {
	hash := sha1.New()
	hash.Write([]byte(template))
	suite.Equal(expectedHash, hex.EncodeToString(hash.Sum(nil)))
}
