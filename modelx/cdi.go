package modelx

const (
	CDIGetCustomerKeyServiceID = "GetCustomerKey"
	CDIGetCitizenIDServiceID   = "GetCitizenID"

	CDICreateCustomerActionHeader = "CREATE_CUSTOMER_PROFILE"
	CDIUpdateCustomerActionHeader = "UPDATE_CUSTOMER_PROFILE"

	//citizen-id, customer-key, passport, cif-cbs, cif-ems
	CDIXHeaderIdentifierTypeCitizenId   = "citizen-id"
	CDIXHeaderIdentifierTypeCustomerKey = "customer-key"
	CDIXHeaderIdentifierTypeCPassport   = "passport"
	CDIXHeaderIdentifierTypeCifCbs      = "cif-cbs"
	CDIXHeaderIdentifierTypeCifEms      = "cif-ems"
)

type CDICreateCustomerProfileRequest struct {
	CustomerBasicInfo     *CreateCustomerBasicInfo     `json:"customerBasicInfo"`
	CustomerSensitiveInfo *CreateCustomerSensitiveInfo `json:"customerSensitiveInfo"`
	CustomerWorkingInfo   *CreateCustomerWorkingInfo   `json:"customerWorkingInfo"`
	CustomerAddressInfo   *[]CreateCustomerAddressInfo `json:"customerAddressInfo"`
	CustomerContactInfo   *[]CreateCustomerContactInfo `json:"customerContactInfo"`
	CustomerFlagInfo      *CreateCustomerFlagInfo      `json:"customerFlagInfo"`
}

type CDIUpdateCustomerProfileRequest struct {
	IsUpdateCustomerBasicInfo     bool                         `json:"isUpdateCustomerBasicInfo"`
	CustomerBasicInfo             *UpdateCustomerBasicInfo     `json:"customerBasicInfo"`
	IsUpdateCustomerSensitiveInfo bool                         `json:"isUpdateCustomerSensitiveInfo"`
	CustomerSensitiveInfo         *UpdateCustomerSensitiveInfo `json:"customerSensitiveInfo"`
	IsUpdateCustomerWorkingInfo   bool                         `json:"isUpdateCustomerWorkingInfo"`
	CustomerWorkingInfo           *UpdateCustomerWorkingInfo   `json:"customerWorkingInfo"`
	IsUpdateCustomerAddressInfo   bool                         `json:"isUpdateCustomerAddressInfo"`
	CustomerAddressInfo           *[]UpdateCustomerAddressInfo `json:"customerAddressInfo"`
	IsUpdateCustomerContactInfo   bool                         `json:"isUpdateCustomerContactInfo"`
	CustomerContactInfo           *[]UpdateCustomerContactInfo `json:"customerContactInfo"`
	IsUpdateCustomerFlagInfo      bool                         `json:"isUpdateCustomerFlagInfo"`
	CustomerFlagInfo              *UpdateCustomerFlagInfo      `json:"customerFlagInfo"`
}

type CreateCustomerBasicInfo struct {
	Identifier        string `json:"identifier"`
	CifEms            string `json:"cifEms"`
	CifCbs            string `json:"cifCbs"`
	CifPaotang        string `json:"cifPaotang"`
	Owner             string `json:"owner"`
	CustomerCode      string `json:"customerCode"`
	TitleCodeTh       string `json:"titleCodeTh"`
	TitleTh           string `json:"titleTh"`
	FirstNameTh       string `json:"firstNameTh"`
	MiddleNameTh      string `json:"middleNameTh"`
	LastNameTh        string `json:"lastNameTh"`
	TitleCodeEn       string `json:"titleCodeEn"`
	TitleEn           string `json:"titleEn"`
	FirstNameEn       string `json:"firstNameEn"`
	MiddleNameEn      string `json:"middleNameEn"`
	LastNameEn        string `json:"lastNameEn"`
	Gender            string `json:"gender"`
	PassportNumber    string `json:"passportNumber"`
	CreatedChannel    string `json:"createdChannel"`
	CreatedProduct    string `json:"createdProduct"`
	CreatedAtChannel  string `json:"createdAtChannel"`
	CreatedBranchCode string `json:"createdBranchCode"`
}
type CreateCustomerSensitiveInfo struct {
	DateOfBirthTh        string `json:"dateOfBirthTh"`
	DateOfBirthEn        string `json:"dateOfBirthEn"`
	Nationality          string `json:"nationality"`
	CardIssueDate        string `json:"cardIssueDate"`
	CardExpireDate       string `json:"cardExpireDate"`
	CardIssueCenter      string `json:"cardIssueCenter"`
	PassportCountryIssue string `json:"passportCountryIssue"`
	PassportExpireDate   string `json:"passportExpireDate"`
	PassportIssueDate    string `json:"passportIssueDate"`
	CountryOfBirth       string `json:"countryOfBirth"`
	CountryOfNation      string `json:"countryOfNation"`
	MaritalStatus        string `json:"maritalStatus"`
	SpouseNameTh         string `json:"spouseNameTh"`
	SpouseNameEn         string `json:"spouseNameEn"`
	CreatedChannel       string `json:"createdChannel"`
	CreatedProduct       string `json:"createdProduct"`
	CreatedAtChannel     string `json:"createdAtChannel"`
	CreatedBranchCode    string `json:"createdBranchCode"`
}

