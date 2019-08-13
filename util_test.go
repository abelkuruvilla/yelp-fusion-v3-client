package yelp

import "testing"

func Test_structToQueryParams(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Business Search Params",
			args: args{
				s: &BusinessSearchParams{
					Term:  "assisted-living",
					Limit: 10,
				},
			},
			want:    "?term=assisted-living&limit=10",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := structToQueryParams(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("structToQueryParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("structToQueryParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
