package lap_data

import (
	"github.com/roryphillips/f1-telemetry-client/internal/common"
)

// Packet lap data
type Packet struct {
	// Header packet header
	Header common.Header
	// LapData for all cars on track
	LapData [22]LapData
}

// LapData per-car lap data
type LapData struct {
	// LastLapTime Last lap time in seconds
	LastLapTime            float32
	// CurrentLapTime Current time around the lap in seconds
	CurrentLapTime         float32
	// Sector1Time Sector 1 time in milliseconds
	Sector1Time            uint16
	// Sector2Time Sector 2 time in milliseconds
	Sector2Time            uint16
	// BestLapTime Best lap time of the session in seconds
	BestLapTime            float32
	// BestLapNum Lap number best time achieved on
	BestLapNum             uint8
	// BestLapSector1Time Sector 1 time of the best lap of the session in milliseconds
	BestLapSector1Time     uint16
	// BestLapSector2Time Sector 2 time of the best lap of the session in milliseconds
	BestLapSector2Time     uint16
	// BestLapSector3Time Sector 3 time of the best lap of the session in milliseconds
	BestLapSector3Time     uint16
	// BestOverallSector1Time Best overall sector 1 time of the session
	BestOverallSector1Time uint16
	// BestOverallSector1Lap Lap number best overall sector 1 time achieved on
	BestOverallSector1Lap  uint8
	// BestOverallSector2Time Best overall sector 2 time of the session
	BestOverallSector2Time uint16
	// BestOverallSector2Lap Lap number best overall sector 2 time achieved on
	BestOverallSector2Lap  uint8
	// BestOverallSector3Time Best overall sector 3 time of the session
	BestOverallSector3Time uint16
	// BestOverallSector3Lap Lap number best overall sector 3 time achieved on
	BestOverallSector3Lap  uint8
	// LapDistance Distance vehicle is around current lap in metres - may be negative if the line isn't crossed yet
	LapDistance            float32
	// TotalDistance Distance travelled in the session in metres - may be negative if the line isn't crossed yet
	TotalDistance          float32
	// SafetyCarDelta Delta in seconds for the safety car
	SafetyCarDelta         float32
	// CarPosition Car race position
	CarPosition            uint8
	// CurrentLapNum Current lap number
	CurrentLapNum          uint8
	//
	PitStatus              PitStatus
	Sector                 Sector
	CurrentLapInvalid      bool
	Penalties              uint8
	GridPosition           uint8
	DriverStatus           DriverStatus
	ResultStatus           ResultStatus
}
