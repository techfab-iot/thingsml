# ThingsML
ThingsML implementation in Go

[ThingsML](https://docs.kpnthings.com/dm/processing/thingsml) is a more compact, fully compatible superset of SenML, best suited for limited connections like LoRaWAN and for reducing data transfer costs.

## Supported formats
- JSON
- CBOR

## Installing
``` sh
go get github.com/techfab-iot/thingsml@latest
```

## Normalization
The main goal of this library is to transform the compact ThingsML structure into a more friendly normalized SenML. You can do this by using `thingsml.NormalizeJSON` or `thingsml.NormalizeCBOR`.

### Example with JSON
``` go
msg := []byte(`
	[
		{
			"bn": "urn:dev:mac:ec4c429e392c:",
			"bt": 1686429251,
			"i_": -24,
			"v": 23.76
		},
		{
			"i_": -23,
			"v": 47.25
		}
	]
`)

records, err := thingsml.NormalizeJSON(msg)
```

### Example with CBOR
``` go
msg := []byte{
	0x82, 0xa4, 0x21, 0x78, 0x19, 0x75, 0x72, 0x6e, 0x3a, 0x64, 0x65, 0x76,
	0x3a, 0x6d, 0x61, 0x63, 0x3a, 0x65, 0x63, 0x34, 0x63, 0x34, 0x32, 0x39,
	0x65, 0x33, 0x39, 0x32, 0x63, 0x3a, 0x22, 0x1a, 0x64, 0x84, 0xde, 0x43,
	0x17, 0x37, 0x02, 0xfb, 0x40, 0x37, 0xc2, 0x8f, 0x5c, 0x28, 0xf5, 0xc3,
	0xa2, 0x17, 0x36, 0x02, 0xfb, 0x40, 0x47, 0xa0, 0x00, 0x00, 0x00, 0x00,
	0x00,
}

records, err := thingsml.NormalizeCBOR(msg)
```
