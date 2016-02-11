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

type Sequence struct {
	DomainResource   `bson:",inline"`
	Type             string                        `bson:"type,omitempty" json:"type,omitempty"`
	VariationID      []CodeableConcept             `bson:"variationID,omitempty" json:"variationID,omitempty"`
	ReferenceSeq     *CodeableConcept              `bson:"referenceSeq,omitempty" json:"referenceSeq,omitempty"`
	Quantity         *Quantity                     `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Coordinate       []SequenceCoordinateComponent `bson:"coordinate,omitempty" json:"coordinate,omitempty"`
	Species          *CodeableConcept              `bson:"species,omitempty" json:"species,omitempty"`
	ObservedAllele   string                        `bson:"observedAllele,omitempty" json:"observedAllele,omitempty"`
	ReferenceAllele  string                        `bson:"referenceAllele,omitempty" json:"referenceAllele,omitempty"`
	Cigar            string                        `bson:"cigar,omitempty" json:"cigar,omitempty"`
	Quality          []SequenceQualityComponent    `bson:"quality,omitempty" json:"quality,omitempty"`
	AllelicState     *CodeableConcept              `bson:"allelicState,omitempty" json:"allelicState,omitempty"`
	AllelicFrequency *float64                      `bson:"allelicFrequency,omitempty" json:"allelicFrequency,omitempty"`
	CopyNumberEvent  *CodeableConcept              `bson:"copyNumberEvent,omitempty" json:"copyNumberEvent,omitempty"`
	ReadCoverage     *int32                        `bson:"readCoverage,omitempty" json:"readCoverage,omitempty"`
	Chip             *SequenceChipComponent        `bson:"chip,omitempty" json:"chip,omitempty"`
	Repository       []SequenceRepositoryComponent `bson:"repository,omitempty" json:"repository,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Sequence) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Sequence"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Sequence), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Sequence) GetBSON() (interface{}, error) {
	x.ResourceType = "Sequence"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "sequence" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type sequence Sequence

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Sequence) UnmarshalJSON(data []byte) (err error) {
	x2 := sequence{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Sequence(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Sequence) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Sequence"
	} else if x.ResourceType != "Sequence" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Sequence, instead received %s", x.ResourceType))
	}
	return nil
}

type SequenceCoordinateComponent struct {
	Chromosome  *CodeableConcept `bson:"chromosome,omitempty" json:"chromosome,omitempty"`
	Start       *int32           `bson:"start,omitempty" json:"start,omitempty"`
	End         *int32           `bson:"end,omitempty" json:"end,omitempty"`
	GenomeBuild string           `bson:"genomeBuild,omitempty" json:"genomeBuild,omitempty"`
}

type SequenceQualityComponent struct {
	Start    *int32    `bson:"start,omitempty" json:"start,omitempty"`
	End      *int32    `bson:"end,omitempty" json:"end,omitempty"`
	Score    *Quantity `bson:"score,omitempty" json:"score,omitempty"`
	Platform string    `bson:"platform,omitempty" json:"platform,omitempty"`
}

type SequenceChipComponent struct {
	ChipId         string `bson:"chipId,omitempty" json:"chipId,omitempty"`
	ManufacturerId string `bson:"manufacturerId,omitempty" json:"manufacturerId,omitempty"`
	Version        string `bson:"version,omitempty" json:"version,omitempty"`
}

type SequenceRepositoryComponent struct {
	Url            string `bson:"url,omitempty" json:"url,omitempty"`
	Name           string `bson:"name,omitempty" json:"name,omitempty"`
	Structure      string `bson:"structure,omitempty" json:"structure,omitempty"`
	VariantId      string `bson:"variantId,omitempty" json:"variantId,omitempty"`
	ReadGroupSetId string `bson:"readGroupSetId,omitempty" json:"readGroupSetId,omitempty"`
}

type SequencePlus struct {
	Sequence                     `bson:",inline"`
	SequencePlusRelatedResources `bson:",inline"`
}

type SequencePlusRelatedResources struct {
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

func (s *SequencePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *s.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *s.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *s.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *s.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if s.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *s.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if s.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *s.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if s.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *s.RevIncludedListResourcesReferencingItem
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if s.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *s.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if s.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *s.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if s.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *s.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if s.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *s.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *s.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if s.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *s.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (s *SequencePlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *s.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (s *SequencePlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (s *SequencePlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *s.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *s.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *s.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *s.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *s.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *s.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *s.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *s.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *s.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *s.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *s.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (s *SequencePlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if s.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *s.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *s.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *s.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *s.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *s.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *s.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *s.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *s.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *s.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *s.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *s.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *s.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *s.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *s.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *s.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if s.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *s.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
