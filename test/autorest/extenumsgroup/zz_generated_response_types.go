//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package extenumsgroup

import "net/http"

// PetAddPetResponse contains the response from method Pet.AddPet.
type PetAddPetResponse struct {
	PetAddPetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PetAddPetResult contains the result from method Pet.AddPet.
type PetAddPetResult struct {
	Pet
}

// PetGetByPetIDResponse contains the response from method Pet.GetByPetID.
type PetGetByPetIDResponse struct {
	PetGetByPetIDResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PetGetByPetIDResult contains the result from method Pet.GetByPetID.
type PetGetByPetIDResult struct {
	Pet
}
