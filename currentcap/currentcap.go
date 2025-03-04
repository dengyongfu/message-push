// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package currentcap

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DelayRedeemRouterDelayedRedeem is an auto generated low-level Go binding around an user-defined struct.
type DelayRedeemRouterDelayedRedeem struct {
	Amount    *big.Int
	CreatedAt uint32
	Token     common.Address
}

// DelayRedeemRouterUserDelayedRedeems is an auto generated low-level Go binding around an user-defined struct.
type DelayRedeemRouterUserDelayedRedeems struct {
	DelayedRedeemsCompleted *big.Int
	DelayedRedeems          []DelayRedeemRouterDelayedRedeem
}

// CurrentcapMetaData contains all meta data concerning the Currentcap contract.
var CurrentcapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"BlacklistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"BlacklistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"BtclistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"BtclistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemFee\",\"type\":\"uint256\"}],\"name\":\"DelayedRedeemCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedAmount\",\"type\":\"uint256\"}],\"name\":\"DelayedRedeemsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delayedRedeemsCompleted\",\"type\":\"uint256\"}],\"name\":\"DelayedRedeemsCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimedAmount\",\"type\":\"uint256\"}],\"name\":\"DelayedRedeemsPrincipalClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delayedRedeemsCompleted\",\"type\":\"uint256\"}],\"name\":\"DelayedRedeemsPrincipalCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ManagementFeeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousQuota\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuota\",\"type\":\"uint256\"}],\"name\":\"MaxQuotaSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousQuota\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuota\",\"type\":\"uint256\"}],\"name\":\"RateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"RedeemDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"RedeemFeeRateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"RedeemPrincipalDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokensPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"TokensUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"WhitelistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"WhitelistEnabledSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"WhitelistRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_REDEEM_FEE_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXCHANGE_RATE_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DAILY_REDEEM_CAP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REDEEM_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_BTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REDEEM_FEE_RATE_RANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_IN_A_DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"addToBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"addToBtclist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"canClaimDelayedRedeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"canClaimDelayedRedeemPrincipal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimDelayedRedeems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxNumberOfDelayedRedeemsToClaim\",\"type\":\"uint256\"}],\"name\":\"claimDelayedRedeems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxNumberOfDelayedRedeemsToClaim\",\"type\":\"uint256\"}],\"name\":\"claimPrincipals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPrincipals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"createDelayedRedeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getClaimableUserDelayedRedeems\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"createdAt\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structDelayRedeemRouter.DelayedRedeem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getCurrentCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserDelayedRedeems\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"createdAt\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structDelayRedeemRouter.DelayedRedeem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_uniBTC\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_redeemDelay\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_whitelistEnabled\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isBtclisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastRebaseTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managementFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxQuotas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"pauseTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"quotaBases\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"quotaRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemPrincipalDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"removeFromBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"removeFromBtclist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_quotas\",\"type\":\"uint256[]\"}],\"name\":\"setMaxQuotaForTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_quotas\",\"type\":\"uint256[]\"}],\"name\":\"setQuotaRates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDelay\",\"type\":\"uint256\"}],\"name\":\"setRedeemDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setRedeemFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDelay\",\"type\":\"uint256\"}],\"name\":\"setRedeemPrincipalDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setWhitelistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenDebts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalDebts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCleared\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniBTC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"unpauseTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"userDelayedRedeemByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"createdAt\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structDelayRedeemRouter.DelayedRedeem\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userRedeems\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayedRedeemsCompleted\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint224\",\"name\":\"amount\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"createdAt\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structDelayRedeemRouter.DelayedRedeem[]\",\"name\":\"delayedRedeems\",\"type\":\"tuple[]\"}],\"internalType\":\"structDelayRedeemRouter.UserDelayedRedeems\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userRedeemsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"withdrawManagementFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// CurrentcapABI is the input ABI used to generate the binding from.
// Deprecated: Use CurrentcapMetaData.ABI instead.
var CurrentcapABI = CurrentcapMetaData.ABI

// Currentcap is an auto generated Go binding around an Ethereum contract.
type Currentcap struct {
	CurrentcapCaller     // Read-only binding to the contract
	CurrentcapTransactor // Write-only binding to the contract
	CurrentcapFilterer   // Log filterer for contract events
}

// CurrentcapCaller is an auto generated read-only Go binding around an Ethereum contract.
type CurrentcapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurrentcapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CurrentcapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurrentcapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CurrentcapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CurrentcapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CurrentcapSession struct {
	Contract     *Currentcap       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CurrentcapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CurrentcapCallerSession struct {
	Contract *CurrentcapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CurrentcapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CurrentcapTransactorSession struct {
	Contract     *CurrentcapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CurrentcapRaw is an auto generated low-level Go binding around an Ethereum contract.
type CurrentcapRaw struct {
	Contract *Currentcap // Generic contract binding to access the raw methods on
}

// CurrentcapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CurrentcapCallerRaw struct {
	Contract *CurrentcapCaller // Generic read-only contract binding to access the raw methods on
}

// CurrentcapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CurrentcapTransactorRaw struct {
	Contract *CurrentcapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCurrentcap creates a new instance of Currentcap, bound to a specific deployed contract.
func NewCurrentcap(address common.Address, backend bind.ContractBackend) (*Currentcap, error) {
	contract, err := bindCurrentcap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Currentcap{CurrentcapCaller: CurrentcapCaller{contract: contract}, CurrentcapTransactor: CurrentcapTransactor{contract: contract}, CurrentcapFilterer: CurrentcapFilterer{contract: contract}}, nil
}

// NewCurrentcapCaller creates a new read-only instance of Currentcap, bound to a specific deployed contract.
func NewCurrentcapCaller(address common.Address, caller bind.ContractCaller) (*CurrentcapCaller, error) {
	contract, err := bindCurrentcap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CurrentcapCaller{contract: contract}, nil
}

// NewCurrentcapTransactor creates a new write-only instance of Currentcap, bound to a specific deployed contract.
func NewCurrentcapTransactor(address common.Address, transactor bind.ContractTransactor) (*CurrentcapTransactor, error) {
	contract, err := bindCurrentcap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CurrentcapTransactor{contract: contract}, nil
}

// NewCurrentcapFilterer creates a new log filterer instance of Currentcap, bound to a specific deployed contract.
func NewCurrentcapFilterer(address common.Address, filterer bind.ContractFilterer) (*CurrentcapFilterer, error) {
	contract, err := bindCurrentcap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CurrentcapFilterer{contract: contract}, nil
}

// bindCurrentcap binds a generic wrapper to an already deployed contract.
func bindCurrentcap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CurrentcapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Currentcap *CurrentcapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Currentcap.Contract.CurrentcapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Currentcap *CurrentcapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Currentcap.Contract.CurrentcapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Currentcap *CurrentcapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Currentcap.Contract.CurrentcapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Currentcap *CurrentcapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Currentcap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Currentcap *CurrentcapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Currentcap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Currentcap *CurrentcapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Currentcap.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Currentcap *CurrentcapCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Currentcap *CurrentcapSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Currentcap.Contract.DEFAULTADMINROLE(&_Currentcap.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Currentcap *CurrentcapCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Currentcap.Contract.DEFAULTADMINROLE(&_Currentcap.CallOpts)
}

// DEFAULTREDEEMFEERATE is a free data retrieval call binding the contract method 0xc71aba6c.
//
// Solidity: function DEFAULT_REDEEM_FEE_RATE() view returns(uint256)
func (_Currentcap *CurrentcapCaller) DEFAULTREDEEMFEERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "DEFAULT_REDEEM_FEE_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTREDEEMFEERATE is a free data retrieval call binding the contract method 0xc71aba6c.
//
// Solidity: function DEFAULT_REDEEM_FEE_RATE() view returns(uint256)
func (_Currentcap *CurrentcapSession) DEFAULTREDEEMFEERATE() (*big.Int, error) {
	return _Currentcap.Contract.DEFAULTREDEEMFEERATE(&_Currentcap.CallOpts)
}

// DEFAULTREDEEMFEERATE is a free data retrieval call binding the contract method 0xc71aba6c.
//
// Solidity: function DEFAULT_REDEEM_FEE_RATE() view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) DEFAULTREDEEMFEERATE() (*big.Int, error) {
	return _Currentcap.Contract.DEFAULTREDEEMFEERATE(&_Currentcap.CallOpts)
}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_Currentcap *CurrentcapCaller) EXCHANGERATEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "EXCHANGE_RATE_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_Currentcap *CurrentcapSession) EXCHANGERATEBASE() (*big.Int, error) {
	return _Currentcap.Contract.EXCHANGERATEBASE(&_Currentcap.CallOpts)
}

// EXCHANGERATEBASE is a free data retrieval call binding the contract method 0xb7b038da.
//
// Solidity: function EXCHANGE_RATE_BASE() view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) EXCHANGERATEBASE() (*big.Int, error) {
	return _Currentcap.Contract.EXCHANGERATEBASE(&_Currentcap.CallOpts)
}

// MAXDAILYREDEEMCAP is a free data retrieval call binding the contract method 0xb0d381c7.
//
// Solidity: function MAX_DAILY_REDEEM_CAP() view returns(uint256)
func (_Currentcap *CurrentcapCaller) MAXDAILYREDEEMCAP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "MAX_DAILY_REDEEM_CAP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDAILYREDEEMCAP is a free data retrieval call binding the contract method 0xb0d381c7.
//
// Solidity: function MAX_DAILY_REDEEM_CAP() view returns(uint256)
func (_Currentcap *CurrentcapSession) MAXDAILYREDEEMCAP() (*big.Int, error) {
	return _Currentcap.Contract.MAXDAILYREDEEMCAP(&_Currentcap.CallOpts)
}

// MAXDAILYREDEEMCAP is a free data retrieval call binding the contract method 0xb0d381c7.
//
// Solidity: function MAX_DAILY_REDEEM_CAP() view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) MAXDAILYREDEEMCAP() (*big.Int, error) {
	return _Currentcap.Contract.MAXDAILYREDEEMCAP(&_Currentcap.CallOpts)
}

// MAXREDEEMDELAY is a free data retrieval call binding the contract method 0xc5a56664.
//
// Solidity: function MAX_REDEEM_DELAY() view returns(uint256)
func (_Currentcap *CurrentcapCaller) MAXREDEEMDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "MAX_REDEEM_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXREDEEMDELAY is a free data retrieval call binding the contract method 0xc5a56664.
//
// Solidity: function MAX_REDEEM_DELAY() view returns(uint256)
func (_Currentcap *CurrentcapSession) MAXREDEEMDELAY() (*big.Int, error) {
	return _Currentcap.Contract.MAXREDEEMDELAY(&_Currentcap.CallOpts)
}

// MAXREDEEMDELAY is a free data retrieval call binding the contract method 0xc5a56664.
//
// Solidity: function MAX_REDEEM_DELAY() view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) MAXREDEEMDELAY() (*big.Int, error) {
	return _Currentcap.Contract.MAXREDEEMDELAY(&_Currentcap.CallOpts)
}

// NATIVEBTC is a free data retrieval call binding the contract method 0x3af02ff3.
//
// Solidity: function NATIVE_BTC() view returns(address)
func (_Currentcap *CurrentcapCaller) NATIVEBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "NATIVE_BTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEBTC is a free data retrieval call binding the contract method 0x3af02ff3.
//
// Solidity: function NATIVE_BTC() view returns(address)
func (_Currentcap *CurrentcapSession) NATIVEBTC() (common.Address, error) {
	return _Currentcap.Contract.NATIVEBTC(&_Currentcap.CallOpts)
}

// NATIVEBTC is a free data retrieval call binding the contract method 0x3af02ff3.
//
// Solidity: function NATIVE_BTC() view returns(address)
func (_Currentcap *CurrentcapCallerSession) NATIVEBTC() (common.Address, error) {
	return _Currentcap.Contract.NATIVEBTC(&_Currentcap.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Currentcap *CurrentcapCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Currentcap *CurrentcapSession) PAUSERROLE() ([32]byte, error) {
	return _Currentcap.Contract.PAUSERROLE(&_Currentcap.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Currentcap *CurrentcapCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Currentcap.Contract.PAUSERROLE(&_Currentcap.CallOpts)
}

// REDEEMFEERATERANGE is a free data retrieval call binding the contract method 0x58f7e664.
//
// Solidity: function REDEEM_FEE_RATE_RANGE() view returns(uint256)
func (_Currentcap *CurrentcapCaller) REDEEMFEERATERANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "REDEEM_FEE_RATE_RANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REDEEMFEERATERANGE is a free data retrieval call binding the contract method 0x58f7e664.
//
// Solidity: function REDEEM_FEE_RATE_RANGE() view returns(uint256)
func (_Currentcap *CurrentcapSession) REDEEMFEERATERANGE() (*big.Int, error) {
	return _Currentcap.Contract.REDEEMFEERATERANGE(&_Currentcap.CallOpts)
}

// REDEEMFEERATERANGE is a free data retrieval call binding the contract method 0x58f7e664.
//
// Solidity: function REDEEM_FEE_RATE_RANGE() view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) REDEEMFEERATERANGE() (*big.Int, error) {
	return _Currentcap.Contract.REDEEMFEERATERANGE(&_Currentcap.CallOpts)
}

// SECONDSINADAY is a free data retrieval call binding the contract method 0xf9cfa06f.
//
// Solidity: function SECONDS_IN_A_DAY() view returns(uint256)
func (_Currentcap *CurrentcapCaller) SECONDSINADAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "SECONDS_IN_A_DAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSINADAY is a free data retrieval call binding the contract method 0xf9cfa06f.
//
// Solidity: function SECONDS_IN_A_DAY() view returns(uint256)
func (_Currentcap *CurrentcapSession) SECONDSINADAY() (*big.Int, error) {
	return _Currentcap.Contract.SECONDSINADAY(&_Currentcap.CallOpts)
}

// SECONDSINADAY is a free data retrieval call binding the contract method 0xf9cfa06f.
//
// Solidity: function SECONDS_IN_A_DAY() view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) SECONDSINADAY() (*big.Int, error) {
	return _Currentcap.Contract.SECONDSINADAY(&_Currentcap.CallOpts)
}

