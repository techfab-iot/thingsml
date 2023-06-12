package thingsml_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/techfab-iot/thingsml"
)

func TestNormalize(t *testing.T) {
	temperatureExample := 23.76
	humidityExample := 47.25

	pack := thingsml.Pack{
		{
			Name:  "urn:dev:mac:abcd1234:temperature",
			Time:  1686429251,
			Unit:  "Cel",
			Value: &temperatureExample,
		},
		{
			Name:  "urn:dev:mac:abcd1234:humidity",
			Time:  1686429251,
			Unit:  "%RH",
			Value: &humidityExample,
		},
	}

	tests := []struct {
		name    string
		payload []byte
		want    thingsml.Pack
	}{
		{
			name: "Normalized JSON SenML",
			payload: []byte(`[{
				"n": "urn:dev:mac:abcd1234:temperature",
				"t": 1686429251,
				"u": "Cel",
				"v": 23.76
			}, {
				"n": "urn:dev:mac:abcd1234:humidity",
				"t": 1686429251,
				"u": "%RH",
				"v": 47.25
			}]`),
			want: pack,
		},
		{
			name: "Denormalized JSON SenML",
			payload: []byte(`[{
				"bn": "urn:dev:mac:abcd1234:",
				"bt": 1686429250,
				"n": "temperature",
				"u": "Cel",
				"v": 23.76,
				"t": 1
			}, {
				"n": "humidity",
				"u": "%RH",
				"v": 47.25,
				"t": 1
			}]`),
			want: pack,
		},
		{
			name: "Denormalized JSON ThingsML",
			payload: []byte(`[{
				"bn": "urn:dev:mac:abcd1234:",
				"bt": 1686429251,
				"i_": -24,
				"v": 23.76
			}, {
				"i_": -23,
				"v": 47.25
			}]`),
			want: pack,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf(tt.name)
		t.Run(testname, func(t *testing.T) {
			pack, err := thingsml.NormalizeJSON(tt.payload)
			if err != nil {
				t.Fatalf("Fatal: %v", err)
			}

			if !cmp.Equal(pack, tt.want) {
				t.Errorf("Got %v, want %v", pack, tt.want)
			}
		})
	}

	tests = []struct {
		name    string
		payload []byte
		want    thingsml.Pack
	}{
		{
			name: "Normalized CBOR SenML",
			payload: []byte{
				0x82, 0xa4, 0x00, 0x78, 0x20, 0x75, 0x72, 0x6e, 0x3a, 0x64, 0x65, 0x76,
				0x3a, 0x6d, 0x61, 0x63, 0x3a, 0x61, 0x62, 0x63, 0x64, 0x31, 0x32, 0x33,
				0x34, 0x3a, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
				0x65, 0x06, 0x1a, 0x64, 0x84, 0xde, 0x43, 0x01, 0x63, 0x43, 0x65, 0x6c,
				0x02, 0xfb, 0x40, 0x37, 0xc2, 0x8f, 0x5c, 0x28, 0xf5, 0xc3, 0xa4, 0x00,
				0x78, 0x1d, 0x75, 0x72, 0x6e, 0x3a, 0x64, 0x65, 0x76, 0x3a, 0x6d, 0x61,
				0x63, 0x3a, 0x61, 0x62, 0x63, 0x64, 0x31, 0x32, 0x33, 0x34, 0x3a, 0x68,
				0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x06, 0x1a, 0x64, 0x84, 0xde,
				0x43, 0x01, 0x63, 0x25, 0x52, 0x48, 0x02, 0xfb, 0x40, 0x47, 0xa0, 0x00,
				0x00, 0x00, 0x00, 0x00,
			},
			want: pack,
		},
		{
			name: "Denormalized CBOR SenML",
			payload: []byte{
				0x82, 0xa6, 0x21, 0x75, 0x75, 0x72, 0x6e, 0x3a, 0x64, 0x65, 0x76, 0x3a,
				0x6d, 0x61, 0x63, 0x3a, 0x61, 0x62, 0x63, 0x64, 0x31, 0x32, 0x33, 0x34,
				0x3a, 0x22, 0x1a, 0x64, 0x84, 0xde, 0x42, 0x00, 0x6b, 0x74, 0x65, 0x6d,
				0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x01, 0x63, 0x43, 0x65,
				0x6c, 0x02, 0xfb, 0x40, 0x37, 0xc2, 0x8f, 0x5c, 0x28, 0xf5, 0xc3, 0x06,
				0x01, 0xa4, 0x00, 0x68, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79,
				0x01, 0x63, 0x25, 0x52, 0x48, 0x02, 0xfb, 0x40, 0x47, 0xa0, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x06, 0x01,
			},
			want: pack,
		},
		{
			name: "Denormalized CBOR ThingsML",
			payload: []byte{
				0x82, 0xa5, 0x21, 0x75, 0x75, 0x72, 0x6e, 0x3a, 0x64, 0x65, 0x76, 0x3a,
				0x6d, 0x61, 0x63, 0x3a, 0x61, 0x62, 0x63, 0x64, 0x31, 0x32, 0x33, 0x34,
				0x3a, 0x22, 0x1a, 0x64, 0x84, 0xde, 0x42, 0x02, 0xfb, 0x40, 0x37, 0xc2,
				0x8f, 0x5c, 0x28, 0xf5, 0xc3, 0x06, 0x01, 0x17, 0x37, 0xa3, 0x02, 0xfb,
				0x40, 0x47, 0xa0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x01, 0x17, 0x36,
			},
			want: pack,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf(tt.name)
		t.Run(testname, func(t *testing.T) {
			pack, err := thingsml.NormalizeCBOR([]byte(tt.payload))
			if err != nil {
				t.Fatalf("Fatal: %v", err)
			}

			if !cmp.Equal(pack, tt.want) {
				t.Errorf("Got %v, want %v", pack, tt.want)
			}
		})
	}
}

func TestNormalizeErrors(t *testing.T) {
	tests := []struct {
		name    string
		payload []byte
		wantErr error
	}{
		{
			name:    "Invalid JSON",
			payload: []byte(`{`),
			wantErr: thingsml.ErrUnmarshal,
		},
		{
			name: "Missing name",
			payload: []byte(`
				[
					{
						"t": 1686429251,
						"u": "Cel",
						"v": 23.76
					},
					{
						"n": "humidity",
						"t": 1686429251,
						"u": "%RH",
						"v": 47.25
					}
				]
			`),
			wantErr: thingsml.ErrEmptyName,
		},
		{
			name: "Missing value",
			payload: []byte(`
				[
					{
						"bn": "urn:dev:mac:abcd1234:",
						"bt": 1686429251,
						"i_": -24
					},
					{
						"i_": -23
					}
				]
			`),
			wantErr: thingsml.ErrNoValues,
		},
		{
			name: "Too many values",
			payload: []byte(`
				[
					{
						"bn": "urn:dev:mac:abcd1234:",
						"bt": 1686429251,
						"i_": -24,
						"v": 23.76
					},
					{
						"i_": -23,
						"v": 47.25,
						"vb": true
					}
				]
			`),
			wantErr: thingsml.ErrTooManyValues,
		},
		{
			name: "Bad characters",
			payload: []byte(`
				[
					{
						"bn": "urn:dev:mac:abcd1234:",
						"bt": 1686429251,
						"i_": -24,
						"v": 23.76
					},
					{
						"n": "rotação",
						"v": 3.1415,
						"u": "rad"
					}
				]
			`),
			wantErr: thingsml.ErrBadChar,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf(tt.name)
		t.Run(testname, func(t *testing.T) {
			_, err := thingsml.NormalizeJSON([]byte(tt.payload))
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Expected error %v, got %v", tt.wantErr, err)
			}
		})
	}

	tests = []struct {
		name    string
		payload []byte
		wantErr error
	}{
		{
			name:    "Invalid CBOR",
			payload: []byte{0x01, 0x02, 0x03, 0x04},
			wantErr: thingsml.ErrUnmarshal,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf(tt.name)
		t.Run(testname, func(t *testing.T) {
			_, err := thingsml.NormalizeCBOR([]byte(tt.payload))
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}
