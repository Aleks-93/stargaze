package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

// Default parameter namespace
const (
	DefaultParamspace                 = ModuleName
	DefaultVotingPeriod time.Duration = time.Hour * 24 * 3
)

// 50%
var (
	DefaultRewardPoolAllocation = sdk.NewDecWithPrec(50, 2)
)

// Parameter store keys
var (
	KeyVotingPeriod         = []byte("VotingPeriod")
	KeyRewardPoolAllocation = []byte("RewardPoolAllocation")
)

// Params - used for initializing default parameter for stake at genesis
type Params struct {
	VotingPeriod         time.Duration `json:"voting_period" yaml:"voting_period"`
	RewardPoolAllocation sdk.Dec       `json:"reward_pool_allocation" yaml:"reward_pool_allocation"`
}

// NewParams creates a new Params object
func NewParams(votingPeriod time.Duration, rewardPoolAllocation sdk.Dec) Params {
	return Params{
		VotingPeriod:         votingPeriod,
		RewardPoolAllocation: rewardPoolAllocation,
	}
}

// String implements the stringer interface for Params
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyVotingPeriod, &p.VotingPeriod, validateVotingPeriod),
		paramtypes.NewParamSetPair(KeyRewardPoolAllocation, &p.RewardPoolAllocation, validateRewardPoolAllocation),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(DefaultVotingPeriod, DefaultRewardPoolAllocation)
}

func (p Params) Validate() error {
	if err := validateVotingPeriod(p.VotingPeriod); err != nil {
		return err
	}

	return nil
}

func validateVotingPeriod(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("voting period must be positive: %d", v)
	}

	return nil
}

func validateRewardPoolAllocation(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsZero() {
		return fmt.Errorf("reward pool allocation can't be zero: %d", v)
	}

	return nil
}