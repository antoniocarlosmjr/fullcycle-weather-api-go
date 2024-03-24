package adapters

import (
	"reflect"
	"testing"

	"github.com/antoniocarlosmjr/weather-api-go/src/api/models/entities"
)

func TestGetWeather(t *testing.T) {
	type args struct {
		locale string
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.Weather
		wantErr bool
	}{
		{
			name: "Test GetWeather Error Invalid Locale",
			args: args{
				locale: "",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWeather(tt.args.locale)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWeather() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWeather() got = %v, want %v", got, tt.want)
			}
		})
	}
}
