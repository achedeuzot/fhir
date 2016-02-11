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

type ExpansionProfile struct {
	DomainResource         `bson:",inline"`
	Url                    string                                `bson:"url,omitempty" json:"url,omitempty"`
	Identifier             *Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version                string                                `bson:"version,omitempty" json:"version,omitempty"`
	Name                   string                                `bson:"name,omitempty" json:"name,omitempty"`
	Status                 string                                `bson:"status,omitempty" json:"status,omitempty"`
	Experimental           *bool                                 `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher              string                                `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact                []ExpansionProfileContactComponent    `bson:"contact,omitempty" json:"contact,omitempty"`
	Date                   *FHIRDateTime                         `bson:"date,omitempty" json:"date,omitempty"`
	Description            string                                `bson:"description,omitempty" json:"description,omitempty"`
	CodeSystem             *ExpansionProfileCodeSystemComponent  `bson:"codeSystem,omitempty" json:"codeSystem,omitempty"`
	IncludeDesignations    *bool                                 `bson:"includeDesignations,omitempty" json:"includeDesignations,omitempty"`
	Designation            *ExpansionProfileDesignationComponent `bson:"designation,omitempty" json:"designation,omitempty"`
	IncludeDefinition      *bool                                 `bson:"includeDefinition,omitempty" json:"includeDefinition,omitempty"`
	IncludeInactive        *bool                                 `bson:"includeInactive,omitempty" json:"includeInactive,omitempty"`
	ExcludeNested          *bool                                 `bson:"excludeNested,omitempty" json:"excludeNested,omitempty"`
	ExcludeNotForUI        *bool                                 `bson:"excludeNotForUI,omitempty" json:"excludeNotForUI,omitempty"`
	ExcludePostCoordinated *bool                                 `bson:"excludePostCoordinated,omitempty" json:"excludePostCoordinated,omitempty"`
	DisplayLanguage        string                                `bson:"displayLanguage,omitempty" json:"displayLanguage,omitempty"`
	LimitedExpansion       *bool                                 `bson:"limitedExpansion,omitempty" json:"limitedExpansion,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ExpansionProfile) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ExpansionProfile"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ExpansionProfile), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ExpansionProfile) GetBSON() (interface{}, error) {
	x.ResourceType = "ExpansionProfile"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "expansionProfile" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type expansionProfile ExpansionProfile

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ExpansionProfile) UnmarshalJSON(data []byte) (err error) {
	x2 := expansionProfile{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ExpansionProfile(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ExpansionProfile) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ExpansionProfile"
	} else if x.ResourceType != "ExpansionProfile" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ExpansionProfile, instead received %s", x.ResourceType))
	}
	return nil
}

type ExpansionProfileContactComponent struct {
	Name    string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type ExpansionProfileCodeSystemComponent struct {
	Include *ExpansionProfileCodeSystemIncludeComponent `bson:"include,omitempty" json:"include,omitempty"`
	Exclude *ExpansionProfileCodeSystemExcludeComponent `bson:"exclude,omitempty" json:"exclude,omitempty"`
}

type ExpansionProfileCodeSystemIncludeComponent struct {
	CodeSystem []ExpansionProfileCodeSystemIncludeCodeSystemComponent `bson:"codeSystem,omitempty" json:"codeSystem,omitempty"`
}

type ExpansionProfileCodeSystemIncludeCodeSystemComponent struct {
	System  string `bson:"system,omitempty" json:"system,omitempty"`
	Version string `bson:"version,omitempty" json:"version,omitempty"`
}

type ExpansionProfileCodeSystemExcludeComponent struct {
	CodeSystem []ExpansionProfileCodeSystemExcludeCodeSystemComponent `bson:"codeSystem,omitempty" json:"codeSystem,omitempty"`
}

type ExpansionProfileCodeSystemExcludeCodeSystemComponent struct {
	System  string `bson:"system,omitempty" json:"system,omitempty"`
	Version string `bson:"version,omitempty" json:"version,omitempty"`
}

type ExpansionProfileDesignationComponent struct {
	Include *ExpansionProfileDesignationIncludeComponent `bson:"include,omitempty" json:"include,omitempty"`
	Exclude *ExpansionProfileDesignationExcludeComponent `bson:"exclude,omitempty" json:"exclude,omitempty"`
}

type ExpansionProfileDesignationIncludeComponent struct {
	Designation []ExpansionProfileDesignationIncludeDesignationComponent `bson:"designation,omitempty" json:"designation,omitempty"`
}

type ExpansionProfileDesignationIncludeDesignationComponent struct {
	Language string  `bson:"language,omitempty" json:"language,omitempty"`
	Use      *Coding `bson:"use,omitempty" json:"use,omitempty"`
}

type ExpansionProfileDesignationExcludeComponent struct {
	Designation []ExpansionProfileDesignationExcludeDesignationComponent `bson:"designation,omitempty" json:"designation,omitempty"`
}

type ExpansionProfileDesignationExcludeDesignationComponent struct {
	Language string  `bson:"language,omitempty" json:"language,omitempty"`
	Use      *Coding `bson:"use,omitempty" json:"use,omitempty"`
}

type ExpansionProfilePlus struct {
	ExpansionProfile                     `bson:",inline"`
	ExpansionProfilePlusRelatedResources `bson:",inline"`
}

type ExpansionProfilePlusRelatedResources struct {
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

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if e.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *e.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if e.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *e.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if e.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *e.RevIncludedListResourcesReferencingItem
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if e.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *e.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if e.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *e.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if e.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *e.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *e.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if e.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *e.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if e.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *e.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (e *ExpansionProfilePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (e *ExpansionProfilePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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

func (e *ExpansionProfilePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
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
