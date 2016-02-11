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

type EligibilityResponse struct {
	DomainResource      `bson:",inline"`
	Identifier          []Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Request             *Reference                             `bson:"request,omitempty" json:"request,omitempty"`
	Outcome             string                                 `bson:"outcome,omitempty" json:"outcome,omitempty"`
	Disposition         string                                 `bson:"disposition,omitempty" json:"disposition,omitempty"`
	Ruleset             *Coding                                `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset     *Coding                                `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created             *FHIRDateTime                          `bson:"created,omitempty" json:"created,omitempty"`
	Organization        *Reference                             `bson:"organization,omitempty" json:"organization,omitempty"`
	RequestProvider     *Reference                             `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
	RequestOrganization *Reference                             `bson:"requestOrganization,omitempty" json:"requestOrganization,omitempty"`
	Inforce             *bool                                  `bson:"inforce,omitempty" json:"inforce,omitempty"`
	Contract            *Reference                             `bson:"contract,omitempty" json:"contract,omitempty"`
	Form                *Coding                                `bson:"form,omitempty" json:"form,omitempty"`
	BenefitBalance      []EligibilityResponseBenefitsComponent `bson:"benefitBalance,omitempty" json:"benefitBalance,omitempty"`
	Error               []EligibilityResponseErrorsComponent   `bson:"error,omitempty" json:"error,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *EligibilityResponse) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "EligibilityResponse"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to EligibilityResponse), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *EligibilityResponse) GetBSON() (interface{}, error) {
	x.ResourceType = "EligibilityResponse"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "eligibilityResponse" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type eligibilityResponse EligibilityResponse

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *EligibilityResponse) UnmarshalJSON(data []byte) (err error) {
	x2 := eligibilityResponse{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = EligibilityResponse(x2)
		return x.checkResourceType()
	}
	return
}

func (x *EligibilityResponse) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "EligibilityResponse"
	} else if x.ResourceType != "EligibilityResponse" {
		return errors.New(fmt.Sprintf("Expected resourceType to be EligibilityResponse, instead received %s", x.ResourceType))
	}
	return nil
}

type EligibilityResponseBenefitsComponent struct {
	Category    *Coding                               `bson:"category,omitempty" json:"category,omitempty"`
	SubCategory *Coding                               `bson:"subCategory,omitempty" json:"subCategory,omitempty"`
	Network     *Coding                               `bson:"network,omitempty" json:"network,omitempty"`
	Unit        *Coding                               `bson:"unit,omitempty" json:"unit,omitempty"`
	Term        *Coding                               `bson:"term,omitempty" json:"term,omitempty"`
	Financial   []EligibilityResponseBenefitComponent `bson:"financial,omitempty" json:"financial,omitempty"`
}

type EligibilityResponseBenefitComponent struct {
	Type                   *Coding   `bson:"type,omitempty" json:"type,omitempty"`
	BenefitUnsignedInt     *uint32   `bson:"benefitUnsignedInt,omitempty" json:"benefitUnsignedInt,omitempty"`
	BenefitMoney           *Quantity `bson:"benefitMoney,omitempty" json:"benefitMoney,omitempty"`
	BenefitUsedUnsignedInt *uint32   `bson:"benefitUsedUnsignedInt,omitempty" json:"benefitUsedUnsignedInt,omitempty"`
	BenefitUsedMoney       *Quantity `bson:"benefitUsedMoney,omitempty" json:"benefitUsedMoney,omitempty"`
}

type EligibilityResponseErrorsComponent struct {
	Code *Coding `bson:"code,omitempty" json:"code,omitempty"`
}

type EligibilityResponsePlus struct {
	EligibilityResponse                     `bson:",inline"`
	EligibilityResponsePlusRelatedResources `bson:",inline"`
}

