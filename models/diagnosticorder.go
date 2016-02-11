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

type DiagnosticOrder struct {
	DomainResource        `bson:",inline"`
	Identifier            []Identifier                    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status                string                          `bson:"status,omitempty" json:"status,omitempty"`
	Priority              string                          `bson:"priority,omitempty" json:"priority,omitempty"`
	Subject               *Reference                      `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter             *Reference                      `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Orderer               *Reference                      `bson:"orderer,omitempty" json:"orderer,omitempty"`
	Reason                []CodeableConcept               `bson:"reason,omitempty" json:"reason,omitempty"`
	SupportingInformation []Reference                     `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	Specimen              []Reference                     `bson:"specimen,omitempty" json:"specimen,omitempty"`
	Event                 []DiagnosticOrderEventComponent `bson:"event,omitempty" json:"event,omitempty"`
	Item                  []DiagnosticOrderItemComponent  `bson:"item,omitempty" json:"item,omitempty"`
	Note                  []Annotation                    `bson:"note,omitempty" json:"note,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DiagnosticOrder) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DiagnosticOrder"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DiagnosticOrder), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DiagnosticOrder) GetBSON() (interface{}, error) {
	x.ResourceType = "DiagnosticOrder"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "diagnosticOrder" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type diagnosticOrder DiagnosticOrder

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DiagnosticOrder) UnmarshalJSON(data []byte) (err error) {
	x2 := diagnosticOrder{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DiagnosticOrder(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DiagnosticOrder) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DiagnosticOrder"
	} else if x.ResourceType != "DiagnosticOrder" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DiagnosticOrder, instead received %s", x.ResourceType))
	}
	return nil
}

type DiagnosticOrderEventComponent struct {
	Status      string           `bson:"status,omitempty" json:"status,omitempty"`
	Description *CodeableConcept `bson:"description,omitempty" json:"description,omitempty"`
	DateTime    *FHIRDateTime    `bson:"dateTime,omitempty" json:"dateTime,omitempty"`
	Actor       *Reference       `bson:"actor,omitempty" json:"actor,omitempty"`
}

