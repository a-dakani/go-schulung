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
	vin := "1HGCM82633A123456"
	v, err := FromString(vin)
	if err != nil {
		t.Error("Expected no error")
	}
	if v.region != "United States" {
		t.Error("Expected region to be United States")
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
			vin:            "2HGFC2F58LH123456",
			wantErr:        false,
			expectedRegion: "Canada",
		},
		{
			name:           "invalid region ",
			vin:            "IBAUN73538VF12345",
			wantErr:        true,
			expectedRegion: "Canada",
		},
		{
			name:           "vin too long",
			vin:            "2HGFC2F58LH1234562",
			wantErr:        true,
			expectedRegion: "Canada",
		},
		{
			name:           "vin too short",
			vin:            "WBAUN73538VF1234",
			wantErr:        true,
			expectedRegion: "Germany",
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
