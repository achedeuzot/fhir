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

type Measure struct {
	DomainResource   `bson:",inline"`
	Identifier       []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version          string                       `bson:"version,omitempty" json:"version,omitempty"`
	ModuleMetadata   *ModuleMetadata              `bson:"moduleMetadata,omitempty" json:"moduleMetadata,omitempty"`
	Library          []Reference                  `bson:"library,omitempty" json:"library,omitempty"`
	Population       []MeasurePopulationComponent `bson:"population,omitempty" json:"population,omitempty"`
	Stratifier       []string                     `bson:"stratifier,omitempty" json:"stratifier,omitempty"`
	SupplementalData []string                     `bson:"supplementalData,omitempty" json:"supplementalData,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Measure) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Measure"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Measure), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Measure) GetBSON() (interface{}, error) {
	x.ResourceType = "Measure"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "measure" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type measure Measure

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Measure) UnmarshalJSON(data []byte) (err error) {
	x2 := measure{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Measure(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Measure) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Measure"
	} else if x.ResourceType != "Measure" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Measure, instead received %s", x.ResourceType))
	}
	return nil
}

type MeasurePopulationComponent struct {
	Type        string `bson:"type,omitempty" json:"type,omitempty"`
	Name        string `bson:"name,omitempty" json:"name,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
	Criteria    string `bson:"criteria,omitempty" json:"criteria,omitempty"`
}

type MeasurePlus struct {
	Measure                     `bson:",inline"`
	MeasurePlusRelatedResources `bson:",inline"`
}

type MeasurePlusRelatedResources struct {
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

func (m *MeasurePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *m.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *m.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if m.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *m.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if m.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *m.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if m.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *m.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if m.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *m.RevIncludedListResourcesReferencingItem
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if m.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *m.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if m.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *m.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if m.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *m.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if m.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *m.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *m.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if m.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *m.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (m *MeasurePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if m.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *m.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (m *MeasurePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (m *MeasurePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *m.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *m.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *m.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *m.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *m.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *m.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *m.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *m.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *m.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *m.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *m.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *m.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *m.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (m *MeasurePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if m.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *m.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *m.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *m.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *m.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *m.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *m.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *m.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *m.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *m.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *m.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *m.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *m.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *m.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *m.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *m.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *m.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *m.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if m.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *m.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
