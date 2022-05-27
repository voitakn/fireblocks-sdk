package fireblocks

const (
	ApiKeyAuthScopes      = "ApiKeyAuth.Scopes"
	BearerTokenAuthScopes = "bearerTokenAuth.Scopes"
)

// Defines values for AuthorizationInfoLogic.
const (
	AuthorizationInfoLogicAND AuthorizationInfoLogic = "AND"

	AuthorizationInfoLogicOR AuthorizationInfoLogic = "OR"
)

// Defines values for SignedMessageAlgorithm.
const (
	SignedMessageAlgorithmMPCECDSASECP256K1 SignedMessageAlgorithm = "MPC_ECDSA_SECP256K1"

	SignedMessageAlgorithmMPCECDSASECP256R1 SignedMessageAlgorithm = "MPC_ECDSA_SECP256R1"

	SignedMessageAlgorithmMPCEDDSAED25519 SignedMessageAlgorithm = "MPC_EDDSA_ED25519"
)

// Defines values for TransactionOperation.
const (
	TransactionOperationBURN TransactionOperation = "BURN"

	TransactionOperationCONTRACTCALL TransactionOperation = "CONTRACT_CALL"

	TransactionOperationINTERNALLEDGERTRANSFER TransactionOperation = "INTERNAL_LEDGER_TRANSFER"

	TransactionOperationMINT TransactionOperation = "MINT"

	TransactionOperationRAW TransactionOperation = "RAW"

	TransactionOperationREDEEMFROMCOMPOUND TransactionOperation = "REDEEM_FROM_COMPOUND"

	TransactionOperationSUPPLYTOCOMPOUND TransactionOperation = "SUPPLY_TO_COMPOUND"

	TransactionOperationTRANSFER TransactionOperation = "TRANSFER"

	TransactionOperationTYPEDMESSAGE TransactionOperation = "TYPED_MESSAGE"
)

// Defines values for TransactionRequestFeeLevel.
const (
	TransactionRequestFeeLevelHIGH TransactionRequestFeeLevel = "HIGH"

	TransactionRequestFeeLevelLOW TransactionRequestFeeLevel = "LOW"

	TransactionRequestFeeLevelMEDIUM TransactionRequestFeeLevel = "MEDIUM"
)

// Defines values for TransactionResponseStatus.
const (
	TransactionResponseStatusBLOCKED TransactionResponseStatus = "BLOCKED"

	TransactionResponseStatusBROADCASTING TransactionResponseStatus = "BROADCASTING"

	TransactionResponseStatusCANCELLED TransactionResponseStatus = "CANCELLED"

	TransactionResponseStatusCANCELLING TransactionResponseStatus = "CANCELLING"

	TransactionResponseStatusCOMPLETED TransactionResponseStatus = "COMPLETED"

	TransactionResponseStatusCONFIRMING TransactionResponseStatus = "CONFIRMING"

	TransactionResponseStatusFAILED TransactionResponseStatus = "FAILED"

	TransactionResponseStatusPARTIALLYCOMPLETED TransactionResponseStatus = "PARTIALLY_COMPLETED"

	TransactionResponseStatusPENDING3RDPARTY TransactionResponseStatus = "PENDING_3RD_PARTY"

	TransactionResponseStatusPENDING3RDPARTYMANUALAPPROVAL TransactionResponseStatus = "PENDING_3RD_PARTY_MANUAL_APPROVAL"

	TransactionResponseStatusPENDINGAMLSCREENING TransactionResponseStatus = "PENDING_AML_SCREENING"

	TransactionResponseStatusPENDINGAUTHORIZATION TransactionResponseStatus = "PENDING_AUTHORIZATION"

	TransactionResponseStatusPENDINGSIGNATURE TransactionResponseStatus = "PENDING_SIGNATURE"

	TransactionResponseStatusQUEUED TransactionResponseStatus = "QUEUED"

	TransactionResponseStatusREJECTED TransactionResponseStatus = "REJECTED"

	TransactionResponseStatusSUBMITTED TransactionResponseStatus = "SUBMITTED"

	TransactionResponseStatusTIMEOUT TransactionResponseStatus = "TIMEOUT"
)

