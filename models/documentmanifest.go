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

type DocumentManifest struct {
	DomainResource   `bson:",inline"`
	MasterIdentifier *Identifier                        `bson:"masterIdentifier,omitempty" json:"masterIdentifier,omitempty"`
	Identifier       []Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject          *Reference                         `bson:"subject,omitempty" json:"subject,omitempty"`
	Recipient        []Reference                        `bson:"recipient,omitempty" json:"recipient,omitempty"`
	Type             *CodeableConcept                   `bson:"type,omitempty" json:"type,omitempty"`
	Author           []Reference                        `bson:"author,omitempty" json:"author,omitempty"`
	Created          *FHIRDateTime                      `bson:"created,omitempty" json:"created,omitempty"`
	Source           string                             `bson:"source,omitempty" json:"source,omitempty"`
	Status           string                             `bson:"status,omitempty" json:"status,omitempty"`
	Description      string                             `bson:"description,omitempty" json:"description,omitempty"`
	Content          []DocumentManifestContentComponent `bson:"content,omitempty" json:"content,omitempty"`
	Related          []DocumentManifestRelatedComponent `bson:"related,omitempty" json:"related,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DocumentManifest) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DocumentManifest"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DocumentManifest), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DocumentManifest) GetBSON() (interface{}, error) {
	x.ResourceType = "DocumentManifest"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "documentManifest" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type documentManifest DocumentManifest

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *DocumentManifest) UnmarshalJSON(data []byte) (err error) {
	x2 := documentManifest{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = DocumentManifest(x2)
		return x.checkResourceType()
	}
	return
}

func (x *DocumentManifest) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "DocumentManifest"
	} else if x.ResourceType != "DocumentManifest" {
		return errors.New(fmt.Sprintf("Expected resourceType to be DocumentManifest, instead received %s", x.ResourceType))
	}
	return nil
}

type DocumentManifestContentComponent struct {
	PAttachment *Attachment `bson:"pAttachment,omitempty" json:"pAttachment,omitempty"`
	PReference  *Reference  `bson:"pReference,omitempty" json:"pReference,omitempty"`
}

type DocumentManifestRelatedComponent struct {
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ref        *Reference  `bson:"ref,omitempty" json:"ref,omitempty"`
}

type DocumentManifestPlus struct {
	DocumentManifest                     `bson:",inline"`
	DocumentManifestPlusRelatedResources `bson:",inline"`
}

type DocumentManifestPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedBySubject            *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedBySubject,omitempty"`
	IncludedGroupResourcesReferencedBySubject                   *[]Group                 `bson:"_includedGroupResourcesReferencedBySubject,omitempty"`
	IncludedDeviceResourcesReferencedBySubject                  *[]Device                `bson:"_includedDeviceResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedPractitionerResourcesReferencedByAuthor             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByAuthor,omitempty"`
	IncludedOrganizationResourcesReferencedByAuthor             *[]Organization          `bson:"_includedOrganizationResourcesReferencedByAuthor,omitempty"`
	IncludedDeviceResourcesReferencedByAuthor                   *[]Device                `bson:"_includedDeviceResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByAuthor                  *[]Patient               `bson:"_includedPatientResourcesReferencedByAuthor,omitempty"`
	IncludedRelatedPersonResourcesReferencedByAuthor            *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByAuthor,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByRecipient          *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByRecipient,omitempty"`
	IncludedOrganizationResourcesReferencedByRecipient          *[]Organization          `bson:"_includedOrganizationResourcesReferencedByRecipient,omitempty"`
	IncludedPatientResourcesReferencedByRecipient               *[]Patient               `bson:"_includedPatientResourcesReferencedByRecipient,omitempty"`
	IncludedRelatedPersonResourcesReferencedByRecipient         *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByRecipient,omitempty"`
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

func (d *DocumentManifestPlusRelatedResources) GetIncludedPractitionerResourceReferencedBySubject() (practitioner *Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedBySubject == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*d.IncludedPractitionerResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*d.IncludedPractitionerResourcesReferencedBySubject))
	} else if len(*d.IncludedPractitionerResourcesReferencedBySubject) == 1 {
		practitioner = &(*d.IncludedPractitionerResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetIncludedGroupResourceReferencedBySubject() (group *Group, err error) {
	if d.IncludedGroupResourcesReferencedBySubject == nil {
		err = errors.New("Included groups not requested")
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*d.IncludedGroupResourcesReferencedBySubject))
	} else if len(*d.IncludedGroupResourcesReferencedBySubject) == 1 {
		group = &(*d.IncludedGroupResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetIncludedDeviceResourceReferencedBySubject() (device *Device, err error) {
	if d.IncludedDeviceResourcesReferencedBySubject == nil {
		err = errors.New("Included devices not requested")
	} else if len(*d.IncludedDeviceResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*d.IncludedDeviceResourcesReferencedBySubject))
	} else if len(*d.IncludedDeviceResourcesReferencedBySubject) == 1 {
		device = &(*d.IncludedDeviceResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetIncludedPatientResourceReferencedBySubject() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedBySubject))
	} else if len(*d.IncludedPatientResourcesReferencedBySubject) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedBySubject)[0]
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByAuthor() (practitioners []Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedByAuthor == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *d.IncludedPractitionerResourcesReferencedByAuthor
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByAuthor() (organizations []Organization, err error) {
	if d.IncludedOrganizationResourcesReferencedByAuthor == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *d.IncludedOrganizationResourcesReferencedByAuthor
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetIncludedDeviceResourcesReferencedByAuthor() (devices []Device, err error) {
	if d.IncludedDeviceResourcesReferencedByAuthor == nil {
		err = errors.New("Included devices not requested")
	} else {
		devices = *d.IncludedDeviceResourcesReferencedByAuthor
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetIncludedPatientResourcesReferencedByAuthor() (patients []Patient, err error) {
	if d.IncludedPatientResourcesReferencedByAuthor == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *d.IncludedPatientResourcesReferencedByAuthor
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByAuthor() (relatedPeople []RelatedPerson, err error) {
	if d.IncludedRelatedPersonResourcesReferencedByAuthor == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *d.IncludedRelatedPersonResourcesReferencedByAuthor
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if d.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*d.IncludedPatientResourcesReferencedByPatient))
	} else if len(*d.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*d.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetIncludedPractitionerResourcesReferencedByRecipient() (practitioners []Practitioner, err error) {
	if d.IncludedPractitionerResourcesReferencedByRecipient == nil {
		err = errors.New("Included practitioners not requested")
	} else {
		practitioners = *d.IncludedPractitionerResourcesReferencedByRecipient
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetIncludedOrganizationResourcesReferencedByRecipient() (organizations []Organization, err error) {
	if d.IncludedOrganizationResourcesReferencedByRecipient == nil {
		err = errors.New("Included organizations not requested")
	} else {
		organizations = *d.IncludedOrganizationResourcesReferencedByRecipient
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetIncludedPatientResourcesReferencedByRecipient() (patients []Patient, err error) {
	if d.IncludedPatientResourcesReferencedByRecipient == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *d.IncludedPatientResourcesReferencedByRecipient
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetIncludedRelatedPersonResourcesReferencedByRecipient() (relatedPeople []RelatedPerson, err error) {
	if d.IncludedRelatedPersonResourcesReferencedByRecipient == nil {
		err = errors.New("Included relatedPeople not requested")
	} else {
		relatedPeople = *d.IncludedRelatedPersonResourcesReferencedByRecipient
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if d.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *d.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if d.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *d.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if d.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *d.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if d.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *d.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if d.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *d.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if d.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *d.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if d.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *d.RevIncludedListResourcesReferencingItem
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if d.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *d.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if d.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *d.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if d.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *d.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if d.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *d.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if d.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *d.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if d.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *d.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if d.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *d.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if d.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *d.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (d *DocumentManifestPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedPractitionerResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedPractitionerResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedGroupResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedGroupResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedDeviceResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedDeviceResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for _, r := range *d.IncludedPractitionerResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for _, r := range *d.IncludedOrganizationResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedDeviceResourcesReferencedByAuthor != nil {
		for _, r := range *d.IncludedDeviceResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResourcesReferencedByAuthor != nil {
		for _, r := range *d.IncludedPatientResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for _, r := range *d.IncludedRelatedPersonResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *d.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPractitionerResourcesReferencedByRecipient != nil {
		for _, r := range *d.IncludedPractitionerResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedOrganizationResourcesReferencedByRecipient != nil {
		for _, r := range *d.IncludedOrganizationResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResourcesReferencedByRecipient != nil {
		for _, r := range *d.IncludedPatientResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedRelatedPersonResourcesReferencedByRecipient != nil {
		for _, r := range *d.IncludedRelatedPersonResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (d *DocumentManifestPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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

func (d *DocumentManifestPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if d.IncludedPractitionerResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedPractitionerResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedGroupResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedGroupResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedDeviceResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedDeviceResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *d.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPractitionerResourcesReferencedByAuthor != nil {
		for _, r := range *d.IncludedPractitionerResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedOrganizationResourcesReferencedByAuthor != nil {
		for _, r := range *d.IncludedOrganizationResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedDeviceResourcesReferencedByAuthor != nil {
		for _, r := range *d.IncludedDeviceResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResourcesReferencedByAuthor != nil {
		for _, r := range *d.IncludedPatientResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedRelatedPersonResourcesReferencedByAuthor != nil {
		for _, r := range *d.IncludedRelatedPersonResourcesReferencedByAuthor {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *d.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPractitionerResourcesReferencedByRecipient != nil {
		for _, r := range *d.IncludedPractitionerResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedOrganizationResourcesReferencedByRecipient != nil {
		for _, r := range *d.IncludedOrganizationResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedPatientResourcesReferencedByRecipient != nil {
		for _, r := range *d.IncludedPatientResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
	if d.IncludedRelatedPersonResourcesReferencedByRecipient != nil {
		for _, r := range *d.IncludedRelatedPersonResourcesReferencedByRecipient {
			resourceMap[r.Id] = &r
		}
	}
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
