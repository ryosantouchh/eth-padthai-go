package main

import (
	"ryosantouchh/eth-padthai-go/ledger"
	"ryosantouchh/eth-padthai-go/liquidity"
)

func main() {
	mockLedger := ledger.MockLedger()
	mockLedger.Transfer("1", "2", "ASSET_A", 20.0)

  mockTradePair := liquidity.MockTradePair()
  mockTradePair.Trade("ETH", 10.0)
}
