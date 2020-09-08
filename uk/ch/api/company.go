package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/appinesshq/globire-go/uk/ch/api/enum"
	"github.com/pkg/errors"
)

// CompanyType represents the company type (legal form)
type CompanyType string

// String implements the Stringer interface to get a human readable string from the CH enums
func (f CompanyType) String() string {
	return enum.Constants.Get("company_type", string(f))
}

// AccountType represents a type of company accounts
type AccountType string

// String implements the Stringer interface to get a human readable string from the CH enums
func (f AccountType) String() string {
	return enum.Constants.Get("account_type", string(f))
}

// ForeignAccountType represents a type of company accounts for a foreign company
type ForeignAccountType string

// String implements the Stringer interface to get a human readable string from the CH enums
func (f ForeignAccountType) String() string {
	return enum.Constants.Get("foreign_account_type", string(f))
}

// CompanyStatus represents the status of a company in the register
type CompanyStatus string

// String implements the Stringer interface to get a human readable string from the CH enums
func (f CompanyStatus) String() string {
	return enum.Constants.Get("company_status", string(f))
}

// CompanyStatusDetail represents extra information on the company's status
type CompanyStatusDetail string

// String implements the Stringer interface to get a human readable string from the CH enums
func (f CompanyStatusDetail) String() string {
	return enum.Constants.Get("company_status_detail", string(f))
}

// TermsOfAccountPublication represents the terms related to a company's accounts
type TermsOfAccountPublication string

// String implements the Stringer interface to get a human readable string from the CH enums
func (f TermsOfAccountPublication) String() string {
	return enum.Constants.Get("terms_of_account_publication", string(f))
}

// Jurisdiction represents the jurisdiction of registration of a company
type Jurisdiction string

// String implements the Stringer interface to get a human readable string from the CH enums
func (f Jurisdiction) String() string {
	return enum.Constants.Get("jurisdiction", string(f))
}

// PartialDataAvailable represents partial (not yet processed) data
type PartialDataAvailable string

// String implements the Stringer interface to get a human readable string from the CH enums
func (f PartialDataAvailable) String() string {
	return enum.Constants.Get("partial_data_available", string(f))
}

// SICCode represents a Standard Industrial Classification code
type SICCode string

// String implements the Stringer interface to get a human readable string from the CH enums
func (f SICCode) String() string {
	desc := enum.Constants.Get("sic_descriptions", string(f))
	if desc == "" {
		return fmt.Sprintf("%s - Unknown", string(f))
	}
	return fmt.Sprintf("%s - %s", string(f), desc)
}