// Defines values for TransactionSubStatus.
const (
	TransactionSubStatusACTUALFEETOOHIGH TransactionSubStatus = "ACTUAL_FEE_TOO_HIGH"

	TransactionSubStatusADDRESSNOTWHITELISTED TransactionSubStatus = "ADDRESS_NOT_WHITELISTED"

	TransactionSubStatusAMLSCREENINGREJECTED TransactionSubStatus = "AML_SCREENING_REJECTED"

	TransactionSubStatusAMOUNTTOOSMALL TransactionSubStatus = "AMOUNT_TOO_SMALL"

	TransactionSubStatusAPICALLLIMIT TransactionSubStatus = "API_CALL_LIMIT"

	TransactionSubStatusAPIINVALIDSIGNATURE TransactionSubStatus = "API_INVALID_SIGNATURE"

	TransactionSubStatusAUTHORIZERNOTFOUND TransactionSubStatus = "AUTHORIZER_NOT_FOUND"

	TransactionSubStatusBLOCKEDBYPOLICY TransactionSubStatus = "BLOCKED_BY_POLICY"

	TransactionSubStatusCANCELLEDBYUSER TransactionSubStatus = "CANCELLED_BY_USER"

	TransactionSubStatusCANCELLEDEXTERNALLY TransactionSubStatus = "CANCELLED_EXTERNALLY"

	TransactionSubStatusCONFIRMED TransactionSubStatus = "CONFIRMED"

	TransactionSubStatusCONNECTIVITYERROR TransactionSubStatus = "CONNECTIVITY_ERROR"

	TransactionSubStatusDROPPEDBYBLOCKCHAIN TransactionSubStatus = "DROPPED_BY_BLOCKCHAIN"

	TransactionSubStatusENVUNSUPPORTEDASSET TransactionSubStatus = "ENV_UNSUPPORTED_ASSET"

	TransactionSubStatusERRORUNSUPPORTEDTRANSACTIONTYPE TransactionSubStatus = "ERROR_UNSUPPORTED_TRANSACTION_TYPE"

	TransactionSubStatusFAILEDAMLSCREENING TransactionSubStatus = "FAILED_AML_SCREENING"

	TransactionSubStatusFAILONLOWFEE TransactionSubStatus = "FAIL_ON_LOW_FEE"

	TransactionSubStatusGASLIMITTOOLOW TransactionSubStatus = "GAS_LIMIT_TOO_LOW"

	TransactionSubStatusGASPRICETOOLOWFORRBF TransactionSubStatus = "GAS_PRICE_TOO_LOW_FOR_RBF"

	TransactionSubStatusINCOMPLETEUSERSETUP TransactionSubStatus = "INCOMPLETE_USER_SETUP"

	TransactionSubStatusINSUFFICIENTFUNDS TransactionSubStatus = "INSUFFICIENT_FUNDS"

	TransactionSubStatusINSUFFICIENTFUNDSFORFEE TransactionSubStatus = "INSUFFICIENT_FUNDS_FOR_FEE"

	TransactionSubStatusINSUFFICIENTRESERVEDFUNDING TransactionSubStatus = "INSUFFICIENT_RESERVED_FUNDING"

	TransactionSubStatusINTERNALERROR TransactionSubStatus = "INTERNAL_ERROR"

	TransactionSubStatusINVALIDADDRESS TransactionSubStatus = "INVALID_ADDRESS"

	TransactionSubStatusINVALIDCONTRACTCALLDATA TransactionSubStatus = "INVALID_CONTRACT_CALL_DATA"

	TransactionSubStatusINVALIDEXCHANGEACCOUNT TransactionSubStatus = "INVALID_EXCHANGE_ACCOUNT"

	TransactionSubStatusINVALIDFEE TransactionSubStatus = "INVALID_FEE"

	TransactionSubStatusINVALIDFEEPARAMS TransactionSubStatus = "INVALID_FEE_PARAMS"

	TransactionSubStatusINVALIDNONCEFORRBF TransactionSubStatus = "INVALID_NONCE_FOR_RBF"

	TransactionSubStatusINVALIDNONCETOOHIGH TransactionSubStatus = "INVALID_NONCE_TOO_HIGH"

	TransactionSubStatusINVALIDNONCETOOLOW TransactionSubStatus = "INVALID_NONCE_TOO_LOW"

	TransactionSubStatusINVALIDSIGNATURE TransactionSubStatus = "INVALID_SIGNATURE"

	TransactionSubStatusINVALIDTAGORMEMO TransactionSubStatus = "INVALID_TAG_OR_MEMO"

	TransactionSubStatusINVALIDTHIRDPARTYRESPONSE TransactionSubStatus = "INVALID_THIRD_PARTY_RESPONSE"

	TransactionSubStatusINVALIDUNMANAGEDWALLET TransactionSubStatus = "INVALID_UNMANAGED_WALLET"

	TransactionSubStatusMANUALDEPOSITADDRESSREQUIRED TransactionSubStatus = "MANUAL_DEPOSIT_ADDRESS_REQUIRED"

	TransactionSubStatusMAXFEEEXCEEDED TransactionSubStatus = "MAX_FEE_EXCEEDED"

	TransactionSubStatusMISSINGDEPOSITADDRESS TransactionSubStatus = "MISSING_DEPOSIT_ADDRESS"

	TransactionSubStatusMISSINGTAGORMEMO TransactionSubStatus = "MISSING_TAG_OR_MEMO"

	TransactionSubStatusN3RDPARTYCANCELLED TransactionSubStatus = "3RD_PARTY_CANCELLED"

	TransactionSubStatusN3RDPARTYCOMPLETED TransactionSubStatus = "3RD_PARTY_COMPLETED"

	TransactionSubStatusN3RDPARTYCONFIRMING TransactionSubStatus = "3RD_PARTY_CONFIRMING"

	TransactionSubStatusN3RDPARTYFAILED TransactionSubStatus = "3RD_PARTY_FAILED"

	TransactionSubStatusN3RDPARTYPROCESSING TransactionSubStatus = "3RD_PARTY_PROCESSING"

	TransactionSubStatusN3RDPARTYREJECTED TransactionSubStatus = "3RD_PARTY_REJECTED"

	TransactionSubStatusNEEDMORETOCREATEDESTINATION TransactionSubStatus = "NEED_MORE_TO_CREATE_DESTINATION"

	TransactionSubStatusNONEXISTINGACCOUNTNAME TransactionSubStatus = "NON_EXISTING_ACCOUNT_NAME"

	TransactionSubStatusONETIMEADDRESSDISABLED TransactionSubStatus = "ONE_TIME_ADDRESS_DISABLED"

	TransactionSubStatusPARTIALLYFAILED TransactionSubStatus = "PARTIALLY_FAILED"

	TransactionSubStatusPENDINGBLOCKCHAINCONFIRMATIONS TransactionSubStatus = "PENDING_BLOCKCHAIN_CONFIRMATIONS"

	TransactionSubStatusSIGNERNOTFOUND TransactionSubStatus = "SIGNER_NOT_FOUND"

	TransactionSubStatusSIGNINGERROR TransactionSubStatus = "SIGNING_ERROR"

	TransactionSubStatusTHIRDPARTYINTERNALERROR TransactionSubStatus = "THIRD_PARTY_INTERNAL_ERROR"

	TransactionSubStatusTIMEOUT TransactionSubStatus = "TIMEOUT"

	TransactionSubStatusTOOLONGMEMPOOLCHAIN TransactionSubStatus = "TOO_LONG_MEMPOOL_CHAIN"

	TransactionSubStatusTOOMANYINPUTS TransactionSubStatus = "TOO_MANY_INPUTS"

	TransactionSubStatusTXOUTDATED TransactionSubStatus = "TX_OUTDATED"

	TransactionSubStatusUNAUTHORISEDDEVICE TransactionSubStatus = "UNAUTHORISED__DEVICE"

	TransactionSubStatusUNAUTHORISEDMISSINGCREDENTIALS TransactionSubStatus = "UNAUTHORISED__MISSING_CREDENTIALS"

	TransactionSubStatusUNAUTHORISEDMISSINGPERMISSION TransactionSubStatus = "UNAUTHORISED__MISSING_PERMISSION"

	TransactionSubStatusUNAUTHORISEDUSER TransactionSubStatus = "UNAUTHORISED__USER"

	TransactionSubStatusUNKNOWNERROR TransactionSubStatus = "UNKNOWN_ERROR"

	TransactionSubStatusUNSUPPORTEDASSET TransactionSubStatus = "UNSUPPORTED_ASSET"

	TransactionSubStatusUNSUPPORTEDOPERATION TransactionSubStatus = "UNSUPPORTED_OPERATION"

	TransactionSubStatusVAULTWALLETNOTREADY TransactionSubStatus = "VAULT_WALLET_NOT_READY"

	TransactionSubStatusWITHDRAWLIMIT TransactionSubStatus = "WITHDRAW_LIMIT"

	TransactionSubStatusZEROBALANCEINPERMANENTADDRESS TransactionSubStatus = "ZERO_BALANCE_IN_PERMANENT_ADDRESS"
)

