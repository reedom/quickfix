package derivativeinstrument

import (
	"github.com/quickfixgo/quickfix/fix50sp2/derivativeinstrumentpartysubidsgrp"
	"time"
)

//NoDerivativeSecurityAltID is a repeating group in DerivativeInstrument
type NoDerivativeSecurityAltID struct {
	//DerivativeSecurityAltID is a non-required field for NoDerivativeSecurityAltID.
	DerivativeSecurityAltID *string `fix:"1219"`
	//DerivativeSecurityAltIDSource is a non-required field for NoDerivativeSecurityAltID.
	DerivativeSecurityAltIDSource *string `fix:"1220"`
}

//NoDerivativeEvents is a repeating group in DerivativeInstrument
type NoDerivativeEvents struct {
	//DerivativeEventType is a non-required field for NoDerivativeEvents.
	DerivativeEventType *int `fix:"1287"`
	//DerivativeEventDate is a non-required field for NoDerivativeEvents.
	DerivativeEventDate *string `fix:"1288"`
	//DerivativeEventTime is a non-required field for NoDerivativeEvents.
	DerivativeEventTime *time.Time `fix:"1289"`
	//DerivativeEventPx is a non-required field for NoDerivativeEvents.
	DerivativeEventPx *float64 `fix:"1290"`
	//DerivativeEventText is a non-required field for NoDerivativeEvents.
	DerivativeEventText *string `fix:"1291"`
}

//NoDerivativeInstrumentParties is a repeating group in DerivativeInstrument
type NoDerivativeInstrumentParties struct {
	//DerivativeInstrumentPartyID is a non-required field for NoDerivativeInstrumentParties.
	DerivativeInstrumentPartyID *string `fix:"1293"`
	//DerivativeInstrumentPartyIDSource is a non-required field for NoDerivativeInstrumentParties.
	DerivativeInstrumentPartyIDSource *string `fix:"1294"`
	//DerivativeInstrumentPartyRole is a non-required field for NoDerivativeInstrumentParties.
	DerivativeInstrumentPartyRole *int `fix:"1295"`
	//DerivativeInstrumentPartySubIDsGrp Component
	DerivativeInstrumentPartySubIDsGrp derivativeinstrumentpartysubidsgrp.Component
}

