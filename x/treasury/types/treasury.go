package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewTreasury(mintGenesisSupply bool, targetSupply sdk.Coins, distributed sdk.Coins, burned sdk.Coins) Treasury {
	return Treasury{
		MintGenesisSupply: mintGenesisSupply,
		TargetSupply:      targetSupply,
		Distributed:       distributed,
		Burned:            burned,
	}
}

func InitialTreasury(mintGenesisSupply bool, targetSupply sdk.Coins, distributed sdk.Coins, burned sdk.Coins) Treasury {
	return NewTreasury(
		mintGenesisSupply,
		targetSupply,
		distributed,
		burned,
	)
}

func DefaultInitialTreasury() Treasury {
	targetMelcSupply, _ := sdk.ConvertCoin(sdk.NewInt64Coin("melc", 200000000), "umelc")
	targetMelgSupply, _ := sdk.ConvertCoin(sdk.NewInt64Coin("melg", 200000000), "umelg")

	return InitialTreasury(
		true,
		sdk.NewCoins(targetMelcSupply, targetMelgSupply),
		sdk.NewCoins(),
		sdk.NewCoins(),
	)
}
