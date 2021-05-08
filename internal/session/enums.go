package session

// SessionType type of session
type SessionType uint8

const (
	// SessionTypeUnknown unknown session type
	SessionTypeUnknown SessionType = 0
	// SessionTypePractice1 practice one
	SessionTypePractice1 SessionType = 1
	// SessionTypePractice2 practice two
	SessionTypePractice2 SessionType = 2
	// SessionTypePractice3 practice three
	SessionTypePractice3 SessionType = 3
	// SessionTypeShortPractice short practice
	SessionTypeShortPractice SessionType = 4
	// SessionTypeQualifying1 qualifying one
	SessionTypeQualifying1 SessionType = 5
	// SessionTypeQualifying2 qualifying two
	SessionTypeQualifying2 SessionType = 6
	// SessionTypeQualifying3 qualifying three
	SessionTypeQualifying3 SessionType = 7
	// SessionTypeShortQualifying short qualifying
	SessionTypeShortQualifying SessionType = 8
	// SessionTypeOneShotQualifying one shot qualifying
	SessionTypeOneShotQualifying SessionType = 9
	// SessionTypeRace1 race one
	SessionTypeRace1 SessionType = 10
	// SessionTypeRace2 race two
	SessionTypeRace2 SessionType = 11
	// SessionTypeTimeTrial time trial
	SessionTypeTimeTrial SessionType = 12
)

// WeatherType type of weather
type WeatherType uint8

const (
	//WeatherTypeClear clear
	WeatherTypeClear WeatherType = 0
	//WeatherTypeLightCloud light cloud
	WeatherTypeLightCloud WeatherType = 1
	//WeatherTypeOvercast overcast
	WeatherTypeOvercast WeatherType = 2
	//WeatherTypeLightRain light rain
	WeatherTypeLightRain WeatherType = 3
	//WeatherTypeHeavyRain heavy rain
	WeatherTypeHeavyRain WeatherType = 4
	//WeatherTypeStorm storm
	WeatherTypeStorm WeatherType = 5
)

type TrackType int8

const (
	// TrackTypeUnknown Unknown track
	TrackTypeUnknown TrackType = -1
)

type FormulaType uint8

const (
	// FormulaTypeF1Modern F1 Modern
	FormulaTypeF1Modern FormulaType = 0
	// FormulaTypeF1Classic F1 Classic
	FormulaTypeF1Classic FormulaType = 1
	// FormulaTypeF2 F2
	FormulaTypeF2 FormulaType = 2
	// FormulaTypeF1Generic F1 Generic
	FormulaTypeF1Generic FormulaType = 3
)

// ZoneFlag flag flown
type ZoneFlag int8

const (
	// ZoneFlagUnknown unknown flag
	ZoneFlagUnknown ZoneFlag = -1
	// ZoneFlagNone no flag type
	ZoneFlagNone ZoneFlag = 0
	// ZoneFlagGreen green flag
	ZoneFlagGreen ZoneFlag = 1
	// ZoneFlagBlue blue flag
	ZoneFlagBlue ZoneFlag = 2
	// ZoneFlagYellow yellow flag
	ZoneFlagYellow ZoneFlag = 3
	// ZoneFlagRed red flag
	ZoneFlagRed ZoneFlag = 4
)
