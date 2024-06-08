package liquidity

import "fmt"

type Pool struct {
	Pairs map[string]*TradePair
}

type TradePair struct {
	AssetX   string
	AssetY   string
	ReserveX float64
	ReserveY float64
	Constant float64
}

func (p *TradePair) Trade(asset string, amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("trade amount must be positive")
	}

	var newReserveX, newReserveY float64
	switch asset {
	case p.AssetX:
		if p.ReserveX-amount <= 0 {
			return fmt.Errorf("insufficient %v reserve for trade", p.AssetX)
		}
		newReserveX = p.ReserveX - amount
		newReserveY = p.Constant / newReserveX
	case p.AssetY:
		if p.ReserveY-amount <= 0 {
			return fmt.Errorf("insufficient %v reserve for trade", p.AssetY)
		}
		newReserveY = p.ReserveY - amount
		newReserveX = p.Constant / newReserveY
	default:
		return fmt.Errorf("unsupported asset: %v", asset)
	}

	if newReserveX <= 0 || newReserveY <= 0 {
		return fmt.Errorf("trade would result in zero reserve")
	}

	p.ReserveX = newReserveX
	p.ReserveY = newReserveY

	return nil
}

func MockTradePair() *TradePair {
	tradePair := TradePair{
		AssetX:   "ETH",
		AssetY:   "SOL",
		ReserveX: 100.0,
		ReserveY: 100.0,
		Constant: 10000,
	}
	return &tradePair
}
