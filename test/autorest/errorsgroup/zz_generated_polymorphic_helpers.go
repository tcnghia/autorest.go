//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package errorsgroup

import "encoding/json"

type notFoundErrorBase struct {
	wrapped NotFoundErrorBaseClassification
}

func (n *notFoundErrorBase) UnmarshalJSON(data []byte) (err error) {
	n.wrapped, err = unmarshalNotFoundErrorBaseClassification(data)
	return
}

func unmarshalNotFoundErrorBaseClassification(rawMsg json.RawMessage) (NotFoundErrorBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b NotFoundErrorBaseClassification
	switch m["whatNotFound"] {
	case "AnimalNotFound":
		b = &AnimalNotFound{}
	case "InvalidResourceLink":
		b = &LinkNotFound{}
	default:
		b = &NotFoundErrorBase{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalNotFoundErrorBaseClassificationArray(rawMsg json.RawMessage) ([]NotFoundErrorBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]NotFoundErrorBaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalNotFoundErrorBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalNotFoundErrorBaseClassificationMap(rawMsg json.RawMessage) (map[string]NotFoundErrorBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]NotFoundErrorBaseClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalNotFoundErrorBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

type petActionError struct {
	wrapped PetActionErrorClassification
}

func (p *petActionError) UnmarshalJSON(data []byte) (err error) {
	p.wrapped, err = unmarshalPetActionErrorClassification(data)
	return
}

func unmarshalPetActionErrorClassification(rawMsg json.RawMessage) (PetActionErrorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PetActionErrorClassification
	switch m["errorType"] {
	case "PetHungryOrThirstyError":
		b = &PetHungryOrThirstyError{}
	case "PetSadError":
		b = &PetSadError{}
	default:
		b = &PetActionError{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalPetActionErrorClassificationArray(rawMsg json.RawMessage) ([]PetActionErrorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]PetActionErrorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalPetActionErrorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalPetActionErrorClassificationMap(rawMsg json.RawMessage) (map[string]PetActionErrorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]PetActionErrorClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalPetActionErrorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}

func unmarshalPetSadErrorClassification(rawMsg json.RawMessage) (PetSadErrorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b PetSadErrorClassification
	switch m["errorType"] {
	case "PetHungryOrThirstyError":
		b = &PetHungryOrThirstyError{}
	default:
		b = &PetSadError{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalPetSadErrorClassificationArray(rawMsg json.RawMessage) ([]PetSadErrorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]PetSadErrorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalPetSadErrorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalPetSadErrorClassificationMap(rawMsg json.RawMessage) (map[string]PetSadErrorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages map[string]json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fMap := make(map[string]PetSadErrorClassification, len(rawMessages))
	for key, rawMessage := range rawMessages {
		f, err := unmarshalPetSadErrorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fMap[key] = f
	}
	return fMap, nil
}
