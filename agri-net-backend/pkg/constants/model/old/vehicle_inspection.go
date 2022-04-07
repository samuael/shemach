package model

// Inspection ...
type Inspection struct {
	ID                       uint                 `json:"id" gorm:"primaryKey;autoIncrement:true"`
	GarageID                 uint64               `json:"garage_id"`
	InspectorID              uint                 `json:"inspector_id"`
	Drivername               string               `json:"driver_name"`
	VehicleModel             string               `json:"vehicle_model"`
	VehicleYear              string               `json:"vehicle_year"`
	VehicleMake              string               `json:"vehicle_make"`
	VehicleColor             string               `json:"vehicle_color"`
	LicensePlate             string               `json:"license_plate"`
	FrontImage               string               `json:"front_image"`
	LeftSideImage            string               `json:"left_side_image"`
	RightSideImage           string               `json:"right_side_image"`
	BackImage                string               `json:"back_image"`
	SignatureImage           string               `json:"signature"`
	VinNumber                string               `json:"vin_number"`
	
	HandBrake                *FunctionalityResult `json:"hand_brake"  					gorm:"hand_brake,	foreignkey:HandBrakeRefer"`
	SteeringSystem           *FunctionalityResult `json:"steering_system"  				gorm:"steering_system,	foreignkey:SteeringSystemRefer"`
	BrakeSystem              *FunctionalityResult `json:"brake_system" 					gorm:"brake_system,	foreignkey:BrakeSystemRefer"`
	SeatBelt                 *FunctionalityResult `json:"seat_belt" 					gorm:"seat_belt,	foreignkey:SeatBeltRefer"`
	DoorAndWindow            *FunctionalityResult `json:"door_and_window"   			gorm:"door_and_window,	foreignkey:DoorAndWindowRefer"`
	DashBoardLight           *FunctionalityResult `json:"dash_board_light"  			gorm:"dash_board_light,	foreignkey:DashBoardLightRefer"`
	WindShield               *FunctionalityResult `json:"wind_shield" 					gorm:"wind_shield,	foreignkey:WindShieldRefer"`
	BaggageDoorWindow        *FunctionalityResult `json:"baggage_door_window" 			gorm:"baggage_door_window,	foreignkey:BaggageDoorWindowRefer"`
	GearBox                  *FunctionalityResult `json:"gear_box" 						gorm:"gear_box,	foreignkey:GearBoxRefer"`
	ShockAbsorber            *FunctionalityResult `json:"shock_absorber" 				gorm:"shock_absorber,	foreignkey:ShockAbsorberRefer"`
	FrontHighAndLowBeamLight *FunctionalityResult `json:"high_and_low_beam_light" 		gorm:"high_and_low_beam_light,	foreignkey:FrontHighAndLowBeamLightRefer"`
	RearLightAndBrakeLight   *FunctionalityResult `json:"rear_light_and_break_light" 	gorm:"rear_light_and_break_light,	foreignkey:RearLightAndBrakeLightRefer"`
	WiperOperation           *FunctionalityResult `json:"wiper_operation" 				gorm:"wiper_operation,	foreignkey:WiperOperationRefer"`
	CarHorn                  *FunctionalityResult `json:"car_horn" 						gorm:"car_horn,	foreignkey:CarHornRefer"`
	SideMirrors              *FunctionalityResult `json:"side_mirrors" 					gorm:"side_mirrors,	foreignkey:SideMirrorsRefer"`
	GeneralBodyCondition     *FunctionalityResult `json:"general_body_condition" 		gorm:"general_body_condition,	foreignkey:GeneralBodyConditionRefer"`
	DriverPerformance        bool                 `json:"driver_performance"`
	Balancing                bool                 `json:"balancing"`
	Hazard                   bool                 `json:"hazard"`
	SignalLightUsage         bool                 `json:"signal_light_usage"` // turn indicator
	Passed                   bool                 `json:"passed"`
}

// InspectionUpdate ... this datastructure is created to be used for parsing the input json for update
type InspectionUpdate struct {
	ID           int    `json:"id"`
	Drivername   string `json:"driver_name"`
	VehicleModel string `json:"vehicle_model"`
	VehicleYear  string `json:"vehicle_year"`
	VehicleMake  string `json:"vehicle_make"`
	VehicleColor string `json:"vehicle_color"`

	HandBrake                string `json:"hand_brake"`
	SteeringSystem           string `json:"steering_system"`
	BrakeSystem              string `json:"brake_system"`
	SeatBelt                 string `json:"seat_belt"`
	DoorAndWindow            string `json:"door_and_window"`
	DashBoardLight           string `json:"dash_board_light"`
	WindShield               string `json:"wind_shield"`
	BaggageDoorWindow        string `json:"baggage_door_window"`
	GearBox                  string `json:"gear_box"`
	ShockAbsorber            string `json:"shock_absorber"`
	FrontHighAndLowBeamLight string `json:"high_and_low_beam_light"`
	RearLightAndBrakeLight   string `json:"rear_light_and_break_light"`
	WiperOperation           string `json:"wiper_operation"`
	CarHorn                  string `json:"car_horn"`
	SideMirrors              string `json:"side_mirrors"`
	GeneralBodyCondition     string `json:"general_body_condition"`
	DriverPerformance        bool   `json:"driver_performance"`
	Balancing                bool   `json:"balancing"`
	Hazard                   bool   `json:"hazard"`
	SignalLightUsage         bool   `json:"signal_light_usage"`
}

/*
	InspectionUpdate JSON Value
	{
		"id":0,
		"driver_name":"",
		"vehicle_model":"",
		"vehicle_year":"",
		"vehicle_make":"",
		"vehicle_color":"",
		"hand_brake":"",
		"steering_system":"",
		"brake_system":"",
		"seat_belt":"",
		"door_and_window":"",
		"dash_board_light":"",
		"wind_shield":"",
		"baggage_door_window":"",
		"gear_box":"",
		"shock_absorber":"",
		"high_and_low_beam_light":"",
		"rear_light_and_break_light":"",
		"wiper_operation":"",
		"car_horn":"",
		"side_mirrors":"",
		"general_body_condition":"",
		"driver_performance":"",
		"balancing":"",
		"hazard":"",
		"signal_light_usage":"",
}



*/

// FunctionalityResult to be used by a list of functionality parameters and their reasons
type FunctionalityResult struct {
	ID     uint   `gorm:"primaryKey;autoIncrement:true"`
	Result bool   `json:"result" gorm:"boolean;not null;"`           // representing the functionality result
	Reason string `json:"reason" gorm:"type:varchar(255);not null;"` // To represent the failure reason
}