// CanClaimDelayedRedeem is a free data retrieval call binding the contract method 0x8b745eae.
//
// Solidity: function canClaimDelayedRedeem(address user, uint256 index) view returns(bool)
func (_Currentcap *CurrentcapCaller) CanClaimDelayedRedeem(opts *bind.CallOpts, user common.Address, index *big.Int) (bool, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "canClaimDelayedRedeem", user, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaimDelayedRedeem is a free data retrieval call binding the contract method 0x8b745eae.
//
// Solidity: function canClaimDelayedRedeem(address user, uint256 index) view returns(bool)
func (_Currentcap *CurrentcapSession) CanClaimDelayedRedeem(user common.Address, index *big.Int) (bool, error) {
	return _Currentcap.Contract.CanClaimDelayedRedeem(&_Currentcap.CallOpts, user, index)
}

// CanClaimDelayedRedeem is a free data retrieval call binding the contract method 0x8b745eae.
//
// Solidity: function canClaimDelayedRedeem(address user, uint256 index) view returns(bool)
func (_Currentcap *CurrentcapCallerSession) CanClaimDelayedRedeem(user common.Address, index *big.Int) (bool, error) {
	return _Currentcap.Contract.CanClaimDelayedRedeem(&_Currentcap.CallOpts, user, index)
}

// CanClaimDelayedRedeemPrincipal is a free data retrieval call binding the contract method 0x7ac221fd.
//
// Solidity: function canClaimDelayedRedeemPrincipal(address user, uint256 index) view returns(bool)
func (_Currentcap *CurrentcapCaller) CanClaimDelayedRedeemPrincipal(opts *bind.CallOpts, user common.Address, index *big.Int) (bool, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "canClaimDelayedRedeemPrincipal", user, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaimDelayedRedeemPrincipal is a free data retrieval call binding the contract method 0x7ac221fd.
//
// Solidity: function canClaimDelayedRedeemPrincipal(address user, uint256 index) view returns(bool)
func (_Currentcap *CurrentcapSession) CanClaimDelayedRedeemPrincipal(user common.Address, index *big.Int) (bool, error) {
	return _Currentcap.Contract.CanClaimDelayedRedeemPrincipal(&_Currentcap.CallOpts, user, index)
}

// CanClaimDelayedRedeemPrincipal is a free data retrieval call binding the contract method 0x7ac221fd.
//
// Solidity: function canClaimDelayedRedeemPrincipal(address user, uint256 index) view returns(bool)
func (_Currentcap *CurrentcapCallerSession) CanClaimDelayedRedeemPrincipal(user common.Address, index *big.Int) (bool, error) {
	return _Currentcap.Contract.CanClaimDelayedRedeemPrincipal(&_Currentcap.CallOpts, user, index)
}

// GetClaimableUserDelayedRedeems is a free data retrieval call binding the contract method 0x4929b254.
//
// Solidity: function getClaimableUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_Currentcap *CurrentcapCaller) GetClaimableUserDelayedRedeems(opts *bind.CallOpts, user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "getClaimableUserDelayedRedeems", user)

	if err != nil {
		return *new([]DelayRedeemRouterDelayedRedeem), err
	}

	out0 := *abi.ConvertType(out[0], new([]DelayRedeemRouterDelayedRedeem)).(*[]DelayRedeemRouterDelayedRedeem)

	return out0, err

}

// GetClaimableUserDelayedRedeems is a free data retrieval call binding the contract method 0x4929b254.
//
// Solidity: function getClaimableUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_Currentcap *CurrentcapSession) GetClaimableUserDelayedRedeems(user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	return _Currentcap.Contract.GetClaimableUserDelayedRedeems(&_Currentcap.CallOpts, user)
}

// GetClaimableUserDelayedRedeems is a free data retrieval call binding the contract method 0x4929b254.
//
// Solidity: function getClaimableUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_Currentcap *CurrentcapCallerSession) GetClaimableUserDelayedRedeems(user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	return _Currentcap.Contract.GetClaimableUserDelayedRedeems(&_Currentcap.CallOpts, user)
}

// GetCurrentCap is a free data retrieval call binding the contract method 0x2c612832.
//
// Solidity: function getCurrentCap(address token) view returns(uint256)
func (_Currentcap *CurrentcapCaller) GetCurrentCap(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "getCurrentCap", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentCap is a free data retrieval call binding the contract method 0x2c612832.
//
// Solidity: function getCurrentCap(address token) view returns(uint256)
func (_Currentcap *CurrentcapSession) GetCurrentCap(token common.Address) (*big.Int, error) {
	return _Currentcap.Contract.GetCurrentCap(&_Currentcap.CallOpts, token)
}

// GetCurrentCap is a free data retrieval call binding the contract method 0x2c612832.
//
// Solidity: function getCurrentCap(address token) view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) GetCurrentCap(token common.Address) (*big.Int, error) {
	return _Currentcap.Contract.GetCurrentCap(&_Currentcap.CallOpts, token)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Currentcap *CurrentcapCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Currentcap *CurrentcapSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Currentcap.Contract.GetRoleAdmin(&_Currentcap.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Currentcap *CurrentcapCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Currentcap.Contract.GetRoleAdmin(&_Currentcap.CallOpts, role)
}

// GetUserDelayedRedeems is a free data retrieval call binding the contract method 0x99a49065.
//
// Solidity: function getUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_Currentcap *CurrentcapCaller) GetUserDelayedRedeems(opts *bind.CallOpts, user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "getUserDelayedRedeems", user)

	if err != nil {
		return *new([]DelayRedeemRouterDelayedRedeem), err
	}

	out0 := *abi.ConvertType(out[0], new([]DelayRedeemRouterDelayedRedeem)).(*[]DelayRedeemRouterDelayedRedeem)

	return out0, err

}

// GetUserDelayedRedeems is a free data retrieval call binding the contract method 0x99a49065.
//
// Solidity: function getUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_Currentcap *CurrentcapSession) GetUserDelayedRedeems(user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	return _Currentcap.Contract.GetUserDelayedRedeems(&_Currentcap.CallOpts, user)
}

// GetUserDelayedRedeems is a free data retrieval call binding the contract method 0x99a49065.
//
// Solidity: function getUserDelayedRedeems(address user) view returns((uint224,uint32,address)[])
func (_Currentcap *CurrentcapCallerSession) GetUserDelayedRedeems(user common.Address) ([]DelayRedeemRouterDelayedRedeem, error) {
	return _Currentcap.Contract.GetUserDelayedRedeems(&_Currentcap.CallOpts, user)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Currentcap *CurrentcapCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Currentcap *CurrentcapSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Currentcap.Contract.HasRole(&_Currentcap.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Currentcap *CurrentcapCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Currentcap.Contract.HasRole(&_Currentcap.CallOpts, role, account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_Currentcap *CurrentcapCaller) IsBlacklisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "isBlacklisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_Currentcap *CurrentcapSession) IsBlacklisted(account common.Address) (bool, error) {
	return _Currentcap.Contract.IsBlacklisted(&_Currentcap.CallOpts, account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account) view returns(bool)
func (_Currentcap *CurrentcapCallerSession) IsBlacklisted(account common.Address) (bool, error) {
	return _Currentcap.Contract.IsBlacklisted(&_Currentcap.CallOpts, account)
}

// IsBtclisted is a free data retrieval call binding the contract method 0x362aada3.
//
// Solidity: function isBtclisted(address token) view returns(bool)
func (_Currentcap *CurrentcapCaller) IsBtclisted(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "isBtclisted", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBtclisted is a free data retrieval call binding the contract method 0x362aada3.
//
// Solidity: function isBtclisted(address token) view returns(bool)
func (_Currentcap *CurrentcapSession) IsBtclisted(token common.Address) (bool, error) {
	return _Currentcap.Contract.IsBtclisted(&_Currentcap.CallOpts, token)
}

// IsBtclisted is a free data retrieval call binding the contract method 0x362aada3.
//
// Solidity: function isBtclisted(address token) view returns(bool)
func (_Currentcap *CurrentcapCallerSession) IsBtclisted(token common.Address) (bool, error) {
	return _Currentcap.Contract.IsBtclisted(&_Currentcap.CallOpts, token)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Currentcap *CurrentcapCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "isWhitelisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Currentcap *CurrentcapSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Currentcap.Contract.IsWhitelisted(&_Currentcap.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Currentcap *CurrentcapCallerSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Currentcap.Contract.IsWhitelisted(&_Currentcap.CallOpts, account)
}

// LastRebaseTimestamps is a free data retrieval call binding the contract method 0x7364a714.
//
// Solidity: function lastRebaseTimestamps(address ) view returns(uint256)
func (_Currentcap *CurrentcapCaller) LastRebaseTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "lastRebaseTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRebaseTimestamps is a free data retrieval call binding the contract method 0x7364a714.
//
// Solidity: function lastRebaseTimestamps(address ) view returns(uint256)
func (_Currentcap *CurrentcapSession) LastRebaseTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Currentcap.Contract.LastRebaseTimestamps(&_Currentcap.CallOpts, arg0)
}

// LastRebaseTimestamps is a free data retrieval call binding the contract method 0x7364a714.
//
// Solidity: function lastRebaseTimestamps(address ) view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) LastRebaseTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Currentcap.Contract.LastRebaseTimestamps(&_Currentcap.CallOpts, arg0)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Currentcap *CurrentcapCaller) ManagementFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "managementFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Currentcap *CurrentcapSession) ManagementFee() (*big.Int, error) {
	return _Currentcap.Contract.ManagementFee(&_Currentcap.CallOpts)
}

// ManagementFee is a free data retrieval call binding the contract method 0xa6f7f5d6.
//
// Solidity: function managementFee() view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) ManagementFee() (*big.Int, error) {
	return _Currentcap.Contract.ManagementFee(&_Currentcap.CallOpts)
}

// MaxQuotas is a free data retrieval call binding the contract method 0x4fa581f9.
//
// Solidity: function maxQuotas(address ) view returns(uint256)
func (_Currentcap *CurrentcapCaller) MaxQuotas(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "maxQuotas", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxQuotas is a free data retrieval call binding the contract method 0x4fa581f9.
//
// Solidity: function maxQuotas(address ) view returns(uint256)
func (_Currentcap *CurrentcapSession) MaxQuotas(arg0 common.Address) (*big.Int, error) {
	return _Currentcap.Contract.MaxQuotas(&_Currentcap.CallOpts, arg0)
}

// MaxQuotas is a free data retrieval call binding the contract method 0x4fa581f9.
//
// Solidity: function maxQuotas(address ) view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) MaxQuotas(arg0 common.Address) (*big.Int, error) {
	return _Currentcap.Contract.MaxQuotas(&_Currentcap.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Currentcap *CurrentcapCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Currentcap *CurrentcapSession) Paused() (bool, error) {
	return _Currentcap.Contract.Paused(&_Currentcap.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Currentcap *CurrentcapCallerSession) Paused() (bool, error) {
	return _Currentcap.Contract.Paused(&_Currentcap.CallOpts)
}

// QuotaBases is a free data retrieval call binding the contract method 0x7bd8ba9d.
//
// Solidity: function quotaBases(address ) view returns(uint256)
func (_Currentcap *CurrentcapCaller) QuotaBases(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "quotaBases", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuotaBases is a free data retrieval call binding the contract method 0x7bd8ba9d.
//
// Solidity: function quotaBases(address ) view returns(uint256)
func (_Currentcap *CurrentcapSession) QuotaBases(arg0 common.Address) (*big.Int, error) {
	return _Currentcap.Contract.QuotaBases(&_Currentcap.CallOpts, arg0)
}

// QuotaBases is a free data retrieval call binding the contract method 0x7bd8ba9d.
//
// Solidity: function quotaBases(address ) view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) QuotaBases(arg0 common.Address) (*big.Int, error) {
	return _Currentcap.Contract.QuotaBases(&_Currentcap.CallOpts, arg0)
}

// QuotaRates is a free data retrieval call binding the contract method 0xc0bcde5d.
//
// Solidity: function quotaRates(address ) view returns(uint256)
func (_Currentcap *CurrentcapCaller) QuotaRates(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "quotaRates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuotaRates is a free data retrieval call binding the contract method 0xc0bcde5d.
//
// Solidity: function quotaRates(address ) view returns(uint256)
func (_Currentcap *CurrentcapSession) QuotaRates(arg0 common.Address) (*big.Int, error) {
	return _Currentcap.Contract.QuotaRates(&_Currentcap.CallOpts, arg0)
}

// QuotaRates is a free data retrieval call binding the contract method 0xc0bcde5d.
//
// Solidity: function quotaRates(address ) view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) QuotaRates(arg0 common.Address) (*big.Int, error) {
	return _Currentcap.Contract.QuotaRates(&_Currentcap.CallOpts, arg0)
}

// RedeemDelay is a free data retrieval call binding the contract method 0xd2adf402.
//
// Solidity: function redeemDelay() view returns(uint256)
func (_Currentcap *CurrentcapCaller) RedeemDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "redeemDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemDelay is a free data retrieval call binding the contract method 0xd2adf402.
//
// Solidity: function redeemDelay() view returns(uint256)
func (_Currentcap *CurrentcapSession) RedeemDelay() (*big.Int, error) {
	return _Currentcap.Contract.RedeemDelay(&_Currentcap.CallOpts)
}

// RedeemDelay is a free data retrieval call binding the contract method 0xd2adf402.
//
// Solidity: function redeemDelay() view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) RedeemDelay() (*big.Int, error) {
	return _Currentcap.Contract.RedeemDelay(&_Currentcap.CallOpts)
}

// RedeemFeeRate is a free data retrieval call binding the contract method 0x5872e6fa.
//
// Solidity: function redeemFeeRate() view returns(uint256)
func (_Currentcap *CurrentcapCaller) RedeemFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "redeemFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemFeeRate is a free data retrieval call binding the contract method 0x5872e6fa.
//
// Solidity: function redeemFeeRate() view returns(uint256)
func (_Currentcap *CurrentcapSession) RedeemFeeRate() (*big.Int, error) {
	return _Currentcap.Contract.RedeemFeeRate(&_Currentcap.CallOpts)
}

// RedeemFeeRate is a free data retrieval call binding the contract method 0x5872e6fa.
//
// Solidity: function redeemFeeRate() view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) RedeemFeeRate() (*big.Int, error) {
	return _Currentcap.Contract.RedeemFeeRate(&_Currentcap.CallOpts)
}

