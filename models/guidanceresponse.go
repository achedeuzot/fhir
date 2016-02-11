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

type GuidanceResponse struct {
	DomainResource    `bson:",inline"`
	RequestId         string                            `bson:"requestId,omitempty" json:"requestId,omitempty"`
	Module            *Reference                        `bson:"module,omitempty" json:"module,omitempty"`
	Status            string                            `bson:"status,omitempty" json:"status,omitempty"`
	EvaluationMessage []Reference                       `bson:"evaluationMessage,omitempty" json:"evaluationMessage,omitempty"`
	OutputParameters  *Reference                        `bson:"outputParameters,omitempty" json:"outputParameters,omitempty"`
	Action            []GuidanceResponseActionComponent `bson:"action,omitempty" json:"action,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *GuidanceResponse) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "GuidanceResponse"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to GuidanceResponse), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *GuidanceResponse) GetBSON() (interface{}, error) {
	x.ResourceType = "GuidanceResponse"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "guidanceResponse" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type guidanceResponse GuidanceResponse

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *GuidanceResponse) UnmarshalJSON(data []byte) (err error) {
	x2 := guidanceResponse{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = GuidanceResponse(x2)
		return x.checkResourceType()
	}
	return
}

func (x *GuidanceResponse) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "GuidanceResponse"
	} else if x.ResourceType != "GuidanceResponse" {
		return errors.New(fmt.Sprintf("Expected resourceType to be GuidanceResponse, instead received %s", x.ResourceType))
	}
	return nil
}

type GuidanceResponseActionComponent struct {
	ActionIdentifier   *Identifier                       `bson:"actionIdentifier,omitempty" json:"actionIdentifier,omitempty"`
	Number             string                            `bson:"number,omitempty" json:"number,omitempty"`
	SupportingEvidence []Attachment                      `bson:"supportingEvidence,omitempty" json:"supportingEvidence,omitempty"`
	Documentation      []Attachment                      `bson:"documentation,omitempty" json:"documentation,omitempty"`
	Participant        []Reference                       `bson:"participant,omitempty" json:"participant,omitempty"`
	Title              string                            `bson:"title,omitempty" json:"title,omitempty"`
	Description        string                            `bson:"description,omitempty" json:"description,omitempty"`
	TextEquivalent     string                            `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty"`
	Concept            []CodeableConcept                 `bson:"concept,omitempty" json:"concept,omitempty"`
	Type               string                            `bson:"type,omitempty" json:"type,omitempty"`
	Resource           *Reference                        `bson:"resource,omitempty" json:"resource,omitempty"`
	Actions            []GuidanceResponseActionComponent `bson:"actions,omitempty" json:"actions,omitempty"`
}

type GuidanceResponsePlus struct {
	GuidanceResponse                     `bson:",inline"`
	GuidanceResponsePlusRelatedResources `bson:",inline"`
}

type GuidanceResponsePlusRelatedResources struct {
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

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if g.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *g.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if g.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *g.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if g.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *g.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if g.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *g.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if g.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *g.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if g.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *g.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if g.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *g.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if g.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *g.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if g.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *g.RevIncludedListResourcesReferencingItem
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if g.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *g.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if g.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *g.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if g.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *g.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if g.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *g.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if g.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *g.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if g.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *g.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if g.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *g.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if g.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *g.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if g.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *g.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (g *GuidanceResponsePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (g *GuidanceResponsePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *g.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *g.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *g.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *g.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *g.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *g.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *g.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *g.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *g.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *g.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *g.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *g.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *g.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *g.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *g.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (g *GuidanceResponsePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if g.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *g.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *g.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *g.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *g.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *g.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *g.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *g.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *g.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *g.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *g.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *g.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *g.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *g.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *g.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *g.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if g.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *g.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
