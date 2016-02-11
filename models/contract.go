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

type Contract struct {
	DomainResource    `bson:",inline"`
	Identifier        *Identifier                           `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Issued            *FHIRDateTime                         `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies           *Period                               `bson:"applies,omitempty" json:"applies,omitempty"`
	Subject           []Reference                           `bson:"subject,omitempty" json:"subject,omitempty"`
	Authority         []Reference                           `bson:"authority,omitempty" json:"authority,omitempty"`
	Domain            []Reference                           `bson:"domain,omitempty" json:"domain,omitempty"`
	Type              *CodeableConcept                      `bson:"type,omitempty" json:"type,omitempty"`
	SubType           []CodeableConcept                     `bson:"subType,omitempty" json:"subType,omitempty"`
	Action            []CodeableConcept                     `bson:"action,omitempty" json:"action,omitempty"`
	ActionReason      []CodeableConcept                     `bson:"actionReason,omitempty" json:"actionReason,omitempty"`
	Party             []ContractPartyComponent              `bson:"party,omitempty" json:"party,omitempty"`
	ValuedItem        []ContractValuedItemComponent         `bson:"valuedItem,omitempty" json:"valuedItem,omitempty"`
	Signer            []ContractSignatoryComponent          `bson:"signer,omitempty" json:"signer,omitempty"`
	Term              []ContractTermComponent               `bson:"term,omitempty" json:"term,omitempty"`
	BindingAttachment *Attachment                           `bson:"bindingAttachment,omitempty" json:"bindingAttachment,omitempty"`
	BindingReference  *Reference                            `bson:"bindingReference,omitempty" json:"bindingReference,omitempty"`
	Friendly          []ContractFriendlyLanguageComponent   `bson:"friendly,omitempty" json:"friendly,omitempty"`
	Legal             []ContractLegalLanguageComponent      `bson:"legal,omitempty" json:"legal,omitempty"`
	Rule              []ContractComputableLanguageComponent `bson:"rule,omitempty" json:"rule,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *Contract) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "Contract"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to Contract), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *Contract) GetBSON() (interface{}, error) {
	x.ResourceType = "Contract"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "contract" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type contract Contract

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *Contract) UnmarshalJSON(data []byte) (err error) {
	x2 := contract{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = Contract(x2)
		return x.checkResourceType()
	}
	return
}

func (x *Contract) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "Contract"
	} else if x.ResourceType != "Contract" {
		return errors.New(fmt.Sprintf("Expected resourceType to be Contract, instead received %s", x.ResourceType))
	}
	return nil
}

type ContractPartyComponent struct {
	Entity *Reference        `bson:"entity,omitempty" json:"entity,omitempty"`
	Role   []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type ContractValuedItemComponent struct {
	EntityCodeableConcept *CodeableConcept `bson:"entityCodeableConcept,omitempty" json:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `bson:"entityReference,omitempty" json:"entityReference,omitempty"`
	Identifier            *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	EffectiveTime         *FHIRDateTime    `bson:"effectiveTime,omitempty" json:"effectiveTime,omitempty"`
	Quantity              *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice             *Quantity        `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                *float64         `bson:"factor,omitempty" json:"factor,omitempty"`
	Points                *float64         `bson:"points,omitempty" json:"points,omitempty"`
	Net                   *Quantity        `bson:"net,omitempty" json:"net,omitempty"`
}

type ContractSignatoryComponent struct {
	Type      *Coding    `bson:"type,omitempty" json:"type,omitempty"`
	Party     *Reference `bson:"party,omitempty" json:"party,omitempty"`
	Signature string     `bson:"signature,omitempty" json:"signature,omitempty"`
}

type ContractTermComponent struct {
	Identifier   *Identifier                       `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Issued       *FHIRDateTime                     `bson:"issued,omitempty" json:"issued,omitempty"`
	Applies      *Period                           `bson:"applies,omitempty" json:"applies,omitempty"`
	Type         *CodeableConcept                  `bson:"type,omitempty" json:"type,omitempty"`
	SubType      *CodeableConcept                  `bson:"subType,omitempty" json:"subType,omitempty"`
	Topic        *Reference                        `bson:"topic,omitempty" json:"topic,omitempty"`
	Subject      *Reference                        `bson:"subject,omitempty" json:"subject,omitempty"`
	Action       []CodeableConcept                 `bson:"action,omitempty" json:"action,omitempty"`
	ActionReason []CodeableConcept                 `bson:"actionReason,omitempty" json:"actionReason,omitempty"`
	Party        []ContractTermPartyComponent      `bson:"party,omitempty" json:"party,omitempty"`
	Text         string                            `bson:"text,omitempty" json:"text,omitempty"`
	ValuedItem   []ContractTermValuedItemComponent `bson:"valuedItem,omitempty" json:"valuedItem,omitempty"`
	Group        []ContractTermComponent           `bson:"group,omitempty" json:"group,omitempty"`
}

