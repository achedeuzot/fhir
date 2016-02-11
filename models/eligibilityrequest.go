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

type EligibilityRequest struct {
	DomainResource      `bson:",inline"`
	Identifier          []Identifier  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ruleset             *Coding       `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset     *Coding       `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created             *FHIRDateTime `bson:"created,omitempty" json:"created,omitempty"`
	Target              *Reference    `bson:"target,omitempty" json:"target,omitempty"`
	Provider            *Reference    `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization        *Reference    `bson:"organization,omitempty" json:"organization,omitempty"`
	Priority            *Coding       `bson:"priority,omitempty" json:"priority,omitempty"`
	Enterer             *Reference    `bson:"enterer,omitempty" json:"enterer,omitempty"`
	Facility            *Reference    `bson:"facility,omitempty" json:"facility,omitempty"`
	Patient             *Reference    `bson:"patient,omitempty" json:"patient,omitempty"`
	Coverage            *Reference    `bson:"coverage,omitempty" json:"coverage,omitempty"`
	BusinessArrangement string        `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	Relationship        *Coding       `bson:"relationship,omitempty" json:"relationship,omitempty"`
	ServicedDate        *FHIRDateTime `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod      *Period       `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	BenefitCategory     *Coding       `bson:"benefitCategory,omitempty" json:"benefitCategory,omitempty"`
	BenefitSubCategory  *Coding       `bson:"benefitSubCategory,omitempty" json:"benefitSubCategory,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *EligibilityRequest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "EligibilityRequest"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to EligibilityRequest), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *EligibilityRequest) GetBSON() (interface{}, error) {
	x.ResourceType = "EligibilityRequest"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "eligibilityRequest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type eligibilityRequest EligibilityRequest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *EligibilityRequest) UnmarshalJSON(data []byte) (err error) {
	x2 := eligibilityRequest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = EligibilityRequest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *EligibilityRequest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "EligibilityRequest"
	} else if x.ResourceType != "EligibilityRequest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be EligibilityRequest, instead received %s", x.ResourceType))
	}
	return nil
}

type EligibilityRequestPlus struct {
	EligibilityRequest                     `bson:",inline"`
	EligibilityRequestPlusRelatedResources `bson:",inline"`
}

type EligibilityRequestPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByProvider           *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByProvider,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization       *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedLocationResourcesReferencedByFacility               *[]Location              `bson:"_includedLocationResourcesReferencedByFacility,omitempty"`
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
	RevIncludedEligibilityResponseResourcesReferencingRequest   *[]EligibilityResponse   `bson:"_revIncludedEligibilityResponseResourcesReferencingRequest,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                 *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference          *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject           *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry             *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated      *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest       *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger    *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
}

func (e *EligibilityRequestPlusRelatedResources) GetIncludedPractitionerResourceReferencedByProvider() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByProvider == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByProvider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByProvider))
	} else if len(*e.IncludedPractitionerResourcesReferencedByProvider) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByProvider)[0]
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if e.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResourcesReferencedByPatient))
	} else if len(*e.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*e.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*e.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetIncludedLocationResourceReferencedByFacility() (location *Location, err error) {
	if e.IncludedLocationResourcesReferencedByFacility == nil {
		err = errors.New("Included locations not requested")
	} else if len(*e.IncludedLocationResourcesReferencedByFacility) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*e.IncludedLocationResourcesReferencedByFacility))
	} else if len(*e.IncludedLocationResourcesReferencedByFacility) == 1 {
		location = &(*e.IncludedLocationResourcesReferencedByFacility)[0]
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if e.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *e.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if e.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *e.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if e.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *e.RevIncludedListResourcesReferencingItem
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if e.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *e.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedEligibilityResponseResourcesReferencingRequest() (eligibilityResponses []EligibilityResponse, err error) {
	if e.RevIncludedEligibilityResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded eligibilityResponses not requested")
	} else {
		eligibilityResponses = *e.RevIncludedEligibilityResponseResourcesReferencingRequest
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if e.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *e.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if e.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *e.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *e.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if e.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *e.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if e.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *e.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (e *EligibilityRequestPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedPractitionerResourcesReferencedByProvider != nil {
		for _, r := range *e.IncludedPractitionerResourcesReferencedByProvider {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *e.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for _, r := range *e.IncludedOrganizationResourcesReferencedByOrganization {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedLocationResourcesReferencedByFacility != nil {
		for _, r := range *e.IncludedLocationResourcesReferencedByFacility {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (e *EligibilityRequestPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *e.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *e.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *e.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *e.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *e.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *e.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *e.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedEligibilityResponseResourcesReferencingRequest != nil {
		for _, r := range *e.RevIncludedEligibilityResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *e.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *e.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *e.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *e.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (e *EligibilityRequestPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedPractitionerResourcesReferencedByProvider != nil {
		for _, r := range *e.IncludedPractitionerResourcesReferencedByProvider {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *e.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for _, r := range *e.IncludedOrganizationResourcesReferencedByOrganization {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedLocationResourcesReferencedByFacility != nil {
		for _, r := range *e.IncludedLocationResourcesReferencedByFacility {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *e.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *e.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *e.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *e.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *e.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *e.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *e.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedEligibilityResponseResourcesReferencingRequest != nil {
		for _, r := range *e.RevIncludedEligibilityResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *e.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *e.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *e.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *e.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
