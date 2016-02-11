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

type DecisionSupportRule struct {
	DomainResource `bson:",inline"`
	ModuleMetadata *ModuleMetadata                       `bson:"moduleMetadata,omitempty" json:"moduleMetadata,omitempty"`
	Library        []Reference                           `bson:"library,omitempty" json:"library,omitempty"`
	Trigger        []DecisionSupportRuleTriggerComponent `bson:"trigger,omitempty" json:"trigger,omitempty"`
	Condition      string                                `bson:"condition,omitempty" json:"condition,omitempty"`
	Action         []DecisionSupportRuleActionComponent  `bson:"action,omitempty" json:"action,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DecisionSupportRule) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DecisionSupportRule"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DecisionSupportRule), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DecisionSupportRule) GetBSON() (interface{}, error) {
	x.ResourceType = "DecisionSupportRule"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "decisionSupportRule" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type decisionSupportRule DecisionSupportRule

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DecisionSupportRule) UnmarshalJSON(data []byte) (err error) {
	x2 := decisionSupportRule{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DecisionSupportRule(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DecisionSupportRule) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DecisionSupportRule"
	} else if x.ResourceType != "DecisionSupportRule" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DecisionSupportRule, instead received %s", x.ResourceType))
	}
	return nil
}

type DecisionSupportRuleTriggerComponent struct {
	Type                 string        `bson:"type,omitempty" json:"type,omitempty"`
	EventName            string        `bson:"eventName,omitempty" json:"eventName,omitempty"`
	EventTimingTiming    *Timing       `bson:"eventTimingTiming,omitempty" json:"eventTimingTiming,omitempty"`
	EventTimingReference *Reference    `bson:"eventTimingReference,omitempty" json:"eventTimingReference,omitempty"`
	EventTimingDate      *FHIRDateTime `bson:"eventTimingDate,omitempty" json:"eventTimingDate,omitempty"`
	EventTimingDateTime  *FHIRDateTime `bson:"eventTimingDateTime,omitempty" json:"eventTimingDateTime,omitempty"`
}

type DecisionSupportRuleActionComponent struct {
	ActionIdentifier   *Identifier                                       `bson:"actionIdentifier,omitempty" json:"actionIdentifier,omitempty"`
	Number             string                                            `bson:"number,omitempty" json:"number,omitempty"`
	SupportingEvidence []Attachment                                      `bson:"supportingEvidence,omitempty" json:"supportingEvidence,omitempty"`
	Documentation      []Attachment                                      `bson:"documentation,omitempty" json:"documentation,omitempty"`
	ParticipantType    []string                                          `bson:"participantType,omitempty" json:"participantType,omitempty"`
	Title              string                                            `bson:"title,omitempty" json:"title,omitempty"`
	Description        string                                            `bson:"description,omitempty" json:"description,omitempty"`
	TextEquivalent     string                                            `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty"`
	Concept            []CodeableConcept                                 `bson:"concept,omitempty" json:"concept,omitempty"`
	Type               string                                            `bson:"type,omitempty" json:"type,omitempty"`
	Resource           *Reference                                        `bson:"resource,omitempty" json:"resource,omitempty"`
	Customization      []DecisionSupportRuleActionCustomizationComponent `bson:"customization,omitempty" json:"customization,omitempty"`
	Actions            []DecisionSupportRuleActionComponent              `bson:"actions,omitempty" json:"actions,omitempty"`
}

type DecisionSupportRuleActionCustomizationComponent struct {
	Path       string `bson:"path,omitempty" json:"path,omitempty"`
	Expression string `bson:"expression,omitempty" json:"expression,omitempty"`
}

type DecisionSupportRulePlus struct {
	DecisionSupportRule                     `bson:",inline"`
	DecisionSupportRulePlusRelatedResources `bson:",inline"`
}

type DecisionSupportRulePlusRelatedResources struct {
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

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if d.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *d.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingItem
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if d.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *d.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if d.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *d.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *d.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if d.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *d.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (d *DecisionSupportRulePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (d *DecisionSupportRulePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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
	return resourceMap
}

func (d *DecisionSupportRulePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
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
	return resourceMap
}