// RedeemPrincipalDelay is a free data retrieval call binding the contract method 0x83c0894b.
//
// Solidity: function redeemPrincipalDelay() view returns(uint256)
func (_Currentcap *CurrentcapCaller) RedeemPrincipalDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "redeemPrincipalDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RedeemPrincipalDelay is a free data retrieval call binding the contract method 0x83c0894b.
//
// Solidity: function redeemPrincipalDelay() view returns(uint256)
func (_Currentcap *CurrentcapSession) RedeemPrincipalDelay() (*big.Int, error) {
	return _Currentcap.Contract.RedeemPrincipalDelay(&_Currentcap.CallOpts)
}

// RedeemPrincipalDelay is a free data retrieval call binding the contract method 0x83c0894b.
//
// Solidity: function redeemPrincipalDelay() view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) RedeemPrincipalDelay() (*big.Int, error) {
	return _Currentcap.Contract.RedeemPrincipalDelay(&_Currentcap.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Currentcap *CurrentcapCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Currentcap *CurrentcapSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Currentcap.Contract.SupportsInterface(&_Currentcap.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Currentcap *CurrentcapCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Currentcap.Contract.SupportsInterface(&_Currentcap.CallOpts, interfaceId)
}

// TokenDebts is a free data retrieval call binding the contract method 0xf190439e.
//
// Solidity: function tokenDebts(address ) view returns(uint256 totalDebts, uint256 totalCleared)
func (_Currentcap *CurrentcapCaller) TokenDebts(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalDebts   *big.Int
	TotalCleared *big.Int
}, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "tokenDebts", arg0)

	outstruct := new(struct {
		TotalDebts   *big.Int
		TotalCleared *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalDebts = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalCleared = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenDebts is a free data retrieval call binding the contract method 0xf190439e.
//
// Solidity: function tokenDebts(address ) view returns(uint256 totalDebts, uint256 totalCleared)
func (_Currentcap *CurrentcapSession) TokenDebts(arg0 common.Address) (struct {
	TotalDebts   *big.Int
	TotalCleared *big.Int
}, error) {
	return _Currentcap.Contract.TokenDebts(&_Currentcap.CallOpts, arg0)
}

// TokenDebts is a free data retrieval call binding the contract method 0xf190439e.
//
// Solidity: function tokenDebts(address ) view returns(uint256 totalDebts, uint256 totalCleared)
func (_Currentcap *CurrentcapCallerSession) TokenDebts(arg0 common.Address) (struct {
	TotalDebts   *big.Int
	TotalCleared *big.Int
}, error) {
	return _Currentcap.Contract.TokenDebts(&_Currentcap.CallOpts, arg0)
}

// UniBTC is a free data retrieval call binding the contract method 0x59f3d39b.
//
// Solidity: function uniBTC() view returns(address)
func (_Currentcap *CurrentcapCaller) UniBTC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "uniBTC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniBTC is a free data retrieval call binding the contract method 0x59f3d39b.
//
// Solidity: function uniBTC() view returns(address)
func (_Currentcap *CurrentcapSession) UniBTC() (common.Address, error) {
	return _Currentcap.Contract.UniBTC(&_Currentcap.CallOpts)
}

// UniBTC is a free data retrieval call binding the contract method 0x59f3d39b.
//
// Solidity: function uniBTC() view returns(address)
func (_Currentcap *CurrentcapCallerSession) UniBTC() (common.Address, error) {
	return _Currentcap.Contract.UniBTC(&_Currentcap.CallOpts)
}

// UserDelayedRedeemByIndex is a free data retrieval call binding the contract method 0xc1541631.
//
// Solidity: function userDelayedRedeemByIndex(address user, uint256 index) view returns((uint224,uint32,address))
func (_Currentcap *CurrentcapCaller) UserDelayedRedeemByIndex(opts *bind.CallOpts, user common.Address, index *big.Int) (DelayRedeemRouterDelayedRedeem, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "userDelayedRedeemByIndex", user, index)

	if err != nil {
		return *new(DelayRedeemRouterDelayedRedeem), err
	}

	out0 := *abi.ConvertType(out[0], new(DelayRedeemRouterDelayedRedeem)).(*DelayRedeemRouterDelayedRedeem)

	return out0, err

}

// UserDelayedRedeemByIndex is a free data retrieval call binding the contract method 0xc1541631.
//
// Solidity: function userDelayedRedeemByIndex(address user, uint256 index) view returns((uint224,uint32,address))
func (_Currentcap *CurrentcapSession) UserDelayedRedeemByIndex(user common.Address, index *big.Int) (DelayRedeemRouterDelayedRedeem, error) {
	return _Currentcap.Contract.UserDelayedRedeemByIndex(&_Currentcap.CallOpts, user, index)
}

// UserDelayedRedeemByIndex is a free data retrieval call binding the contract method 0xc1541631.
//
// Solidity: function userDelayedRedeemByIndex(address user, uint256 index) view returns((uint224,uint32,address))
func (_Currentcap *CurrentcapCallerSession) UserDelayedRedeemByIndex(user common.Address, index *big.Int) (DelayRedeemRouterDelayedRedeem, error) {
	return _Currentcap.Contract.UserDelayedRedeemByIndex(&_Currentcap.CallOpts, user, index)
}

// UserRedeems is a free data retrieval call binding the contract method 0x8d18e24b.
//
// Solidity: function userRedeems(address user) view returns((uint256,(uint224,uint32,address)[]))
func (_Currentcap *CurrentcapCaller) UserRedeems(opts *bind.CallOpts, user common.Address) (DelayRedeemRouterUserDelayedRedeems, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "userRedeems", user)

	if err != nil {
		return *new(DelayRedeemRouterUserDelayedRedeems), err
	}

	out0 := *abi.ConvertType(out[0], new(DelayRedeemRouterUserDelayedRedeems)).(*DelayRedeemRouterUserDelayedRedeems)

	return out0, err

}

// UserRedeems is a free data retrieval call binding the contract method 0x8d18e24b.
//
// Solidity: function userRedeems(address user) view returns((uint256,(uint224,uint32,address)[]))
func (_Currentcap *CurrentcapSession) UserRedeems(user common.Address) (DelayRedeemRouterUserDelayedRedeems, error) {
	return _Currentcap.Contract.UserRedeems(&_Currentcap.CallOpts, user)
}

// UserRedeems is a free data retrieval call binding the contract method 0x8d18e24b.
//
// Solidity: function userRedeems(address user) view returns((uint256,(uint224,uint32,address)[]))
func (_Currentcap *CurrentcapCallerSession) UserRedeems(user common.Address) (DelayRedeemRouterUserDelayedRedeems, error) {
	return _Currentcap.Contract.UserRedeems(&_Currentcap.CallOpts, user)
}

// UserRedeemsLength is a free data retrieval call binding the contract method 0x5449c33c.
//
// Solidity: function userRedeemsLength(address user) view returns(uint256)
func (_Currentcap *CurrentcapCaller) UserRedeemsLength(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "userRedeemsLength", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRedeemsLength is a free data retrieval call binding the contract method 0x5449c33c.
//
// Solidity: function userRedeemsLength(address user) view returns(uint256)
func (_Currentcap *CurrentcapSession) UserRedeemsLength(user common.Address) (*big.Int, error) {
	return _Currentcap.Contract.UserRedeemsLength(&_Currentcap.CallOpts, user)
}

// UserRedeemsLength is a free data retrieval call binding the contract method 0x5449c33c.
//
// Solidity: function userRedeemsLength(address user) view returns(uint256)
func (_Currentcap *CurrentcapCallerSession) UserRedeemsLength(user common.Address) (*big.Int, error) {
	return _Currentcap.Contract.UserRedeemsLength(&_Currentcap.CallOpts, user)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Currentcap *CurrentcapCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Currentcap *CurrentcapSession) Vault() (common.Address, error) {
	return _Currentcap.Contract.Vault(&_Currentcap.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Currentcap *CurrentcapCallerSession) Vault() (common.Address, error) {
	return _Currentcap.Contract.Vault(&_Currentcap.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Currentcap *CurrentcapCaller) WhitelistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Currentcap.contract.Call(opts, &out, "whitelistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Currentcap *CurrentcapSession) WhitelistEnabled() (bool, error) {
	return _Currentcap.Contract.WhitelistEnabled(&_Currentcap.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Currentcap *CurrentcapCallerSession) WhitelistEnabled() (bool, error) {
	return _Currentcap.Contract.WhitelistEnabled(&_Currentcap.CallOpts)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x935eb35f.
//
// Solidity: function addToBlacklist(address[] _accounts) returns()
func (_Currentcap *CurrentcapTransactor) AddToBlacklist(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "addToBlacklist", _accounts)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x935eb35f.
//
// Solidity: function addToBlacklist(address[] _accounts) returns()
func (_Currentcap *CurrentcapSession) AddToBlacklist(_accounts []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.AddToBlacklist(&_Currentcap.TransactOpts, _accounts)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x935eb35f.
//
// Solidity: function addToBlacklist(address[] _accounts) returns()
func (_Currentcap *CurrentcapTransactorSession) AddToBlacklist(_accounts []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.AddToBlacklist(&_Currentcap.TransactOpts, _accounts)
}

// AddToBtclist is a paid mutator transaction binding the contract method 0x0ba15afb.
//
// Solidity: function addToBtclist(address[] _tokens) returns()
func (_Currentcap *CurrentcapTransactor) AddToBtclist(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "addToBtclist", _tokens)
}

// AddToBtclist is a paid mutator transaction binding the contract method 0x0ba15afb.
//
// Solidity: function addToBtclist(address[] _tokens) returns()
func (_Currentcap *CurrentcapSession) AddToBtclist(_tokens []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.AddToBtclist(&_Currentcap.TransactOpts, _tokens)
}

// AddToBtclist is a paid mutator transaction binding the contract method 0x0ba15afb.
//
// Solidity: function addToBtclist(address[] _tokens) returns()
func (_Currentcap *CurrentcapTransactorSession) AddToBtclist(_tokens []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.AddToBtclist(&_Currentcap.TransactOpts, _tokens)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _accounts) returns()
func (_Currentcap *CurrentcapTransactor) AddToWhitelist(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "addToWhitelist", _accounts)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _accounts) returns()
func (_Currentcap *CurrentcapSession) AddToWhitelist(_accounts []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.AddToWhitelist(&_Currentcap.TransactOpts, _accounts)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] _accounts) returns()
func (_Currentcap *CurrentcapTransactorSession) AddToWhitelist(_accounts []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.AddToWhitelist(&_Currentcap.TransactOpts, _accounts)
}

// ClaimDelayedRedeems is a paid mutator transaction binding the contract method 0xf2881130.
//
// Solidity: function claimDelayedRedeems() returns()
func (_Currentcap *CurrentcapTransactor) ClaimDelayedRedeems(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "claimDelayedRedeems")
}

// ClaimDelayedRedeems is a paid mutator transaction binding the contract method 0xf2881130.
//
// Solidity: function claimDelayedRedeems() returns()
func (_Currentcap *CurrentcapSession) ClaimDelayedRedeems() (*types.Transaction, error) {
	return _Currentcap.Contract.ClaimDelayedRedeems(&_Currentcap.TransactOpts)
}

// ClaimDelayedRedeems is a paid mutator transaction binding the contract method 0xf2881130.
//
// Solidity: function claimDelayedRedeems() returns()
func (_Currentcap *CurrentcapTransactorSession) ClaimDelayedRedeems() (*types.Transaction, error) {
	return _Currentcap.Contract.ClaimDelayedRedeems(&_Currentcap.TransactOpts)
}

// ClaimDelayedRedeems0 is a paid mutator transaction binding the contract method 0xf33fca17.
//
// Solidity: function claimDelayedRedeems(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_Currentcap *CurrentcapTransactor) ClaimDelayedRedeems0(opts *bind.TransactOpts, maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "claimDelayedRedeems0", maxNumberOfDelayedRedeemsToClaim)
}

// ClaimDelayedRedeems0 is a paid mutator transaction binding the contract method 0xf33fca17.
//
// Solidity: function claimDelayedRedeems(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_Currentcap *CurrentcapSession) ClaimDelayedRedeems0(maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.ClaimDelayedRedeems0(&_Currentcap.TransactOpts, maxNumberOfDelayedRedeemsToClaim)
}

// ClaimDelayedRedeems0 is a paid mutator transaction binding the contract method 0xf33fca17.
//
// Solidity: function claimDelayedRedeems(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_Currentcap *CurrentcapTransactorSession) ClaimDelayedRedeems0(maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.ClaimDelayedRedeems0(&_Currentcap.TransactOpts, maxNumberOfDelayedRedeemsToClaim)
}

// ClaimPrincipals is a paid mutator transaction binding the contract method 0x7849ace2.
//
// Solidity: function claimPrincipals(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_Currentcap *CurrentcapTransactor) ClaimPrincipals(opts *bind.TransactOpts, maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "claimPrincipals", maxNumberOfDelayedRedeemsToClaim)
}

// ClaimPrincipals is a paid mutator transaction binding the contract method 0x7849ace2.
//
// Solidity: function claimPrincipals(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_Currentcap *CurrentcapSession) ClaimPrincipals(maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.ClaimPrincipals(&_Currentcap.TransactOpts, maxNumberOfDelayedRedeemsToClaim)
}