type (
	// PreviousName struct contains data of a company's previous names and the time of use
	PreviousName struct {
		Name          string `json:"name"`
		EffectiveFrom ChDate `json:"effective_from"`
		CeasedOn      ChDate `json:"ceased_on"`
	}

	// RefDate struct consists of Day and Month
	RefDate struct {
		Day   string `json:"day"`
		Month string `json:"month"`
	}

	// Accounts struct contains a company's last and next filing info for the Annual Accounts
	Accounts struct {
		AccountingReferenceDate RefDate `json:"accounting_reference_date"`
		LastAccounts            struct {
			MadeUpTo      ChDate      `json:"made_up_to"`
			Type          AccountType `json:"type"`
			PeriodEndOn   ChDate      `json:"period_end_on"`
			PeriodStartOn ChDate      `json:"period_start_on"`
		} `json:"last_accounts"`
		NextAccounts struct {
			DueOn         ChDate `json:"due_on"`
			Overdue       bool   `json:"overdue"`
			PeriodEndOn   ChDate `json:"period_end_on"`
			PeriodStartOn ChDate `json:"period_start_on"`
		} `json:"next_accounts"`
		NextDue      ChDate `json:"next_due"`        // Deprecated. Please use next_accounts.due_on.
		NextMadeUpTo ChDate `json:"next_made_up_to"` // Deprecated. Please use next_accounts.period_end_on.
		Overdue      bool   `json:"overdue"`         // Deprecated. Please use next_accounts.overdue
	}

	// AnnualReturn struct contains a company's the last and next filing dates for the Annual Return
	AnnualReturn struct {
		LastMadeUpTo ChDate `json:"last_made_up_to"`
		NextDue      ChDate `json:"next_due"`
		NextMadeUpTo ChDate `json:"next_made_up_to"`
		Overdue      bool   `json:"overdue"`
	}

	// Branch struct contains data of a Branch
	Branch struct {
		BusinessActivity    string `json:"business_activity"`
		ParentCompanyName   string `json:"parent_company_name"`
		ParentCompanyNumber string `json:"parent_company_number"`
	}

	// ForeignCompanyDetails struct contains data of Foreign Companies
	ForeignCompanyDetails struct {
		AccountingRequirement struct {
			ForeignAccountType        ForeignAccountType        `json:"foreign_account_type"`
			TermsOfAccountPublication TermsOfAccountPublication `json:"terms_of_account_publication"`
		} `json:"accounting_requirement"`
		Accounts struct {
			From RefDate `json:"account_period_from"`
			To   RefDate `json:"account_period_to"`
			Term struct {
				Months string `json:"months"`
			} `json:"must_file_within"`
		} `json:"accounts"`
		LegalForm                     string `json:"legal_form"`
		CompanyType                   string `json:"company_type"`
		BusinessActivity              string `json:"business_activity"`
		GovernedBy                    string `json:"governed_by"`
		IsACreditFinancialInstitution bool   `json:"is_a_credit_financial_institution"`
		RegistrationNumber            string `json:"registration_number"`
		OriginatingRegistry           struct {
			Country string `json:"country"`
			Name    string `json:"name"`
		} `json:"originating_registry"`
	}

	// Company struct contains basic company data
	Company struct {
		api                   *API
		Accounts              Accounts              `json:"accounts"`
		AnnualReturn          AnnualReturn          `json:"annual_return"`
		BranchCompanyDetails  Branch                `json:"branch_company_details"`
		CanFile               bool                  `json:"can_file"`
		Name                  string                `json:"company_name"`
		CompanyNumber         string                `json:"company_number"`
		CompanyStatus         CompanyStatus         `json:"company_status"`
		CompanyStatusDetail   CompanyStatusDetail   `json:"company_status_detail"`
		ConfirmationStatement AnnualReturn          `json:"confirmation_statement"`
		DateOfCessation       ChDate                `json:"date_of_cessation"`
		DateOfCreation        ChDate                `json:"date_of_creation"`
		Etag                  string                `json:"etag"`
		ForeignCompanyDetails ForeignCompanyDetails `json:"foreign_company_details"`
		HasBeenLiquidated     bool                  `json:"has_been_liquidated"`
		HasCharges            bool                  `json:"has_charges"`
		HasInsolvencyHistory  bool                  `json:"has_insolvency_history"`
		// IsCommunityInterestCompany bool                  `json:"is_community_interest_company"`
		Jurisdiction            Jurisdiction `json:"jurisdiction"`
		LastFullMembersListDate ChDate       `json:"last_full_members_list_date"`
		Links                   struct {
			Charges                                 string `json:"charges"`
			FilingHistory                           string `json:"filing_history"`
			Insolvency                              string `json:"insolvency"`
			Officers                                string `json:"officers"`
			PersonsWithSignificantControl           string `json:"persons_with_significant_control"`
			PersonsWithSignificantControlStatements string `json:"persons_with_significant_control_statements"`
			Registers                               string `json:"registers"`
			Self                                    string `json:"self"`
		} `json:"links"`
		PartialDataAvailable        PartialDataAvailable `json:"partial_data_available"`
		PreviousCompanyNames        []PreviousName       `json:"previous_company_names"`
		RegisteredOfficeAddress     Address              `json:"registered_office_address"`
		RegisteredOfficeIsInDispute bool                 `json:"registered_office_is_in_dispute"`
		SICCodes                    []SICCode            `json:"sic_codes"`
		// TODO: Deal with subtype
		SubType                              string      `json:"subtype"`
		Type                                 CompanyType `json:"type"`
		UndeliverableRegisteredOfficeAddress bool        `json:"undeliverable_registered_office_address"`
	}
)

// HasTasks returns a boolean value representing whether the company has outstanding tasks
func (c *Company) HasTasks() bool {
	return c.AnnualReturn.Overdue || c.ConfirmationStatement.Overdue || c.Accounts.Overdue
}

func (a *API) GetCompany(companyNumber string) (*Company, error) {
	c := Company{api: a}

	if err := a.Do(context.Background(), http.MethodGet, "/company/"+companyNumber, nil, nil, &c); err != nil {
		return nil, errors.Wrapf(err, "getting company")
	}

	return &c, nil
}
