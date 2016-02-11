// Copyright (c) 2011-2015, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type VisionPrescription struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	DateWritten           *FHIRDateTime                         `bson:"dateWritten,omitempty" json:"dateWritten,omitempty"`
	Patient               *Reference                            `bson:"patient,omitempty" json:"patient,omitempty"`
	Prescriber            *Reference                            `bson:"prescriber,omitempty" json:"prescriber,omitempty"`
	Encounter             *Reference                            `bson:"encounter,omitempty" json:"encounter,omitempty"`
	ReasonCodeableConcept *CodeableConcept                      `bson:"reasonCodeableConcept,omitempty" json:"reasonCodeableConcept,omitempty"`
	ReasonReference       *Reference                            `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Dispense              []VisionPrescriptionDispenseComponent `bson:"dispense,omitempty" json:"dispense,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *VisionPrescription) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "VisionPrescription"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to VisionPrescription), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *VisionPrescription) GetBSON() (interface{}, error) {
	x.ResourceType = "VisionPrescription"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "visionPrescription" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type visionPrescription VisionPrescription

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *VisionPrescription) UnmarshalJSON(data []byte) (err error) {
	x2 := visionPrescription{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = VisionPrescription(x2)
		return x.checkResourceType()
	}
	return
}

func (x *VisionPrescription) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "VisionPrescription"
	} else if x.ResourceType != "VisionPrescription" {
		return errors.New(fmt.Sprintf("Expected resourceType to be VisionPrescription, instead received %s", x.ResourceType))
	}
	return nil
}

type VisionPrescriptionDispenseComponent struct {
	Product   *Coding   `bson:"product,omitempty" json:"product,omitempty"`
	Eye       string    `bson:"eye,omitempty" json:"eye,omitempty"`
	Sphere    *float64  `bson:"sphere,omitempty" json:"sphere,omitempty"`
	Cylinder  *float64  `bson:"cylinder,omitempty" json:"cylinder,omitempty"`
	Axis      *int32    `bson:"axis,omitempty" json:"axis,omitempty"`
	Prism     *float64  `bson:"prism,omitempty" json:"prism,omitempty"`
	Base      string    `bson:"base,omitempty" json:"base,omitempty"`
	Add       *float64  `bson:"add,omitempty" json:"add,omitempty"`
	Power     *float64  `bson:"power,omitempty" json:"power,omitempty"`
	BackCurve *float64  `bson:"backCurve,omitempty" json:"backCurve,omitempty"`
	Diameter  *float64  `bson:"diameter,omitempty" json:"diameter,omitempty"`
	Duration  *Quantity `bson:"duration,omitempty" json:"duration,omitempty"`
	Color     string    `bson:"color,omitempty" json:"color,omitempty"`
	Brand     string    `bson:"brand,omitempty" json:"brand,omitempty"`
	Notes     string    `bson:"notes,omitempty" json:"notes,omitempty"`
}

type VisionPrescriptionPlus struct {
	VisionPrescription                     `bson:",inline"`
	VisionPrescriptionPlusRelatedResources `bson:",inline"`
}

type VisionPrescriptionPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByPrescriber         *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByPrescriber,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter             *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse        *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment     *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedCarePlanResourcesReferencingActivityreference    *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                  *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference          *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPlan       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingPlan,omitempty"`
}

func (v *VisionPrescriptionPlusRelatedResources) GetIncludedPractitionerResourceReferencedByPrescriber() (practitioner *Practitioner, err error) {
	if v.IncludedPractitionerResourcesReferencedByPrescriber == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*v.IncludedPractitionerResourcesReferencedByPrescriber) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*v.IncludedPractitionerResourcesReferencedByPrescriber))
	} else if len(*v.IncludedPractitionerResourcesReferencedByPrescriber) == 1 {
		practitioner = &(*v.IncludedPractitionerResourcesReferencedByPrescriber)[0]
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if v.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*v.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*v.IncludedPatientResourcesReferencedByPatient))
	} else if len(*v.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*v.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if v.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*v.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*v.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*v.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*v.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if v.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *v.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if v.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *v.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if v.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *v.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if v.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *v.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if v.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *v.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if v.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *v.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if v.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *v.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if v.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *v.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if v.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *v.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if v.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *v.RevIncludedListResourcesReferencingItem
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if v.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *v.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if v.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *v.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if v.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *v.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if v.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *v.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if v.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *v.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if v.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *v.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if v.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *v.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if v.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *v.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if v.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *v.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingPlan() (clinicalImpressions []ClinicalImpression, err error) {
	if v.RevIncludedClinicalImpressionResourcesReferencingPlan == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *v.RevIncludedClinicalImpressionResourcesReferencingPlan
	}
	return
}

func (v *VisionPrescriptionPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if v.IncludedPractitionerResourcesReferencedByPrescriber != nil {
		for _, r := range *v.IncludedPractitionerResourcesReferencedByPrescriber {
			resourceMap[r.Id] = &r
		}
	}
	if v.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *v.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if v.IncludedEncounterResourcesReferencedByEncounter != nil {
		for _, r := range *v.IncludedEncounterResourcesReferencedByEncounter {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (v *VisionPrescriptionPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if v.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *v.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *v.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *v.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *v.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *v.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *v.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *v.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *v.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for _, r := range *v.RevIncludedCarePlanResourcesReferencingActivityreference {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *v.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *v.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *v.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *v.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *v.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *v.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *v.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *v.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *v.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *v.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for _, r := range *v.RevIncludedClinicalImpressionResourcesReferencingPlan {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (v *VisionPrescriptionPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if v.IncludedPractitionerResourcesReferencedByPrescriber != nil {
		for _, r := range *v.IncludedPractitionerResourcesReferencedByPrescriber {
			resourceMap[r.Id] = &r
		}
	}
	if v.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *v.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if v.IncludedEncounterResourcesReferencedByEncounter != nil {
		for _, r := range *v.IncludedEncounterResourcesReferencedByEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *v.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *v.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *v.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *v.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *v.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *v.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *v.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *v.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for _, r := range *v.RevIncludedCarePlanResourcesReferencingActivityreference {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *v.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *v.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *v.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *v.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *v.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *v.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *v.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *v.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *v.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *v.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if v.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for _, r := range *v.RevIncludedClinicalImpressionResourcesReferencingPlan {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
