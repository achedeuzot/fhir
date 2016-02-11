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

type OrderSet struct {
	DomainResource `bson:",inline"`
	ModuleMetadata *ModuleMetadata         `bson:"moduleMetadata,omitempty" json:"moduleMetadata,omitempty"`
	Library        []Reference             `bson:"library,omitempty" json:"library,omitempty"`
	Item           []OrderSetItemComponent `bson:"item,omitempty" json:"item,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *OrderSet) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "OrderSet"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to OrderSet), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *OrderSet) GetBSON() (interface{}, error) {
	x.ResourceType = "OrderSet"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "orderSet" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type orderSet OrderSet

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *OrderSet) UnmarshalJSON(data []byte) (err error) {
	x2 := orderSet{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = OrderSet(x2)
		return x.checkResourceType()
	}
	return
}

func (x *OrderSet) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "OrderSet"
	} else if x.ResourceType != "OrderSet" {
		return errors.New(fmt.Sprintf("Expected resourceType to be OrderSet, instead received %s", x.ResourceType))
	}
	return nil
}

type OrderSetItemComponent struct {
	Identifier          *Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Number              string                               `bson:"number,omitempty" json:"number,omitempty"`
	Title               string                               `bson:"title,omitempty" json:"title,omitempty"`
	Description         string                               `bson:"description,omitempty" json:"description,omitempty"`
	TextEquivalent      string                               `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty"`
	SupportingEvidence  []Attachment                         `bson:"supportingEvidence,omitempty" json:"supportingEvidence,omitempty"`
	Documentation       []Attachment                         `bson:"documentation,omitempty" json:"documentation,omitempty"`
	ParticipantType     []string                             `bson:"participantType,omitempty" json:"participantType,omitempty"`
	Concept             []CodeableConcept                    `bson:"concept,omitempty" json:"concept,omitempty"`
	Type                string                               `bson:"type,omitempty" json:"type,omitempty"`
	GroupingBehavior    string                               `bson:"groupingBehavior,omitempty" json:"groupingBehavior,omitempty"`
	SelectionBehavior   string                               `bson:"selectionBehavior,omitempty" json:"selectionBehavior,omitempty"`
	RequiredBehavior    string                               `bson:"requiredBehavior,omitempty" json:"requiredBehavior,omitempty"`
	PrecheckBehavior    string                               `bson:"precheckBehavior,omitempty" json:"precheckBehavior,omitempty"`
	CardinalityBehavior string                               `bson:"cardinalityBehavior,omitempty" json:"cardinalityBehavior,omitempty"`
	Resource            *Reference                           `bson:"resource,omitempty" json:"resource,omitempty"`
	Customization       []OrderSetItemCustomizationComponent `bson:"customization,omitempty" json:"customization,omitempty"`
	Items               []OrderSetItemComponent              `bson:"items,omitempty" json:"items,omitempty"`
}

type OrderSetItemCustomizationComponent struct {
	Path       string `bson:"path,omitempty" json:"path,omitempty"`
	Expression string `bson:"expression,omitempty" json:"expression,omitempty"`
}

type OrderSetPlus struct {
	OrderSet                     `bson:",inline"`
	OrderSetPlusRelatedResources `bson:",inline"`
}

type OrderSetPlusRelatedResources struct {
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

func (o *OrderSetPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *o.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *o.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *o.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if o.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *o.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if o.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *o.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if o.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *o.RevIncludedListResourcesReferencingItem
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if o.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *o.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if o.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *o.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if o.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *o.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if o.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *o.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *o.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if o.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *o.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *o.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (o *OrderSetPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (o *OrderSetPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *o.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *o.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *o.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *o.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *o.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *o.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *o.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *o.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *o.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *o.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *o.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (o *OrderSetPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if o.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *o.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *o.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *o.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *o.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *o.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *o.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *o.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *o.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *o.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *o.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *o.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *o.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *o.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *o.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *o.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if o.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *o.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
