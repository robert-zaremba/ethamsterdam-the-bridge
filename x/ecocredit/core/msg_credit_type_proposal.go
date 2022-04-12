package core

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

var _ govtypes.Content = &CreditTypeProposal{}

const (
	ProposalType = "CreditTypeProposal"
	RouterKey    = ecocredit.ModuleName
)

func init() {
	govtypes.RegisterProposalType(ProposalType)
}

func NewCreditTypeProposal(title, description string, ct CreditType) govtypes.Content {
	return &CreditTypeProposal{title, description, &ct}
}

func (m *CreditTypeProposal) ProposalRoute() string { return RouterKey }

func (m *CreditTypeProposal) ProposalType() string { return ProposalType }

func (m *CreditTypeProposal) ValidateBasic() error {
	if m.CreditType == nil {
		return sdkerrors.ErrInvalidRequest.Wrap("credit type cannot be nil")
	}
	if err := m.CreditType.Validate(); err != nil {
		return err
	}
	return govtypes.ValidateAbstract(m)
}

func (m *CreditTypeProposal) String() string {
	return fmt.Sprintf(`Credit Type Proposal:
  Title:       %s
  Description: %s
  Credit Type: %v
`, m.Title, m.Description, m.CreditType)
}