type DiagnosticOrderItemComponent struct {
	Code     *CodeableConcept                `bson:"code,omitempty" json:"code,omitempty"`
	Specimen []Reference                     `bson:"specimen,omitempty" json:"specimen,omitempty"`
	BodySite *CodeableConcept                `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	Status   string                          `bson:"status,omitempty" json:"status,omitempty"`
	Event    []DiagnosticOrderEventComponent `bson:"event,omitempty" json:"event,omitempty"`
}

type DiagnosticOrderPlus struct {
	DiagnosticOrder                     `bson:",inline"`
	DiagnosticOrderPlusRelatedResources `bson:",inline"`
}

type DiagnosticOrderPlusRelatedResources struct {
	IncludedGroupResourcesReferencedBySubject                   *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                  *[]Device                `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedLocationResourcesReferencedBySubject                *[]Location              `bson:"_includedLocationResourcesReferencedBySubject,omitempty"`
	IncludedEncounterResourcesReferencedByEncounter             *[]Encounter             `bson:"_includedEncounterResourcesReferencedByEncounter,omitempty"`
	IncludedPractitionerResourcesReferencedByActorPath1         *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByActorPath1,omitempty"`
	IncludedPractitionerResourcesReferencedByActorPath2         *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByActorPath2,omitempty"`
	IncludedDeviceResourcesReferencedByActorPath1               *[]Device                `bson:"_includedDeviceResourcesReferencedByActorPath1,omitempty"`
	IncludedDeviceResourcesReferencedByActorPath2               *[]Device                `bson:"_includedDeviceResourcesReferencedByActorPath2,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByOrderer            *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByOrderer,omitempty"`
	IncludedSpecimenResourcesReferencedBySpecimenPath1          *[]Specimen              `bson:"_includedSpecimenResourcesReferencedBySpecimenPath1,omitempty"`
	IncludedSpecimenResourcesReferencedBySpecimenPath2          *[]Specimen              `bson:"_includedSpecimenResourcesReferencedBySpecimenPath2,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse        *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedImagingStudyResourcesReferencingOrder            *[]ImagingStudy          `bson:"_revIncludedImagingStudyResourcesReferencingOrder,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment     *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedCarePlanResourcesReferencingActivityreference    *[]CarePlan              `bson:"_revIncludedCarePlanResourcesReferencingActivityreference,omitempty"`
	RevIncludedListResourcesReferencingItem                     *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                  *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedDiagnosticReportResourcesReferencingRequest      *[]DiagnosticReport      `bson:"_revIncludedDiagnosticReportResourcesReferencingRequest,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference          *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingAction     *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingAction,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingPlan       *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingPlan,omitempty"`
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if d.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*d.IncludedGroupResourcesReferencedBySubject))
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*d.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if d.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*d.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*d.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedBySubject))
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedLocationResourceReferencedBySubject() (location *Location, err error) {
	if d.IncludedLocationResourcesReferencedBySubject == nil {
		err = errors.New("Included locations not requested")
	} else if len(*d.IncludedLocationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*d.IncludedLocationResourcesReferencedBySubject))
	} else if len(*d.IncludedLocationResourcesReferencedBySubject) == 1 {
		location = &(*d.IncludedLocationResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedEncounterResourceReferencedByEncounter() (encounter *Encounter, err error) {
	if d.IncludedEncounterResourcesReferencedByEncounter == nil {
		err = errors.New("Included encounters not requested")
	} else if len(*d.IncludedEncounterResourcesReferencedByEncounter) > 1 {
		err = fmt.Errorf("Expected 0 or 1 encounter, but found %d", len(*d.IncludedEncounterResourcesReferencedByEncounter))
	} else if len(*d.IncludedEncounterResourcesReferencedByEncounter) == 1 {
		encounter = &(*d.IncludedEncounterResourcesReferencedByEncounter)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedPractitionerResourceReferencedByActorPath1() (practitioner *Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedByActorPath1 == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedPractitionerResourcesReferencedByActorPath1) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedPractitionerResourcesReferencedByActorPath1))
	} else if len(*d.IncludedPractitionerResourcesReferencedByActorPath1) == 1 {
		practitioner = &(*d.IncludedPractitionerResourcesReferencedByActorPath1)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedPractitionerResourceReferencedByActorPath2() (practitioner *Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedByActorPath2 == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedPractitionerResourcesReferencedByActorPath2) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedPractitionerResourcesReferencedByActorPath2))
	} else if len(*d.IncludedPractitionerResourcesReferencedByActorPath2) == 1 {
		practitioner = &(*d.IncludedPractitionerResourcesReferencedByActorPath2)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedDeviceResourceReferencedByActorPath1() (device *Device, err error) {
	if d.IncludedDeviceResourcesReferencedByActorPath1 == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResourcesReferencedByActorPath1) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResourcesReferencedByActorPath1))
	} else if len(*d.IncludedDeviceResourcesReferencedByActorPath1) == 1 {
		device = &(*d.IncludedDeviceResourcesReferencedByActorPath1)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedDeviceResourceReferencedByActorPath2() (device *Device, err error) {
	if d.IncludedDeviceResourcesReferencedByActorPath2 == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResourcesReferencedByActorPath2) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResourcesReferencedByActorPath2))
	} else if len(*d.IncludedDeviceResourcesReferencedByActorPath2) == 1 {
		device = &(*d.IncludedDeviceResourcesReferencedByActorPath2)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedByPatient))
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedPractitionerResourceReferencedByOrderer() (practitioner *Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedByOrderer == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedPractitionerResourcesReferencedByOrderer) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedPractitionerResourcesReferencedByOrderer))
	} else if len(*d.IncludedPractitionerResourcesReferencedByOrderer) == 1 {
		practitioner = &(*d.IncludedPractitionerResourcesReferencedByOrderer)[0]
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedSpecimenResourcesReferencedBySpecimenPath1() (specimen []Specimen, err error) {
	if d.IncludedSpecimenResourcesReferencedBySpecimenPath1 == nil {
		err = errors.New("Included specimen not requested")
	} else {
		specimen = *d.IncludedSpecimenResourcesReferencedBySpecimenPath1
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedSpecimenResourcesReferencedBySpecimenPath2() (specimen []Specimen, err error) {
	if d.IncludedSpecimenResourcesReferencedBySpecimenPath2 == nil {
		err = errors.New("Included specimen not requested")
	} else {
		specimen = *d.IncludedSpecimenResourcesReferencedBySpecimenPath2
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedImagingStudyResourcesReferencingOrder() (imagingStudies []ImagingStudy, err error) {
	if d.RevIncludedImagingStudyResourcesReferencingOrder == nil {
		err = errors.New("RevIncluded imagingStudies not requested")
	} else {
		imagingStudies = *d.RevIncludedImagingStudyResourcesReferencingOrder
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if d.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *d.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedCarePlanResourcesReferencingActivityreference() (carePlans []CarePlan, err error) {
	if d.RevIncludedCarePlanResourcesReferencingActivityreference == nil {
		err = errors.New("RevIncluded carePlans not requested")
	} else {
		carePlans = *d.RevIncludedCarePlanResourcesReferencingActivityreference
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingItem
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if d.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *d.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if d.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *d.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedDiagnosticReportResourcesReferencingRequest() (diagnosticReports []DiagnosticReport, err error) {
	if d.RevIncludedDiagnosticReportResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded diagnosticReports not requested")
	} else {
		diagnosticReports = *d.RevIncludedDiagnosticReportResourcesReferencingRequest
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *d.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if d.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *d.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingAction() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingAction == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingAction
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingPlan() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingPlan == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingPlan
	}
	return
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedGroupResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedGroupResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedDeviceResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedDeviceResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedLocationResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedLocationResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedEncounterResourcesReferencedByEncounter != nil {
		for _, r := range *d.IncludedEncounterResourcesReferencedByEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPractitionerResourcesReferencedByActorPath1 != nil {
		for _, r := range *d.IncludedPractitionerResourcesReferencedByActorPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPractitionerResourcesReferencedByActorPath2 != nil {
		for _, r := range *d.IncludedPractitionerResourcesReferencedByActorPath2 {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedDeviceResourcesReferencedByActorPath1 != nil {
		for _, r := range *d.IncludedDeviceResourcesReferencedByActorPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedDeviceResourcesReferencedByActorPath2 != nil {
		for _, r := range *d.IncludedDeviceResourcesReferencedByActorPath2 {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *d.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPractitionerResourcesReferencedByOrderer != nil {
		for _, r := range *d.IncludedPractitionerResourcesReferencedByOrderer {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSpecimenResourcesReferencedBySpecimenPath1 != nil {
		for _, r := range *d.IncludedSpecimenResourcesReferencedBySpecimenPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSpecimenResourcesReferencedBySpecimenPath2 != nil {
		for _, r := range *d.IncludedSpecimenResourcesReferencedBySpecimenPath2 {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (d *DiagnosticOrderPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *d.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *d.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *d.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedImagingStudyResourcesReferencingOrder != nil {
		for _, r := range *d.RevIncludedImagingStudyResourcesReferencingOrder {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *d.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *d.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for _, r := range *d.RevIncludedCarePlanResourcesReferencingActivityreference {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *d.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *d.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDiagnosticReportResourcesReferencingRequest != nil {
		for _, r := range *d.RevIncludedDiagnosticReportResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *d.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *d.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *d.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *d.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for _, r := range *d.RevIncludedClinicalImpressionResourcesReferencingAction {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for _, r := range *d.RevIncludedClinicalImpressionResourcesReferencingPlan {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (d *DiagnosticOrderPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedGroupResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedGroupResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedDeviceResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedDeviceResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedLocationResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedLocationResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedEncounterResourcesReferencedByEncounter != nil {
		for _, r := range *d.IncludedEncounterResourcesReferencedByEncounter {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPractitionerResourcesReferencedByActorPath1 != nil {
		for _, r := range *d.IncludedPractitionerResourcesReferencedByActorPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPractitionerResourcesReferencedByActorPath2 != nil {
		for _, r := range *d.IncludedPractitionerResourcesReferencedByActorPath2 {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedDeviceResourcesReferencedByActorPath1 != nil {
		for _, r := range *d.IncludedDeviceResourcesReferencedByActorPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedDeviceResourcesReferencedByActorPath2 != nil {
		for _, r := range *d.IncludedDeviceResourcesReferencedByActorPath2 {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *d.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPractitionerResourcesReferencedByOrderer != nil {
		for _, r := range *d.IncludedPractitionerResourcesReferencedByOrderer {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSpecimenResourcesReferencedBySpecimenPath1 != nil {
		for _, r := range *d.IncludedSpecimenResourcesReferencedBySpecimenPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedSpecimenResourcesReferencedBySpecimenPath2 != nil {
		for _, r := range *d.IncludedSpecimenResourcesReferencedBySpecimenPath2 {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *d.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *d.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *d.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *d.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedImagingStudyResourcesReferencingOrder != nil {
		for _, r := range *d.RevIncludedImagingStudyResourcesReferencingOrder {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *d.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *d.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *d.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCarePlanResourcesReferencingActivityreference != nil {
		for _, r := range *d.RevIncludedCarePlanResourcesReferencingActivityreference {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *d.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *d.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDiagnosticReportResourcesReferencingRequest != nil {
		for _, r := range *d.RevIncludedDiagnosticReportResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *d.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *d.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *d.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *d.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *d.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingAction != nil {
		for _, r := range *d.RevIncludedClinicalImpressionResourcesReferencingAction {
			resourceMap[r.Id] = &r
		}
	}
	if d.RevIncludedClinicalImpressionResourcesReferencingPlan != nil {
		for _, r := range *d.RevIncludedClinicalImpressionResourcesReferencingPlan {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
