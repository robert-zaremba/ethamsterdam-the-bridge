package basket_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gotest.tools/v3/assert"

	basketv1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/server/basket/mocks"
)

type suite struct {
	t               *testing.T
	db              ormdb.ModuleDB
	stateStore      basketv1.StateStore
	ctx             context.Context
	k               basket.Keeper
	ctrl            *gomock.Controller
	acct            sdk.AccAddress
	bankKeeper      *mocks.MockBankKeeper
	ecocreditKeeper *mocks.MockEcocreditKeeper
	fooBasketId     uint64
	barBasketId     uint64
	storeKey        *sdk.KVStoreKey
	sdkCtx          sdk.Context
}

func setup(t *testing.T) *suite {
	// prepare database
	s := &suite{t: t}
	var err error
	s.db, err = ormdb.NewModuleDB(server.ModuleSchema, ormdb.ModuleDBOptions{})
	assert.NilError(t, err)
	s.stateStore, err = basketv1.NewStateStore(s.db)
	assert.NilError(t, err)

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	s.storeKey = sdk.NewKVStoreKey("test")
	cms.MountStoreWithDB(s.storeKey, sdk.StoreTypeIAVL, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	s.sdkCtx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)
	s.ctx = sdk.WrapSDKContext(s.sdkCtx)

	// add some data
	s.fooBasketId, err = s.stateStore.BasketStore().InsertReturningID(s.ctx, &basketv1.Basket{
		BasketDenom:       "foo",
		DisableAutoRetire: false,
		CreditTypeName:    "C",
		Exponent:          6,
	})
	assert.NilError(t, err)

	assert.NilError(t, s.stateStore.BasketBalanceStore().Insert(s.ctx, &basketv1.BasketBalance{
		BasketId:       s.fooBasketId,
		BatchDenom:     "C1",
		Balance:        "3.0",
		BatchStartDate: timestamppb.New(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))
	s.setTradableSupply("C1", "3.0")

	assert.NilError(t, s.stateStore.BasketBalanceStore().Insert(s.ctx, &basketv1.BasketBalance{
		BasketId:       s.fooBasketId,
		BatchDenom:     "C2",
		Balance:        "5.0",
		BatchStartDate: timestamppb.New(time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))
	s.setTradableSupply("C2", "5.0")

	s.barBasketId, err = s.stateStore.BasketStore().InsertReturningID(s.ctx, &basketv1.Basket{
		BasketDenom:       "bar",
		DisableAutoRetire: true,
		CreditTypeName:    "C",
		Exponent:          6,
	})
	assert.NilError(t, err)

	assert.NilError(t, s.stateStore.BasketBalanceStore().Insert(s.ctx, &basketv1.BasketBalance{
		BasketId:       s.barBasketId,
		BatchDenom:     "C3",
		Balance:        "7.0",
		BatchStartDate: timestamppb.New(time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))
	s.setTradableSupply("C3", "7.0")

	assert.NilError(t, s.stateStore.BasketBalanceStore().Insert(s.ctx, &basketv1.BasketBalance{
		BasketId:       s.barBasketId,
		BatchDenom:     "C4",
		Balance:        "4.0",
		BatchStartDate: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
	}))
	s.setTradableSupply("C4", "4.0")

	// setup test keeper
	s.ctrl = gomock.NewController(t)
	assert.NilError(t, err)
	s.bankKeeper = mocks.NewMockBankKeeper(s.ctrl)
	s.ecocreditKeeper = mocks.NewMockEcocreditKeeper(s.ctrl)
	s.k = basket.NewKeeper(s.db, s.ecocreditKeeper, s.bankKeeper, s.storeKey)

	s.acct = sdk.AccAddress{0, 1, 2, 3, 4, 5}

	return s
}

func TestTakeMustRetire(t *testing.T) {
	t.Parallel()
	s := setup(t)

	// foo requires RetireOnTake
	_, err := s.k.Take(s.ctx, &baskettypes.MsgTake{
		Owner:              s.acct.String(),
		BasketDenom:        "foo",
		Amount:             "6.0",
		RetirementLocation: "",
		RetireOnTake:       false,
	})
	assert.ErrorIs(t, err, basket.ErrCantDisableRetire)
}

func TestTakeRetire(t *testing.T) {
	t.Parallel()
	s := setup(t)

	fooCoins := sdk.NewCoins(sdk.NewCoin("foo", sdk.NewInt(6000000)))
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), s.acct, baskettypes.BasketSubModuleName, fooCoins)
	s.bankKeeper.EXPECT().BurnCoins(gomock.Any(), baskettypes.BasketSubModuleName, fooCoins)

	res, err := s.k.Take(s.ctx, &baskettypes.MsgTake{
		Owner:              s.acct.String(),
		BasketDenom:        "foo",
		Amount:             "6000000",
		RetirementLocation: "US",
		RetireOnTake:       true,
	})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Credits))
	assert.Equal(t, "C2", res.Credits[0].BatchDenom)
	assertDecStringEqual(t, "5.0", res.Credits[0].Amount)
	assert.Equal(t, "C1", res.Credits[1].BatchDenom)
	assertDecStringEqual(t, "1.0", res.Credits[1].Amount)
	found, err := s.stateStore.BasketBalanceStore().Has(s.ctx, s.fooBasketId, "C2")
	assert.NilError(t, err)
	assert.Assert(t, !found)
	balance, err := s.stateStore.BasketBalanceStore().Get(s.ctx, s.fooBasketId, "C1")
	assert.NilError(t, err)
	assertDecStringEqual(t, "2.0", balance.Balance)

	s.expectTradableBalance("C1", "0")
	s.expectTradableBalance("C2", "0")
	s.expectRetiredBalance("C1", "1")
	s.expectRetiredBalance("C2", "5")
	s.expectTradableSupply("C1", "2")
	s.expectTradableSupply("C2", "0")
	s.expectRetiredSupply("C1", "1")
	s.expectRetiredSupply("C2", "5")
}

func TestTakeTradable(t *testing.T) {
	t.Parallel()
	s := setup(t)

	barCoins := sdk.NewCoins(sdk.NewCoin("bar", sdk.NewInt(10000000)))
	s.bankKeeper.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), s.acct, baskettypes.BasketSubModuleName, barCoins)
	s.bankKeeper.EXPECT().BurnCoins(gomock.Any(), baskettypes.BasketSubModuleName, barCoins)

	res, err := s.k.Take(s.ctx, &baskettypes.MsgTake{
		Owner:        s.acct.String(),
		BasketDenom:  "bar",
		Amount:       "10000000",
		RetireOnTake: false,
	})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Credits))
	assert.Equal(t, "C3", res.Credits[0].BatchDenom)
	assertDecStringEqual(t, "7.0", res.Credits[0].Amount)
	assert.Equal(t, "C4", res.Credits[1].BatchDenom)
	assertDecStringEqual(t, "3.0", res.Credits[1].Amount)
	found, err := s.stateStore.BasketBalanceStore().Has(s.ctx, s.barBasketId, "C3")
	assert.NilError(t, err)
	assert.Assert(t, !found)
	balance, err := s.stateStore.BasketBalanceStore().Get(s.ctx, s.barBasketId, "C4")
	assert.NilError(t, err)
	assertDecStringEqual(t, "1.0", balance.Balance)

	s.expectTradableBalance("C3", "7")
	s.expectTradableBalance("C4", "3")
	s.expectRetiredBalance("C3", "0")
	s.expectRetiredBalance("C4", "0")
	s.expectTradableSupply("C3", "7")
	s.expectTradableSupply("C4", "4")
	s.expectRetiredSupply("C3", "0")
	s.expectRetiredSupply("C4", "0")
}

