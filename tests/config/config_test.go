package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	configv1 "github.com/go-goim/api/config/v1"
)

func Test_Config_Validate(t *testing.T) {
	type args struct {
		config *configv1.Service
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "validate config",
			args: args{
				config: &configv1.Service{
					Name:    "goim.service.test",
					Version: "v0.0.1",
				},
			},
			wantErr: false,
		},
		{
			name: "validate config with empty name",
			args: args{
				config: &configv1.Service{
					Name:    "",
					Version: "v0.0.1",
				},
			},
			wantErr: true,
		},
		{
			name: "validate config with invalid name",
			args: args{
				config: &configv1.Service{
					Name:    "goim.xxx.test",
					Version: "v0.0.1",
				},
			},
			wantErr: true,
		},
		{
			name: "validate config with empty version",
			args: args{
				config: &configv1.Service{
					Name:    "goim.service.test",
					Version: "",
				},
			},
			wantErr: true,
		},
		{
			name: "validate config with invalid version",
			args: args{
				config: &configv1.Service{
					Name:    "goim.service.test",
					Version: "v0.0.1.0",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.args.config.Validate()
			if tt.wantErr {
				assert.Error(t, gotErr)
			} else {
				assert.NoError(t, gotErr)
			}
		},
		)
	}
}
