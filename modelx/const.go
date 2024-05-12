package modelx

const (
	// Redis Key
	KeyBiometricCheckDup          = "biometric_check_dup#%s"
	KeyBiometricFaceCompareState  = "biometric_face_compare_state#%s"
	KeyBiometricFaceCompareResult = "biometric_face_compare_result#%s"

	// Response code
	CodeSuccess                 = "0000"
	CodeInternalError           = "1001"
	CodeProcessing              = "1002"
	CodeNotFound                = "1003"
	CodeBadRequest              = "1004"
	CodeDuplicatedRefID         = "1005"
	CodeDipchipNotFound         = "1006"
	CodeIALNotFound             = "1007"
	CodeCustomerProfileNotFound = "1008"
	CodeFlagsRequired           = "1009"
	CodeDipchipInfoNotFound     = "1010"
	CodeFATCAInfoNotFound       = "1011"
	CodeChannelNotRegistered    = "1012"

	CodeMDMInternalError       = "2000"
	CodeMDMFail                = "2001"
	CodeMDMUnAuthenticated     = "2002"
	CodeMDMTimeout             = "2003"
	CodeMDMMQConnectionBroken  = "2004"
	CodeBRSTimeout             = "2005"
	CodeMDMMQUnknownObjectName = "2006"

	CodeUnableToGetMasterData = "3000"

	CodeInputErrorValidationData                  = "4000"
	CodeBadRequestData                            = "4001"
	CodeNotFounndData                             = "4002"
	CodeNotFoundBiometricCompareFaceResultData    = "4003"
	CodeBiometricCompareFaceResultFormatErrorData = "4004"
	CodeBiometricCompareFaceResultNotSuccessData  = "4405"

	CodeInternalServerError = "5001"
	CodeKafkaError          = "5100"
	CodeRedisError          = "5200"
	// Begin Error Code 5500 About DB Error
	CodeMySQLError   = "5500"
	CodeMongoDBError = "5600"

	CodeUpdateCBSCustomerInfoSegment1DataFail = "6001"
	CodeUpdateCBSCustomerInfoSegment2DataFail = "6002"
	CodeUpdateCBSLegalAddressInfoDataFail     = "6003"
	CodeUpdateCBSMailingAddressInfoDataFail   = "6004"
	CodeUpdateCBSOfficeAddressInfoDataFail    = "6005"
	CodeUpdateCBSContactInfoDataFail          = "6006"
	CodeUpdateCBSKycAmlInfoDataFail           = "6007"

	CodeUpdateCBSCustomerInfoSegment1DataTimeout = "6011"
	CodeUpdateCBSCustomerInfoSegment2DataTimeout = "6012"
	CodeUpdateCBSLegalAddressInfoDataTimeout     = "6013"
	CodeUpdateCBSMailingAddressInfoDataTimeout   = "6014"
	CodeUpdateCBSOfficeAddressInfoDataTimeout    = "6015"
	CodeUpdateCBSContactInfoDataTimeout          = "6016"
	CodeUpdateCBSKycAmlInfoDataTimeout           = "6017"

	CodeCBSBusinessError = "6100"
	CodeCBSTimeoutError  = "6200"
	CodeCBSServerError   = "6300"
	// State
	StateProcess               = "PROCESSING"
	StateSuccess               = "SUCCESS"
	StateFail                  = "FAIL"
	StateError                 = "ERROR"
	StateCallbackFail          = "CALLBACK_FAIL"
	StateCallbackEKYCKafkaFail = "CALLBACK_EKYC_KAFKA_FAIL"

	// Source type
	SourceKTB   = "KTB"
	SourceGOV   = "GOV"
	SourceOther = "OTHER"

	// Platform name
	PlatformEKYC = "E-KYC"

	MQEKYCRequesterName = "eKYC"

	// Service name
	ServiceBiometricsFaceCompare = "BIOMETRICS-FACE-COMPARE"
	ServiceCustomerFlags         = "flags"

	//Headers
	HeaderXService        = "X-Service"
	HeaderXRefID          = "X-Ref-Id"
	HeaderXPlatform       = "X-Platform"
	HeaderXTrustedSource  = "X-Trusted-Source"
	HeaderXCallbackURL    = "X-Callback-Url"
	HeaderXChannel        = "X-Channel"
	HeaderXProduct        = "X-Product"
	HeaderXSource         = "X-Source"
	HeaderXApiKey         = "X-Api-Key"
	HeaderXReqId          = "X-Request-Id"
	HeaderXTopicId        = "X-Topic-Id"
	HeaderXStreamKafka    = "X-Stream-Kafka"
	HeaderXSessionId      = "X-Session-Id"
	HeaderXPairId         = "X-Pair-Id"
	HeaderXAdditionalInfo = "X-Additional-Info"
	HeaderXIdentifierType = "X-Identifier-Type"
	HeaderXIdentifier     = "X-Identifier"
	HeaderXAction         = "X-Action"
	HeaderXCreateType     = "X-Create-Type"
	HeaderXDevopsSrc      = "x-devops-src"
	HeaderXDevopsDest     = "x-devops-dest"
	HeaderXDevopsKey      = "x-devops-key"
	HeaderAppID           = "appId"
	HeaderAppSecret       = "appSecret"
	HeaderReferenceId     = "referenceId"
	HeaderServiceId       = "serviceId"
	HeaderAccessToken     = "accessToken"
	HeaderConsentName     = "consentName"

	PaotangChannel      = "PAOTANG"
	NextChannel         = "NEXT"
	LottoChannel        = "LOTTO"
	GlodgChannel        = "GLODG"
	InsuranceChannel    = "INSURANCE"
	KtxChannel          = "KTX"
	CardPlatformChannel = "CARD_PLATFORM"

	NextKtx                  = "NEXT_KTX"
	NextMutualFund           = "NEXT_MUTUAL_FUND"
	NextMutualFundOAPlatform = "MFOA"
	NextDigitalLending       = "NEXT_DIGITAL_LENDING"
	NextExistingToBank       = "NEXT_ETB"
	NextNewToBank            = "NEXT_NTB"
	NextTravelCard           = "NEXT_TRAVEL_CARD"
	NextEMoney               = "NEXT_EMONEY"
	NextInterWallet          = "NEXT_INTERWALLET"
	NextDebitCard            = "NEXT_DEBITCARD"
	NextVirtualDebitCard     = "NEXT_VIRTUALDEBITCARD"
	NextCasaETB              = "NEXT_CASAETB"
	PaotangGoldWallet        = "PT_GOLDWALLET"
	PaotangPDMO1stMarket     = "PT_PDMO1stMKT"
	PaotangPDMO2ndMarket     = "PT_PDMO2ndMKT"
	PaotangPDMOLG            = "PT_PDMOLG"
	PaotangCorpBond          = "PT_CORPBOND"
	PaotangGWallet           = "PT_GWALLET"
	NdidEnrollment           = "NDID-ENROLLMENT"
	PaotangPay               = "PT_PAY"
	NextKTX                  = "KTX"

	CodeNotFoundCustomerAMLOnCloud = 0
	CodeFoundCustomerAMLOnCloud    = 1

	FatcaStatusHasQA       = "Y"
	FatcaStatusHasNotQA    = "N"
	FatcaClass01           = "01"
	FatcaIrsFormTypeNoForm = "NOFORM"

	AmlRegulation       = "AML"
	AmlKycRegulation    = "AMLKYC"
	IalRegulation       = "IAL"
	FatcaRegulation     = "FATCA"
	FraudRegulation     = "FRAUD"
	SuspectRegulation   = "SUSPECT"
	WhitelistRegulation = "WHITELIST"
)

