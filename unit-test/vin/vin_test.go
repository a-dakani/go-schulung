package vin

import "testing"

func Test_parse_lengthCheck(t *testing.T) {
	vin := "123456789012345678"
	_, err := FromString(vin)
	if err == nil {
		t.Error("Expected error")
	}
}

func Test_parse_region(t *testing.T) {
	vin := "12345678901234567"
	v, err := FromString(vin)
	if err != nil {
		t.Error("Expected no error")
	}
	if v.region != "123" {
		t.Error("Expected region to be 123")
	}
}

func Test_table_parse(t *testing.T) {
	tests := []struct {
		name           string
		vin            string
		wantErr        bool
		expectedRegion string
	}{
		{
			name:           "valid vin",
			vin:            "12345678901234567",
			wantErr:        false,
			expectedRegion: "123",
		},
		{
			name:           "vin too long",
			vin:            "123456789012345678",
			wantErr:        true,
			expectedRegion: "123",
		},
		{
			name:           "vin too short",
			vin:            "1234567890123456",
			wantErr:        true,
			expectedRegion: "123",
		},
		{
			name:           "vin empty",
			vin:            "",
			wantErr:        true,
			expectedRegion: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := FromString(tt.vin)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if v != nil && v.region != tt.expectedRegion {
				t.Errorf("parse() region = %v, expectedRegion %v", v.region, tt.expectedRegion)
			}
		})
	}
}
