package model

import "github.com/SCR-NETWORK/SCR_Network/domain/consensus/model/externalapi"

// HeadersSelectedTipManager manages the state of the headers selected tip
type HeadersSelectedTipManager interface {
	AddHeaderTip(stagingArea *StagingArea, hash *externalapi.DomainHash) error
}