// Defines values for TransferPeerPathType.
const (
	TransferPeerPathTypeCOMPOUND TransferPeerPathType = "COMPOUND"

	TransferPeerPathTypeEXCHANGEACCOUNT TransferPeerPathType = "EXCHANGE_ACCOUNT"

	TransferPeerPathTypeEXTERNALWALLET TransferPeerPathType = "EXTERNAL_WALLET"

	TransferPeerPathTypeFIATACCOUNT TransferPeerPathType = "FIAT_ACCOUNT"

	TransferPeerPathTypeGASSTATION TransferPeerPathType = "GAS_STATION"

	TransferPeerPathTypeINTERNALWALLET TransferPeerPathType = "INTERNAL_WALLET"

	TransferPeerPathTypeNETWORKCONNECTION TransferPeerPathType = "NETWORK_CONNECTION"

	TransferPeerPathTypeONETIMEADDRESS TransferPeerPathType = "ONE_TIME_ADDRESS"

	TransferPeerPathTypeUNKNOWN TransferPeerPathType = "UNKNOWN"

	TransferPeerPathTypeVAULTACCOUNT TransferPeerPathType = "VAULT_ACCOUNT"
)

// Defines values for TransferPeerPathResponseVirtualType.
const (
	TransferPeerPathResponseVirtualTypeDEFAULT TransferPeerPathResponseVirtualType = "DEFAULT"

	TransferPeerPathResponseVirtualTypeOFFEXCHANGE TransferPeerPathResponseVirtualType = "OFF_EXCHANGE"

	TransferPeerPathResponseVirtualTypeUNKNOWN TransferPeerPathResponseVirtualType = "UNKNOWN"
)

// Defines values for VaultWalletAddressAddressFormat.
const (
	VaultWalletAddressAddressFormatLEGACY VaultWalletAddressAddressFormat = "LEGACY"

	VaultWalletAddressAddressFormatSEGWIT VaultWalletAddressAddressFormat = "SEGWIT"
)
