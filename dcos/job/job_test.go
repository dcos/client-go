package job

import (
	"net/http"
	"testing"

	"github.com/mesosphere/dcos-go/dcos/config"
	"github.com/stretchr/testify/assert"
)

func TestNewJobService(t *testing.T) {
	testurl := "http://foobar.dcos.io"
	testtoken := "token"
	conf := &config.Config{
		Core: &config.Core{
			DcosUrl:      testurl,
			DcosAcsToken: testtoken,
		},
	}
	j := NewJobService(conf, http.DefaultClient)
	assert.Equal(t, "http://foobar.dcos.io/service/metronome", j.GetServiceURL("").String())
	assert.Equal(t, "http://foobar.dcos.io/service/metronome/foobar", j.GetServiceURL("foobar").String())

}