type CreateCustomerWorkingInfo struct {
	OccupationCode    string `json:"occupationCode"`
	OccupationSector  string `json:"occupationSector"`
	SubOccupationCode string `json:"subOccupationCode"`
	IncomeCode        string `json:"incomeCode"`
	EducationCode     string `json:"educationCode"`
	CreatedChannel    string `json:"createdChannel"`
	CreatedProduct    string `json:"createdProduct"`
	CreatedAtChannel  string `json:"createdAtChannel"`
	CreatedBranchCode string `json:"createdBranchCode"`
}
type CreateCustomerAddressInfo struct {
	AddressType       string `json:"addressType"`
	AddressName       string `json:"addressName"`
	AddressLine1      string `json:"addressLine1"`
	AddressLine2      string `json:"addressLine2"`
	AddressLine3      string `json:"addressLine3"`
	SubDistrictCode   string `json:"subDistrictCode"`
	DistrictCode      string `json:"districtCode"`
	ProvinceCode      string `json:"provinceCode"`
	PostalCode        string `json:"postalCode"`
	Country           string `json:"country"`
	PhoneNo           string `json:"phoneNo"`
	PhoneExt          string `json:"phoneExt"`
	CreatedChannel    string `json:"createdChannel"`
	CreatedProduct    string `json:"createdProduct"`
	CreatedAtChannel  string `json:"createdAtChannel"`
	CreatedBranchCode string `json:"createdBranchCode"`
}

type CreateCustomerContactInfo struct {
	Channel           string `json:"channel"`
	MobileNo          string `json:"mobileNo"`
	Email             string `json:"email"`
	CreatedChannel    string `json:"createdChannel"`
	CreatedProduct    string `json:"createdProduct"`
	CreatedAtChannel  string `json:"createdAtChannel"`
	CreatedBranchCode string `json:"createdBranchCode"`
}
type CreateCustomerFlagInfo struct {
	KycBranch                 string `json:"kycBranch"`
	KycLevel                  string `json:"kycLevel"`
	KycRiskScore              string `json:"kycRiskScore"`
	KycCommentCode            string `json:"kycCommentCode"`
	KycReviewCode             string `json:"kycReviewCode"`
	KycStatusLastUpdate       string `json:"kycStatusLastUpdate"`
	AmlStatus                 string `json:"amlStatus"`
	AmlSubListType            string `json:"amlSubListType"`
	PoliticianRelateFlag      string `json:"politicianRelateFlag"`
	InterPoliticianRelateFlag string `json:"interPoliticianRelateFlag"`
	ResponseUnit              string `json:"responseUnit"`
	NdidFlag                  string `json:"ndidFlag"`
	CreatedChannel            string `json:"createdChannel"`
	CreatedProduct            string `json:"createdProduct"`
	CreatedAtChannel          string `json:"createdAtChannel"`
	CreatedBranchCode         string `json:"createdBranchCode"`
}

type UpdateCustomerBasicInfo struct {
	Identifier        string `json:"identifier"`
	CifEms            string `json:"cifEms"`
	CifCbs            string `json:"cifCbs"`
	CifPaotang        string `json:"cifPaotang"`
	Owner             string `json:"owner"`
	CustomerCode      string `json:"customerCode"`
	TitleCodeTh       string `json:"titleCodeTh"`
	TitleTh           string `json:"titleTh"`
	FirstNameTh       string `json:"firstNameTh"`
	MiddleNameTh      string `json:"middleNameTh"`
	LastNameTh        string `json:"lastNameTh"`
	TitleCodeEn       string `json:"titleCodeEn"`
	TitleEn           string `json:"titleEn"`
	FirstNameEn       string `json:"firstNameEn"`
	MiddleNameEn      string `json:"middleNameEn"`
	LastNameEn        string `json:"lastNameEn"`
	Gender            string `json:"gender"`
	PassportNumber    string `json:"passportNumber"`
	UpdatedChannel    string `json:"updatedChannel"`
	UpdatedProduct    string `json:"updatedProduct"`
	UpdatedAtChannel  string `json:"updatedAtChannel"`
	UpdatedBranchCode string `json:"updatedBranchCode"`
}
type UpdateCustomerSensitiveInfo struct {
	DateOfBirthTh        string `json:"dateOfBirthTh"`
	DateOfBirthEn        string `json:"dateOfBirthEn"`
	Nationality          string `json:"nationality"`
	CardIssueDate        string `json:"cardIssueDate"`
	CardExpireDate       string `json:"cardExpireDate"`
	CardIssueCenter      string `json:"cardIssueCenter"`
	PassportCountryIssue string `json:"passportCountryIssue"`
	PassportExpireDate   string `json:"passportExpireDate"`
	PassportIssueDate    string `json:"passportIssueDate"`
	CountryOfBirth       string `json:"countryOfBirth"`
	CountryOfNation      string `json:"countryOfNation"`
	MaritalStatus        string `json:"maritalStatus"`
	SpouseNameTh         string `json:"spouseNameTh"`
	SpouseNameEn         string `json:"spouseNameEn"`
	UpdatedChannel       string `json:"updatedChannel"`
	UpdatedProduct       string `json:"updatedProduct"`
	UpdatedAtChannel     string `json:"updatedAtChannel"`
	UpdatedBranchCode    string `json:"updatedBranchCode"`
}