var Description = map[string]string{
	CodeSuccess:                                   "Success",
	CodeInternalError:                             "Internal server error",
	CodeProcessing:                                "Processing",
	CodeNotFound:                                  "Not Found",
	CodeBadRequest:                                "Bad Request",
	CodeDuplicatedRefID:                           "Duplicated RefId",
	CodeDipchipNotFound:                           "Dipchip not found",
	CodeIALNotFound:                               "Ial not found",
	CodeCustomerProfileNotFound:                   "Customer profile not found.",
	CodeFlagsRequired:                             "Flags are required.",
	CodeMDMInternalError:                          "MDM System Internal Error.",
	CodeMDMFail:                                   "MDM System Fail with Params.",
	CodeMDMUnAuthenticated:                        "UnAuthenticated from MDM System.",
	CodeMDMTimeout:                                "MDM System Timeout.",
	CodeMDMMQConnectionBroken:                     "MDM MQ ESB Connection Broken.",
	CodeDipchipInfoNotFound:                       "Dipchip information not found",
	CodeMDMMQUnknownObjectName:                    "MDM need to add queue name configuration.",
	CodeFATCAInfoNotFound:                         "FATCA information not found",
	CodeUnableToGetMasterData:                     "Unable to get master data",
	CodeInputErrorValidationData:                  "Input Validation Error. (%s)",
	CodeBadRequestData:                            "Bad Request Data or Request Data Required",
	CodeNotFounndData:                             "Error Not Found Path or Data",
	CodeInternalServerError:                       "Internal Server Error.",
	CodeChannelNotRegistered:                      "This channel is not registered MDM Config.",
	CodeKafkaError:                                "Kafka Cluster Error.",
	CodeRedisError:                                "Redis Error.",
	CodeMySQLError:                                "MySQL Database Error.",
	CodeMongoDBError:                              "MongoDB Error.",
	CodeBRSTimeout:                                "BRS System Timeout",
	CodeNotFoundBiometricCompareFaceResultData:    "Not Found Biometric Compare-Face Result.",
	CodeBiometricCompareFaceResultFormatErrorData: "Biometric Compare-Face Result Data Format Error.",
	CodeBiometricCompareFaceResultNotSuccessData:  "Biometric Compare-Face Result Not Success.",
	CodeUpdateCBSCustomerInfoSegment1DataFail:     "Error Update Customer Info Segment1 to CBS.",
	CodeUpdateCBSCustomerInfoSegment2DataFail:     "Error Update Customer Info Segment2 to CBS.",
	CodeUpdateCBSLegalAddressInfoDataFail:         "Error Update Legal Address Info to CBS.",
	CodeUpdateCBSMailingAddressInfoDataFail:       "Error Update Mailing Address Info to CBS.",
	CodeUpdateCBSOfficeAddressInfoDataFail:        "Error Update Office Addres Info to CBS.",
	CodeUpdateCBSContactInfoDataFail:              "Error Update Contact Info to CBS.",
	CodeUpdateCBSKycAmlInfoDataFail:               "Error Update Kyc Aml Info to CBS.",
	CodeUpdateCBSCustomerInfoSegment1DataTimeout:  "Update Customer Info Segment1 to CBS Timeout.",
	CodeUpdateCBSCustomerInfoSegment2DataTimeout:  "Update Customer Info Segment2 to CBS Timeout.",
	CodeUpdateCBSLegalAddressInfoDataTimeout:      "Update Legal Address Info to CBS Timeout.",
	CodeUpdateCBSMailingAddressInfoDataTimeout:    "Update Mailing Address Info to CBS Timeout.",
	CodeUpdateCBSOfficeAddressInfoDataTimeout:     "Update Office Addres Info to CBS Timeout.",
	CodeUpdateCBSContactInfoDataTimeout:           "Update Contact Info to CBS Timeout.",
	CodeUpdateCBSKycAmlInfoDataTimeout:            "Update Kyc Aml Info to CBS Timeout.",
	CodeCBSBusinessError:                          "CBS Business Error (%s)",
	CodeCBSTimeoutError:                           "CBS Timeout Error.",
	CodeCBSServerError:                            "CBS Server Error.",
}

var MQProductChannel = map[string]string{
	NextMutualFund:           "MF",
	NextMutualFundOAPlatform: "MFOA",
	NextTravelCard:           "TRC",
	NextEMoney:               "EM",
	NextExistingToBank:       "ETB",
	NextNewToBank:            "NTB",
	NextInterWallet:          "INTERWALLET",
	NextDebitCard:            "NEXT_DEBITCARD",
	NextVirtualDebitCard:     "NEXT_VIRTUALDEBITCARD",
	NextCasaETB:              "NEXT_CASAETB",
	PaotangGoldWallet:        "PT_GOLDWALLET",
	PaotangPDMO1stMarket:     "PT_PDMO1stMKT",
	PaotangPDMO2ndMarket:     "PT_PDMO2ndMKT",
	PaotangPDMOLG:            "PT_PDMOLG",
	PaotangCorpBond:          "PT_CORPBOND",
	PaotangGWallet:           "PT_GWALLET",
	NdidEnrollment:           "N_NDID",
	PaotangPay:               "PT_PAY",
	NextKTX:                  "KTX",
}