//Component is a fix50sp2 DerivativeInstrument Component
type Component struct {
	//DerivativeSymbol is a non-required field for DerivativeInstrument.
	DerivativeSymbol *string `fix:"1214"`
	//DerivativeSymbolSfx is a non-required field for DerivativeInstrument.
	DerivativeSymbolSfx *string `fix:"1215"`
	//DerivativeSecurityID is a non-required field for DerivativeInstrument.
	DerivativeSecurityID *string `fix:"1216"`
	//DerivativeSecurityIDSource is a non-required field for DerivativeInstrument.
	DerivativeSecurityIDSource *string `fix:"1217"`
	//NoDerivativeSecurityAltID is a non-required field for DerivativeInstrument.
	NoDerivativeSecurityAltID []NoDerivativeSecurityAltID `fix:"1218,omitempty"`
	//DerivativeProduct is a non-required field for DerivativeInstrument.
	DerivativeProduct *int `fix:"1246"`
	//DerivativeProductComplex is a non-required field for DerivativeInstrument.
	DerivativeProductComplex *string `fix:"1228"`
	//DerivFlexProductEligibilityIndicator is a non-required field for DerivativeInstrument.
	DerivFlexProductEligibilityIndicator *bool `fix:"1243"`
	//DerivativeSecurityGroup is a non-required field for DerivativeInstrument.
	DerivativeSecurityGroup *string `fix:"1247"`
	//DerivativeCFICode is a non-required field for DerivativeInstrument.
	DerivativeCFICode *string `fix:"1248"`
	//DerivativeSecurityType is a non-required field for DerivativeInstrument.
	DerivativeSecurityType *string `fix:"1249"`
	//DerivativeSecuritySubType is a non-required field for DerivativeInstrument.
	DerivativeSecuritySubType *string `fix:"1250"`
	//DerivativeMaturityMonthYear is a non-required field for DerivativeInstrument.
	DerivativeMaturityMonthYear *string `fix:"1251"`
	//DerivativeMaturityDate is a non-required field for DerivativeInstrument.
	DerivativeMaturityDate *string `fix:"1252"`
	//DerivativeMaturityTime is a non-required field for DerivativeInstrument.
	DerivativeMaturityTime *string `fix:"1253"`
	//DerivativeSettleOnOpenFlag is a non-required field for DerivativeInstrument.
	DerivativeSettleOnOpenFlag *string `fix:"1254"`
	//DerivativeInstrmtAssignmentMethod is a non-required field for DerivativeInstrument.
	DerivativeInstrmtAssignmentMethod *string `fix:"1255"`
	//DerivativeSecurityStatus is a non-required field for DerivativeInstrument.
	DerivativeSecurityStatus *string `fix:"1256"`
	//DerivativeIssueDate is a non-required field for DerivativeInstrument.
	DerivativeIssueDate *string `fix:"1276"`
	//DerivativeInstrRegistry is a non-required field for DerivativeInstrument.
	DerivativeInstrRegistry *string `fix:"1257"`
	//DerivativeCountryOfIssue is a non-required field for DerivativeInstrument.
	DerivativeCountryOfIssue *string `fix:"1258"`
	//DerivativeStateOrProvinceOfIssue is a non-required field for DerivativeInstrument.
	DerivativeStateOrProvinceOfIssue *string `fix:"1259"`
	//DerivativeStrikePrice is a non-required field for DerivativeInstrument.
	DerivativeStrikePrice *float64 `fix:"1261"`
	//DerivativeLocaleOfIssue is a non-required field for DerivativeInstrument.
	DerivativeLocaleOfIssue *string `fix:"1260"`
	//DerivativeStrikeCurrency is a non-required field for DerivativeInstrument.
	DerivativeStrikeCurrency *string `fix:"1262"`
	//DerivativeStrikeMultiplier is a non-required field for DerivativeInstrument.
	DerivativeStrikeMultiplier *float64 `fix:"1263"`
	//DerivativeStrikeValue is a non-required field for DerivativeInstrument.
	DerivativeStrikeValue *float64 `fix:"1264"`
	//DerivativeOptAttribute is a non-required field for DerivativeInstrument.
	DerivativeOptAttribute *string `fix:"1265"`
	//DerivativeContractMultiplier is a non-required field for DerivativeInstrument.
	DerivativeContractMultiplier *float64 `fix:"1266"`
	//DerivativeMinPriceIncrement is a non-required field for DerivativeInstrument.
	DerivativeMinPriceIncrement *float64 `fix:"1267"`
	//DerivativeMinPriceIncrementAmount is a non-required field for DerivativeInstrument.
	DerivativeMinPriceIncrementAmount *float64 `fix:"1268"`
	//DerivativeUnitOfMeasure is a non-required field for DerivativeInstrument.
	DerivativeUnitOfMeasure *string `fix:"1269"`
	//DerivativeUnitOfMeasureQty is a non-required field for DerivativeInstrument.
	DerivativeUnitOfMeasureQty *float64 `fix:"1270"`
	//DerivativePriceUnitOfMeasure is a non-required field for DerivativeInstrument.
	DerivativePriceUnitOfMeasure *string `fix:"1315"`
	//DerivativePriceUnitOfMeasureQty is a non-required field for DerivativeInstrument.
	DerivativePriceUnitOfMeasureQty *float64 `fix:"1316"`
	//DerivativeExerciseStyle is a non-required field for DerivativeInstrument.
	DerivativeExerciseStyle *string `fix:"1299"`
	//DerivativeOptPayAmount is a non-required field for DerivativeInstrument.
	DerivativeOptPayAmount *float64 `fix:"1225"`
	//DerivativeTimeUnit is a non-required field for DerivativeInstrument.
	DerivativeTimeUnit *string `fix:"1271"`
	//DerivativeSecurityExchange is a non-required field for DerivativeInstrument.
	DerivativeSecurityExchange *string `fix:"1272"`
	//DerivativePositionLimit is a non-required field for DerivativeInstrument.
	DerivativePositionLimit *int `fix:"1273"`
	//DerivativeNTPositionLimit is a non-required field for DerivativeInstrument.
	DerivativeNTPositionLimit *int `fix:"1274"`
	//DerivativeIssuer is a non-required field for DerivativeInstrument.
	DerivativeIssuer *string `fix:"1275"`
	//DerivativeEncodedIssuerLen is a non-required field for DerivativeInstrument.
	DerivativeEncodedIssuerLen *int `fix:"1277"`
	//DerivativeEncodedIssuer is a non-required field for DerivativeInstrument.
	DerivativeEncodedIssuer *string `fix:"1278"`
	//DerivativeSecurityDesc is a non-required field for DerivativeInstrument.
	DerivativeSecurityDesc *string `fix:"1279"`
	//DerivativeEncodedSecurityDescLen is a non-required field for DerivativeInstrument.
	DerivativeEncodedSecurityDescLen *int `fix:"1280"`
	//DerivativeEncodedSecurityDesc is a non-required field for DerivativeInstrument.
	DerivativeEncodedSecurityDesc *string `fix:"1281"`
	//DerivativeContractSettlMonth is a non-required field for DerivativeInstrument.
	DerivativeContractSettlMonth *string `fix:"1285"`
	//NoDerivativeEvents is a non-required field for DerivativeInstrument.
	NoDerivativeEvents []NoDerivativeEvents `fix:"1286,omitempty"`
	//NoDerivativeInstrumentParties is a non-required field for DerivativeInstrument.
	NoDerivativeInstrumentParties []NoDerivativeInstrumentParties `fix:"1292,omitempty"`
	//DerivativeSettlMethod is a non-required field for DerivativeInstrument.
	DerivativeSettlMethod *string `fix:"1317"`
	//DerivativePriceQuoteMethod is a non-required field for DerivativeInstrument.
	DerivativePriceQuoteMethod *string `fix:"1318"`
	//DerivativeValuationMethod is a non-required field for DerivativeInstrument.
	DerivativeValuationMethod *string `fix:"1319"`
	//DerivativeListMethod is a non-required field for DerivativeInstrument.
	DerivativeListMethod *int `fix:"1320"`
	//DerivativeCapPrice is a non-required field for DerivativeInstrument.
	DerivativeCapPrice *float64 `fix:"1321"`
	//DerivativeFloorPrice is a non-required field for DerivativeInstrument.
	DerivativeFloorPrice *float64 `fix:"1322"`
	//DerivativePutOrCall is a non-required field for DerivativeInstrument.
	DerivativePutOrCall *int `fix:"1323"`
	//DerivativeSecurityXMLLen is a non-required field for DerivativeInstrument.
	DerivativeSecurityXMLLen *int `fix:"1282"`
	//DerivativeSecurityXML is a non-required field for DerivativeInstrument.
	DerivativeSecurityXML *string `fix:"1283"`
	//DerivativeSecurityXMLSchema is a non-required field for DerivativeInstrument.
	DerivativeSecurityXMLSchema *string `fix:"1284"`
	//DerivativeContractMultiplierUnit is a non-required field for DerivativeInstrument.
	DerivativeContractMultiplierUnit *int `fix:"1438"`
	//DerivativeFlowScheduleType is a non-required field for DerivativeInstrument.
	DerivativeFlowScheduleType *int `fix:"1442"`
}

func New() *Component { return new(Component) }