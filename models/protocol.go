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

type Protocol struct {
	DomainResource `bson:",inline"`
	Identifier     []Identifier            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Title          string                  `bson:"title,omitempty" json:"title,omitempty"`
	Status         string                  `bson:"status,omitempty" json:"status,omitempty"`
	Type           string                  `bson:"type,omitempty" json:"type,omitempty"`
	Subject        *Reference              `bson:"subject,omitempty" json:"subject,omitempty"`
	Group          *Reference              `bson:"group,omitempty" json:"group,omitempty"`
	Purpose        string                  `bson:"purpose,omitempty" json:"purpose,omitempty"`
	Author         *Reference              `bson:"author,omitempty" json:"author,omitempty"`
	Step           []ProtocolStepComponent `bson:"step,omitempty" json:"step,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Protocol) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Protocol"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Protocol), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Protocol) GetBSON() (interface{}, error) {
	x.ResourceType = "Protocol"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "protocol" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type protocol Protocol

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Protocol) UnmarshalJSON(data []byte) (err error) {
	x2 := protocol{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Protocol(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Protocol) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Protocol"
	} else if x.ResourceType != "Protocol" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Protocol, instead received %s", x.ResourceType))
	}
	return nil
}

type ProtocolStepComponent struct {
	Name          string                             `bson:"name,omitempty" json:"name,omitempty"`
	Description   string                             `bson:"description,omitempty" json:"description,omitempty"`
	Duration      *Quantity                          `bson:"duration,omitempty" json:"duration,omitempty"`
	Precondition  *ProtocolStepPreconditionComponent `bson:"precondition,omitempty" json:"precondition,omitempty"`
	Exit          *ProtocolStepPreconditionComponent `bson:"exit,omitempty" json:"exit,omitempty"`
	FirstActivity string                             `bson:"firstActivity,omitempty" json:"firstActivity,omitempty"`
	Activity      []ProtocolStepActivityComponent    `bson:"activity,omitempty" json:"activity,omitempty"`
	Next          []ProtocolStepNextComponent        `bson:"next,omitempty" json:"next,omitempty"`
}

type ProtocolStepPreconditionComponent struct {
	Description  string                                      `bson:"description,omitempty" json:"description,omitempty"`
	Condition    *ProtocolStepPreconditionConditionComponent `bson:"condition,omitempty" json:"condition,omitempty"`
	Intersection []ProtocolStepPreconditionComponent         `bson:"intersection,omitempty" json:"intersection,omitempty"`
	Union        []ProtocolStepPreconditionComponent         `bson:"union,omitempty" json:"union,omitempty"`
	Exclude      []ProtocolStepPreconditionComponent         `bson:"exclude,omitempty" json:"exclude,omitempty"`
}

type ProtocolStepPreconditionConditionComponent struct {
	Type                 *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	ValueCodeableConcept *CodeableConcept `bson:"valueCodeableConcept,omitempty" json:"valueCodeableConcept,omitempty"`
	ValueBoolean         *bool            `bson:"valueBoolean,omitempty" json:"valueBoolean,omitempty"`
	ValueSimpleQuantity  *Quantity        `bson:"valueSimpleQuantity,omitempty" json:"valueSimpleQuantity,omitempty"`
	ValueRange           *Range           `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
}

type ProtocolStepActivityComponent struct {
	Alternative []string                                 `bson:"alternative,omitempty" json:"alternative,omitempty"`
	Component   []ProtocolStepActivityComponentComponent `bson:"component,omitempty" json:"component,omitempty"`
	Following   []string                                 `bson:"following,omitempty" json:"following,omitempty"`
	Wait        *Quantity                                `bson:"wait,omitempty" json:"wait,omitempty"`
	Detail      *ProtocolStepActivityDetailComponent     `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ProtocolStepActivityComponentComponent struct {
	Sequence *int32 `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Activity string `bson:"activity,omitempty" json:"activity,omitempty"`
}

