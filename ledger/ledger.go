package ledger

import "fmt"

type Ledger struct {
	NodeList map[string]*Node
}

type Node struct {
	ID     string
	Assets map[string]*Asset
}

type Asset struct {
	ID      string
	Balance float64
}

func MockLedger() Ledger {
	AssetNode1 := Asset{
		ID:      "ASSET_A",
		Balance: 100.0,
	}
	node1 := Node{
		ID: "1",
		Assets: map[string]*Asset{
			"ASSET_A": &AssetNode1,
		},
	}

	AssetNode2 := Asset{
		ID:      "ASSET_A",
		Balance: 50.0,
	}
	node2 := Node{
		ID: "2",
		Assets: map[string]*Asset{
			"ASSET_A": &AssetNode2,
		},
	}

	ledger := Ledger{
		NodeList: map[string]*Node{
			"1": &node1,
			"2": &node2,
		},
	}

	return ledger
}

func (l *Ledger) Transfer(sender string, recipient string, assetId string, amount float64) error {
	if sender == recipient {
		err := fmt.Errorf("sender and recipient is the same address, cannot sent the asset to %v\n", sender)
		fmt.Println(err)
		return err
	}

	senderNode := l.NodeList[sender]
	recipientNode := l.NodeList[recipient]

	if !senderNode.AssetIDExist(assetId) && !recipientNode.AssetIDExist(assetId) {
		err := fmt.Errorf("asset id = %v is not recognized", assetId)
		fmt.Println(err)
		return err
	}

	if !senderNode.CheckSufficientBalance(assetId, amount) {
		err := fmt.Errorf("sender : %v with asset id : %v for transfer is not sufficient", sender, assetId)
		fmt.Println(err)
		return err
	}

	senderNode.Assets[assetId].Balance -= amount
	recipientNode.Assets[assetId].Balance += amount

	return nil
}

func (n *Node) CheckSufficientBalance(assetId string, amount float64) bool {
	if n.Assets[assetId].Balance > amount {
		return true
	}
	return false
}

func (n *Node) AssetIDExist(assetId string) bool {
	if _, ok := n.Assets[assetId]; ok {
		return true
	}
	return false
}