type UpdateCustomerWorkingInfo struct {
	OccupationCode    string `json:"occupationCode"`
	OccupationSector  string `json:"occupationSector"`
	SubOccupationCode string `json:"subOccupationCode"`
	IncomeCode        string `json:"incomeCode"`
	EducationCode     string `json:"educationCode"`
	UpdatedChannel    string `json:"updatedChannel"`
	UpdatedProduct    string `json:"updatedProduct"`
	UpdatedAtChannel  string `json:"updatedAtChannel"`
	UpdatedBranchCode string `json:"updatedBranchCode"`
}
type UpdateCustomerAddressInfo struct {
	AddressType       string `json:"addressType"`
	AddressName       string `json:"addressName"`
	AddressLine1      string `json:"addressLine1"`
	AddressLine2      string `json:"addressLine2"`
	AddressLine3      string `json:"addressLine3"`
	SubDistrictCode   string `json:"subDistrictCode"`
	DistrictCode      string `json:"districtCode"`
	ProvinceCode      string `json:"provinceCode"`
	PostalCode        string `json:"postalCode"`
	Country           string `json:"country"`
	PhoneNo           string `json:"phoneNo"`
	PhoneExt          string `json:"phoneExt"`
	UpdatedChannel    string `json:"updatedChannel"`
	UpdatedProduct    string `json:"updatedProduct"`
	UpdatedAtChannel  string `json:"updatedAtChannel"`
	UpdatedBranchCode string `json:"updatedBranchCode"`
}

type UpdateCustomerContactInfo struct {
	Channel           string `json:"channel"`
	MobileNo          string `json:"mobileNo"`
	Email             string `json:"email"`
	UpdatedChannel    string `json:"updatedChannel"`
	UpdatedProduct    string `json:"updatedProduct"`
	UpdatedAtChannel  string `json:"updatedAtChannel"`
	UpdatedBranchCode string `json:"updatedBranchCode"`
}
type UpdateCustomerFlagInfo struct {
	KycBranch                 string `json:"kycBranch"`
	KycLevel                  string `json:"kycLevel"`
	KycRiskScore              string `json:"kycRiskScore"`
	KycCommentCode            string `json:"kycCommentCode"`
	KycReviewCode             string `json:"kycReviewCode"`
	KycStatusLastUpdate       string `json:"kycStatusLastUpdate"`
	AmlStatus                 string `json:"amlStatus"`
	AmlSubListType            string `json:"amlSubListType"`
	PoliticianRelateFlag      string `json:"politicianRelateFlag"`
	InterPoliticianRelateFlag string `json:"interPoliticianRelateFlag"`
	ResponseUnit              string `json:"responseUnit"`
	NdidFlag                  string `json:"ndidFlag"`
	UpdatedChannel            string `json:"updatedChannel"`
	UpdatedProduct            string `json:"updatedProduct"`
	UpdatedAtChannel          string `json:"updatedAtChannel"`
	UpdatedBranchCode         string `json:"updatedBranchCode"`
}
type CDICustomerProfileResponse struct {
	Code        string                         `json:"code"`
	Description string                         `json:"description"`
	Data        CDICustomerProfileResponseData `json:"data"`
}

type CDICustomerProfileResponseData struct {
	CustomerKey                 int          `json:"customerKey"`
	Identifier                  string       `json:"identifier"`
	CifEms                      string       `json:"cifEms"`
	CifCbs                      string       `json:"cifCbs"`
	CustomerBasicInfoStatus     *CDIResponse `json:"customerBasicInfo"`
	CustomerSensitiveInfoStatus *CDIResponse `json:"customerSensitiveInfo"`
	CustomerWorkingInfoStatus   *CDIResponse `json:"customerWorkingInfo"`
	CustomerAddressInfoStatus   *CDIResponse `json:"customerAddressInfo"`
	CustomerContactInfoStatus   *CDIResponse `json:"customerContactInfo"`
	CustomerFlagInfoStatus      *CDIResponse `json:"customerFlagInfo"`
}

type CDIResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
