package tron

type EthereumError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BlockByNumberResp struct {
	BaseFeePerGas    interface{}    `json:"baseFeePerGas"`
	Difficulty       interface{}    `json:"difficulty"`
	ExtraData        interface{}    `json:"extraData"`
	GasLimit         string         `json:"gasLimit"`
	GasUsed          string         `json:"gasUsed"`
	Hash             string         `json:"hash"`
	LogsBloom        string         `json:"logsBloom"`
	Miner            string         `json:"miner"`
	MixHash          interface{}    `json:"mixHash"`
	Nonce            interface{}    `json:"nonce"`
	Number           string         `json:"number"`
	ParentHash       string         `json:"parentHash"`
	ReceiptsRoot     interface{}    `json:"receiptsRoot"`
	Sha3Uncles       interface{}    `json:"sha3Uncles"`
	Size             string         `json:"size"`
	StateRoot        string         `json:"stateRoot"`
	Timestamp        string         `json:"timestamp"`
	TotalDifficulty  interface{}    `json:"totalDifficulty"`
	Transactions     []Transactions `json:"transactions"`
	TransactionsRoot string         `json:"transactionsRoot"`
	Uncles           []interface{}  `json:"uncles"`
}

type Transactions struct {
	BlockHash        string      `json:"blockHash"`
	BlockNumber      string      `json:"blockNumber"`
	From             string      `json:"from"`
	Gas              string      `json:"gas"`
	GasPrice         string      `json:"gasPrice"`
	Hash             string      `json:"hash"`
	Input            string      `json:"input"`
	Nonce            interface{} `json:"nonce"`
	R                string      `json:"r"`
	S                string      `json:"s"`
	To               string      `json:"to"`
	TransactionIndex string      `json:"transactionIndex"`
	Type             string      `json:"type"`
	V                string      `json:"v"`
	Value            string      `json:"value"`
}

type TransactionReceipt struct {
	BlockHash         string      `json:"blockHash"`
	BlockNumber       string      `json:"blockNumber"`
	ContractAddress   interface{} `json:"contractAddress"`
	CumulativeGasUsed string      `json:"cumulativeGasUsed"`
	EffectiveGasPrice string      `json:"effectiveGasPrice"`
	From              string      `json:"from"`
	GasUsed           string      `json:"gasUsed"`
	Logs              []struct {
		Address          string   `json:"address"`
		BlockHash        string   `json:"blockHash"`
		BlockNumber      string   `json:"blockNumber"`
		Data             string   `json:"data"`
		LogIndex         string   `json:"logIndex"`
		Removed          bool     `json:"removed"`
		Topics           []string `json:"topics"`
		TransactionHash  string   `json:"transactionHash"`
		TransactionIndex string   `json:"transactionIndex"`
	} `json:"logs"`
	LogsBloom        string `json:"logsBloom"`
	Status           string `json:"status"`
	To               string `json:"to"`
	TransactionHash  string `json:"transactionHash"`
	TransactionIndex string `json:"transactionIndex"`
	Type             string `json:"type"`
}

type Resp struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      string `json:"id"`
	Result  string `json:"result,omitempty"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error,omitempty"`
}
