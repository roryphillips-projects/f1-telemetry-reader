package session

import (
	"github.com/roryphillips/f1-telemetry-client/internal/common"
)

// Packet session data
type Packet struct {
	// Header packet header
	Header common.Header `json:"header"`
	// Weather current weather
	Weather WeatherType `json:"weather" packet:"0"`
	// TrackTemperature in degrees celsius
	TrackTemperature int8 `json:"track_temperature" packet:"1"`
	// AirTemperature in degrees celsius
	AirTemperature int8 `json:"air_temperature" packet:"2"`
	// TotalLaps total number of laps in this race
	TotalLaps uint8 `json:"total_laps" packet:"3"`
	// TrackLength track length in metres
	TrackLength uint16 `json:"track_length" packet:"4"`
	// Session type of the current session
	Session SessionType `json:"session" packet:"5"`
	// Track track this session is in
	Track TrackType `json:"track" packet:"6"`
	// Formula formula this session uses
	Formula FormulaType `json:"formula" packet:"7"`
	// SessionTimeLeft time (in seconds) remaining for this session
	SessionTimeLeft uint16 `json:"session_time_left" packet:"8"`
	// SessionDuration time (in seconds) this session lasts for
	SessionDuration uint16 `json:"session_duration" packet:"9"`
	// PitSpeedLimit limit (in km/h) for the pit lane
	PitSpeedLimit uint8 `json:"pit_speed_limit" packet:"10"`
	// GamePaused whether the game is paused
	GamePaused bool `json:"game_paused" packet:"11"`
	// IsSpectating whether the player is spectating
	IsSpectating bool `json:"is_spectating" packet:"12"`
	// SpectatorCarIndex index of the car being spectated
	SpectatorCarIndex uint8 `json:"spectator_car_index" packet:"13"`
	// SLIProNativeSupport whether SLI pro is supported
	SLIProNativeSupport bool `json:"sli_pro_native_support" packet:"14"`
	// NumMarshalZones number of marshal zones to follow
	NumMarshalZones uint8 `json:"num_marshal_zones" packet:"15"`
	// MarshalZones list of marshal zones - max 21
	MarshalZones [21]MarshalZone `json:"marshal_zones" packet:"16"`
	// NetworkGame whether the game is online (true) or offline
	NetworkGame bool `json:"network_game" packet:"17"`
	// NumWeatherForecastSamples number of forecast samples
	NumWeatherForecastSamples uint8 `json:"num_weather_forecast_samples" packet:"18"`
	// WeatherForecastSamples list of forecast samples - max 20
	WeatherForecastSamples [20]WeatherForecastSample `json:"weather_forecast_samples" packet:"19"`
}

// MarshalZone marshal zone data
type MarshalZone struct {
	// ZoneStart Fraction (0..1) of way through the lap the marshal zone starts
	ZoneStart float32 `json:"zone_start" packet:"0"`
	// ZoneFlag Flag flown in the marshal zone
	ZoneFlag ZoneFlag `json:"zone_flag" packet:"1"`
}

// WeatherForecastSample weather forecast sample data
type WeatherForecastSample struct {
	// SessionType type of session the forecast is for
	Session SessionType `json:"session" packet:"0"`
	// TimeOffset time in minutes this forecast is for
	TimeOffset uint8 `json:"time_offset" packet:"1"`
	// Weather type of weather
	Weather WeatherType `json:"weather" packet:"2"`
	// TrackTemperature in degrees celsius
	TrackTemperature int8 `json:"track_temperature" packet:"3"`
	// AirTemperature in degrees celsius
	AirTemperature int8 `json:"air_temperature" packet:"4"`
}