type EligibilityResponsePlusRelatedResources struct {
	IncludedEligibilityRequestResourcesReferencedByRequest       *[]EligibilityRequest    `bson:"_includedEligibilityRequestResourcesReferencedByRequest,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization        *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedPractitionerResourcesReferencedByRequestprovider     *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByRequestprovider,omitempty"`
	IncludedOrganizationResourcesReferencedByRequestorganization *[]Organization          `bson:"_includedOrganizationResourcesReferencedByRequestorganization,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref    *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref   *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingRequest          *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingRequest,omitempty"`
	RevIncludedPaymentNoticeResourcesReferencingResponse         *[]PaymentNotice         `bson:"_revIncludedPaymentNoticeResourcesReferencingResponse,omitempty"`
	RevIncludedOrderResponseResourcesReferencingFulfillment      *[]OrderResponse         `bson:"_revIncludedOrderResponseResourcesReferencingFulfillment,omitempty"`
	RevIncludedMessageHeaderResourcesReferencingData             *[]MessageHeader         `bson:"_revIncludedMessageHeaderResourcesReferencingData,omitempty"`
	RevIncludedProvenanceResourcesReferencingTarget              *[]Provenance            `bson:"_revIncludedProvenanceResourcesReferencingTarget,omitempty"`
	RevIncludedListResourcesReferencingItem                      *[]List                  `bson:"_revIncludedListResourcesReferencingItem,omitempty"`
	RevIncludedOrderResourcesReferencingDetail                   *[]Order                 `bson:"_revIncludedOrderResourcesReferencingDetail,omitempty"`
	RevIncludedBasicResourcesReferencingSubject                  *[]Basic                 `bson:"_revIncludedBasicResourcesReferencingSubject,omitempty"`
	RevIncludedAuditEventResourcesReferencingReference           *[]AuditEvent            `bson:"_revIncludedAuditEventResourcesReferencingReference,omitempty"`
	RevIncludedCompositionResourcesReferencingSubject            *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingSubject,omitempty"`
	RevIncludedCompositionResourcesReferencingEntry              *[]Composition           `bson:"_revIncludedCompositionResourcesReferencingEntry,omitempty"`
	RevIncludedDetectedIssueResourcesReferencingImplicated       *[]DetectedIssue         `bson:"_revIncludedDetectedIssueResourcesReferencingImplicated,omitempty"`
	RevIncludedQuestionnaireResponseResourcesReferencingSubject  *[]QuestionnaireResponse `bson:"_revIncludedQuestionnaireResponseResourcesReferencingSubject,omitempty"`
	RevIncludedProcessResponseResourcesReferencingRequest        *[]ProcessResponse       `bson:"_revIncludedProcessResponseResourcesReferencingRequest,omitempty"`
	RevIncludedClinicalImpressionResourcesReferencingTrigger     *[]ClinicalImpression    `bson:"_revIncludedClinicalImpressionResourcesReferencingTrigger,omitempty"`
}

func (e *EligibilityResponsePlusRelatedResources) GetIncludedEligibilityRequestResourceReferencedByRequest() (eligibilityRequest *EligibilityRequest, err error) {
	if e.IncludedEligibilityRequestResourcesReferencedByRequest == nil {
		err = errors.New("Included eligibilityrequests not requested")
	} else if len(*e.IncludedEligibilityRequestResourcesReferencedByRequest) > 1 {
		err = fmt.Errorf("Expected 0 or 1 eligibilityRequest, but found %d", len(*e.IncludedEligibilityRequestResourcesReferencedByRequest))
	} else if len(*e.IncludedEligibilityRequestResourcesReferencedByRequest) == 1 {
		eligibilityRequest = &(*e.IncludedEligibilityRequestResourcesReferencedByRequest)[0]
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*e.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetIncludedPractitionerResourceReferencedByRequestprovider() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByRequestprovider == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByRequestprovider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByRequestprovider))
	} else if len(*e.IncludedPractitionerResourcesReferencedByRequestprovider) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByRequestprovider)[0]
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetIncludedOrganizationResourceReferencedByRequestorganization() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByRequestorganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByRequestorganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByRequestorganization))
	} else if len(*e.IncludedOrganizationResourcesReferencedByRequestorganization) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByRequestorganization)[0]
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if e.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *e.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if e.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *e.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if e.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *e.RevIncludedListResourcesReferencingItem
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if e.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *e.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if e.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *e.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if e.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *e.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *e.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if e.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *e.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if e.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *e.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (e *EligibilityResponsePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedEligibilityRequestResourcesReferencedByRequest != nil {
		for _, r := range *e.IncludedEligibilityRequestResourcesReferencedByRequest {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for _, r := range *e.IncludedOrganizationResourcesReferencedByOrganization {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPractitionerResourcesReferencedByRequestprovider != nil {
		for _, r := range *e.IncludedPractitionerResourcesReferencedByRequestprovider {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedOrganizationResourcesReferencedByRequestorganization != nil {
		for _, r := range *e.IncludedOrganizationResourcesReferencedByRequestorganization {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (e *EligibilityResponsePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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

func (e *EligibilityResponsePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedEligibilityRequestResourcesReferencedByRequest != nil {
		for _, r := range *e.IncludedEligibilityRequestResourcesReferencedByRequest {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for _, r := range *e.IncludedOrganizationResourcesReferencedByOrganization {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPractitionerResourcesReferencedByRequestprovider != nil {
		for _, r := range *e.IncludedPractitionerResourcesReferencedByRequestprovider {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedOrganizationResourcesReferencedByRequestorganization != nil {
		for _, r := range *e.IncludedOrganizationResourcesReferencedByRequestorganization {
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