// ClaimPrincipals is a paid mutator transaction binding the contract method 0x7849ace2.
//
// Solidity: function claimPrincipals(uint256 maxNumberOfDelayedRedeemsToClaim) returns()
func (_Currentcap *CurrentcapTransactorSession) ClaimPrincipals(maxNumberOfDelayedRedeemsToClaim *big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.ClaimPrincipals(&_Currentcap.TransactOpts, maxNumberOfDelayedRedeemsToClaim)
}

// ClaimPrincipals0 is a paid mutator transaction binding the contract method 0xa60c7b3a.
//
// Solidity: function claimPrincipals() returns()
func (_Currentcap *CurrentcapTransactor) ClaimPrincipals0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "claimPrincipals0")
}

// ClaimPrincipals0 is a paid mutator transaction binding the contract method 0xa60c7b3a.
//
// Solidity: function claimPrincipals() returns()
func (_Currentcap *CurrentcapSession) ClaimPrincipals0() (*types.Transaction, error) {
	return _Currentcap.Contract.ClaimPrincipals0(&_Currentcap.TransactOpts)
}

// ClaimPrincipals0 is a paid mutator transaction binding the contract method 0xa60c7b3a.
//
// Solidity: function claimPrincipals() returns()
func (_Currentcap *CurrentcapTransactorSession) ClaimPrincipals0() (*types.Transaction, error) {
	return _Currentcap.Contract.ClaimPrincipals0(&_Currentcap.TransactOpts)
}

// CreateDelayedRedeem is a paid mutator transaction binding the contract method 0xddc1bdea.
//
// Solidity: function createDelayedRedeem(address token, uint256 amount) returns()
func (_Currentcap *CurrentcapTransactor) CreateDelayedRedeem(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "createDelayedRedeem", token, amount)
}

// CreateDelayedRedeem is a paid mutator transaction binding the contract method 0xddc1bdea.
//
// Solidity: function createDelayedRedeem(address token, uint256 amount) returns()
func (_Currentcap *CurrentcapSession) CreateDelayedRedeem(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.CreateDelayedRedeem(&_Currentcap.TransactOpts, token, amount)
}

// CreateDelayedRedeem is a paid mutator transaction binding the contract method 0xddc1bdea.
//
// Solidity: function createDelayedRedeem(address token, uint256 amount) returns()
func (_Currentcap *CurrentcapTransactorSession) CreateDelayedRedeem(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.CreateDelayedRedeem(&_Currentcap.TransactOpts, token, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Currentcap *CurrentcapTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Currentcap *CurrentcapSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.GrantRole(&_Currentcap.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Currentcap *CurrentcapTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.GrantRole(&_Currentcap.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xeea39f79.
//
// Solidity: function initialize(address _defaultAdmin, address _uniBTC, address _vault, uint256 _redeemDelay, bool _whitelistEnabled) returns()
func (_Currentcap *CurrentcapTransactor) Initialize(opts *bind.TransactOpts, _defaultAdmin common.Address, _uniBTC common.Address, _vault common.Address, _redeemDelay *big.Int, _whitelistEnabled bool) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "initialize", _defaultAdmin, _uniBTC, _vault, _redeemDelay, _whitelistEnabled)
}

// Initialize is a paid mutator transaction binding the contract method 0xeea39f79.
//
// Solidity: function initialize(address _defaultAdmin, address _uniBTC, address _vault, uint256 _redeemDelay, bool _whitelistEnabled) returns()
func (_Currentcap *CurrentcapSession) Initialize(_defaultAdmin common.Address, _uniBTC common.Address, _vault common.Address, _redeemDelay *big.Int, _whitelistEnabled bool) (*types.Transaction, error) {
	return _Currentcap.Contract.Initialize(&_Currentcap.TransactOpts, _defaultAdmin, _uniBTC, _vault, _redeemDelay, _whitelistEnabled)
}

// Initialize is a paid mutator transaction binding the contract method 0xeea39f79.
//
// Solidity: function initialize(address _defaultAdmin, address _uniBTC, address _vault, uint256 _redeemDelay, bool _whitelistEnabled) returns()
func (_Currentcap *CurrentcapTransactorSession) Initialize(_defaultAdmin common.Address, _uniBTC common.Address, _vault common.Address, _redeemDelay *big.Int, _whitelistEnabled bool) (*types.Transaction, error) {
	return _Currentcap.Contract.Initialize(&_Currentcap.TransactOpts, _defaultAdmin, _uniBTC, _vault, _redeemDelay, _whitelistEnabled)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Currentcap *CurrentcapTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Currentcap *CurrentcapSession) Pause() (*types.Transaction, error) {
	return _Currentcap.Contract.Pause(&_Currentcap.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Currentcap *CurrentcapTransactorSession) Pause() (*types.Transaction, error) {
	return _Currentcap.Contract.Pause(&_Currentcap.TransactOpts)
}

// PauseTokens is a paid mutator transaction binding the contract method 0xc609684a.
//
// Solidity: function pauseTokens(address[] _tokens) returns()
func (_Currentcap *CurrentcapTransactor) PauseTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "pauseTokens", _tokens)
}

// PauseTokens is a paid mutator transaction binding the contract method 0xc609684a.
//
// Solidity: function pauseTokens(address[] _tokens) returns()
func (_Currentcap *CurrentcapSession) PauseTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.PauseTokens(&_Currentcap.TransactOpts, _tokens)
}

// PauseTokens is a paid mutator transaction binding the contract method 0xc609684a.
//
// Solidity: function pauseTokens(address[] _tokens) returns()
func (_Currentcap *CurrentcapTransactorSession) PauseTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.PauseTokens(&_Currentcap.TransactOpts, _tokens)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x89daf799.
//
// Solidity: function removeFromBlacklist(address[] _accounts) returns()
func (_Currentcap *CurrentcapTransactor) RemoveFromBlacklist(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "removeFromBlacklist", _accounts)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x89daf799.
//
// Solidity: function removeFromBlacklist(address[] _accounts) returns()
func (_Currentcap *CurrentcapSession) RemoveFromBlacklist(_accounts []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.RemoveFromBlacklist(&_Currentcap.TransactOpts, _accounts)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x89daf799.
//
// Solidity: function removeFromBlacklist(address[] _accounts) returns()
func (_Currentcap *CurrentcapTransactorSession) RemoveFromBlacklist(_accounts []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.RemoveFromBlacklist(&_Currentcap.TransactOpts, _accounts)
}

// RemoveFromBtclist is a paid mutator transaction binding the contract method 0xdd954363.
//
// Solidity: function removeFromBtclist(address[] _tokens) returns()
func (_Currentcap *CurrentcapTransactor) RemoveFromBtclist(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "removeFromBtclist", _tokens)
}

// RemoveFromBtclist is a paid mutator transaction binding the contract method 0xdd954363.
//
// Solidity: function removeFromBtclist(address[] _tokens) returns()
func (_Currentcap *CurrentcapSession) RemoveFromBtclist(_tokens []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.RemoveFromBtclist(&_Currentcap.TransactOpts, _tokens)
}

// RemoveFromBtclist is a paid mutator transaction binding the contract method 0xdd954363.
//
// Solidity: function removeFromBtclist(address[] _tokens) returns()
func (_Currentcap *CurrentcapTransactorSession) RemoveFromBtclist(_tokens []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.RemoveFromBtclist(&_Currentcap.TransactOpts, _tokens)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _accounts) returns()
func (_Currentcap *CurrentcapTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "removeFromWhitelist", _accounts)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _accounts) returns()
func (_Currentcap *CurrentcapSession) RemoveFromWhitelist(_accounts []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.RemoveFromWhitelist(&_Currentcap.TransactOpts, _accounts)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] _accounts) returns()
func (_Currentcap *CurrentcapTransactorSession) RemoveFromWhitelist(_accounts []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.RemoveFromWhitelist(&_Currentcap.TransactOpts, _accounts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Currentcap *CurrentcapTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Currentcap *CurrentcapSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.RenounceRole(&_Currentcap.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Currentcap *CurrentcapTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.RenounceRole(&_Currentcap.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Currentcap *CurrentcapTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Currentcap *CurrentcapSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.RevokeRole(&_Currentcap.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Currentcap *CurrentcapTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.RevokeRole(&_Currentcap.TransactOpts, role, account)
}

// SetMaxQuotaForTokens is a paid mutator transaction binding the contract method 0x3a413394.
//
// Solidity: function setMaxQuotaForTokens(address[] _tokens, uint256[] _quotas) returns()
func (_Currentcap *CurrentcapTransactor) SetMaxQuotaForTokens(opts *bind.TransactOpts, _tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "setMaxQuotaForTokens", _tokens, _quotas)
}

// SetMaxQuotaForTokens is a paid mutator transaction binding the contract method 0x3a413394.
//
// Solidity: function setMaxQuotaForTokens(address[] _tokens, uint256[] _quotas) returns()
func (_Currentcap *CurrentcapSession) SetMaxQuotaForTokens(_tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.SetMaxQuotaForTokens(&_Currentcap.TransactOpts, _tokens, _quotas)
}

// SetMaxQuotaForTokens is a paid mutator transaction binding the contract method 0x3a413394.
//
// Solidity: function setMaxQuotaForTokens(address[] _tokens, uint256[] _quotas) returns()
func (_Currentcap *CurrentcapTransactorSession) SetMaxQuotaForTokens(_tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.SetMaxQuotaForTokens(&_Currentcap.TransactOpts, _tokens, _quotas)
}

// SetQuotaRates is a paid mutator transaction binding the contract method 0x377ef35a.
//
// Solidity: function setQuotaRates(address[] _tokens, uint256[] _quotas) returns()
func (_Currentcap *CurrentcapTransactor) SetQuotaRates(opts *bind.TransactOpts, _tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "setQuotaRates", _tokens, _quotas)
}

// SetQuotaRates is a paid mutator transaction binding the contract method 0x377ef35a.
//
// Solidity: function setQuotaRates(address[] _tokens, uint256[] _quotas) returns()
func (_Currentcap *CurrentcapSession) SetQuotaRates(_tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.SetQuotaRates(&_Currentcap.TransactOpts, _tokens, _quotas)
}

// SetQuotaRates is a paid mutator transaction binding the contract method 0x377ef35a.
//
// Solidity: function setQuotaRates(address[] _tokens, uint256[] _quotas) returns()
func (_Currentcap *CurrentcapTransactorSession) SetQuotaRates(_tokens []common.Address, _quotas []*big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.SetQuotaRates(&_Currentcap.TransactOpts, _tokens, _quotas)
}

// SetRedeemDelay is a paid mutator transaction binding the contract method 0x668282a0.
//
// Solidity: function setRedeemDelay(uint256 _newDelay) returns()
func (_Currentcap *CurrentcapTransactor) SetRedeemDelay(opts *bind.TransactOpts, _newDelay *big.Int) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "setRedeemDelay", _newDelay)
}

// SetRedeemDelay is a paid mutator transaction binding the contract method 0x668282a0.
//
// Solidity: function setRedeemDelay(uint256 _newDelay) returns()
func (_Currentcap *CurrentcapSession) SetRedeemDelay(_newDelay *big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.SetRedeemDelay(&_Currentcap.TransactOpts, _newDelay)
}

// SetRedeemDelay is a paid mutator transaction binding the contract method 0x668282a0.
//
// Solidity: function setRedeemDelay(uint256 _newDelay) returns()
func (_Currentcap *CurrentcapTransactorSession) SetRedeemDelay(_newDelay *big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.SetRedeemDelay(&_Currentcap.TransactOpts, _newDelay)
}

// SetRedeemFeeRate is a paid mutator transaction binding the contract method 0x21e822c5.
//
// Solidity: function setRedeemFeeRate(uint256 _newFeeRate) returns()
func (_Currentcap *CurrentcapTransactor) SetRedeemFeeRate(opts *bind.TransactOpts, _newFeeRate *big.Int) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "setRedeemFeeRate", _newFeeRate)
}

// SetRedeemFeeRate is a paid mutator transaction binding the contract method 0x21e822c5.
//
// Solidity: function setRedeemFeeRate(uint256 _newFeeRate) returns()
func (_Currentcap *CurrentcapSession) SetRedeemFeeRate(_newFeeRate *big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.SetRedeemFeeRate(&_Currentcap.TransactOpts, _newFeeRate)
}

// SetRedeemFeeRate is a paid mutator transaction binding the contract method 0x21e822c5.
//
// Solidity: function setRedeemFeeRate(uint256 _newFeeRate) returns()
func (_Currentcap *CurrentcapTransactorSession) SetRedeemFeeRate(_newFeeRate *big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.SetRedeemFeeRate(&_Currentcap.TransactOpts, _newFeeRate)
}

// SetRedeemPrincipalDelay is a paid mutator transaction binding the contract method 0x98c0e0f0.
//
// Solidity: function setRedeemPrincipalDelay(uint256 _newDelay) returns()
func (_Currentcap *CurrentcapTransactor) SetRedeemPrincipalDelay(opts *bind.TransactOpts, _newDelay *big.Int) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "setRedeemPrincipalDelay", _newDelay)
}

// SetRedeemPrincipalDelay is a paid mutator transaction binding the contract method 0x98c0e0f0.
//
// Solidity: function setRedeemPrincipalDelay(uint256 _newDelay) returns()
func (_Currentcap *CurrentcapSession) SetRedeemPrincipalDelay(_newDelay *big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.SetRedeemPrincipalDelay(&_Currentcap.TransactOpts, _newDelay)
}

// SetRedeemPrincipalDelay is a paid mutator transaction binding the contract method 0x98c0e0f0.
//
// Solidity: function setRedeemPrincipalDelay(uint256 _newDelay) returns()
func (_Currentcap *CurrentcapTransactorSession) SetRedeemPrincipalDelay(_newDelay *big.Int) (*types.Transaction, error) {
	return _Currentcap.Contract.SetRedeemPrincipalDelay(&_Currentcap.TransactOpts, _newDelay)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _enabled) returns()
