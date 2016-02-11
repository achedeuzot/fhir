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

type Linkage struct {
	DomainResource `bson:",inline"`
	Author         *Reference             `bson:"author,omitempty" json:"author,omitempty"`
	Item           []LinkageItemComponent `bson:"item,omitempty" json:"item,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Linkage) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Linkage"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Linkage), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Linkage) GetBSON() (interface{}, error) {
	x.ResourceType = "Linkage"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "linkage" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type linkage Linkage

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Linkage) UnmarshalJSON(data []byte) (err error) {
	x2 := linkage{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Linkage(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Linkage) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Linkage"
	} else if x.ResourceType != "Linkage" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Linkage, instead received %s", x.ResourceType))
	}
	return nil
}

type LinkageItemComponent struct {
	Type     string     `bson:"type,omitempty" json:"type,omitempty"`
	Resource *Reference `bson:"resource,omitempty" json:"resource,omitempty"`
}

type LinkagePlus struct {
	Linkage                     `bson:",inline"`
	LinkagePlusRelatedResources `bson:",inline"`
}

type LinkagePlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByAuthor             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedOrganizationResourcesReferencedByAuthor             *[]Organization          `bson:"_includedOrganizationResourcesReferencedByAuthor,omitempty"`
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

func (l *LinkagePlusRelatedResources) GetIncludedPractitionerResourceReferencedByAuthor() (practitioner *Practitioner, err error) {
	if l.IncludedPractitionerResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*l.IncludedPractitionerResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*l.IncludedPractitionerResourcesReferencedByAuthor))
	} else if len(*l.IncludedPractitionerResourcesReferencedByAuthor) == 1 {
		practitioner = &(*l.IncludedPractitionerResourcesReferencedByAuthor)[0]
	}
	return
}

func (l *LinkagePlusRelatedResources) GetIncludedOrganizationResourceReferencedByAuthor() (organization *Organization, err error) {
	if l.IncludedOrganizationResourcesReferencedByAuthor == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*l.IncludedOrganizationResourcesReferencedByAuthor) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*l.IncludedOrganizationResourcesReferencedByAuthor))
	} else if len(*l.IncludedOrganizationResourcesReferencedByAuthor) == 1 {
		organization = &(*l.IncludedOrganizationResourcesReferencedByAuthor)[0]
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if l.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *l.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *l.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *l.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *l.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if l.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *l.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if l.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *l.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if l.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *l.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if l.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *l.RevIncludedListResourcesReferencingItem
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if l.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *l.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if l.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *l.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if l.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *l.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if l.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *l.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if l.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *l.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *l.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if l.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *l.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (l *LinkagePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if l.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *l.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (l *LinkagePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for _, r := range *l.IncludedPractitionerResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for _, r := range *l.IncludedOrganizationResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (l *LinkagePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *l.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *l.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *l.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *l.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *l.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *l.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *l.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *l.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *l.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *l.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *l.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *l.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *l.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *l.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *l.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *l.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (l *LinkagePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if l.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for _, r := range *l.IncludedPractitionerResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if l.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for _, r := range *l.IncludedOrganizationResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *l.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *l.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *l.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *l.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *l.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *l.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *l.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *l.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *l.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *l.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *l.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *l.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *l.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *l.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *l.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *l.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *l.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if l.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *l.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
