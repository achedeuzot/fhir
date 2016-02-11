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

type ExplanationOfBenefit struct {
	DomainResource            `bson:",inline"`
	Identifier                []Identifier                                  `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Claim                     *Reference                                    `bson:"claim,omitempty" json:"claim,omitempty"`
	ClaimResponse             *Reference                                    `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
	Ruleset                   *Coding                                       `bson:"ruleset,omitempty" json:"ruleset,omitempty"`
	OriginalRuleset           *Coding                                       `bson:"originalRuleset,omitempty" json:"originalRuleset,omitempty"`
	Created                   *FHIRDateTime                                 `bson:"created,omitempty" json:"created,omitempty"`
	BillablePeriod            *Period                                       `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty"`
	Disposition               string                                        `bson:"disposition,omitempty" json:"disposition,omitempty"`
	Provider                  *Reference                                    `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization              *Reference                                    `bson:"organization,omitempty" json:"organization,omitempty"`
	Facility                  *Reference                                    `bson:"facility,omitempty" json:"facility,omitempty"`
	RelatedClaim              []Reference                                   `bson:"relatedClaim,omitempty" json:"relatedClaim,omitempty"`
	Prescription              *Reference                                    `bson:"prescription,omitempty" json:"prescription,omitempty"`
	OriginalPrescription      *Reference                                    `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	Payee                     *ExplanationOfBenefitPayeeComponent           `bson:"payee,omitempty" json:"payee,omitempty"`
	Referral                  *Reference                                    `bson:"referral,omitempty" json:"referral,omitempty"`
	Diagnosis                 []ExplanationOfBenefitDiagnosisComponent      `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	SpecialCondition          []Coding                                      `bson:"specialCondition,omitempty" json:"specialCondition,omitempty"`
	Patient                   *Reference                                    `bson:"patient,omitempty" json:"patient,omitempty"`
	Precedence                *uint32                                       `bson:"precedence,omitempty" json:"precedence,omitempty"`
	Coverage                  *ExplanationOfBenefitCoverageComponent        `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Exception                 []Coding                                      `bson:"exception,omitempty" json:"exception,omitempty"`
	School                    string                                        `bson:"school,omitempty" json:"school,omitempty"`
	AccidentDate              *FHIRDateTime                                 `bson:"accidentDate,omitempty" json:"accidentDate,omitempty"`
	AccidentType              *Coding                                       `bson:"accidentType,omitempty" json:"accidentType,omitempty"`
	AccidentLocationString    string                                        `bson:"accidentLocationString,omitempty" json:"accidentLocationString,omitempty"`
	AccidentLocationAddress   *Address                                      `bson:"accidentLocationAddress,omitempty" json:"accidentLocationAddress,omitempty"`
	AccidentLocationReference *Reference                                    `bson:"accidentLocationReference,omitempty" json:"accidentLocationReference,omitempty"`
	InterventionException     []Coding                                      `bson:"interventionException,omitempty" json:"interventionException,omitempty"`
	OnsetDate                 *FHIRDateTime                                 `bson:"onsetDate,omitempty" json:"onsetDate,omitempty"`
	OnsetPeriod               *Period                                       `bson:"onsetPeriod,omitempty" json:"onsetPeriod,omitempty"`
	EmploymentImpacted        *Period                                       `bson:"employmentImpacted,omitempty" json:"employmentImpacted,omitempty"`
	Hospitalization           *Period                                       `bson:"hospitalization,omitempty" json:"hospitalization,omitempty"`
	Item                      []ExplanationOfBenefitItemsComponent          `bson:"item,omitempty" json:"item,omitempty"`
	AddItem                   []ExplanationOfBenefitAddedItemComponent      `bson:"addItem,omitempty" json:"addItem,omitempty"`
	ClaimTotal                *Quantity                                     `bson:"claimTotal,omitempty" json:"claimTotal,omitempty"`
	MissingTeeth              []ExplanationOfBenefitMissingTeethComponent   `bson:"missingTeeth,omitempty" json:"missingTeeth,omitempty"`
	UnallocDeductable         *Quantity                                     `bson:"unallocDeductable,omitempty" json:"unallocDeductable,omitempty"`
	TotalBenefit              *Quantity                                     `bson:"totalBenefit,omitempty" json:"totalBenefit,omitempty"`
	PaymentAdjustment         *Quantity                                     `bson:"paymentAdjustment,omitempty" json:"paymentAdjustment,omitempty"`
	PaymentAdjustmentReason   *Coding                                       `bson:"paymentAdjustmentReason,omitempty" json:"paymentAdjustmentReason,omitempty"`
	PaymentDate               *FHIRDateTime                                 `bson:"paymentDate,omitempty" json:"paymentDate,omitempty"`
	PaymentAmount             *Quantity                                     `bson:"paymentAmount,omitempty" json:"paymentAmount,omitempty"`
	PaymentRef                *Identifier                                   `bson:"paymentRef,omitempty" json:"paymentRef,omitempty"`
	Reserved                  *Coding                                       `bson:"reserved,omitempty" json:"reserved,omitempty"`
	Form                      *Coding                                       `bson:"form,omitempty" json:"form,omitempty"`
	Note                      []ExplanationOfBenefitNotesComponent          `bson:"note,omitempty" json:"note,omitempty"`
	BenefitBalance            []ExplanationOfBenefitBenefitBalanceComponent `bson:"benefitBalance,omitempty" json:"benefitBalance,omitempty"`
}

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *ExplanationOfBenefit) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "ExplanationOfBenefit"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to ExplanationOfBenefit), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *ExplanationOfBenefit) GetBSON() (interface{}, error) {
	x.ResourceType = "ExplanationOfBenefit"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "explanationOfBenefit" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type explanationOfBenefit ExplanationOfBenefit

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
func (x *ExplanationOfBenefit) UnmarshalJSON(data []byte) (err error) {
	x2 := explanationOfBenefit{}
	if err = json.Unmarshal(data, &x2); err == nil {
		if x2.Contained != nil {
			for i := range x2.Contained {
				x2.Contained[i] = MapToResource(x2.Contained[i], true)
			}
		}
		*x = ExplanationOfBenefit(x2)
		return x.checkResourceType()
	}
	return
}

func (x *ExplanationOfBenefit) checkResourceType() error {
	if x.ResourceType == "" {
		x.ResourceType = "ExplanationOfBenefit"
	} else if x.ResourceType != "ExplanationOfBenefit" {
		return errors.New(fmt.Sprintf("Expected resourceType to be ExplanationOfBenefit, instead received %s", x.ResourceType))
	}
	return nil
}

type ExplanationOfBenefitPayeeComponent struct {
	Type         *Coding    `bson:"type,omitempty" json:"type,omitempty"`
	Provider     *Reference `bson:"provider,omitempty" json:"provider,omitempty"`
	Organization *Reference `bson:"organization,omitempty" json:"organization,omitempty"`
	Person       *Reference `bson:"person,omitempty" json:"person,omitempty"`
}

type ExplanationOfBenefitDiagnosisComponent struct {
	Sequence  *uint32 `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Diagnosis *Coding `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
}

type ExplanationOfBenefitCoverageComponent struct {
	Coverage     *Reference `bson:"coverage,omitempty" json:"coverage,omitempty"`
	Relationship *Coding    `bson:"relationship,omitempty" json:"relationship,omitempty"`
	PreAuthRef   []string   `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
}

type ExplanationOfBenefitItemsComponent struct {
	Sequence        *uint32                                         `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type            *Coding                                         `bson:"type,omitempty" json:"type,omitempty"`
	Provider        *Reference                                      `bson:"provider,omitempty" json:"provider,omitempty"`
	DiagnosisLinkId []uint32                                        `bson:"diagnosisLinkId,omitempty" json:"diagnosisLinkId,omitempty"`
	Service         *Coding                                         `bson:"service,omitempty" json:"service,omitempty"`
	ServicedDate    *FHIRDateTime                                   `bson:"servicedDate,omitempty" json:"servicedDate,omitempty"`
	ServicedPeriod  *Period                                         `bson:"servicedPeriod,omitempty" json:"servicedPeriod,omitempty"`
	Place           *Coding                                         `bson:"place,omitempty" json:"place,omitempty"`
	Quantity        *Quantity                                       `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice       *Quantity                                       `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor          *float64                                        `bson:"factor,omitempty" json:"factor,omitempty"`
	Points          *float64                                        `bson:"points,omitempty" json:"points,omitempty"`
	Net             *Quantity                                       `bson:"net,omitempty" json:"net,omitempty"`
	Udi             *Coding                                         `bson:"udi,omitempty" json:"udi,omitempty"`
	BodySite        *Coding                                         `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	SubSite         []Coding                                        `bson:"subSite,omitempty" json:"subSite,omitempty"`
	Modifier        []Coding                                        `bson:"modifier,omitempty" json:"modifier,omitempty"`
	NoteNumber      []uint32                                        `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	Adjudication    []ExplanationOfBenefitItemAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail          []ExplanationOfBenefitDetailComponent           `bson:"detail,omitempty" json:"detail,omitempty"`
	Prosthesis      *ExplanationOfBenefitProsthesisComponent        `bson:"prosthesis,omitempty" json:"prosthesis,omitempty"`
}

type ExplanationOfBenefitItemAdjudicationComponent struct {
	Category *Coding   `bson:"category,omitempty" json:"category,omitempty"`
	Reason   *Coding   `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount   *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value    *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ExplanationOfBenefitDetailComponent struct {
	Sequence     *uint32                                           `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type         *Coding                                           `bson:"type,omitempty" json:"type,omitempty"`
	Service      *Coding                                           `bson:"service,omitempty" json:"service,omitempty"`
	Quantity     *Quantity                                         `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice    *Quantity                                         `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor       *float64                                          `bson:"factor,omitempty" json:"factor,omitempty"`
	Points       *float64                                          `bson:"points,omitempty" json:"points,omitempty"`
	Net          *Quantity                                         `bson:"net,omitempty" json:"net,omitempty"`
	Udi          *Coding                                           `bson:"udi,omitempty" json:"udi,omitempty"`
	Adjudication []ExplanationOfBenefitDetailAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	SubDetail    []ExplanationOfBenefitSubDetailComponent          `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ExplanationOfBenefitDetailAdjudicationComponent struct {
	Code   *Coding   `bson:"code,omitempty" json:"code,omitempty"`
	Reason *Coding   `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value  *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ExplanationOfBenefitSubDetailComponent struct {
	Sequence     *uint32                                              `bson:"sequence,omitempty" json:"sequence,omitempty"`
	Type         *Coding                                              `bson:"type,omitempty" json:"type,omitempty"`
	Service      *Coding                                              `bson:"service,omitempty" json:"service,omitempty"`
	Quantity     *Quantity                                            `bson:"quantity,omitempty" json:"quantity,omitempty"`
	UnitPrice    *Quantity                                            `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	Factor       *float64                                             `bson:"factor,omitempty" json:"factor,omitempty"`
	Points       *float64                                             `bson:"points,omitempty" json:"points,omitempty"`
	Net          *Quantity                                            `bson:"net,omitempty" json:"net,omitempty"`
	Udi          *Coding                                              `bson:"udi,omitempty" json:"udi,omitempty"`
	Adjudication []ExplanationOfBenefitSubDetailAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ExplanationOfBenefitSubDetailAdjudicationComponent struct {
	Code   *Coding   `bson:"code,omitempty" json:"code,omitempty"`
	Reason *Coding   `bson:"reason,omitempty" json:"reason,omitempty"`
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value  *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ExplanationOfBenefitProsthesisComponent struct {
	Initial       *bool         `bson:"initial,omitempty" json:"initial,omitempty"`
	PriorDate     *FHIRDateTime `bson:"priorDate,omitempty" json:"priorDate,omitempty"`
	PriorMaterial *Coding       `bson:"priorMaterial,omitempty" json:"priorMaterial,omitempty"`
}

type ExplanationOfBenefitAddedItemComponent struct {
	SequenceLinkId   []uint32                                             `bson:"sequenceLinkId,omitempty" json:"sequenceLinkId,omitempty"`
	Service          *Coding                                              `bson:"service,omitempty" json:"service,omitempty"`
	Fee              *Quantity                                            `bson:"fee,omitempty" json:"fee,omitempty"`
	NoteNumberLinkId []uint32                                             `bson:"noteNumberLinkId,omitempty" json:"noteNumberLinkId,omitempty"`
	Adjudication     []ExplanationOfBenefitAddedItemAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	Detail           []ExplanationOfBenefitAddedItemsDetailComponent      `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ExplanationOfBenefitAddedItemAdjudicationComponent struct {
	Code   *Coding   `bson:"code,omitempty" json:"code,omitempty"`
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value  *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ExplanationOfBenefitAddedItemsDetailComponent struct {
	Service      *Coding                                                    `bson:"service,omitempty" json:"service,omitempty"`
	Fee          *Quantity                                                  `bson:"fee,omitempty" json:"fee,omitempty"`
	Adjudication []ExplanationOfBenefitAddedItemDetailAdjudicationComponent `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ExplanationOfBenefitAddedItemDetailAdjudicationComponent struct {
	Code   *Coding   `bson:"code,omitempty" json:"code,omitempty"`
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	Value  *float64  `bson:"value,omitempty" json:"value,omitempty"`
}

type ExplanationOfBenefitMissingTeethComponent struct {
	Tooth          *Coding       `bson:"tooth,omitempty" json:"tooth,omitempty"`
	Reason         *Coding       `bson:"reason,omitempty" json:"reason,omitempty"`
	ExtractionDate *FHIRDateTime `bson:"extractionDate,omitempty" json:"extractionDate,omitempty"`
}

type ExplanationOfBenefitNotesComponent struct {
	Number *uint32 `bson:"number,omitempty" json:"number,omitempty"`
	Type   *Coding `bson:"type,omitempty" json:"type,omitempty"`
	Text   string  `bson:"text,omitempty" json:"text,omitempty"`
}

type ExplanationOfBenefitBenefitBalanceComponent struct {
	Category    *Coding                                `bson:"category,omitempty" json:"category,omitempty"`
	SubCategory *Coding                                `bson:"subCategory,omitempty" json:"subCategory,omitempty"`
	Network     *Coding                                `bson:"network,omitempty" json:"network,omitempty"`
	Unit        *Coding                                `bson:"unit,omitempty" json:"unit,omitempty"`
	Term        *Coding                                `bson:"term,omitempty" json:"term,omitempty"`
	Financial   []ExplanationOfBenefitBenefitComponent `bson:"financial,omitempty" json:"financial,omitempty"`
}

type ExplanationOfBenefitBenefitComponent struct {
	Type                   *Coding   `bson:"type,omitempty" json:"type,omitempty"`
	BenefitUnsignedInt     *uint32   `bson:"benefitUnsignedInt,omitempty" json:"benefitUnsignedInt,omitempty"`
	BenefitMoney           *Quantity `bson:"benefitMoney,omitempty" json:"benefitMoney,omitempty"`
	BenefitUsedUnsignedInt *uint32   `bson:"benefitUsedUnsignedInt,omitempty" json:"benefitUsedUnsignedInt,omitempty"`
	BenefitUsedMoney       *Quantity `bson:"benefitUsedMoney,omitempty" json:"benefitUsedMoney,omitempty"`
}

type ExplanationOfBenefitPlus struct {
	ExplanationOfBenefit                     `bson:",inline"`
	ExplanationOfBenefitPlusRelatedResources `bson:",inline"`
}

type ExplanationOfBenefitPlusRelatedResources struct {
	IncludedPractitionerResourcesReferencedByProvider           *[]Practitioner          `bson:"_includedPractitionerResourcesReferencedByProvider,omitempty"`
	IncludedPatientResourcesReferencedByPatient                 *[]Patient               `bson:"_includedPatientResourcesReferencedByPatient,omitempty"`
	IncludedOrganizationResourcesReferencedByOrganization       *[]Organization          `bson:"_includedOrganizationResourcesReferencedByOrganization,omitempty"`
	IncludedClaimResourcesReferencedByClaim                     *[]Claim                 `bson:"_includedClaimResourcesReferencedByClaim,omitempty"`
	IncludedLocationResourcesReferencedByFacility               *[]Location              `bson:"_includedLocationResourcesReferencedByFacility,omitempty"`
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

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPractitionerResourceReferencedByProvider() (practitioner *Practitioner, err error) {
	if e.IncludedPractitionerResourcesReferencedByProvider == nil {
		err = errors.New("Included practitioners not requested")
	} else if len(*e.IncludedPractitionerResourcesReferencedByProvider) > 1 {
		err = fmt.Errorf("Expected 0 or 1 practitioner, but found %d", len(*e.IncludedPractitionerResourcesReferencedByProvider))
	} else if len(*e.IncludedPractitionerResourcesReferencedByProvider) == 1 {
		practitioner = &(*e.IncludedPractitionerResourcesReferencedByProvider)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedPatientResourceReferencedByPatient() (patient *Patient, err error) {
	if e.IncludedPatientResourcesReferencedByPatient == nil {
		err = errors.New("Included patients not requested")
	} else if len(*e.IncludedPatientResourcesReferencedByPatient) > 1 {
		err = fmt.Errorf("Expected 0 or 1 patient, but found %d", len(*e.IncludedPatientResourcesReferencedByPatient))
	} else if len(*e.IncludedPatientResourcesReferencedByPatient) == 1 {
		patient = &(*e.IncludedPatientResourcesReferencedByPatient)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedOrganizationResourceReferencedByOrganization() (organization *Organization, err error) {
	if e.IncludedOrganizationResourcesReferencedByOrganization == nil {
		err = errors.New("Included organizations not requested")
	} else if len(*e.IncludedOrganizationResourcesReferencedByOrganization) > 1 {
		err = fmt.Errorf("Expected 0 or 1 organization, but found %d", len(*e.IncludedOrganizationResourcesReferencedByOrganization))
	} else if len(*e.IncludedOrganizationResourcesReferencedByOrganization) == 1 {
		organization = &(*e.IncludedOrganizationResourcesReferencedByOrganization)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedClaimResourceReferencedByClaim() (claim *Claim, err error) {
	if e.IncludedClaimResourcesReferencedByClaim == nil {
		err = errors.New("Included claims not requested")
	} else if len(*e.IncludedClaimResourcesReferencedByClaim) > 1 {
		err = fmt.Errorf("Expected 0 or 1 claim, but found %d", len(*e.IncludedClaimResourcesReferencedByClaim))
	} else if len(*e.IncludedClaimResourcesReferencedByClaim) == 1 {
		claim = &(*e.IncludedClaimResourcesReferencedByClaim)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedLocationResourceReferencedByFacility() (location *Location, err error) {
	if e.IncludedLocationResourcesReferencedByFacility == nil {
		err = errors.New("Included locations not requested")
	} else if len(*e.IncludedLocationResourcesReferencedByFacility) > 1 {
		err = fmt.Errorf("Expected 0 or 1 location, but found %d", len(*e.IncludedLocationResourcesReferencedByFacility))
	} else if len(*e.IncludedLocationResourcesReferencedByFacility) == 1 {
		location = &(*e.IncludedLocationResourcesReferencedByFacility)[0]
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingContentref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingContentref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingContentref
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDocumentManifestResourcesReferencingRelatedref() (documentManifests []DocumentManifest, err error) {
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentManifests not requested")
	} else {
		documentManifests = *e.RevIncludedDocumentManifestResourcesReferencingRelatedref
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDocumentReferenceResourcesReferencingRelatedref() (documentReferences []DocumentReference, err error) {
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref == nil {
		err = errors.New("RevIncluded documentReferences not requested")
	} else {
		documentReferences = *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingRequest() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingRequest
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedPaymentNoticeResourcesReferencingResponse() (paymentNotices []PaymentNotice, err error) {
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse == nil {
		err = errors.New("RevIncluded paymentNotices not requested")
	} else {
		paymentNotices = *e.RevIncludedPaymentNoticeResourcesReferencingResponse
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedOrderResponseResourcesReferencingFulfillment() (orderResponses []OrderResponse, err error) {
	if e.RevIncludedOrderResponseResourcesReferencingFulfillment == nil {
		err = errors.New("RevIncluded orderResponses not requested")
	} else {
		orderResponses = *e.RevIncludedOrderResponseResourcesReferencingFulfillment
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedMessageHeaderResourcesReferencingData() (messageHeaders []MessageHeader, err error) {
	if e.RevIncludedMessageHeaderResourcesReferencingData == nil {
		err = errors.New("RevIncluded messageHeaders not requested")
	} else {
		messageHeaders = *e.RevIncludedMessageHeaderResourcesReferencingData
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedProvenanceResourcesReferencingTarget() (provenances []Provenance, err error) {
	if e.RevIncludedProvenanceResourcesReferencingTarget == nil {
		err = errors.New("RevIncluded provenances not requested")
	} else {
		provenances = *e.RevIncludedProvenanceResourcesReferencingTarget
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedListResourcesReferencingItem() (lists []List, err error) {
	if e.RevIncludedListResourcesReferencingItem == nil {
		err = errors.New("RevIncluded lists not requested")
	} else {
		lists = *e.RevIncludedListResourcesReferencingItem
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedOrderResourcesReferencingDetail() (orders []Order, err error) {
	if e.RevIncludedOrderResourcesReferencingDetail == nil {
		err = errors.New("RevIncluded orders not requested")
	} else {
		orders = *e.RevIncludedOrderResourcesReferencingDetail
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedBasicResourcesReferencingSubject() (basics []Basic, err error) {
	if e.RevIncludedBasicResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded basics not requested")
	} else {
		basics = *e.RevIncludedBasicResourcesReferencingSubject
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedAuditEventResourcesReferencingReference() (auditEvents []AuditEvent, err error) {
	if e.RevIncludedAuditEventResourcesReferencingReference == nil {
		err = errors.New("RevIncluded auditEvents not requested")
	} else {
		auditEvents = *e.RevIncludedAuditEventResourcesReferencingReference
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingSubject() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingSubject
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedCompositionResourcesReferencingEntry() (compositions []Composition, err error) {
	if e.RevIncludedCompositionResourcesReferencingEntry == nil {
		err = errors.New("RevIncluded compositions not requested")
	} else {
		compositions = *e.RevIncludedCompositionResourcesReferencingEntry
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedDetectedIssueResourcesReferencingImplicated() (detectedIssues []DetectedIssue, err error) {
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated == nil {
		err = errors.New("RevIncluded detectedIssues not requested")
	} else {
		detectedIssues = *e.RevIncludedDetectedIssueResourcesReferencingImplicated
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedQuestionnaireResponseResourcesReferencingSubject() (questionnaireResponses []QuestionnaireResponse, err error) {
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject == nil {
		err = errors.New("RevIncluded questionnaireResponses not requested")
	} else {
		questionnaireResponses = *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedProcessResponseResourcesReferencingRequest() (processResponses []ProcessResponse, err error) {
	if e.RevIncludedProcessResponseResourcesReferencingRequest == nil {
		err = errors.New("RevIncluded processResponses not requested")
	} else {
		processResponses = *e.RevIncludedProcessResponseResourcesReferencingRequest
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedClinicalImpressionResourcesReferencingTrigger() (clinicalImpressions []ClinicalImpression, err error) {
	if e.RevIncludedClinicalImpressionResourcesReferencingTrigger == nil {
		err = errors.New("RevIncluded clinicalImpressions not requested")
	} else {
		clinicalImpressions = *e.RevIncludedClinicalImpressionResourcesReferencingTrigger
	}
	return
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedPractitionerResourcesReferencedByProvider != nil {
		for _, r := range *e.IncludedPractitionerResourcesReferencedByProvider {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *e.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for _, r := range *e.IncludedOrganizationResourcesReferencedByOrganization {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedClaimResourcesReferencedByClaim != nil {
		for _, r := range *e.IncludedClaimResourcesReferencedByClaim {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedLocationResourcesReferencedByFacility != nil {
		for _, r := range *e.IncludedLocationResourcesReferencedByFacility {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *e.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *e.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *e.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *e.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *e.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *e.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *e.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *e.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *e.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *e.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *e.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}

func (e *ExplanationOfBenefitPlusRelatedResources) GetIncludedAndRevIncludedResources() map[string]interface{} {
	resourceMap := make(map[string]interface{})
	if e.IncludedPractitionerResourcesReferencedByProvider != nil {
		for _, r := range *e.IncludedPractitionerResourcesReferencedByProvider {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedPatientResourcesReferencedByPatient != nil {
		for _, r := range *e.IncludedPatientResourcesReferencedByPatient {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedOrganizationResourcesReferencedByOrganization != nil {
		for _, r := range *e.IncludedOrganizationResourcesReferencedByOrganization {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedClaimResourcesReferencedByClaim != nil {
		for _, r := range *e.IncludedClaimResourcesReferencedByClaim {
			resourceMap[r.Id] = &r
		}
	}
	if e.IncludedLocationResourcesReferencedByFacility != nil {
		for _, r := range *e.IncludedLocationResourcesReferencedByFacility {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingContentref != nil {
		for _, r := range *e.RevIncludedDocumentManifestResourcesReferencingContentref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentManifestResourcesReferencingRelatedref != nil {
		for _, r := range *e.RevIncludedDocumentManifestResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDocumentReferenceResourcesReferencingRelatedref != nil {
		for _, r := range *e.RevIncludedDocumentReferenceResourcesReferencingRelatedref {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingRequest != nil {
		for _, r := range *e.RevIncludedPaymentNoticeResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedPaymentNoticeResourcesReferencingResponse != nil {
		for _, r := range *e.RevIncludedPaymentNoticeResourcesReferencingResponse {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedOrderResponseResourcesReferencingFulfillment != nil {
		for _, r := range *e.RevIncludedOrderResponseResourcesReferencingFulfillment {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedMessageHeaderResourcesReferencingData != nil {
		for _, r := range *e.RevIncludedMessageHeaderResourcesReferencingData {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProvenanceResourcesReferencingTarget != nil {
		for _, r := range *e.RevIncludedProvenanceResourcesReferencingTarget {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedListResourcesReferencingItem != nil {
		for _, r := range *e.RevIncludedListResourcesReferencingItem {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedOrderResourcesReferencingDetail != nil {
		for _, r := range *e.RevIncludedOrderResourcesReferencingDetail {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedBasicResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedBasicResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedAuditEventResourcesReferencingReference != nil {
		for _, r := range *e.RevIncludedAuditEventResourcesReferencingReference {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCompositionResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedCompositionResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedCompositionResourcesReferencingEntry != nil {
		for _, r := range *e.RevIncludedCompositionResourcesReferencingEntry {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedDetectedIssueResourcesReferencingImplicated != nil {
		for _, r := range *e.RevIncludedDetectedIssueResourcesReferencingImplicated {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedQuestionnaireResponseResourcesReferencingSubject != nil {
		for _, r := range *e.RevIncludedQuestionnaireResponseResourcesReferencingSubject {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedProcessResponseResourcesReferencingRequest != nil {
		for _, r := range *e.RevIncludedProcessResponseResourcesReferencingRequest {
			resourceMap[r.Id] = &r
		}
	}
	if e.RevIncludedClinicalImpressionResourcesReferencingTrigger != nil {
		for _, r := range *e.RevIncludedClinicalImpressionResourcesReferencingTrigger {
			resourceMap[r.Id] = &r
		}
	}
	return resourceMap
}