func (_Currentcap *CurrentcapTransactor) SetWhitelistEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "setWhitelistEnabled", _enabled)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _enabled) returns()
func (_Currentcap *CurrentcapSession) SetWhitelistEnabled(_enabled bool) (*types.Transaction, error) {
	return _Currentcap.Contract.SetWhitelistEnabled(&_Currentcap.TransactOpts, _enabled)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _enabled) returns()
func (_Currentcap *CurrentcapTransactorSession) SetWhitelistEnabled(_enabled bool) (*types.Transaction, error) {
	return _Currentcap.Contract.SetWhitelistEnabled(&_Currentcap.TransactOpts, _enabled)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Currentcap *CurrentcapTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Currentcap *CurrentcapSession) Unpause() (*types.Transaction, error) {
	return _Currentcap.Contract.Unpause(&_Currentcap.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Currentcap *CurrentcapTransactorSession) Unpause() (*types.Transaction, error) {
	return _Currentcap.Contract.Unpause(&_Currentcap.TransactOpts)
}

// UnpauseTokens is a paid mutator transaction binding the contract method 0x7a4df6eb.
//
// Solidity: function unpauseTokens(address[] _tokens) returns()
func (_Currentcap *CurrentcapTransactor) UnpauseTokens(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "unpauseTokens", _tokens)
}

// UnpauseTokens is a paid mutator transaction binding the contract method 0x7a4df6eb.
//
// Solidity: function unpauseTokens(address[] _tokens) returns()
func (_Currentcap *CurrentcapSession) UnpauseTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.UnpauseTokens(&_Currentcap.TransactOpts, _tokens)
}

// UnpauseTokens is a paid mutator transaction binding the contract method 0x7a4df6eb.
//
// Solidity: function unpauseTokens(address[] _tokens) returns()
func (_Currentcap *CurrentcapTransactorSession) UnpauseTokens(_tokens []common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.UnpauseTokens(&_Currentcap.TransactOpts, _tokens)
}

// WithdrawManagementFee is a paid mutator transaction binding the contract method 0xb175968e.
//
// Solidity: function withdrawManagementFee(uint256 _amount, address _recipient) returns()
func (_Currentcap *CurrentcapTransactor) WithdrawManagementFee(opts *bind.TransactOpts, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Currentcap.contract.Transact(opts, "withdrawManagementFee", _amount, _recipient)
}

// WithdrawManagementFee is a paid mutator transaction binding the contract method 0xb175968e.
//
// Solidity: function withdrawManagementFee(uint256 _amount, address _recipient) returns()
func (_Currentcap *CurrentcapSession) WithdrawManagementFee(_amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.WithdrawManagementFee(&_Currentcap.TransactOpts, _amount, _recipient)
}