type ProtocolStepActivityDetailComponent struct {
	Category              string           `bson:"category,omitempty" json:"category,omitempty"`
	Code                  *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	TimingCodeableConcept *CodeableConcept `bson:"timingCodeableConcept,omitempty" json:"timingCodeableConcept,omitempty"`
	TimingTiming          *Timing          `bson:"timingTiming,omitempty" json:"timingTiming,omitempty"`
	Location              *Reference       `bson:"location,omitempty" json:"location,omitempty"`
	Performer             []Reference      `bson:"performer,omitempty" json:"performer,omitempty"`
	Product               *Reference       `bson:"product,omitempty" json:"product,omitempty"`
	Quantity              *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Description           string           `bson:"description,omitempty" json:"description,omitempty"`
}

type ProtocolStepNextComponent struct {
	Description string                             `bson:"description,omitempty" json:"description,omitempty"`
	Reference   string                             `bson:"reference,omitempty" json:"reference,omitempty"`
	Condition   *ProtocolStepPreconditionComponent `bson:"condition,omitempty" json:"condition,omitempty"`
}

type ProtocolPlus struct {
	Protocol                     `bson:",inline"`
	ProtocolPlusRelatedResources `bson:",inline"`
}

type ProtocolPlusRelatedResources struct {
	IncludedConditionResourcesReferencedBySubject               *[]Condition             `bson:"_includedConditionResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                  *[]Device                `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedMedicationResourcesReferencedBySubject              *[]Medication            `bson:"_includedMedicationResourcesReferencedBySubject,omitempty"`
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

func (p *ProtocolPlusRelatedResources) GetIncludedConditionResourceReferencedBySubject() (condition *Condition, err error) {
	if p.IncludedConditionResourcesReferencedBySubject == nil {
		err = errors.New("Included conditions not requested")
	} else if len(*p.IncludedConditionResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 condition, but found %d", len(*p.IncludedConditionResourcesReferencedBySubject))
	} else if len(*p.IncludedConditionResourcesReferencedBySubject) == 1 {
		condition = &(*p.IncludedConditionResourcesReferencedBySubject)[0]
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if p.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*p.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*p.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*p.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*p.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetIncludedMedicationResourceReferencedBySubject() (medication *Medication, err error) {
	if p.IncludedMedicationResourcesReferencedBySubject == nil {
		err = errors.New("Included medications not requested")
	} else if len(*p.IncludedMedicationResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 medication, but found %d", len(*p.IncludedMedicationResourcesReferencedBySubject))
	} else if len(*p.IncludedMedicationResourcesReferencedBySubject) == 1 {
		medication = &(*p.IncludedMedicationResourcesReferencedBySubject)[0]
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *p.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *p.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if p.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *p.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if p.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *p.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if p.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *p.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if p.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *p.RevIncludedListResourcesReferencingItem
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if p.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *p.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if p.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *p.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if p.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *p.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if p.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *p.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *p.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if p.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *p.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if p.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *p.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (p *ProtocolPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedConditionResourcesReferencedBySubject != nil {
		for _, r := range *p.IncludedConditionResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedDeviceResourcesReferencedBySubject != nil {
		for _, r := range *p.IncludedDeviceResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedMedicationResourcesReferencedBySubject != nil {
		for _, r := range *p.IncludedMedicationResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (p *ProtocolPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *p.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *p.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *p.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *p.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *p.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *p.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *p.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *p.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *p.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (p *ProtocolPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if p.IncludedConditionResourcesReferencedBySubject != nil {
		for _, r := range *p.IncludedConditionResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedDeviceResourcesReferencedBySubject != nil {
		for _, r := range *p.IncludedDeviceResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.IncludedMedicationResourcesReferencedBySubject != nil {
		for _, r := range *p.IncludedMedicationResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *p.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *p.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *p.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *p.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *p.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *p.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *p.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *p.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *p.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *p.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *p.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *p.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *p.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *p.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if p.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *p.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
