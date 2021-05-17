// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package extenumsgroup

type Pet struct {
	// REQUIRED
	IntEnum *IntEnum `json:"IntEnum,omitempty"`

	// Type of Pet
	DaysOfWeek *DaysOfWeekExtensibleEnum `json:"DaysOfWeek,omitempty"`

	// name
	Name *string `json:"name,omitempty"`
}

// PetAddPetOptions contains the optional parameters for the Pet.AddPet method.
type PetAddPetOptions struct {
	// pet param
	PetParam *Pet
}

// PetGetByPetIDOptions contains the optional parameters for the Pet.GetByPetID method.
type PetGetByPetIDOptions struct {
	// placeholder for future optional parameters
}
