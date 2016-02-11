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

type CodeSystem struct {
	DomainResource `bson:",inline"`
	Url            string                                 `bson:"url,omitempty" json:"url,omitempty"`
	Identifier     *Identifier                            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version        string                                 `bson:"version,omitempty" json:"version,omitempty"`
	Name           string                                 `bson:"name,omitempty" json:"name,omitempty"`
	Status         string                                 `bson:"status,omitempty" json:"status,omitempty"`
	Experimental   *bool                                  `bson:"experimental,omitempty" json:"experimental,omitempty"`
	Publisher      string                                 `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact        []CodeSystemContactComponent           `bson:"contact,omitempty" json:"contact,omitempty"`
	Date           *FHIRDateTime                          `bson:"date,omitempty" json:"date,omitempty"`
	Description    string                                 `bson:"description,omitempty" json:"description,omitempty"`
	UseContext     []CodeableConcept                      `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Requirements   string                                 `bson:"requirements,omitempty" json:"requirements,omitempty"`
	Copyright      string                                 `bson:"copyright,omitempty" json:"copyright,omitempty"`
	CaseSensitive  *bool                                  `bson:"caseSensitive,omitempty" json:"caseSensitive,omitempty"`
	ValueSet       string                                 `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	Compositional  *bool                                  `bson:"compositional,omitempty" json:"compositional,omitempty"`
	VersionNeeded  *bool                                  `bson:"versionNeeded,omitempty" json:"versionNeeded,omitempty"`
	Content        string                                 `bson:"content,omitempty" json:"content,omitempty"`
	Count          *uint32                                `bson:"count,omitempty" json:"count,omitempty"`
	Filter         []CodeSystemFilterComponent            `bson:"filter,omitempty" json:"filter,omitempty"`
	Property       []CodeSystemPropertyComponent          `bson:"property,omitempty" json:"property,omitempty"`
	Concept        []CodeSystemConceptDefinitionComponent `bson:"concept,omitempty" json:"concept,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *CodeSystem) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "CodeSystem"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to CodeSystem), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *CodeSystem) GetBSON() (interface{}, error) {
	x.ResourceType = "CodeSystem"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "codeSystem" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type codeSystem CodeSystem

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *CodeSystem) UnmarshalJSON(data []byte) (err error) {
	x2 := codeSystem{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = CodeSystem(x2)
		return x.checkResourceType()
	}
	return
}

func (x *CodeSystem) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "CodeSystem"
	} else if x.ResourceType != "CodeSystem" {
		return errors.New(fmt.Sprintf("Expected resourceType to be CodeSystem, instead received %s", x.ResourceType))
	}
	return nil
}

type CodeSystemContactComponent struct {
	Name    string         `bson:"name,omitempty" json:"name,omitempty"`
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

type CodeSystemFilterComponent struct {
	Code        string   `bson:"code,omitempty" json:"code,omitempty"`
	Description string   `bson:"description,omitempty" json:"description,omitempty"`
	Operator    []string `bson:"operator,omitempty" json:"operator,omitempty"`
	Value       string   `bson:"value,omitempty" json:"value,omitempty"`
}

type CodeSystemPropertyComponent struct {
	Code        string `bson:"code,omitempty" json:"code,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
	Type        string `bson:"type,omitempty" json:"type,omitempty"`
}

type CodeSystemConceptDefinitionComponent struct {
	Code        string                                            `bson:"code,omitempty" json:"code,omitempty"`
	Display     string                                            `bson:"display,omitempty" json:"display,omitempty"`
	Definition  string                                            `bson:"definition,omitempty" json:"definition,omitempty"`
	Designation []CodeSystemConceptDefinitionDesignationComponent `bson:"designation,omitempty" json:"designation,omitempty"`
	Property    []CodeSystemConceptDefinitionPropertyComponent    `bson:"property,omitempty" json:"property,omitempty"`
	Concept     []CodeSystemConceptDefinitionComponent            `bson:"concept,omitempty" json:"concept,omitempty"`
}

type CodeSystemConceptDefinitionDesignationComponent struct {
	Language string  `bson:"language,omitempty" json:"language,omitempty"`
	Use      *Coding `bson:"use,omitempty" json:"use,omitempty"`
	Value    string  `bson:"value,omitempty" json:"value,omitempty"`
}

type CodeSystemConceptDefinitionPropertyComponent struct {
	Code         string  `bson:"code,omitempty" json:"code,omitempty"`
	ValueCode    string  `bson:"valueCode,omitempty" json:"valueCode,omitempty"`
	ValueCoding  *Coding `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
	ValueString  string  `bson:"valueString,omitempty" json:"valueString,omitempty"`
	ValueInteger *int32  `bson:"valueInteger,omitempty" json:"valueInteger,omitempty"`
}

type CodeSystemPlus struct {
	CodeSystem                     `bson:",inline"`
	CodeSystemPlusRelatedResources `bson:",inline"`
}

type CodeSystemPlusRelatedResources struct {
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

func (c *CodeSystemPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *c.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if c.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *c.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (c *CodeSystemPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	return resourceMap
}

func (c *CodeSystemPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *c.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *c.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *c.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *c.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *c.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *c.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *c.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (c *CodeSystemPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *c.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *c.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *c.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *c.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *c.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *c.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *c.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *c.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *c.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *c.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *c.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *c.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *c.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *c.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *c.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *c.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
