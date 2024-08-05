package bsc

type BlockResponse struct {
	Difficulty      string         `json:"difficulty"`
	ExtraData       string         `json:"extraData"`
	GasLimit        string         `json:"gas_limit"`
	GasUsed         string         `json:"gasUsed"`
	Hash            string         `json:"hash"`
	LogsBloom       string         `json:"logsBloom"`
	Miner           string         `json:"miner"`
	MixHash         string         `json:"mixHash"`
	Nonce           string         `json:"nonce"`
	ParentHash      string         `json:"parentHash"`
	ReceiptsRoot    string         `json:"receiptsRoot"`
	Sha3Uncles      string         `json:"sha3Uncles"`
	Size            string         `json:"size"`
	StateRoot       string         `json:"stateRoot"`
	Timestamp       string         `json:"timestamp"`
	TotalDifficulty string         `json:"totalDifficulty"`
	Transactions    []*Transaction `json:"transactions"`
	TransactionRoot string         `json:"transactionRoot"`
	Uncles          []string       `json:"uncles"`
}

type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	R                string `json:"r"`
	S                string `json:"s"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Type             string `json:"type"`
	V                string `json:"v"`
	Value            string `json:"value"`
}
