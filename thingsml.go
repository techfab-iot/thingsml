package thingsml

import (
	"encoding/json"
	"errors"

	"github.com/fxamacker/cbor/v2"
	"github.com/mainflux/senml"
)

var (
	ErrUnmarshal     = errors.New("cannot unmarshal ThingsML")
	ErrEmptyName     = senml.ErrEmptyName
	ErrBadChar       = senml.ErrBadChar
	ErrTooManyValues = senml.ErrTooManyValues
	ErrNoValues      = senml.ErrNoValues
	ErrVersionChange = senml.ErrVersionChange
)

type MeasurementIndex int

const (
	Temperature MeasurementIndex = iota - 24
	Humidity
	Latitude
	Longitude
	Altitude
	Power
	Pressure
	Angle
	Length
	Breadth
	Height
	Weight
	Thickness
	Distance
	Area
	Volume
	Velocity
	ElectricCurrent
	ElectricPotential
	ElectricResistance
	Illuminance
	AccelerationX
	AccelerationY
	AccelerationZ
	Heading
	COConcentration
	CO2Concentration
	Sound
	Frequency
	BatteryLevel
	BatteryVoltage
	Radius
	BatteryLevelLow
	CompassX
	CompassY
	CompassZ
	ReadSwitch
	Presence
	Counter
)

type Record struct {
	BaseName         string            `json:"bn,omitempty" cbor:"-2,keyasint,omitempty"`
	BaseTime         float64           `json:"bt,omitempty" cbor:"-3,keyasint,omitempty"`
	BaseUnit         string            `json:"bu,omitempty" cbor:"-4,keyasint,omitempty"`
	BaseVersion      uint              `json:"bver,omitempty" cbor:"-1,keyasint,omitempty"`
	BaseValue        float64           `json:"bv,omitempty" cbor:"-5,keyasint,omitempty"`
	BaseSum          float64           `json:"bs,omitempty" cbor:"-6,keyasint,omitempty"`
	Name             string            `json:"n,omitempty" cbor:"0,keyasint,omitempty"`
	Unit             string            `json:"u,omitempty" cbor:"1,keyasint,omitempty"`
	Time             float64           `json:"t,omitempty" cbor:"6,keyasint,omitempty"`
	UpdateTime       float64           `json:"ut,omitempty" cbor:"7,keyasint,omitempty"`
	Value            *float64          `json:"v,omitempty" cbor:"2,keyasint,omitempty"`
	StringValue      *string           `json:"vs,omitempty" cbor:"3,keyasint,omitempty"`
	DataValue        *string           `json:"vd,omitempty" cbor:"8,keyasint,omitempty"`
	BoolValue        *bool             `json:"vb,omitempty" cbor:"4,keyasint,omitempty"`
	Sum              *float64          `json:"s,omitempty" cbor:"5,keyasint,omitempty"`
	MeasurementIndex *MeasurementIndex `json:"i_,omitempty" cbor:"23,keyasint,omitempty"`
}

type Pack []Record