func assertDecStringEqual(t *testing.T, expected, actual string) {
	dx, err := math.NewDecFromString(expected)
	assert.NilError(t, err)
	dy, err := math.NewDecFromString(actual)
	assert.NilError(t, err)
	assert.Assert(t, 0 == dx.Cmp(dy), fmt.Sprintf("%s != %s", expected, actual))
}

func (s suite) expectTradableBalance(batchDenom string, expected string) {
	kvStore := s.sdkCtx.KVStore(s.storeKey)
	bal, err := ecocredit.GetDecimal(kvStore, ecocredit.TradableBalanceKey(s.acct, ecocredit.BatchDenomT(batchDenom)))
	assert.NilError(s.t, err)
	s.expectDec(expected, bal)
}

func (s suite) expectRetiredBalance(batchDenom string, expected string) {
	kvStore := s.sdkCtx.KVStore(s.storeKey)
	bal, err := ecocredit.GetDecimal(kvStore, ecocredit.RetiredBalanceKey(s.acct, ecocredit.BatchDenomT(batchDenom)))
	assert.NilError(s.t, err)
	s.expectDec(expected, bal)
}

func (s suite) expectTradableSupply(batchDenom string, expected string) {
	kvStore := s.sdkCtx.KVStore(s.storeKey)
	bal, err := ecocredit.GetDecimal(kvStore, ecocredit.TradableSupplyKey(ecocredit.BatchDenomT(batchDenom)))
	assert.NilError(s.t, err)
	s.expectDec(expected, bal)
}

func (s suite) expectRetiredSupply(batchDenom string, expected string) {
	kvStore := s.sdkCtx.KVStore(s.storeKey)
	bal, err := ecocredit.GetDecimal(kvStore, ecocredit.RetiredSupplyKey(ecocredit.BatchDenomT(batchDenom)))
	assert.NilError(s.t, err)
	s.expectDec(expected, bal)
}

func (s suite) setTradableSupply(batchDenom string, amount string) {
	kvStore := s.sdkCtx.KVStore(s.storeKey)
	dec, err := math.NewDecFromString(amount)
	assert.NilError(s.t, err)
	ecocredit.SetDecimal(kvStore, ecocredit.TradableSupplyKey(ecocredit.BatchDenomT(batchDenom)), dec)
}

func (s suite) expectDec(expected string, actual math.Dec) {
	dec, err := math.NewDecFromString(expected)
	assert.NilError(s.t, err)
	assert.Assert(s.t, actual.Cmp(dec) == 0)
}
