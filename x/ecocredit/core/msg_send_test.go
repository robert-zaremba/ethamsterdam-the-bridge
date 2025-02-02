package core

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/types/testutil"
)

func TestMsgSend(t *testing.T) {
	t.Parallel()

	addr1 := testutil.GenAddress()
	addr2 := testutil.GenAddress()

	tests := map[string]struct {
		src    MsgSend
		expErr bool
	}{
		"valid msg": {
			src: MsgSend{
				Sender:    addr1,
				Recipient: addr2,
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:             batchDenom,
						TradableAmount:         "10",
						RetiredAmount:          "10",
						RetirementJurisdiction: "ST-UVW XY Z12",
					},
				},
			},
			expErr: false,
		},
		"invalid msg with Credits.RetiredAmount negative value": {
			src: MsgSend{
				Sender:    addr1,
				Recipient: addr2,
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:     "some_denom",
						TradableAmount: "10",
						RetiredAmount:  "-10",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without credits": {
			src: MsgSend{
				Sender:    addr1,
				Recipient: addr2,
			},
			expErr: true,
		},
		"invalid msg without sender": {
			src: MsgSend{
				Recipient: addr2,
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:             "some_denom",
						TradableAmount:         "10",
						RetiredAmount:          "10",
						RetirementJurisdiction: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without recipient": {
			src: MsgSend{
				Sender: addr1,
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:             "some_denom",
						TradableAmount:         "10",
						RetiredAmount:          "10",
						RetirementJurisdiction: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.BatchDenom": {
			src: MsgSend{
				Sender:    addr1,
				Recipient: addr2,
				Credits: []*MsgSend_SendCredits{
					{
						TradableAmount:         "10",
						RetiredAmount:          "10",
						RetirementJurisdiction: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.TradableAmount set": {
			src: MsgSend{
				Sender:    addr1,
				Recipient: addr2,
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:             "some_denom",
						RetiredAmount:          "10",
						RetirementJurisdiction: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.RetiredAmount set": {
			src: MsgSend{
				Sender:    addr1,
				Recipient: addr2,
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:             "some_denom",
						TradableAmount:         "10",
						RetirementJurisdiction: "ST-UVW XY Z12",
					},
				},
			},
			expErr: true,
		},
		"invalid msg without Credits.RetirementJurisdiction": {
			src: MsgSend{
				Sender:    addr1,
				Recipient: addr2,
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:     "some_denom",
						TradableAmount: "10",
						RetiredAmount:  "10",
					},
				},
			},
			expErr: true,
		},
		"valid msg without Credits.RetirementJurisdiction(When RetiredAmount is zero)": {
			src: MsgSend{
				Sender:    addr1,
				Recipient: addr2,
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:     batchDenom,
						TradableAmount: "10",
						RetiredAmount:  "0",
					},
				},
			},
			expErr: false,
		},
		"invalid msg with wrong sender": {
			src: MsgSend{
				Sender:    "wrongSender",
				Recipient: addr2,
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:     "some_denom",
						TradableAmount: "10",
						RetiredAmount:  "10",
					},
				},
			},
			expErr: true,
		},
		"invalid msg with wrong recipient": {
			src: MsgSend{
				Sender:    addr1,
				Recipient: "wrongRecipient",
				Credits: []*MsgSend_SendCredits{
					{
						BatchDenom:     "some_denom",
						TradableAmount: "10",
						RetiredAmount:  "10",
					},
				},
			},
			expErr: true,
		},
	}

	for msg, test := range tests {
		t.Run(msg, func(t *testing.T) {
			t.Parallel()

			err := test.src.ValidateBasic()
			if test.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
