package services

import (
	"reflect"
	"testing"

	"github.com/antoniocarlosmjr/weather-api-go/src/api/models/dto"
)

func TestSearchWeather(t *testing.T) {
	type args struct {
		zipCode string
	}
	tests := []struct {
		name    string
		args    args
		want    *dto.WeatherResponse
		wantErr bool
	}{
		{
			name: "Test SearchWeather Success",
			args: args{
				zipCode: "49045700",
			},
			want: &dto.WeatherResponse{
				TempC: 31,
				TempF: 87.80000000000001,
				TempK: 304,
			},
			wantErr: false,
		},
		{
			name: "Test SearchWeather Not Found Zip Code",
			args: args{
				zipCode: "49045707",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchWeather(tt.args.zipCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchWeather() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchWeather() got = %v, want %v", got, tt.want)
			}
		})
	}
}
