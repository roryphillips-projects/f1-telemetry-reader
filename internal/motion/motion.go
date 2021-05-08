package motion

import (
	"github.com/roryphillips/f1-telemetry-client/internal/common"
)

// Packet motion data
type Packet struct {
	// Header Packet header
	Header common.Header `json:"header"`
	// CarMotionData Data for all cars on track
	CarMotion [22]CarMotionData `json:"car_motion" packet:"0"`
	//PlayerCar Only available for the player
	PlayerCar PlayerCarData `json:"player_car" packet:"1"`
}

// CarMotionData Data related to the car motion
type CarMotionData struct {
	// WorldPosition Position in world space
	WorldPosition Vector3 `json:"world_position" packet:"0"`
	// WorldVelocity Velocity in world space
	WorldVelocity Vector3 `json:"world_velocity" packet:"1"`
	// WorldForwardDir World space forward direction (normalised)
	WorldForwardDir Vector3 `json:"world_forward_dir" packet:"2"`
	// WorldRightDir World space right direction (normalised)
	WorldRightDir float32 `json:"world_right_dir" packet:"3"`

	// GForceLateral Lateral G-Force component
	GForceLateral float32 `json:"g_force_lateral" packet:"4"`
	// GForceLongitudinal Longitudinal G-Force component
	GForceLongitudinal float32 `json:"g_force_longitudinal" packet:"5"`
	// GForceVertical Vertical G-Force component
	GForceVertical float32 `json:"g_force_vertical" packet:"6"`

	// Yaw angle in radians
	Yaw float32 `json:"yaw" packet:"7"`
	// Pitch angle in radians
	Pitch float32 `json:"pitch" packet:"8"`
	// Roll angle in radians
	Roll float32 `json:"roll" packet:"9"`
}

// PlayerCarData Data related to the player's car
type PlayerCarData struct {
	// SuspensionPosition Position of each suspension arm
	SuspensionPosition WheelData `json:"suspension_position" packet:"0"`
	// SuspensionVelocity Velocity of each suspension arm
	SuspensionVelocity WheelData `json:"suspension_velocity" packet:"1"`
	// SuspensionAcceleration Acceleration of each suspension arm
	SuspensionAcceleration WheelData `json:"suspension_acceleration" packet:"2"`

	// WheelSpeed Speed of each wheel
	WheelSpeed WheelData `json:"wheel_speed" packet:"3"`
	// WheelSlip Slip ratio of each wheel
	WheelSlip WheelData `json:"wheel_slip" packet:"4"`

	// LocalVelocity Velocity in local space
	LocalVelocity Vector3 `json:"local_velocity" packet:"5"`
	// AngularVelocity Angular velocity
	AngularVelocity Vector3 `json:"angular_velocity" packet:"6"`
	// AngularAcceleration Angular acceleration
	AngularAcceleration Vector3 `json:"angular_acceleration" packet:"7"`

	// FrontWheelsAngle Current front wheels angle in radians
	FrontWheelsAngle float32 `json:"front_wheels_angle" packet:"8"`
}

// Vector3 3-dimensional set of properties
type Vector3 struct {
	// X dimension
	X float32 `json:"x" packet:"0"`
	// Y dimension
	Y float32 `json:"y" packet:"1"`
	// Z dimension
	Z float32 `json:"z" packet:"2"`
}

// WheelData Data that is associated with all wheels on the car
type WheelData struct {
	// RearLeft Rear left wheel or suspension
	RearLeft float32 `json:"rear_left" packet:"0"`
	// RearRight Rear right wheel or suspension
	RearRight float32 `json:"rear_right" packet:"1"`
	// FrontLeft Front left wheel or suspension
	FrontLeft float32 `json:"front_left" packet:"2"`
	// FrontRight Front right wheel or suspension
	FrontRight float32 `json:"front_right" packet:"3"`
}
