package yelp

import (
	"testing"
)

func TestClient_Search(t *testing.T) {

	client, err := NewClient()
	if err != nil {
		t.Errorf("Error retrieving client, %v", err)
		return
	}
	client.Options.Debug = true

	results, err := client.Search(BusinessSearchParams{
		Term:     "bayside-park-emeryville-3",
		Location: "san fransico",
	})

	if err != nil {
		t.Errorf("Got Error on search %v", err.Error())
		return
	}

	if len(results.Businesses) > 10 {
		t.Errorf("Length of Results should be less than or equal to 50 : %v", len(results.Businesses))
		return
	}

}

func TestClient_GetBusinessDetails(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Errorf("Error retrieving client, %v", err)
		return
	}
	client.Options.Debug = true

	type args struct {
		id      string
		request BusinessDetailParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Buena Vista Manor House",
			args: args{
				id: "Ev92faQVAEx0_46IwC6W0A",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := client.GetBusinessDetails(tt.args.id, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetBusinessDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse.Name != tt.name {
				t.Errorf("Response Name is not the same. Wanted : %s, Got: %s", tt.name, gotResponse.Name)
			}

		})
	}
}
