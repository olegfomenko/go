package io

import (
	"github.com/stellar/go/xdr"
)

// StatsChangeProcessor is a state processors that counts number of changes types
// and entry types.
type StatsChangeProcessor struct {
	results StatsChangeProcessorResults
}

// StatsChangeProcessorResults contains results after running StatsChangeProcessor.
type StatsChangeProcessorResults struct {
	AccountsCreated int64
	AccountsUpdated int64
	AccountsRemoved int64

	DataCreated int64
	DataUpdated int64
	DataRemoved int64

	OffersCreated int64
	OffersUpdated int64
	OffersRemoved int64

	TrustLinesCreated int64
	TrustLinesUpdated int64
	TrustLinesRemoved int64
}

func (p *StatsChangeProcessor) ProcessChange(change Change) error {
	switch change.Type {
	case xdr.LedgerEntryTypeAccount:
		switch change.LedgerEntryChangeType() {
		case xdr.LedgerEntryChangeTypeLedgerEntryCreated:
			p.results.AccountsCreated++
		case xdr.LedgerEntryChangeTypeLedgerEntryUpdated:
			p.results.AccountsUpdated++
		case xdr.LedgerEntryChangeTypeLedgerEntryRemoved:
			p.results.AccountsRemoved++
		}
	case xdr.LedgerEntryTypeData:
		switch change.LedgerEntryChangeType() {
		case xdr.LedgerEntryChangeTypeLedgerEntryCreated:
			p.results.DataCreated++
		case xdr.LedgerEntryChangeTypeLedgerEntryUpdated:
			p.results.DataUpdated++
		case xdr.LedgerEntryChangeTypeLedgerEntryRemoved:
			p.results.DataRemoved++
		}
	case xdr.LedgerEntryTypeOffer:
		switch change.LedgerEntryChangeType() {
		case xdr.LedgerEntryChangeTypeLedgerEntryCreated:
			p.results.OffersCreated++
		case xdr.LedgerEntryChangeTypeLedgerEntryUpdated:
			p.results.OffersUpdated++
		case xdr.LedgerEntryChangeTypeLedgerEntryRemoved:
			p.results.OffersRemoved++
		}
	case xdr.LedgerEntryTypeTrustline:
		switch change.LedgerEntryChangeType() {
		case xdr.LedgerEntryChangeTypeLedgerEntryCreated:
			p.results.TrustLinesCreated++
		case xdr.LedgerEntryChangeTypeLedgerEntryUpdated:
			p.results.TrustLinesUpdated++
		case xdr.LedgerEntryChangeTypeLedgerEntryRemoved:
			p.results.TrustLinesRemoved++
		}
	}

	return nil
}

func (p *StatsChangeProcessor) GetResults() StatsChangeProcessorResults {
	return p.results
}
