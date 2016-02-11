package server

import "github.com/labstack/echo"

func RegisterController(name string, e *echo.Echo, m []echo.Middleware) {
	rc := ResourceController{name}
	rcBase := e.Group("/" + name)
	rcBase.Get("", rc.IndexHandler)
	rcBase.Post("", rc.CreateHandler)

	rcItem := rcBase.Group("/:id")
	rcItem.Get("", rc.ShowHandler)
	rcItem.Put("", rc.UpdateHandler)
	rcItem.Delete("", rc.DeleteHandler)

	if len(m) > 0 {
		rcBase.Use(m...)
	}
}

func RegisterRoutes(e *echo.Echo, config map[string][]echo.Middleware) {

	// Batch Support

	e.Post("/", BatchHandler)

	// Resources

	RegisterController("Appointment", e, config["Appointment"])
	RegisterController("ReferralRequest", e, config["ReferralRequest"])
	RegisterController("Account", e, config["Account"])
	RegisterController("DocumentManifest", e, config["DocumentManifest"])
	RegisterController("Goal", e, config["Goal"])
	RegisterController("EnrollmentRequest", e, config["EnrollmentRequest"])
	RegisterController("Medication", e, config["Medication"])
	RegisterController("Measure", e, config["Measure"])
	RegisterController("Subscription", e, config["Subscription"])
	RegisterController("DocumentReference", e, config["DocumentReference"])
	RegisterController("Conformance", e, config["Conformance"])
	RegisterController("RelatedPerson", e, config["RelatedPerson"])
	RegisterController("SupplyRequest", e, config["SupplyRequest"])
	RegisterController("Practitioner", e, config["Practitioner"])
	RegisterController("ExpansionProfile", e, config["ExpansionProfile"])
	RegisterController("OrderSet", e, config["OrderSet"])
	RegisterController("Slot", e, config["Slot"])
	RegisterController("Person", e, config["Person"])
	RegisterController("Contract", e, config["Contract"])
	RegisterController("RiskAssessment", e, config["RiskAssessment"])
	RegisterController("Group", e, config["Group"])
	RegisterController("PaymentNotice", e, config["PaymentNotice"])
	RegisterController("Organization", e, config["Organization"])
	RegisterController("ImplementationGuide", e, config["ImplementationGuide"])
	RegisterController("ImagingStudy", e, config["ImagingStudy"])
	RegisterController("DeviceComponent", e, config["DeviceComponent"])
	RegisterController("FamilyMemberHistory", e, config["FamilyMemberHistory"])
	RegisterController("Encounter", e, config["Encounter"])
	RegisterController("Substance", e, config["Substance"])
	RegisterController("SearchParameter", e, config["SearchParameter"])
	RegisterController("Protocol", e, config["Protocol"])
	RegisterController("Communication", e, config["Communication"])
	RegisterController("Linkage", e, config["Linkage"])
	RegisterController("OrderResponse", e, config["OrderResponse"])
	RegisterController("DeviceUseStatement", e, config["DeviceUseStatement"])
	RegisterController("MessageHeader", e, config["MessageHeader"])
	RegisterController("ImmunizationRecommendation", e, config["ImmunizationRecommendation"])
	RegisterController("BodySite", e, config["BodySite"])
	RegisterController("Provenance", e, config["Provenance"])
	RegisterController("Questionnaire", e, config["Questionnaire"])
	RegisterController("ExplanationOfBenefit", e, config["ExplanationOfBenefit"])
	RegisterController("Specimen", e, config["Specimen"])
	RegisterController("AllergyIntolerance", e, config["AllergyIntolerance"])
	RegisterController("CarePlan", e, config["CarePlan"])
	RegisterController("StructureDefinition", e, config["StructureDefinition"])
	RegisterController("EpisodeOfCare", e, config["EpisodeOfCare"])
	RegisterController("OperationOutcome", e, config["OperationOutcome"])
	RegisterController("Procedure", e, config["Procedure"])
	RegisterController("List", e, config["List"])
	RegisterController("ConceptMap", e, config["ConceptMap"])
	RegisterController("ValueSet", e, config["ValueSet"])
	RegisterController("OperationDefinition", e, config["OperationDefinition"])
	RegisterController("Order", e, config["Order"])
	RegisterController("Immunization", e, config["Immunization"])
	RegisterController("Device", e, config["Device"])
	RegisterController("VisionPrescription", e, config["VisionPrescription"])
	RegisterController("Media", e, config["Media"])
	RegisterController("ProcedureRequest", e, config["ProcedureRequest"])
	RegisterController("EligibilityResponse", e, config["EligibilityResponse"])
	RegisterController("DeviceUseRequest", e, config["DeviceUseRequest"])
	RegisterController("Sequence", e, config["Sequence"])
	RegisterController("DeviceMetric", e, config["DeviceMetric"])
	RegisterController("Flag", e, config["Flag"])
	RegisterController("CodeSystem", e, config["CodeSystem"])
	RegisterController("AppointmentResponse", e, config["AppointmentResponse"])
	RegisterController("GuidanceResponse", e, config["GuidanceResponse"])
	RegisterController("Observation", e, config["Observation"])
	RegisterController("MedicationAdministration", e, config["MedicationAdministration"])
	RegisterController("EnrollmentResponse", e, config["EnrollmentResponse"])
	RegisterController("Binary", e, config["Binary"])
	RegisterController("Library", e, config["Library"])
	RegisterController("MedicationStatement", e, config["MedicationStatement"])
	RegisterController("CommunicationRequest", e, config["CommunicationRequest"])
	RegisterController("TestScript", e, config["TestScript"])
	RegisterController("Basic", e, config["Basic"])
	RegisterController("ClaimResponse", e, config["ClaimResponse"])
	RegisterController("EligibilityRequest", e, config["EligibilityRequest"])
	RegisterController("ProcessRequest", e, config["ProcessRequest"])
	RegisterController("MedicationDispense", e, config["MedicationDispense"])
	RegisterController("DiagnosticReport", e, config["DiagnosticReport"])
	RegisterController("ImagingObjectSelection", e, config["ImagingObjectSelection"])
	RegisterController("HealthcareService", e, config["HealthcareService"])
	RegisterController("DataElement", e, config["DataElement"])
	RegisterController("NutritionOrder", e, config["NutritionOrder"])
	RegisterController("AuditEvent", e, config["AuditEvent"])
	RegisterController("MedicationOrder", e, config["MedicationOrder"])
	RegisterController("DecisionSupportRule", e, config["DecisionSupportRule"])
	RegisterController("PaymentReconciliation", e, config["PaymentReconciliation"])
	RegisterController("Condition", e, config["Condition"])
	RegisterController("Composition", e, config["Composition"])
	RegisterController("DetectedIssue", e, config["DetectedIssue"])
	RegisterController("Bundle", e, config["Bundle"])
	RegisterController("DiagnosticOrder", e, config["DiagnosticOrder"])
	RegisterController("Patient", e, config["Patient"])
	RegisterController("Coverage", e, config["Coverage"])
	RegisterController("QuestionnaireResponse", e, config["QuestionnaireResponse"])
	RegisterController("ProcessResponse", e, config["ProcessResponse"])
	RegisterController("ModuleDefinition", e, config["ModuleDefinition"])
	RegisterController("NamingSystem", e, config["NamingSystem"])
	RegisterController("DecisionSupportServiceModule", e, config["DecisionSupportServiceModule"])
	RegisterController("Schedule", e, config["Schedule"])
	RegisterController("SupplyDelivery", e, config["SupplyDelivery"])
	RegisterController("ClinicalImpression", e, config["ClinicalImpression"])
	RegisterController("Claim", e, config["Claim"])
	RegisterController("Location", e, config["Location"])
}
