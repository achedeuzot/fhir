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

type AuditEvent struct {
	DomainResource `bson:",inline"`
	Type           *Coding                     `bson:"type,omitempty" json:"type,omitempty"`
	Subtype        []Coding                    `bson:"subtype,omitempty" json:"subtype,omitempty"`
	Action         string                      `bson:"action,omitempty" json:"action,omitempty"`
	Recorded       *FHIRDateTime               `bson:"recorded,omitempty" json:"recorded,omitempty"`
	Outcome        string                      `bson:"outcome,omitempty" json:"outcome,omitempty"`
	OutcomeDesc    string                      `bson:"outcomeDesc,omitempty" json:"outcomeDesc,omitempty"`
	PurposeOfEvent []Coding                    `bson:"purposeOfEvent,omitempty" json:"purposeOfEvent,omitempty"`
	Agent          []AuditEventAgentComponent  `bson:"agent,omitempty" json:"agent,omitempty"`
	Source         *AuditEventSourceComponent  `bson:"source,omitempty" json:"source,omitempty"`
	Entity         []AuditEventEntityComponent `bson:"entity,omitempty" json:"entity,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *AuditEvent) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "AuditEvent"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to AuditEvent), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *AuditEvent) GetBSON() (interface{}, error) {
	x.ResourceType = "AuditEvent"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "auditEvent" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type auditEvent AuditEvent

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *AuditEvent) UnmarshalJSON(data []byte) (err error) {
	x2 := auditEvent{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = AuditEvent(x2)
		return x.checkResourceType()
	}
	return
}

func (x *AuditEvent) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "AuditEvent"
	} else if x.ResourceType != "AuditEvent" {
		return errors.New(fmt.Sprintf("Expected resourceType to be AuditEvent, instead received %s", x.ResourceType))
	}
	return nil
}

type AuditEventAgentComponent struct {
	Role         []CodeableConcept                `bson:"role,omitempty" json:"role,omitempty"`
	Reference    *Reference                       `bson:"reference,omitempty" json:"reference,omitempty"`
	UserId       *Identifier                      `bson:"userId,omitempty" json:"userId,omitempty"`
	AltId        string                           `bson:"altId,omitempty" json:"altId,omitempty"`
	Name         string                           `bson:"name,omitempty" json:"name,omitempty"`
	Requestor    *bool                            `bson:"requestor,omitempty" json:"requestor,omitempty"`
	Location     *Reference                       `bson:"location,omitempty" json:"location,omitempty"`
	Policy       []string                         `bson:"policy,omitempty" json:"policy,omitempty"`
	Media        *Coding                          `bson:"media,omitempty" json:"media,omitempty"`
	Network      *AuditEventAgentNetworkComponent `bson:"network,omitempty" json:"network,omitempty"`
	PurposeOfUse []Coding                         `bson:"purposeOfUse,omitempty" json:"purposeOfUse,omitempty"`
}

type AuditEventAgentNetworkComponent struct {
	Address string `bson:"address,omitempty" json:"address,omitempty"`
	Type    string `bson:"type,omitempty" json:"type,omitempty"`
}

type AuditEventSourceComponent struct {
	Site       string      `bson:"site,omitempty" json:"site,omitempty"`
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type       []Coding    `bson:"type,omitempty" json:"type,omitempty"`
}

type AuditEventEntityComponent struct {
	Identifier    *Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Reference     *Reference                        `bson:"reference,omitempty" json:"reference,omitempty"`
	Type          *Coding                           `bson:"type,omitempty" json:"type,omitempty"`
	Role          *Coding                           `bson:"role,omitempty" json:"role,omitempty"`
	Lifecycle     *Coding                           `bson:"lifecycle,omitempty" json:"lifecycle,omitempty"`
	SecurityLabel []Coding                          `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	Name          string                            `bson:"name,omitempty" json:"name,omitempty"`
	Description   string                            `bson:"description,omitempty" json:"description,omitempty"`
	Query         string                            `bson:"query,omitempty" json:"query,omitempty"`
	Detail        []AuditEventEntityDetailComponent `bson:"detail,omitempty" json:"detail,omitempty"`
}

type AuditEventEntityDetailComponent struct {
	Type  string `bson:"type,omitempty" json:"type,omitempty"`
	Value string `bson:"value,omitempty" json:"value,omitempty"`
}

type AuditEventPlus struct {
	AuditEvent                     `bson:",inline"`
	AuditEventPlusRelatedResources `bson:",inline"`
}

type AuditEventPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByParticipant        *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByParticipant,omitempty"`
	IncludedOrganizationResourcesReferencedByParticipant        *[]Organization          `bson:"_includedOrganizationResourcesReferencedByParticipant,omitempty"`
	IncludedDeviceResourcesReferencedByParticipant              *[]Device                `bson:"_includedDeviceResourcesReferencedByParticipant,omitempty"`
	IncludedPatientResourcesReferencedByParticipant             *[]Patient               `bson:"_includedPatientResourcesReferencedByParticipant,omitempty"`
	IncludedRelatedPersonResourcesReferencedByParticipant       *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByParticipant,omitempty"`
	IncludedPatientResourcesReferencedByPatientPath1            *[]Patient               `bson:"_includedPatientResourcesReferencedByPatientPath1,omitempty"`
	IncludedPatientResourcesReferencedByPatientPath2            *[]Patient               `bson:"_includedPatientResourcesReferencedByPatientPath2,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse        *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment     *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData            *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget             *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
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
}

func (a *AuditEventPlusRelatedResources) GetIncludedPractitionerResourceReferencedByParticipant() (practitioner *Practitioner, err error) {
	if a.IncludedPractitionerResourcesReferencedByParticipant == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*a.IncludedPractitionerResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*a.IncludedPractitionerResourcesReferencedByParticipant))
	} else if len(*a.IncludedPractitionerResourcesReferencedByParticipant) == 1 {
		practitioner = &(*a.IncludedPractitionerResourcesReferencedByParticipant)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedOrganizationResourceReferencedByParticipant() (organization *Organization, err error) {
	if a.IncludedOrganizationResourcesReferencedByParticipant == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*a.IncludedOrganizationResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*a.IncludedOrganizationResourcesReferencedByParticipant))
	} else if len(*a.IncludedOrganizationResourcesReferencedByParticipant) == 1 {
		organization = &(*a.IncludedOrganizationResourcesReferencedByParticipant)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedDeviceResourceReferencedByParticipant() (device *Device, err error) {
	if a.IncludedDeviceResourcesReferencedByParticipant == nil {
		err = errors.New("Included devices not requested")
	} else if len(*a.IncludedDeviceResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*a.IncludedDeviceResourcesReferencedByParticipant))
	} else if len(*a.IncludedDeviceResourcesReferencedByParticipant) == 1 {
		device = &(*a.IncludedDeviceResourcesReferencedByParticipant)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedPatientResourceReferencedByParticipant() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByParticipant == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByParticipant))
	} else if len(*a.IncludedPatientResourcesReferencedByParticipant) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByParticipant)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByParticipant() (relatedPerson *RelatedPerson, err error) {
	if a.IncludedRelatedPersonResourcesReferencedByParticipant == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByParticipant) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*a.IncludedRelatedPersonResourcesReferencedByParticipant))
	} else if len(*a.IncludedRelatedPersonResourcesReferencedByParticipant) == 1 {
		relatedPerson = &(*a.IncludedRelatedPersonResourcesReferencedByParticipant)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedPatientResourceReferencedByPatientPath1() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByPatientPath1 == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByPatientPath1) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByPatientPath1))
	} else if len(*a.IncludedPatientResourcesReferencedByPatientPath1) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByPatientPath1)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedPatientResourceReferencedByPatientPath2() (patient *Patient, err error) {
	if a.IncludedPatientResourcesReferencedByPatientPath2 == nil {
		err = errors.New("Included patients not requested")
	} else if len(*a.IncludedPatientResourcesReferencedByPatientPath2) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*a.IncludedPatientResourcesReferencedByPatientPath2))
	} else if len(*a.IncludedPatientResourcesReferencedByPatientPath2) == 1 {
		patient = &(*a.IncludedPatientResourcesReferencedByPatientPath2)[0]
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *a.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *a.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if a.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *a.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if a.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *a.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if a.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *a.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if a.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *a.RevIncludedListResourcesReferencingItem
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if a.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *a.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if a.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *a.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if a.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *a.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if a.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *a.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *a.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if a.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *a.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if a.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *a.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (a *AuditEventPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByParticipant != nil {
		for _, r := range *a.IncludedPractitionerResourcesReferencedByParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedOrganizationResourcesReferencedByParticipant != nil {
		for _, r := range *a.IncludedOrganizationResourcesReferencedByParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedDeviceResourcesReferencedByParticipant != nil {
		for _, r := range *a.IncludedDeviceResourcesReferencedByParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedPatientResourcesReferencedByParticipant != nil {
		for _, r := range *a.IncludedPatientResourcesReferencedByParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByParticipant != nil {
		for _, r := range *a.IncludedRelatedPersonResourcesReferencedByParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedPatientResourcesReferencedByPatientPath1 != nil {
		for _, r := range *a.IncludedPatientResourcesReferencedByPatientPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedPatientResourcesReferencedByPatientPath2 != nil {
		for _, r := range *a.IncludedPatientResourcesReferencedByPatientPath2 {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (a *AuditEventPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *a.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *a.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *a.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *a.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *a.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *a.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *a.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *a.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *a.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *a.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *a.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *a.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *a.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (a *AuditEventPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if a.IncludedPractitionerResourcesReferencedByParticipant != nil {
		for _, r := range *a.IncludedPractitionerResourcesReferencedByParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedOrganizationResourcesReferencedByParticipant != nil {
		for _, r := range *a.IncludedOrganizationResourcesReferencedByParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedDeviceResourcesReferencedByParticipant != nil {
		for _, r := range *a.IncludedDeviceResourcesReferencedByParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedPatientResourcesReferencedByParticipant != nil {
		for _, r := range *a.IncludedPatientResourcesReferencedByParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedRelatedPersonResourcesReferencedByParticipant != nil {
		for _, r := range *a.IncludedRelatedPersonResourcesReferencedByParticipant {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedPatientResourcesReferencedByPatientPath1 != nil {
		for _, r := range *a.IncludedPatientResourcesReferencedByPatientPath1 {
			resourceMap[r.Id] = &r
		}
	}
	if a.IncludedPatientResourcesReferencedByPatientPath2 != nil {
		for _, r := range *a.IncludedPatientResourcesReferencedByPatientPath2 {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *a.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *a.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *a.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *a.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *a.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *a.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *a.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *a.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *a.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *a.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *a.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *a.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *a.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *a.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *a.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *a.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *a.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if a.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *a.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
