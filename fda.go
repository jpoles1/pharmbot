package main

type DrugData struct {
	Meta struct {
		Results struct {
			Skip  int `json:"skip"`
			Limit int `json:"limit"`
			Total int `json:"total"`
		} `json:"results"`
	} `json:"meta"`
	Results []struct {
		EffectiveTime                string   `json:"effective_time"`
		DrugInteractions             []string `json:"drug_interactions"`
		RecentMajorChanges           []string `json:"recent_major_changes"`
		GeriatricUse                 []string `json:"geriatric_use"`
		Description                  []string `json:"description"`
		NonclinicalToxicology        []string `json:"nonclinical_toxicology"`
		DosageFormsAndStrengths      []string `json:"dosage_forms_and_strengths"`
		StorageAndHandling           []string `json:"storage_and_handling"`
		MechanismOfAction            []string `json:"mechanism_of_action"`
		Pharmacokinetics             []string `json:"pharmacokinetics"`
		ClinicalStudiesTable         []string `json:"clinical_studies_table"`
		IndicationsAndUsage          []string `json:"indications_and_usage"`
		SetID                        string   `json:"set_id"`
		ID                           string   `json:"id"`
		DosageAndAdministrationTable []string `json:"dosage_and_administration_table"`
		PediatricUse                 []string `json:"pediatric_use"`
		Contraindications            []string `json:"contraindications"`
		Pregnancy                    []string `json:"pregnancy"`
		SplProductDataElements       []string `json:"spl_product_data_elements"`
		BoxedWarning                 []string `json:"boxed_warning"`
		AdverseReactionsTable        []string `json:"adverse_reactions_table"`
		WarningsAndCautions          []string `json:"warnings_and_cautions"`
		Openfda                      struct {
			ProductNdc         []string `json:"product_ndc"`
			Nui                []string `json:"nui"`
			IsOriginalPackager []bool   `json:"is_original_packager"`
			PackageNdc         []string `json:"package_ndc"`
			GenericName        []string `json:"generic_name"`
			SplSetID           []string `json:"spl_set_id"`
			BrandName          []string `json:"brand_name"`
			ManufacturerName   []string `json:"manufacturer_name"`
			Unii               []string `json:"unii"`
			Rxcui              []string `json:"rxcui"`
			SplID              []string `json:"spl_id"`
			SubstanceName      []string `json:"substance_name"`
			ProductType        []string `json:"product_type"`
			Route              []string `json:"route"`
			PharmClassMoa      []string `json:"pharm_class_moa"`
			ApplicationNumber  []string `json:"application_number"`
			PharmClassEpc      []string `json:"pharm_class_epc"`
		} `json:"openfda"`
		Version                                              string   `json:"version"`
		RecentMajorChangesTable                              []string `json:"recent_major_changes_table"`
		DosageAndAdministration                              []string `json:"dosage_and_administration"`
		AdverseReactions                                     []string `json:"adverse_reactions"`
		AnimalPharmacologyAndOrToxicology                    []string `json:"animal_pharmacology_and_or_toxicology"`
		SplUnclassifiedSection                               []string `json:"spl_unclassified_section"`
		UseInSpecificPopulations                             []string `json:"use_in_specific_populations"`
		HowSupplied                                          []string `json:"how_supplied"`
		InformationForPatients                               []string `json:"information_for_patients"`
		PackageLabelPrincipalDisplayPanel                    []string `json:"package_label_principal_display_panel"`
		ClinicalStudies                                      []string `json:"clinical_studies"`
		ClinicalPharmacology                                 []string `json:"clinical_pharmacology"`
		CarcinogenesisAndMutagenesisAndImpairmentOfFertility []string `json:"carcinogenesis_and_mutagenesis_and_impairment_of_fertility"`
		Overdosage                                           []string `json:"overdosage"`
	} `json:"results"`
}