func (p *Pack) ToSenML() senml.Pack {
	senmlPack := senml.Pack{
		Records: make([]senml.Record, len(*p)),
	}
	for i, r := range *p {
		name := r.Name
		unit := r.Unit

		if r.MeasurementIndex != nil {
			switch *r.MeasurementIndex {
			case Temperature:
				name = "temperature"
				unit = "Cel"
			case Humidity:
				name = "humidity"
				unit = "%RH"
			case Latitude:
				name = "latitude"
				unit = "lat"
			case Longitude:
				name = "longitude"
				unit = "lon"
			case Altitude:
				name = "altitude"
				unit = "m"
			case Length:
				name = "length"
				unit = "m"
			case Breadth:
				name = "breadth"
				unit = "m"
			case Height:
				name = "height"
				unit = "m"
			case Thickness:
				name = "thickness"
				unit = "m"
			case Distance:
				name = "distance"
				unit = "m"
			case Radius:
				name = "radius"
				unit = "m"
			case Power:
				name = "power"
				unit = "W"
			case Pressure:
				name = "pressure"
				unit = "Pa"
			case Angle:
				name = "angle"
				unit = "rad"
			case Heading:
				name = "heading"
				unit = "rad"
			case Weight:
				name = "weight"
				unit = "kg"
			case Area:
				name = "area"
				unit = "m2"
			case Volume:
				name = "volume"
				unit = "m3"
			case Velocity:
				name = "velocity"
				unit = "m/s"
			case ElectricCurrent:
				name = "electricCurrent"
				unit = "A"
			case ElectricPotential:
				name = "electricPotential"
				unit = "V"
			case BatteryVoltage:
				name = "batteryVoltage"
				unit = "V"
			case ElectricResistance:
				name = "electricResistance"
				unit = "Ohm"
			case Illuminance:
				name = "illuminance"
				unit = "lx"
			case AccelerationX:
				name = "accelerationX"
				unit = "m/s2"
			case AccelerationY:
				name = "accelerationY"
				unit = "m/s2"
			case AccelerationZ:
				name = "accelerationZ"
				unit = "m/s2"
			case COConcentration:
				name = "COConcentration"
				unit = "ppm"
			case CO2Concentration:
				name = "CO2Concentration"
				unit = "ppm"
			case Sound:
				name = "sound"
				unit = "dB"
			case Frequency:
				name = "frequency"
				unit = "Hz"
			case BatteryLevel:
				name = "batteryLevel"
				unit = "%EL"
			case ReadSwitch:
				name = "readSwitch"
				unit = "/"
			case BatteryLevelLow:
				name = "batteryLevelLow"
				unit = "/"
			case CompassX:
				name = "compassX"
				unit = "T"
			case CompassY:
				name = "compassY"
				unit = "T"
			case CompassZ:
				name = "compassZ"
				unit = "T"
			case Presence:
				name = "presence"
			case Counter:
				name = "counter"
			}
		}

		senmlPack.Records[i] = senml.Record{
			Name:        name,
			Unit:        unit,
			BaseName:    r.BaseName,
			BaseTime:    r.BaseTime,
			BaseUnit:    r.BaseUnit,
			BaseValue:   r.BaseValue,
			BaseSum:     r.BaseSum,
			BaseVersion: r.BaseVersion,
			Time:        r.Time,
			UpdateTime:  r.UpdateTime,
			Value:       r.Value,
			StringValue: r.StringValue,
			DataValue:   r.DataValue,
			BoolValue:   r.BoolValue,
			Sum:         r.Sum,
		}
	}

	return senmlPack
}

func normalize(p senml.Pack) (Pack, error) {
	senmlPack, err := senml.Normalize(p)
	if err != nil {
		return nil, err
	}

	pack := make(Pack, len(senmlPack.Records))
	for i, r := range senmlPack.Records {
		pack[i] = Record{
			Name:        r.Name,
			Unit:        r.Unit,
			BaseName:    r.BaseName,
			BaseTime:    r.BaseTime,
			BaseUnit:    r.BaseUnit,
			BaseValue:   r.BaseValue,
			BaseSum:     r.BaseSum,
			BaseVersion: r.BaseVersion,
			Time:        r.Time,
			UpdateTime:  r.UpdateTime,
			Value:       r.Value,
			StringValue: r.StringValue,
			DataValue:   r.DataValue,
			BoolValue:   r.BoolValue,
			Sum:         r.Sum,
		}
	}

	return pack, nil
}

func NormalizeJSON(msg []byte) (Pack, error) {
	var pack Pack
	if err := json.Unmarshal(msg, &pack); err != nil {
		return nil, errors.Join(err, ErrUnmarshal)
	}

	return normalize(pack.ToSenML())
}

func NormalizeCBOR(msg []byte) (Pack, error) {
	var pack Pack
	if err := cbor.Unmarshal(msg, &pack); err != nil {
		return nil, errors.Join(err, ErrUnmarshal)
	}

	return normalize(pack.ToSenML())
}
