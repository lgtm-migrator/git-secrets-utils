package lib

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name string
		want string
	}{
		{
			"basic test",
			"hello golang",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Value(); got != tt.want {
				assert.Equal(tt.want, got)
			}
		})
	}
}

func TestMatchLineString(t *testing.T) {
	type args struct {
		lineStr string
	}
	tests := []struct {
		name string
		args args
		want *LineErrorInformation
	}{
		{
			"basic match test",
			args{lineStr: `../blabla-lib-common/blabla-event/blabla-event.iml:153:    <orderEntry type="library" scope="PROVIDED" name="Maven: com.sap.cloud.security.xsuaa:token-client:2.7.8" level="project" />`},
			&LineErrorInformation{
				Line:    "153",
				File:    "../blabla-lib-common/blabla-event/blabla-event.iml",
				Content: `<orderEntry type="library" scope="PROVIDED" name="Maven: com.sap.cloud.security.xsuaa:token-client:2.7.8" level="project" />`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchLineString(tt.args.lineStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchLineString() = %v, want %v", got, tt.want)
			}
		})
	}
}