// WithdrawManagementFee is a paid mutator transaction binding the contract method 0xb175968e.
//
// Solidity: function withdrawManagementFee(uint256 _amount, address _recipient) returns()
func (_Currentcap *CurrentcapTransactorSession) WithdrawManagementFee(_amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Currentcap.Contract.WithdrawManagementFee(&_Currentcap.TransactOpts, _amount, _recipient)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Currentcap *CurrentcapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Currentcap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Currentcap *CurrentcapSession) Receive() (*types.Transaction, error) {
	return _Currentcap.Contract.Receive(&_Currentcap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Currentcap *CurrentcapTransactorSession) Receive() (*types.Transaction, error) {
	return _Currentcap.Contract.Receive(&_Currentcap.TransactOpts)
}

// CurrentcapBlacklistAddedIterator is returned from FilterBlacklistAdded and is used to iterate over the raw logs and unpacked data for BlacklistAdded events raised by the Currentcap contract.
type CurrentcapBlacklistAddedIterator struct {
	Event *CurrentcapBlacklistAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapBlacklistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapBlacklistAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapBlacklistAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapBlacklistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapBlacklistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapBlacklistAdded represents a BlacklistAdded event raised by the Currentcap contract.
type CurrentcapBlacklistAdded struct {
	Accounts []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlacklistAdded is a free log retrieval operation binding the contract event 0x065786e2f100ecf1a39fd27fb1e6042658f97ff4d9093657ae121f534e46c038.
//
// Solidity: event BlacklistAdded(address[] accounts)
func (_Currentcap *CurrentcapFilterer) FilterBlacklistAdded(opts *bind.FilterOpts) (*CurrentcapBlacklistAddedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "BlacklistAdded")
	if err != nil {
		return nil, err
	}
	return &CurrentcapBlacklistAddedIterator{contract: _Currentcap.contract, event: "BlacklistAdded", logs: logs, sub: sub}, nil
}

// WatchBlacklistAdded is a free log subscription operation binding the contract event 0x065786e2f100ecf1a39fd27fb1e6042658f97ff4d9093657ae121f534e46c038.
//
// Solidity: event BlacklistAdded(address[] accounts)
func (_Currentcap *CurrentcapFilterer) WatchBlacklistAdded(opts *bind.WatchOpts, sink chan<- *CurrentcapBlacklistAdded) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "BlacklistAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapBlacklistAdded)
				if err := _Currentcap.contract.UnpackLog(event, "BlacklistAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlacklistAdded is a log parse operation binding the contract event 0x065786e2f100ecf1a39fd27fb1e6042658f97ff4d9093657ae121f534e46c038.
//
// Solidity: event BlacklistAdded(address[] accounts)
func (_Currentcap *CurrentcapFilterer) ParseBlacklistAdded(log types.Log) (*CurrentcapBlacklistAdded, error) {
	event := new(CurrentcapBlacklistAdded)
	if err := _Currentcap.contract.UnpackLog(event, "BlacklistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapBlacklistRemovedIterator is returned from FilterBlacklistRemoved and is used to iterate over the raw logs and unpacked data for BlacklistRemoved events raised by the Currentcap contract.
type CurrentcapBlacklistRemovedIterator struct {
	Event *CurrentcapBlacklistRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapBlacklistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapBlacklistRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapBlacklistRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapBlacklistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapBlacklistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapBlacklistRemoved represents a BlacklistRemoved event raised by the Currentcap contract.
type CurrentcapBlacklistRemoved struct {
	Accounts []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlacklistRemoved is a free log retrieval operation binding the contract event 0xb1a383a26b5d809f3cf5b9b022fbfe3e4896e6f1ff310ce38c785d5638bc31bf.
//
// Solidity: event BlacklistRemoved(address[] accounts)
func (_Currentcap *CurrentcapFilterer) FilterBlacklistRemoved(opts *bind.FilterOpts) (*CurrentcapBlacklistRemovedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "BlacklistRemoved")
	if err != nil {
		return nil, err
	}
	return &CurrentcapBlacklistRemovedIterator{contract: _Currentcap.contract, event: "BlacklistRemoved", logs: logs, sub: sub}, nil
}

// WatchBlacklistRemoved is a free log subscription operation binding the contract event 0xb1a383a26b5d809f3cf5b9b022fbfe3e4896e6f1ff310ce38c785d5638bc31bf.
//
// Solidity: event BlacklistRemoved(address[] accounts)
func (_Currentcap *CurrentcapFilterer) WatchBlacklistRemoved(opts *bind.WatchOpts, sink chan<- *CurrentcapBlacklistRemoved) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "BlacklistRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapBlacklistRemoved)
				if err := _Currentcap.contract.UnpackLog(event, "BlacklistRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlacklistRemoved is a log parse operation binding the contract event 0xb1a383a26b5d809f3cf5b9b022fbfe3e4896e6f1ff310ce38c785d5638bc31bf.
//
// Solidity: event BlacklistRemoved(address[] accounts)
func (_Currentcap *CurrentcapFilterer) ParseBlacklistRemoved(log types.Log) (*CurrentcapBlacklistRemoved, error) {
	event := new(CurrentcapBlacklistRemoved)
	if err := _Currentcap.contract.UnpackLog(event, "BlacklistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapBtclistAddedIterator is returned from FilterBtclistAdded and is used to iterate over the raw logs and unpacked data for BtclistAdded events raised by the Currentcap contract.
type CurrentcapBtclistAddedIterator struct {
	Event *CurrentcapBtclistAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapBtclistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapBtclistAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapBtclistAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapBtclistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapBtclistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapBtclistAdded represents a BtclistAdded event raised by the Currentcap contract.
type CurrentcapBtclistAdded struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBtclistAdded is a free log retrieval operation binding the contract event 0xa8f863280802460b807e3f117dda44a95e056aee58b5ac4ac75c254003130e73.
//
// Solidity: event BtclistAdded(address[] tokens)
func (_Currentcap *CurrentcapFilterer) FilterBtclistAdded(opts *bind.FilterOpts) (*CurrentcapBtclistAddedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "BtclistAdded")
	if err != nil {
		return nil, err
	}
	return &CurrentcapBtclistAddedIterator{contract: _Currentcap.contract, event: "BtclistAdded", logs: logs, sub: sub}, nil
}

// WatchBtclistAdded is a free log subscription operation binding the contract event 0xa8f863280802460b807e3f117dda44a95e056aee58b5ac4ac75c254003130e73.
//
// Solidity: event BtclistAdded(address[] tokens)
func (_Currentcap *CurrentcapFilterer) WatchBtclistAdded(opts *bind.WatchOpts, sink chan<- *CurrentcapBtclistAdded) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "BtclistAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapBtclistAdded)
				if err := _Currentcap.contract.UnpackLog(event, "BtclistAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBtclistAdded is a log parse operation binding the contract event 0xa8f863280802460b807e3f117dda44a95e056aee58b5ac4ac75c254003130e73.
//
// Solidity: event BtclistAdded(address[] tokens)
func (_Currentcap *CurrentcapFilterer) ParseBtclistAdded(log types.Log) (*CurrentcapBtclistAdded, error) {
	event := new(CurrentcapBtclistAdded)
	if err := _Currentcap.contract.UnpackLog(event, "BtclistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapBtclistRemovedIterator is returned from FilterBtclistRemoved and is used to iterate over the raw logs and unpacked data for BtclistRemoved events raised by the Currentcap contract.
type CurrentcapBtclistRemovedIterator struct {
	Event *CurrentcapBtclistRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapBtclistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapBtclistRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapBtclistRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapBtclistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapBtclistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapBtclistRemoved represents a BtclistRemoved event raised by the Currentcap contract.
type CurrentcapBtclistRemoved struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBtclistRemoved is a free log retrieval operation binding the contract event 0x90c77a347cd0b5c7f69208f770965d69e3af347c0e5ef344bfe02c2edcd3ccdc.
//
// Solidity: event BtclistRemoved(address[] tokens)
func (_Currentcap *CurrentcapFilterer) FilterBtclistRemoved(opts *bind.FilterOpts) (*CurrentcapBtclistRemovedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "BtclistRemoved")
	if err != nil {
		return nil, err
	}
	return &CurrentcapBtclistRemovedIterator{contract: _Currentcap.contract, event: "BtclistRemoved", logs: logs, sub: sub}, nil
}

// WatchBtclistRemoved is a free log subscription operation binding the contract event 0x90c77a347cd0b5c7f69208f770965d69e3af347c0e5ef344bfe02c2edcd3ccdc.
//
// Solidity: event BtclistRemoved(address[] tokens)
func (_Currentcap *CurrentcapFilterer) WatchBtclistRemoved(opts *bind.WatchOpts, sink chan<- *CurrentcapBtclistRemoved) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "BtclistRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapBtclistRemoved)
				if err := _Currentcap.contract.UnpackLog(event, "BtclistRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBtclistRemoved is a log parse operation binding the contract event 0x90c77a347cd0b5c7f69208f770965d69e3af347c0e5ef344bfe02c2edcd3ccdc.
//
// Solidity: event BtclistRemoved(address[] tokens)
func (_Currentcap *CurrentcapFilterer) ParseBtclistRemoved(log types.Log) (*CurrentcapBtclistRemoved, error) {
	event := new(CurrentcapBtclistRemoved)
	if err := _Currentcap.contract.UnpackLog(event, "BtclistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapDelayedRedeemCreatedIterator is returned from FilterDelayedRedeemCreated and is used to iterate over the raw logs and unpacked data for DelayedRedeemCreated events raised by the Currentcap contract.
type CurrentcapDelayedRedeemCreatedIterator struct {
	Event *CurrentcapDelayedRedeemCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapDelayedRedeemCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapDelayedRedeemCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapDelayedRedeemCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapDelayedRedeemCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapDelayedRedeemCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapDelayedRedeemCreated represents a DelayedRedeemCreated event raised by the Currentcap contract.
type CurrentcapDelayedRedeemCreated struct {
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	Index     *big.Int
	RedeemFee *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayedRedeemCreated is a free log retrieval operation binding the contract event 0xe2ede624e6a605222e831f4f91f2930b35c4d10323da5b68923c3086d4f0b3c0.
//
// Solidity: event DelayedRedeemCreated(address recipient, address token, uint256 amount, uint256 index, uint256 redeemFee)
func (_Currentcap *CurrentcapFilterer) FilterDelayedRedeemCreated(opts *bind.FilterOpts) (*CurrentcapDelayedRedeemCreatedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "DelayedRedeemCreated")
	if err != nil {
		return nil, err
	}
	return &CurrentcapDelayedRedeemCreatedIterator{contract: _Currentcap.contract, event: "DelayedRedeemCreated", logs: logs, sub: sub}, nil
}

// WatchDelayedRedeemCreated is a free log subscription operation binding the contract event 0xe2ede624e6a605222e831f4f91f2930b35c4d10323da5b68923c3086d4f0b3c0.
//
// Solidity: event DelayedRedeemCreated(address recipient, address token, uint256 amount, uint256 index, uint256 redeemFee)
func (_Currentcap *CurrentcapFilterer) WatchDelayedRedeemCreated(opts *bind.WatchOpts, sink chan<- *CurrentcapDelayedRedeemCreated) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "DelayedRedeemCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapDelayedRedeemCreated)
				if err := _Currentcap.contract.UnpackLog(event, "DelayedRedeemCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayedRedeemCreated is a log parse operation binding the contract event 0xe2ede624e6a605222e831f4f91f2930b35c4d10323da5b68923c3086d4f0b3c0.
//
// Solidity: event DelayedRedeemCreated(address recipient, address token, uint256 amount, uint256 index, uint256 redeemFee)
func (_Currentcap *CurrentcapFilterer) ParseDelayedRedeemCreated(log types.Log) (*CurrentcapDelayedRedeemCreated, error) {
	event := new(CurrentcapDelayedRedeemCreated)
	if err := _Currentcap.contract.UnpackLog(event, "DelayedRedeemCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapDelayedRedeemsClaimedIterator is returned from FilterDelayedRedeemsClaimed and is used to iterate over the raw logs and unpacked data for DelayedRedeemsClaimed events raised by the Currentcap contract.
type CurrentcapDelayedRedeemsClaimedIterator struct {
	Event *CurrentcapDelayedRedeemsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapDelayedRedeemsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapDelayedRedeemsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapDelayedRedeemsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapDelayedRedeemsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapDelayedRedeemsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapDelayedRedeemsClaimed represents a DelayedRedeemsClaimed event raised by the Currentcap contract.
type CurrentcapDelayedRedeemsClaimed struct {
	Recipient     common.Address
	Token         common.Address
	ClaimedAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelayedRedeemsClaimed is a free log retrieval operation binding the contract event 0xea08d9fa9c1ac98b666df5fdb636c7cda43a9b75c4d0e84088b1510d5d2396ea.
//
// Solidity: event DelayedRedeemsClaimed(address recipient, address token, uint256 claimedAmount)
func (_Currentcap *CurrentcapFilterer) FilterDelayedRedeemsClaimed(opts *bind.FilterOpts) (*CurrentcapDelayedRedeemsClaimedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "DelayedRedeemsClaimed")
	if err != nil {
		return nil, err
	}
	return &CurrentcapDelayedRedeemsClaimedIterator{contract: _Currentcap.contract, event: "DelayedRedeemsClaimed", logs: logs, sub: sub}, nil
}

// WatchDelayedRedeemsClaimed is a free log subscription operation binding the contract event 0xea08d9fa9c1ac98b666df5fdb636c7cda43a9b75c4d0e84088b1510d5d2396ea.
//
// Solidity: event DelayedRedeemsClaimed(address recipient, address token, uint256 claimedAmount)
func (_Currentcap *CurrentcapFilterer) WatchDelayedRedeemsClaimed(opts *bind.WatchOpts, sink chan<- *CurrentcapDelayedRedeemsClaimed) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "DelayedRedeemsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapDelayedRedeemsClaimed)
				if err := _Currentcap.contract.UnpackLog(event, "DelayedRedeemsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayedRedeemsClaimed is a log parse operation binding the contract event 0xea08d9fa9c1ac98b666df5fdb636c7cda43a9b75c4d0e84088b1510d5d2396ea.
//
// Solidity: event DelayedRedeemsClaimed(address recipient, address token, uint256 claimedAmount)
func (_Currentcap *CurrentcapFilterer) ParseDelayedRedeemsClaimed(log types.Log) (*CurrentcapDelayedRedeemsClaimed, error) {
	event := new(CurrentcapDelayedRedeemsClaimed)
	if err := _Currentcap.contract.UnpackLog(event, "DelayedRedeemsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapDelayedRedeemsCompletedIterator is returned from FilterDelayedRedeemsCompleted and is used to iterate over the raw logs and unpacked data for DelayedRedeemsCompleted events raised by the Currentcap contract.
type CurrentcapDelayedRedeemsCompletedIterator struct {
	Event *CurrentcapDelayedRedeemsCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapDelayedRedeemsCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapDelayedRedeemsCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapDelayedRedeemsCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapDelayedRedeemsCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapDelayedRedeemsCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapDelayedRedeemsCompleted represents a DelayedRedeemsCompleted event raised by the Currentcap contract.
type CurrentcapDelayedRedeemsCompleted struct {
	Recipient               common.Address
	BurnedAmount            *big.Int
	DelayedRedeemsCompleted *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterDelayedRedeemsCompleted is a free log retrieval operation binding the contract event 0xe4ba4789a7bd26025ed5932a86697793be11feeb86193a97dd9ef9849531ee60.
//
// Solidity: event DelayedRedeemsCompleted(address recipient, uint256 burnedAmount, uint256 delayedRedeemsCompleted)
func (_Currentcap *CurrentcapFilterer) FilterDelayedRedeemsCompleted(opts *bind.FilterOpts) (*CurrentcapDelayedRedeemsCompletedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "DelayedRedeemsCompleted")
	if err != nil {
		return nil, err
	}
	return &CurrentcapDelayedRedeemsCompletedIterator{contract: _Currentcap.contract, event: "DelayedRedeemsCompleted", logs: logs, sub: sub}, nil
}

// WatchDelayedRedeemsCompleted is a free log subscription operation binding the contract event 0xe4ba4789a7bd26025ed5932a86697793be11feeb86193a97dd9ef9849531ee60.
//
// Solidity: event DelayedRedeemsCompleted(address recipient, uint256 burnedAmount, uint256 delayedRedeemsCompleted)
func (_Currentcap *CurrentcapFilterer) WatchDelayedRedeemsCompleted(opts *bind.WatchOpts, sink chan<- *CurrentcapDelayedRedeemsCompleted) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "DelayedRedeemsCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapDelayedRedeemsCompleted)
				if err := _Currentcap.contract.UnpackLog(event, "DelayedRedeemsCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayedRedeemsCompleted is a log parse operation binding the contract event 0xe4ba4789a7bd26025ed5932a86697793be11feeb86193a97dd9ef9849531ee60.
//
// Solidity: event DelayedRedeemsCompleted(address recipient, uint256 burnedAmount, uint256 delayedRedeemsCompleted)
func (_Currentcap *CurrentcapFilterer) ParseDelayedRedeemsCompleted(log types.Log) (*CurrentcapDelayedRedeemsCompleted, error) {
	event := new(CurrentcapDelayedRedeemsCompleted)
	if err := _Currentcap.contract.UnpackLog(event, "DelayedRedeemsCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapDelayedRedeemsPrincipalClaimedIterator is returned from FilterDelayedRedeemsPrincipalClaimed and is used to iterate over the raw logs and unpacked data for DelayedRedeemsPrincipalClaimed events raised by the Currentcap contract.
type CurrentcapDelayedRedeemsPrincipalClaimedIterator struct {
	Event *CurrentcapDelayedRedeemsPrincipalClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapDelayedRedeemsPrincipalClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapDelayedRedeemsPrincipalClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapDelayedRedeemsPrincipalClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapDelayedRedeemsPrincipalClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapDelayedRedeemsPrincipalClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapDelayedRedeemsPrincipalClaimed represents a DelayedRedeemsPrincipalClaimed event raised by the Currentcap contract.
type CurrentcapDelayedRedeemsPrincipalClaimed struct {
	Recipient     common.Address
	Token         common.Address
	ClaimedAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDelayedRedeemsPrincipalClaimed is a free log retrieval operation binding the contract event 0xbf84d30fb64eb92b5961bba3a9c507953f2ea0fec68d7b78601c455825fb334a.
//
// Solidity: event DelayedRedeemsPrincipalClaimed(address recipient, address token, uint256 claimedAmount)
func (_Currentcap *CurrentcapFilterer) FilterDelayedRedeemsPrincipalClaimed(opts *bind.FilterOpts) (*CurrentcapDelayedRedeemsPrincipalClaimedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "DelayedRedeemsPrincipalClaimed")
	if err != nil {
		return nil, err
	}
	return &CurrentcapDelayedRedeemsPrincipalClaimedIterator{contract: _Currentcap.contract, event: "DelayedRedeemsPrincipalClaimed", logs: logs, sub: sub}, nil
}

// WatchDelayedRedeemsPrincipalClaimed is a free log subscription operation binding the contract event 0xbf84d30fb64eb92b5961bba3a9c507953f2ea0fec68d7b78601c455825fb334a.
//
// Solidity: event DelayedRedeemsPrincipalClaimed(address recipient, address token, uint256 claimedAmount)
func (_Currentcap *CurrentcapFilterer) WatchDelayedRedeemsPrincipalClaimed(opts *bind.WatchOpts, sink chan<- *CurrentcapDelayedRedeemsPrincipalClaimed) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "DelayedRedeemsPrincipalClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapDelayedRedeemsPrincipalClaimed)
				if err := _Currentcap.contract.UnpackLog(event, "DelayedRedeemsPrincipalClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayedRedeemsPrincipalClaimed is a log parse operation binding the contract event 0xbf84d30fb64eb92b5961bba3a9c507953f2ea0fec68d7b78601c455825fb334a.
//
// Solidity: event DelayedRedeemsPrincipalClaimed(address recipient, address token, uint256 claimedAmount)
func (_Currentcap *CurrentcapFilterer) ParseDelayedRedeemsPrincipalClaimed(log types.Log) (*CurrentcapDelayedRedeemsPrincipalClaimed, error) {
	event := new(CurrentcapDelayedRedeemsPrincipalClaimed)
	if err := _Currentcap.contract.UnpackLog(event, "DelayedRedeemsPrincipalClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapDelayedRedeemsPrincipalCompletedIterator is returned from FilterDelayedRedeemsPrincipalCompleted and is used to iterate over the raw logs and unpacked data for DelayedRedeemsPrincipalCompleted events raised by the Currentcap contract.
type CurrentcapDelayedRedeemsPrincipalCompletedIterator struct {
	Event *CurrentcapDelayedRedeemsPrincipalCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapDelayedRedeemsPrincipalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapDelayedRedeemsPrincipalCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapDelayedRedeemsPrincipalCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapDelayedRedeemsPrincipalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapDelayedRedeemsPrincipalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapDelayedRedeemsPrincipalCompleted represents a DelayedRedeemsPrincipalCompleted event raised by the Currentcap contract.
type CurrentcapDelayedRedeemsPrincipalCompleted struct {
	Recipient               common.Address
	PrincipalAmount         *big.Int
	DelayedRedeemsCompleted *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterDelayedRedeemsPrincipalCompleted is a free log retrieval operation binding the contract event 0xd8bb463f2d1d5c972542f10f4a33f937c1042d40970dc2f0a9f927e4f18b89c7.
//
// Solidity: event DelayedRedeemsPrincipalCompleted(address recipient, uint256 principalAmount, uint256 delayedRedeemsCompleted)
func (_Currentcap *CurrentcapFilterer) FilterDelayedRedeemsPrincipalCompleted(opts *bind.FilterOpts) (*CurrentcapDelayedRedeemsPrincipalCompletedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "DelayedRedeemsPrincipalCompleted")
	if err != nil {
		return nil, err
	}
	return &CurrentcapDelayedRedeemsPrincipalCompletedIterator{contract: _Currentcap.contract, event: "DelayedRedeemsPrincipalCompleted", logs: logs, sub: sub}, nil
}

// WatchDelayedRedeemsPrincipalCompleted is a free log subscription operation binding the contract event 0xd8bb463f2d1d5c972542f10f4a33f937c1042d40970dc2f0a9f927e4f18b89c7.
//
// Solidity: event DelayedRedeemsPrincipalCompleted(address recipient, uint256 principalAmount, uint256 delayedRedeemsCompleted)
func (_Currentcap *CurrentcapFilterer) WatchDelayedRedeemsPrincipalCompleted(opts *bind.WatchOpts, sink chan<- *CurrentcapDelayedRedeemsPrincipalCompleted) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "DelayedRedeemsPrincipalCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapDelayedRedeemsPrincipalCompleted)
				if err := _Currentcap.contract.UnpackLog(event, "DelayedRedeemsPrincipalCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayedRedeemsPrincipalCompleted is a log parse operation binding the contract event 0xd8bb463f2d1d5c972542f10f4a33f937c1042d40970dc2f0a9f927e4f18b89c7.
//
// Solidity: event DelayedRedeemsPrincipalCompleted(address recipient, uint256 principalAmount, uint256 delayedRedeemsCompleted)
func (_Currentcap *CurrentcapFilterer) ParseDelayedRedeemsPrincipalCompleted(log types.Log) (*CurrentcapDelayedRedeemsPrincipalCompleted, error) {
	event := new(CurrentcapDelayedRedeemsPrincipalCompleted)
	if err := _Currentcap.contract.UnpackLog(event, "DelayedRedeemsPrincipalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Currentcap contract.
type CurrentcapInitializedIterator struct {
	Event *CurrentcapInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapInitialized represents a Initialized event raised by the Currentcap contract.
type CurrentcapInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Currentcap *CurrentcapFilterer) FilterInitialized(opts *bind.FilterOpts) (*CurrentcapInitializedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CurrentcapInitializedIterator{contract: _Currentcap.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Currentcap *CurrentcapFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CurrentcapInitialized) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapInitialized)
				if err := _Currentcap.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Currentcap *CurrentcapFilterer) ParseInitialized(log types.Log) (*CurrentcapInitialized, error) {
	event := new(CurrentcapInitialized)
	if err := _Currentcap.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapManagementFeeWithdrawnIterator is returned from FilterManagementFeeWithdrawn and is used to iterate over the raw logs and unpacked data for ManagementFeeWithdrawn events raised by the Currentcap contract.
type CurrentcapManagementFeeWithdrawnIterator struct {
	Event *CurrentcapManagementFeeWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapManagementFeeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapManagementFeeWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapManagementFeeWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapManagementFeeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapManagementFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapManagementFeeWithdrawn represents a ManagementFeeWithdrawn event raised by the Currentcap contract.
type CurrentcapManagementFeeWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterManagementFeeWithdrawn is a free log retrieval operation binding the contract event 0x583d1744b1f2b01833f7f10ff38436dd7a76ff50695a487bfb9f0f3d07749b49.
//
// Solidity: event ManagementFeeWithdrawn(address recipient, uint256 amount)
func (_Currentcap *CurrentcapFilterer) FilterManagementFeeWithdrawn(opts *bind.FilterOpts) (*CurrentcapManagementFeeWithdrawnIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "ManagementFeeWithdrawn")
	if err != nil {
		return nil, err
	}
	return &CurrentcapManagementFeeWithdrawnIterator{contract: _Currentcap.contract, event: "ManagementFeeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchManagementFeeWithdrawn is a free log subscription operation binding the contract event 0x583d1744b1f2b01833f7f10ff38436dd7a76ff50695a487bfb9f0f3d07749b49.
//
// Solidity: event ManagementFeeWithdrawn(address recipient, uint256 amount)
func (_Currentcap *CurrentcapFilterer) WatchManagementFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *CurrentcapManagementFeeWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "ManagementFeeWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapManagementFeeWithdrawn)
				if err := _Currentcap.contract.UnpackLog(event, "ManagementFeeWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseManagementFeeWithdrawn is a log parse operation binding the contract event 0x583d1744b1f2b01833f7f10ff38436dd7a76ff50695a487bfb9f0f3d07749b49.
//
// Solidity: event ManagementFeeWithdrawn(address recipient, uint256 amount)
func (_Currentcap *CurrentcapFilterer) ParseManagementFeeWithdrawn(log types.Log) (*CurrentcapManagementFeeWithdrawn, error) {
	event := new(CurrentcapManagementFeeWithdrawn)
	if err := _Currentcap.contract.UnpackLog(event, "ManagementFeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapMaxQuotaSetIterator is returned from FilterMaxQuotaSet and is used to iterate over the raw logs and unpacked data for MaxQuotaSet events raised by the Currentcap contract.
type CurrentcapMaxQuotaSetIterator struct {
	Event *CurrentcapMaxQuotaSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapMaxQuotaSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapMaxQuotaSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapMaxQuotaSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapMaxQuotaSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapMaxQuotaSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapMaxQuotaSet represents a MaxQuotaSet event raised by the Currentcap contract.
type CurrentcapMaxQuotaSet struct {
	Token         common.Address
	PreviousQuota *big.Int
	NewQuota      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxQuotaSet is a free log retrieval operation binding the contract event 0x96fa1888abd5b7a22236027bf87904c3f144b181c9c8016594b72c9b47f94d79.
//
// Solidity: event MaxQuotaSet(address token, uint256 previousQuota, uint256 newQuota)
func (_Currentcap *CurrentcapFilterer) FilterMaxQuotaSet(opts *bind.FilterOpts) (*CurrentcapMaxQuotaSetIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "MaxQuotaSet")
	if err != nil {
		return nil, err
	}
	return &CurrentcapMaxQuotaSetIterator{contract: _Currentcap.contract, event: "MaxQuotaSet", logs: logs, sub: sub}, nil
}

// WatchMaxQuotaSet is a free log subscription operation binding the contract event 0x96fa1888abd5b7a22236027bf87904c3f144b181c9c8016594b72c9b47f94d79.
//
// Solidity: event MaxQuotaSet(address token, uint256 previousQuota, uint256 newQuota)
func (_Currentcap *CurrentcapFilterer) WatchMaxQuotaSet(opts *bind.WatchOpts, sink chan<- *CurrentcapMaxQuotaSet) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "MaxQuotaSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapMaxQuotaSet)
				if err := _Currentcap.contract.UnpackLog(event, "MaxQuotaSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMaxQuotaSet is a log parse operation binding the contract event 0x96fa1888abd5b7a22236027bf87904c3f144b181c9c8016594b72c9b47f94d79.
//
// Solidity: event MaxQuotaSet(address token, uint256 previousQuota, uint256 newQuota)
func (_Currentcap *CurrentcapFilterer) ParseMaxQuotaSet(log types.Log) (*CurrentcapMaxQuotaSet, error) {
	event := new(CurrentcapMaxQuotaSet)
	if err := _Currentcap.contract.UnpackLog(event, "MaxQuotaSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Currentcap contract.
type CurrentcapPausedIterator struct {
	Event *CurrentcapPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapPaused represents a Paused event raised by the Currentcap contract.
type CurrentcapPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Currentcap *CurrentcapFilterer) FilterPaused(opts *bind.FilterOpts) (*CurrentcapPausedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CurrentcapPausedIterator{contract: _Currentcap.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Currentcap *CurrentcapFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CurrentcapPaused) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapPaused)
				if err := _Currentcap.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Currentcap *CurrentcapFilterer) ParsePaused(log types.Log) (*CurrentcapPaused, error) {
	event := new(CurrentcapPaused)
	if err := _Currentcap.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapRateSetIterator is returned from FilterRateSet and is used to iterate over the raw logs and unpacked data for RateSet events raised by the Currentcap contract.
type CurrentcapRateSetIterator struct {
	Event *CurrentcapRateSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapRateSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapRateSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapRateSet represents a RateSet event raised by the Currentcap contract.
type CurrentcapRateSet struct {
	Token         common.Address
	PreviousQuota *big.Int
	NewQuota      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRateSet is a free log retrieval operation binding the contract event 0x9e31cca092b9e764bfc6b1b552d55ad4b035e609318fecc26cd38b34e8dd08bb.
//
// Solidity: event RateSet(address token, uint256 previousQuota, uint256 newQuota)
func (_Currentcap *CurrentcapFilterer) FilterRateSet(opts *bind.FilterOpts) (*CurrentcapRateSetIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "RateSet")
	if err != nil {
		return nil, err
	}
	return &CurrentcapRateSetIterator{contract: _Currentcap.contract, event: "RateSet", logs: logs, sub: sub}, nil
}

// WatchRateSet is a free log subscription operation binding the contract event 0x9e31cca092b9e764bfc6b1b552d55ad4b035e609318fecc26cd38b34e8dd08bb.
//
// Solidity: event RateSet(address token, uint256 previousQuota, uint256 newQuota)
func (_Currentcap *CurrentcapFilterer) WatchRateSet(opts *bind.WatchOpts, sink chan<- *CurrentcapRateSet) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "RateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapRateSet)
				if err := _Currentcap.contract.UnpackLog(event, "RateSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRateSet is a log parse operation binding the contract event 0x9e31cca092b9e764bfc6b1b552d55ad4b035e609318fecc26cd38b34e8dd08bb.
//
// Solidity: event RateSet(address token, uint256 previousQuota, uint256 newQuota)
func (_Currentcap *CurrentcapFilterer) ParseRateSet(log types.Log) (*CurrentcapRateSet, error) {
	event := new(CurrentcapRateSet)
	if err := _Currentcap.contract.UnpackLog(event, "RateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapRedeemDelaySetIterator is returned from FilterRedeemDelaySet and is used to iterate over the raw logs and unpacked data for RedeemDelaySet events raised by the Currentcap contract.
type CurrentcapRedeemDelaySetIterator struct {
	Event *CurrentcapRedeemDelaySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapRedeemDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapRedeemDelaySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapRedeemDelaySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapRedeemDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapRedeemDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapRedeemDelaySet represents a RedeemDelaySet event raised by the Currentcap contract.
type CurrentcapRedeemDelaySet struct {
	PreviousDelay *big.Int
	NewDelay      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRedeemDelaySet is a free log retrieval operation binding the contract event 0x1431e66d652872e09a583cd3d9bb280e0f743dca287cc9344a1d80a1596add3a.
//
// Solidity: event RedeemDelaySet(uint256 previousDelay, uint256 newDelay)
func (_Currentcap *CurrentcapFilterer) FilterRedeemDelaySet(opts *bind.FilterOpts) (*CurrentcapRedeemDelaySetIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "RedeemDelaySet")
	if err != nil {
		return nil, err
	}
	return &CurrentcapRedeemDelaySetIterator{contract: _Currentcap.contract, event: "RedeemDelaySet", logs: logs, sub: sub}, nil
}

// WatchRedeemDelaySet is a free log subscription operation binding the contract event 0x1431e66d652872e09a583cd3d9bb280e0f743dca287cc9344a1d80a1596add3a.
//
// Solidity: event RedeemDelaySet(uint256 previousDelay, uint256 newDelay)
func (_Currentcap *CurrentcapFilterer) WatchRedeemDelaySet(opts *bind.WatchOpts, sink chan<- *CurrentcapRedeemDelaySet) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "RedeemDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapRedeemDelaySet)
				if err := _Currentcap.contract.UnpackLog(event, "RedeemDelaySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeemDelaySet is a log parse operation binding the contract event 0x1431e66d652872e09a583cd3d9bb280e0f743dca287cc9344a1d80a1596add3a.
//
// Solidity: event RedeemDelaySet(uint256 previousDelay, uint256 newDelay)
func (_Currentcap *CurrentcapFilterer) ParseRedeemDelaySet(log types.Log) (*CurrentcapRedeemDelaySet, error) {
	event := new(CurrentcapRedeemDelaySet)
	if err := _Currentcap.contract.UnpackLog(event, "RedeemDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapRedeemFeeRateSetIterator is returned from FilterRedeemFeeRateSet and is used to iterate over the raw logs and unpacked data for RedeemFeeRateSet events raised by the Currentcap contract.
type CurrentcapRedeemFeeRateSetIterator struct {
	Event *CurrentcapRedeemFeeRateSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapRedeemFeeRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapRedeemFeeRateSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapRedeemFeeRateSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapRedeemFeeRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapRedeemFeeRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapRedeemFeeRateSet represents a RedeemFeeRateSet event raised by the Currentcap contract.
type CurrentcapRedeemFeeRateSet struct {
	PreviousFeeRate *big.Int
	NewFeeRate      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRedeemFeeRateSet is a free log retrieval operation binding the contract event 0x4d9ff9777a971177b4810d0c671707ff05f2469b58efd13a676a0eb42fe53528.
//
// Solidity: event RedeemFeeRateSet(uint256 previousFeeRate, uint256 newFeeRate)
func (_Currentcap *CurrentcapFilterer) FilterRedeemFeeRateSet(opts *bind.FilterOpts) (*CurrentcapRedeemFeeRateSetIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "RedeemFeeRateSet")
	if err != nil {
		return nil, err
	}
	return &CurrentcapRedeemFeeRateSetIterator{contract: _Currentcap.contract, event: "RedeemFeeRateSet", logs: logs, sub: sub}, nil
}

// WatchRedeemFeeRateSet is a free log subscription operation binding the contract event 0x4d9ff9777a971177b4810d0c671707ff05f2469b58efd13a676a0eb42fe53528.
//
// Solidity: event RedeemFeeRateSet(uint256 previousFeeRate, uint256 newFeeRate)
func (_Currentcap *CurrentcapFilterer) WatchRedeemFeeRateSet(opts *bind.WatchOpts, sink chan<- *CurrentcapRedeemFeeRateSet) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "RedeemFeeRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapRedeemFeeRateSet)
				if err := _Currentcap.contract.UnpackLog(event, "RedeemFeeRateSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeemFeeRateSet is a log parse operation binding the contract event 0x4d9ff9777a971177b4810d0c671707ff05f2469b58efd13a676a0eb42fe53528.
//
// Solidity: event RedeemFeeRateSet(uint256 previousFeeRate, uint256 newFeeRate)
func (_Currentcap *CurrentcapFilterer) ParseRedeemFeeRateSet(log types.Log) (*CurrentcapRedeemFeeRateSet, error) {
	event := new(CurrentcapRedeemFeeRateSet)
	if err := _Currentcap.contract.UnpackLog(event, "RedeemFeeRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapRedeemPrincipalDelaySetIterator is returned from FilterRedeemPrincipalDelaySet and is used to iterate over the raw logs and unpacked data for RedeemPrincipalDelaySet events raised by the Currentcap contract.
type CurrentcapRedeemPrincipalDelaySetIterator struct {
	Event *CurrentcapRedeemPrincipalDelaySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapRedeemPrincipalDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapRedeemPrincipalDelaySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapRedeemPrincipalDelaySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapRedeemPrincipalDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapRedeemPrincipalDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapRedeemPrincipalDelaySet represents a RedeemPrincipalDelaySet event raised by the Currentcap contract.
type CurrentcapRedeemPrincipalDelaySet struct {
	PreviousDelay *big.Int
	NewDelay      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRedeemPrincipalDelaySet is a free log retrieval operation binding the contract event 0xcf8211b2d9a5296d32fee575ddbc4623ed83733cae205f8c1046b6eaf48dd7b2.
//
// Solidity: event RedeemPrincipalDelaySet(uint256 previousDelay, uint256 newDelay)
func (_Currentcap *CurrentcapFilterer) FilterRedeemPrincipalDelaySet(opts *bind.FilterOpts) (*CurrentcapRedeemPrincipalDelaySetIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "RedeemPrincipalDelaySet")
	if err != nil {
		return nil, err
	}
	return &CurrentcapRedeemPrincipalDelaySetIterator{contract: _Currentcap.contract, event: "RedeemPrincipalDelaySet", logs: logs, sub: sub}, nil
}

// WatchRedeemPrincipalDelaySet is a free log subscription operation binding the contract event 0xcf8211b2d9a5296d32fee575ddbc4623ed83733cae205f8c1046b6eaf48dd7b2.
//
// Solidity: event RedeemPrincipalDelaySet(uint256 previousDelay, uint256 newDelay)
func (_Currentcap *CurrentcapFilterer) WatchRedeemPrincipalDelaySet(opts *bind.WatchOpts, sink chan<- *CurrentcapRedeemPrincipalDelaySet) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "RedeemPrincipalDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapRedeemPrincipalDelaySet)
				if err := _Currentcap.contract.UnpackLog(event, "RedeemPrincipalDelaySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeemPrincipalDelaySet is a log parse operation binding the contract event 0xcf8211b2d9a5296d32fee575ddbc4623ed83733cae205f8c1046b6eaf48dd7b2.
//
// Solidity: event RedeemPrincipalDelaySet(uint256 previousDelay, uint256 newDelay)
func (_Currentcap *CurrentcapFilterer) ParseRedeemPrincipalDelaySet(log types.Log) (*CurrentcapRedeemPrincipalDelaySet, error) {
	event := new(CurrentcapRedeemPrincipalDelaySet)
	if err := _Currentcap.contract.UnpackLog(event, "RedeemPrincipalDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Currentcap contract.
type CurrentcapRoleAdminChangedIterator struct {
	Event *CurrentcapRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapRoleAdminChanged represents a RoleAdminChanged event raised by the Currentcap contract.
type CurrentcapRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Currentcap *CurrentcapFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CurrentcapRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CurrentcapRoleAdminChangedIterator{contract: _Currentcap.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Currentcap *CurrentcapFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CurrentcapRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapRoleAdminChanged)
				if err := _Currentcap.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Currentcap *CurrentcapFilterer) ParseRoleAdminChanged(log types.Log) (*CurrentcapRoleAdminChanged, error) {
	event := new(CurrentcapRoleAdminChanged)
	if err := _Currentcap.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Currentcap contract.
type CurrentcapRoleGrantedIterator struct {
	Event *CurrentcapRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapRoleGranted represents a RoleGranted event raised by the Currentcap contract.
type CurrentcapRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Currentcap *CurrentcapFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CurrentcapRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CurrentcapRoleGrantedIterator{contract: _Currentcap.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Currentcap *CurrentcapFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CurrentcapRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapRoleGranted)
				if err := _Currentcap.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Currentcap *CurrentcapFilterer) ParseRoleGranted(log types.Log) (*CurrentcapRoleGranted, error) {
	event := new(CurrentcapRoleGranted)
	if err := _Currentcap.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Currentcap contract.
type CurrentcapRoleRevokedIterator struct {
	Event *CurrentcapRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapRoleRevoked represents a RoleRevoked event raised by the Currentcap contract.
type CurrentcapRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Currentcap *CurrentcapFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CurrentcapRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CurrentcapRoleRevokedIterator{contract: _Currentcap.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Currentcap *CurrentcapFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CurrentcapRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapRoleRevoked)
				if err := _Currentcap.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Currentcap *CurrentcapFilterer) ParseRoleRevoked(log types.Log) (*CurrentcapRoleRevoked, error) {
	event := new(CurrentcapRoleRevoked)
	if err := _Currentcap.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapTokensPausedIterator is returned from FilterTokensPaused and is used to iterate over the raw logs and unpacked data for TokensPaused events raised by the Currentcap contract.
type CurrentcapTokensPausedIterator struct {
	Event *CurrentcapTokensPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapTokensPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapTokensPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapTokensPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapTokensPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapTokensPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapTokensPaused represents a TokensPaused event raised by the Currentcap contract.
type CurrentcapTokensPaused struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensPaused is a free log retrieval operation binding the contract event 0xa4fbb323c047ef2555d72263081bbb4280ee59d506657303bf5fc991b54204be.
//
// Solidity: event TokensPaused(address[] tokens)
func (_Currentcap *CurrentcapFilterer) FilterTokensPaused(opts *bind.FilterOpts) (*CurrentcapTokensPausedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "TokensPaused")
	if err != nil {
		return nil, err
	}
	return &CurrentcapTokensPausedIterator{contract: _Currentcap.contract, event: "TokensPaused", logs: logs, sub: sub}, nil
}

// WatchTokensPaused is a free log subscription operation binding the contract event 0xa4fbb323c047ef2555d72263081bbb4280ee59d506657303bf5fc991b54204be.
//
// Solidity: event TokensPaused(address[] tokens)
func (_Currentcap *CurrentcapFilterer) WatchTokensPaused(opts *bind.WatchOpts, sink chan<- *CurrentcapTokensPaused) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "TokensPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapTokensPaused)
				if err := _Currentcap.contract.UnpackLog(event, "TokensPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensPaused is a log parse operation binding the contract event 0xa4fbb323c047ef2555d72263081bbb4280ee59d506657303bf5fc991b54204be.
//
// Solidity: event TokensPaused(address[] tokens)
func (_Currentcap *CurrentcapFilterer) ParseTokensPaused(log types.Log) (*CurrentcapTokensPaused, error) {
	event := new(CurrentcapTokensPaused)
	if err := _Currentcap.contract.UnpackLog(event, "TokensPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapTokensUnpausedIterator is returned from FilterTokensUnpaused and is used to iterate over the raw logs and unpacked data for TokensUnpaused events raised by the Currentcap contract.
type CurrentcapTokensUnpausedIterator struct {
	Event *CurrentcapTokensUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapTokensUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapTokensUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapTokensUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapTokensUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapTokensUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapTokensUnpaused represents a TokensUnpaused event raised by the Currentcap contract.
type CurrentcapTokensUnpaused struct {
	Tokens []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensUnpaused is a free log retrieval operation binding the contract event 0x4dd04d346e64df7bcc65df29a0f1f1f84815ff758717f30587cabe38792d7c31.
//
// Solidity: event TokensUnpaused(address[] tokens)
func (_Currentcap *CurrentcapFilterer) FilterTokensUnpaused(opts *bind.FilterOpts) (*CurrentcapTokensUnpausedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "TokensUnpaused")
	if err != nil {
		return nil, err
	}
	return &CurrentcapTokensUnpausedIterator{contract: _Currentcap.contract, event: "TokensUnpaused", logs: logs, sub: sub}, nil
}

// WatchTokensUnpaused is a free log subscription operation binding the contract event 0x4dd04d346e64df7bcc65df29a0f1f1f84815ff758717f30587cabe38792d7c31.
//
// Solidity: event TokensUnpaused(address[] tokens)
func (_Currentcap *CurrentcapFilterer) WatchTokensUnpaused(opts *bind.WatchOpts, sink chan<- *CurrentcapTokensUnpaused) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "TokensUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapTokensUnpaused)
				if err := _Currentcap.contract.UnpackLog(event, "TokensUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensUnpaused is a log parse operation binding the contract event 0x4dd04d346e64df7bcc65df29a0f1f1f84815ff758717f30587cabe38792d7c31.
//
// Solidity: event TokensUnpaused(address[] tokens)
func (_Currentcap *CurrentcapFilterer) ParseTokensUnpaused(log types.Log) (*CurrentcapTokensUnpaused, error) {
	event := new(CurrentcapTokensUnpaused)
	if err := _Currentcap.contract.UnpackLog(event, "TokensUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Currentcap contract.
type CurrentcapUnpausedIterator struct {
	Event *CurrentcapUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapUnpaused represents a Unpaused event raised by the Currentcap contract.
type CurrentcapUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Currentcap *CurrentcapFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CurrentcapUnpausedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CurrentcapUnpausedIterator{contract: _Currentcap.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Currentcap *CurrentcapFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CurrentcapUnpaused) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapUnpaused)
				if err := _Currentcap.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Currentcap *CurrentcapFilterer) ParseUnpaused(log types.Log) (*CurrentcapUnpaused, error) {
	event := new(CurrentcapUnpaused)
	if err := _Currentcap.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapWhitelistAddedIterator is returned from FilterWhitelistAdded and is used to iterate over the raw logs and unpacked data for WhitelistAdded events raised by the Currentcap contract.
type CurrentcapWhitelistAddedIterator struct {
	Event *CurrentcapWhitelistAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapWhitelistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapWhitelistAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapWhitelistAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapWhitelistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapWhitelistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapWhitelistAdded represents a WhitelistAdded event raised by the Currentcap contract.
type CurrentcapWhitelistAdded struct {
	Accounts []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAdded is a free log retrieval operation binding the contract event 0xf74f148a4f930a0f67a2c33ba932a14e3e91b4e6468f21e545932fd825111538.
//
// Solidity: event WhitelistAdded(address[] accounts)
func (_Currentcap *CurrentcapFilterer) FilterWhitelistAdded(opts *bind.FilterOpts) (*CurrentcapWhitelistAddedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "WhitelistAdded")
	if err != nil {
		return nil, err
	}
	return &CurrentcapWhitelistAddedIterator{contract: _Currentcap.contract, event: "WhitelistAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistAdded is a free log subscription operation binding the contract event 0xf74f148a4f930a0f67a2c33ba932a14e3e91b4e6468f21e545932fd825111538.
//
// Solidity: event WhitelistAdded(address[] accounts)
func (_Currentcap *CurrentcapFilterer) WatchWhitelistAdded(opts *bind.WatchOpts, sink chan<- *CurrentcapWhitelistAdded) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "WhitelistAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapWhitelistAdded)
				if err := _Currentcap.contract.UnpackLog(event, "WhitelistAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWhitelistAdded is a log parse operation binding the contract event 0xf74f148a4f930a0f67a2c33ba932a14e3e91b4e6468f21e545932fd825111538.
//
// Solidity: event WhitelistAdded(address[] accounts)
func (_Currentcap *CurrentcapFilterer) ParseWhitelistAdded(log types.Log) (*CurrentcapWhitelistAdded, error) {
	event := new(CurrentcapWhitelistAdded)
	if err := _Currentcap.contract.UnpackLog(event, "WhitelistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapWhitelistEnabledSetIterator is returned from FilterWhitelistEnabledSet and is used to iterate over the raw logs and unpacked data for WhitelistEnabledSet events raised by the Currentcap contract.
type CurrentcapWhitelistEnabledSetIterator struct {
	Event *CurrentcapWhitelistEnabledSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapWhitelistEnabledSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapWhitelistEnabledSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapWhitelistEnabledSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapWhitelistEnabledSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapWhitelistEnabledSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapWhitelistEnabledSet represents a WhitelistEnabledSet event raised by the Currentcap contract.
type CurrentcapWhitelistEnabledSet struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistEnabledSet is a free log retrieval operation binding the contract event 0x411283ae1b0e68089790510eb77ccad9b761295be576637799607c8ae066fe9f.
//
// Solidity: event WhitelistEnabledSet(bool enabled)
func (_Currentcap *CurrentcapFilterer) FilterWhitelistEnabledSet(opts *bind.FilterOpts) (*CurrentcapWhitelistEnabledSetIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "WhitelistEnabledSet")
	if err != nil {
		return nil, err
	}
	return &CurrentcapWhitelistEnabledSetIterator{contract: _Currentcap.contract, event: "WhitelistEnabledSet", logs: logs, sub: sub}, nil
}

// WatchWhitelistEnabledSet is a free log subscription operation binding the contract event 0x411283ae1b0e68089790510eb77ccad9b761295be576637799607c8ae066fe9f.
//
// Solidity: event WhitelistEnabledSet(bool enabled)
func (_Currentcap *CurrentcapFilterer) WatchWhitelistEnabledSet(opts *bind.WatchOpts, sink chan<- *CurrentcapWhitelistEnabledSet) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "WhitelistEnabledSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapWhitelistEnabledSet)
				if err := _Currentcap.contract.UnpackLog(event, "WhitelistEnabledSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWhitelistEnabledSet is a log parse operation binding the contract event 0x411283ae1b0e68089790510eb77ccad9b761295be576637799607c8ae066fe9f.
//
// Solidity: event WhitelistEnabledSet(bool enabled)
func (_Currentcap *CurrentcapFilterer) ParseWhitelistEnabledSet(log types.Log) (*CurrentcapWhitelistEnabledSet, error) {
	event := new(CurrentcapWhitelistEnabledSet)
	if err := _Currentcap.contract.UnpackLog(event, "WhitelistEnabledSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CurrentcapWhitelistRemovedIterator is returned from FilterWhitelistRemoved and is used to iterate over the raw logs and unpacked data for WhitelistRemoved events raised by the Currentcap contract.
type CurrentcapWhitelistRemovedIterator struct {
	Event *CurrentcapWhitelistRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CurrentcapWhitelistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CurrentcapWhitelistRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CurrentcapWhitelistRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CurrentcapWhitelistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CurrentcapWhitelistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CurrentcapWhitelistRemoved represents a WhitelistRemoved event raised by the Currentcap contract.
type CurrentcapWhitelistRemoved struct {
	Accounts []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistRemoved is a free log retrieval operation binding the contract event 0x1d474f57a5c483b47a8bf6006e39086f96dd040a00cb348e22f80a4ca2c6f222.
//
// Solidity: event WhitelistRemoved(address[] accounts)
func (_Currentcap *CurrentcapFilterer) FilterWhitelistRemoved(opts *bind.FilterOpts) (*CurrentcapWhitelistRemovedIterator, error) {

	logs, sub, err := _Currentcap.contract.FilterLogs(opts, "WhitelistRemoved")
	if err != nil {
		return nil, err
	}
	return &CurrentcapWhitelistRemovedIterator{contract: _Currentcap.contract, event: "WhitelistRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistRemoved is a free log subscription operation binding the contract event 0x1d474f57a5c483b47a8bf6006e39086f96dd040a00cb348e22f80a4ca2c6f222.
//
// Solidity: event WhitelistRemoved(address[] accounts)
func (_Currentcap *CurrentcapFilterer) WatchWhitelistRemoved(opts *bind.WatchOpts, sink chan<- *CurrentcapWhitelistRemoved) (event.Subscription, error) {

	logs, sub, err := _Currentcap.contract.WatchLogs(opts, "WhitelistRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CurrentcapWhitelistRemoved)
				if err := _Currentcap.contract.UnpackLog(event, "WhitelistRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWhitelistRemoved is a log parse operation binding the contract event 0x1d474f57a5c483b47a8bf6006e39086f96dd040a00cb348e22f80a4ca2c6f222.
//
// Solidity: event WhitelistRemoved(address[] accounts)
func (_Currentcap *CurrentcapFilterer) ParseWhitelistRemoved(log types.Log) (*CurrentcapWhitelistRemoved, error) {
	event := new(CurrentcapWhitelistRemoved)
	if err := _Currentcap.contract.UnpackLog(event, "WhitelistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
