package lap_data

type PitStatus uint8

const (
	// PitStatusNone Not in the pits
	PitStatusNone PitStatus = 0
	// PitStatusPitting Entering the pit area
	PitStatusPitting PitStatus = 1
	// PitStatusInPitArea In the pits
	PitStatusInPitArea PitStatus = 2
)

type Sector uint8

const (
	// Sector1 First sector
	Sector1 Sector = 0
	// Sector2 Second sector
	Sector2 Sector = 1
	// Sector3 Third sector
	Sector3 Sector = 2
)

type DriverStatus uint8

const (
	// DriverStatusInGarage in the garage
	DriverStatusInGarage DriverStatus = 0
	// DriverStatusFlyingLap flying lap
	DriverStatusFlyingLap DriverStatus = 1
	// DriverStatusInLap in lap
	DriverStatusInLap DriverStatus = 2
	// DriverStatusOutLap Out Lap
	DriverStatusOutLap DriverStatus = 3
	// DriverStatusOnTrack On track
	DriverStatusOnTrack DriverStatus = 4
)

type ResultStatus uint8

const (
	//ResultStatusInvalid Invalid
	ResultStatusInvalid ResultStatus = 0
	//ResultStatusInactive Inactive
	ResultStatusInactive ResultStatus = 1
	//ResultStatusActive Active
	ResultStatusActive ResultStatus = 2
	//ResultStatusFinished Finished
	ResultStatusFinished ResultStatus = 3
	//ResultStatusDisqualified Disqualified
	ResultStatusDisqualified ResultStatus = 4
	//ResultStatusNotClassified Not Classified
	ResultStatusNotClassified ResultStatus = 5
	//ResultStatusRetired Retired
	ResultStatusRetired ResultStatus = 6
)