type ContractTermPartyComponent struct {
	Entity *Reference        `bson:"entity,omitempty" json:"entity,omitempty"`
	Role   []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type ContractTermValuedItemComponent struct {
	EntityCodeableConcept *CodeableConcept `bson:"entityCodeableConcept,omitempty" json:"entityCodeableConcept,omitempty"`
	EntityReference       *Reference       `bson:"entityReference,omitempty" json:"entityReference,omitempty"`
	Identifier            *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	EffectiveTime         *FHIRDateTime    `bson:"effectiveTime,omitempty" json:"effectiveTime,omitempty"`
	Quantity              *Quantity        `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice             *Quantity        `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor                *float64         `bson:"factor,omitempty" json:"factor,omitempty"`
	Points                *float64         `bson:"points,omitempty" json:"points,omitempty"`
	Net                   *Quantity        `bson:"net,omitempty" json:"net,omitempty"`
}

type ContractFriendlyLanguageComponent struct {
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type ContractLegalLanguageComponent struct {
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type ContractComputableLanguageComponent struct {
	ContentAttachment *Attachment `bson:"contentAttachment,omitempty" json:"contentAttachment,omitempty"`
	ContentReference  *Reference  `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
}

type ContractPlus struct {
	Contract                     `bson:",inline"`
	ContractPlusRelatedResources `bson:",inline"`
}

type ContractPlusRelatedResources struct {
	IncludedPatientResourcesReferencedBySubject                 *[]Patient               `bson:"_includedPatientResourcesReferencedBySubject,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedPractitionerResourcesReferencedByParty              *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByParty,omitempty"`
	IncludedGroupResourcesReferencedByParty                     *[]Group                 `bson:"_includedGroupResourcesReferencedByParty,omitempty"`
	IncludedOrganizationResourcesReferencedByParty              *[]Organization          `bson:"_includedOrganizationResourcesReferencedByParty,omitempty"`
	IncludedDeviceResourcesReferencedByParty                    *[]Device                `bson:"_includedDeviceResourcesReferencedByParty,omitempty"`
	IncludedPatientResourcesReferencedByParty                   *[]Patient               `bson:"_includedPatientResourcesReferencedByParty,omitempty"`
	IncludedSubstanceResourcesReferencedByParty                 *[]Substance             `bson:"_includedSubstanceResourcesReferencedByParty,omitempty"`
	IncludedContractResourcesReferencedByParty                  *[]Contract              `bson:"_includedContractResourcesReferencedByParty,omitempty"`
	IncludedRelatedPersonResourcesReferencedByParty             *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedByParty,omitempty"`
	IncludedLocationResourcesReferencedByParty                  *[]Location              `bson:"_includedLocationResourcesReferencedByParty,omitempty"`
	IncludedPractitionerResourcesReferencedBySigner             *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedBySigner,omitempty"`
	IncludedOrganizationResourcesReferencedBySigner             *[]Organization          `bson:"_includedOrganizationResourcesReferencedBySigner,omitempty"`
	IncludedPatientResourcesReferencedBySigner                  *[]Patient               `bson:"_includedPatientResourcesReferencedBySigner,omitempty"`
	IncludedRelatedPersonResourcesReferencedBySigner            *[]RelatedPerson         `bson:"_includedRelatedPersonResourcesReferencedBySigner,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingContentref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingContentref,omitempty"`
	RevIncludedDocumentManifestResourcesReferencingRelatedref   *[]DocumentManifest      `bson:"_revIncludedDocumentManifestResourcesReferencingRelatedref,omitempty"`
	RevIncludedDocumentReferenceResourcesReferencingRelatedref  *[]DocumentReference     `bson:"_revIncludedDocumentReferenceResourcesReferencingRelatedref,omitempty"`
	RevIncludedContractResourcesReferencingParty                *[]Contract              `bson:"_revIncludedContractResourcesReferencingParty,omitempty"`
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

func (c *ContractPlusRelatedResources) GetIncludedPatientResourcesReferencedBySubject() (patients []Patient, err error) {
	if c.IncludedPatientResourcesReferencedBySubject == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedPatientResourcesReferencedBySubject
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedPatientResourcesReferencedByPatient() (patients []Patient, err error) {
	if c.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else {
		patients = *c.IncludedPatientResourcesReferencedByPatient
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedPractitionerResourceReferencedByParty() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedByParty == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedByParty) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedByParty))
	} else if len(*c.IncludedPractitionerResourcesReferencedByParty) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedByParty)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedGroupResourceReferencedByParty() (group *Group, err error) {
	if c.IncludedGroupResourcesReferencedByParty == nil {
		err = errors.New("Included groups not requested")
	} else if len(*c.IncludedGroupResourcesReferencedByParty) > 1 {
		err = fmt.Errorf("Expected 0 or 1 group, but found %d", len(*c.IncludedGroupResourcesReferencedByParty))
	} else if len(*c.IncludedGroupResourcesReferencedByParty) == 1 {
		group = &(*c.IncludedGroupResourcesReferencedByParty)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedOrganizationResourceReferencedByParty() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedByParty == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedByParty) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedByParty))
	} else if len(*c.IncludedOrganizationResourcesReferencedByParty) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedByParty)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedDeviceResourceReferencedByParty() (device *Device, err error) {
	if c.IncludedDeviceResourcesReferencedByParty == nil {
		err = errors.New("Included devices not requested")
	} else if len(*c.IncludedDeviceResourcesReferencedByParty) > 1 {
		err = fmt.Errorf("Expected 0 or 1 device, but found %d", len(*c.IncludedDeviceResourcesReferencedByParty))
	} else if len(*c.IncludedDeviceResourcesReferencedByParty) == 1 {
		device = &(*c.IncludedDeviceResourcesReferencedByParty)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedPatientResourceReferencedByParty() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedByParty == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedByParty) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedByParty))
	} else if len(*c.IncludedPatientResourcesReferencedByParty) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedByParty)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedSubstanceResourceReferencedByParty() (substance *Substance, err error) {
	if c.IncludedSubstanceResourcesReferencedByParty == nil {
		err = errors.New("Included substances not requested")
	} else if len(*c.IncludedSubstanceResourcesReferencedByParty) > 1 {
		err = fmt.Errorf("Expected 0 or 1 substance, but found %d", len(*c.IncludedSubstanceResourcesReferencedByParty))
	} else if len(*c.IncludedSubstanceResourcesReferencedByParty) == 1 {
		substance = &(*c.IncludedSubstanceResourcesReferencedByParty)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedContractResourceReferencedByParty() (contract *Contract, err error) {
	if c.IncludedContractResourcesReferencedByParty == nil {
		err = errors.New("Included contracts not requested")
	} else if len(*c.IncludedContractResourcesReferencedByParty) > 1 {
		err = fmt.Errorf("Expected 0 or 1 contract, but found %d", len(*c.IncludedContractResourcesReferencedByParty))
	} else if len(*c.IncludedContractResourcesReferencedByParty) == 1 {
		contract = &(*c.IncludedContractResourcesReferencedByParty)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedByParty() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedByParty == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByParty) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRelatedPersonResourcesReferencedByParty))
	} else if len(*c.IncludedRelatedPersonResourcesReferencedByParty) == 1 {
		relatedPerson = &(*c.IncludedRelatedPersonResourcesReferencedByParty)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedLocationResourceReferencedByParty() (location *Location, err error) {
	if c.IncludedLocationResourcesReferencedByParty == nil {
		err = errors.New("Included locations not requested")
	} else if len(*c.IncludedLocationResourcesReferencedByParty) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*c.IncludedLocationResourcesReferencedByParty))
	} else if len(*c.IncludedLocationResourcesReferencedByParty) == 1 {
		location = &(*c.IncludedLocationResourcesReferencedByParty)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedPractitionerResourceReferencedBySigner() (practitioner *Practitioner, err error) {
	if c.IncludedPractitionerResourcesReferencedBySigner == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*c.IncludedPractitionerResourcesReferencedBySigner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*c.IncludedPractitionerResourcesReferencedBySigner))
	} else if len(*c.IncludedPractitionerResourcesReferencedBySigner) == 1 {
		practitioner = &(*c.IncludedPractitionerResourcesReferencedBySigner)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedOrganizationResourceReferencedBySigner() (organization *Organization, err error) {
	if c.IncludedOrganizationResourcesReferencedBySigner == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*c.IncludedOrganizationResourcesReferencedBySigner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*c.IncludedOrganizationResourcesReferencedBySigner))
	} else if len(*c.IncludedOrganizationResourcesReferencedBySigner) == 1 {
		organization = &(*c.IncludedOrganizationResourcesReferencedBySigner)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedPatientResourceReferencedBySigner() (patient *Patient, err error) {
	if c.IncludedPatientResourcesReferencedBySigner == nil {
		err = errors.New("Included patients not requested")
	} else if len(*c.IncludedPatientResourcesReferencedBySigner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*c.IncludedPatientResourcesReferencedBySigner))
	} else if len(*c.IncludedPatientResourcesReferencedBySigner) == 1 {
		patient = &(*c.IncludedPatientResourcesReferencedBySigner)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedRelatedPersonResourceReferencedBySigner() (relatedPerson *RelatedPerson, err error) {
	if c.IncludedRelatedPersonResourcesReferencedBySigner == nil {
		err = errors.New("Included relatedpeople not requested")
	} else if len(*c.IncludedRelatedPersonResourcesReferencedBySigner) > 1 {
		err = fmt.Errorf("Expected 0 or 1 relatedPerson, but found %d", len(*c.IncludedRelatedPersonResourcesReferencedBySigner))
	} else if len(*c.IncludedRelatedPersonResourcesReferencedBySigner) == 1 {
		relatedPerson = &(*c.IncludedRelatedPersonResourcesReferencedBySigner)[0]
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if c.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *c.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if c.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *c.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedContractResourcesReferencingParty() (contracts []Contract, err error) {
	if c.RevIncludedContractResourcesReferencingParty == nil {
		err = errors.New("RevIncluded contracts not requested")
	} else {
		contracts = *c.RevIncludedContractResourcesReferencingParty
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if c.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *c.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if c.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *c.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if c.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *c.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if c.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *c.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if c.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *c.RevIncludedListResourcesReferencingItem
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if c.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *c.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if c.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *c.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if c.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *c.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if c.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *c.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if c.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *c.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if c.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *c.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if c.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *c.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (c *ContractPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if c.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *c.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (c *ContractPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPractitionerResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedPractitionerResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedGroupResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedGroupResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedOrganizationResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedOrganizationResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedDeviceResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedDeviceResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSubstanceResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedSubstanceResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedContractResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedContractResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedRelatedPersonResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedLocationResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedLocationResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPractitionerResourcesReferencedBySigner != nil {
		for _, r := range *c.IncludedPractitionerResourcesReferencedBySigner {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedOrganizationResourcesReferencedBySigner != nil {
		for _, r := range *c.IncludedOrganizationResourcesReferencedBySigner {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResourcesReferencedBySigner != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedBySigner {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedBySigner != nil {
		for _, r := range *c.IncludedRelatedPersonResourcesReferencedBySigner {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (c *ContractPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
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
	if c.RevIncludedContractResourcesReferencingParty != nil {
		for _, r := range *c.RevIncludedContractResourcesReferencingParty {
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

func (c *ContractPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if c.IncludedPatientResourcesReferencedBySubject != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedBySubject {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPractitionerResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedPractitionerResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedGroupResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedGroupResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedOrganizationResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedOrganizationResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedDeviceResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedDeviceResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedSubstanceResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedSubstanceResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedContractResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedContractResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedRelatedPersonResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedLocationResourcesReferencedByParty != nil {
		for _, r := range *c.IncludedLocationResourcesReferencedByParty {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPractitionerResourcesReferencedBySigner != nil {
		for _, r := range *c.IncludedPractitionerResourcesReferencedBySigner {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedOrganizationResourcesReferencedBySigner != nil {
		for _, r := range *c.IncludedOrganizationResourcesReferencedBySigner {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedPatientResourcesReferencedBySigner != nil {
		for _, r := range *c.IncludedPatientResourcesReferencedBySigner {
			resourceMap[r.Id] = &r
		}
	}
	if c.IncludedRelatedPersonResourcesReferencedBySigner != nil {
		for _, r := range *c.IncludedRelatedPersonResourcesReferencedBySigner {
			resourceMap[r.Id] = &r
		}
	}
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
	if c.RevIncludedContractResourcesReferencingParty != nil {
		for _, r := range *c.RevIncludedContractResourcesReferencingParty {
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
