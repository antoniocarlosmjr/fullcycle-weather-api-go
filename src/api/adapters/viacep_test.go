package adapters

import (
	"reflect"
	"testing"

	"github.com/antoniocarlosmjr/weather-api-go/src/api/models/entities"
)

func TestSearchZipCode(t *testing.T) {
	type args struct {
		zipCode string
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.ZipCode
		wantErr bool
	}{
		{
			name: "Test SearchZipCode Success",
			args: args{
				zipCode: "49045700",
			},
			want: &entities.ZipCode{
				Cep:          "49045-700",
				Log:          "Estrada da Luzia",
				Neighborhood: "Luzia",
				Locale:       "Aracaju",
				UF:           "SE",
				IBGE:         "2800308",
				DDD:          "79",
			},
			wantErr: false,
		},
		{
			name: "Test SearchZipCode When Error Invalid Zip Code",
			args: args{
				zipCode: "12345678",
			},
			want:    &entities.ZipCode{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchZipCode(tt.args.zipCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchZipCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchZipCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
