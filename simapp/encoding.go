package simapp

import (
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/public-awesome/stakebird/simapp/params"
)

// MakeTestEncodingConfig creates an EncodingConfig for testing
func MakeTestEncodingConfig() params.EncodingConfig {
	encodingConfig := params.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
