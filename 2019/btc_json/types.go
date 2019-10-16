package js

type PubKeyType uint8

const (
	NullData    PubKeyType = 1 //nulldata
	PubKey      PubKeyType = 2 //pubkey
	PubKeyHash  PubKeyType = 3 //pubkeyhash
	ScriptHash  PubKeyType = 4 //scripthash
	MultiSig    PubKeyType = 5 //multisig
	NonStandard PubKeyType = 6 //nonstandard
)

func (p PubKeyType) String() string {
	switch p {
	case NullData:
		return "nulldata"
	case PubKey:
		return "pubkey"
	case PubKeyHash:
		return "pubkeyhash"
	case ScriptHash:
		return "scripthash"
	case MultiSig:
		return "multisig"
	case NonStandard:
		return "nonstandard"
	default:
		return "unknown"
	}
}

func PubKeyTypeFromString(typ string) PubKeyType {
	switch typ {
	case "nulldata":
		return NullData
	case "pubkey":
		return PubKey
	case "pubkeyhash":
		return PubKeyHash
	case "scripthash":
		return ScriptHash
	case "multisig":
		return MultiSig
	case "nonstandard":
		return NonStandard
	default:
		return 0
	}
}

type Header struct {
	Height        uint64  `json:"height"            bson:"height"`                      //区块高度
	Confirmations int64   `json:"confirmations"     bson:"confirmations"`               //确认数
	Hash          string  `json:"hash"              bson:"hash"`                        //哈希
	StrippedSize  uint64  `json:"strippedsize"      bson:"strippedsize,omitempty"`      //剔除隔离见证数据后的区块字节数
	PrevBlockHash string  `json:"previousblockhash" bson:"previousblockhash,omitempty"` //前置哈希
	NextBlockHash string  `json:"nextblockhash"     bson:"nextblockhash,omitempty"`     //后置哈希
	Timestamp     int64   `json:"time"              bson:"time"`                        //区块打包时间戳
	MedianTime    int64   `json:"mediantime"        bson:"mediantime,omitempty"`        //区块中值时间戳
	Size          uint64  `json:"size"              bson:"size,omitempty"`              //字节大小
	Weight        uint64  `json:"weight"            bson:"weight,omitempty"`            //BIP141定义的区块权重
	Version       int32   `json:"version"           bson:"version"`                     //版本号
	VersionHex    string  `json:"versionHex"        bson:"versionHex,omitempty"`        //版本号16进制字符串
	MerkleRoot    string  `json:"merkleroot"        bson:"merkleroot,omitempty"`        //区块的默克尔树根
	Nonce         uint32  `json:"nonce"             bson:"nonce,omitempty"`             //nonce值
	Bits          string  `json:"bits"              bson:"bits"`                        //
	Difficulty    float64 `json:"difficulty"        bson:"difficulty,omitempty"`        //难度
	ChainWork     string  `json:"chainwork"         bson:"chainwork,omitempty"`         //
	NTx           uint64  `json:"n_tx"              bson:"n_tx,omitempty"`              //交易数量
}

type Block struct {
	Header
	Txs []*Transaction `json:"tx"` //交易集合
}

type Transaction struct {
	TxId     string `json:"txid"      bson:"txid"`               //交易哈希，包含隔离验证的数据
	Hash     string `json:"hash"      bson:"hash"`               //交易哈希，不包含隔离验证的数据
	Version  int32  `json:"version"   bson:"version"`            //版本号
	Size     int    `json:"size"      bson:"size"`               //交易大小
	VSize    int    `json:"vsize"     bson:"vsize,omitempty"`    //
	Weight   uint64 `json:"weight"    bson:"weight,omitempty"`   //BIP141定义的区块权重
	LockTime uint32 `json:"locktime"  bson:"locktime,omitempty"` //0表示立即加入区块；0-5亿表示区块高度；大于5亿表示Unix时间戳，交易被创建早于该时间，则不会发送到比特币网络；
	Hex      string `json:"hex"       bson:"hex,omitempty"`      //交易原始数据
	VIn      []VIn  `json:"vin"       bson:"vin"`                //输入
	VOut     []VOut `json:"vout"      bson:"vout"`               //输出

	BlockHeight   uint64 `json:"height"        bson:"height"`         //区块高度
	Fee           uint64 `json:"fee"           bson:"fee"`            //输入减去输出的值
	BlockHash     string `json:"blockhash"     bson:"blockhash"`      //区块哈希
	BlockTime     int64  `json:"blocktime"     bson:"blocktime"`      //区块打包时间
	Time          int64  `json:"time"          bson:"time,omitempty"` //区块打包时间
	Confirmations int64  `json:"confirmations" bson:"confirmations"`  //区块确认数
}

type ScriptSig struct {
	Asm string `json:"asm" bson:"asm,omitempty"` //
	Hex string `json:"hex" bson:"hex,omitempty"` //
}
type VIn struct {
	TxID      string     `json:"txid,omitempty"      bson:"txid,omitempty"`      //交易哈希
	VOut      uint32     `json:"vout"                bson:"vout"`                //out的索引量
	ScriptSig *ScriptSig `json:"scriptSig,omitempty" bson:"scriptsig,omitempty"` //输入脚本
	Sequence  uint32     `json:"sequence"            bson:"sequence"`            //配合lockTime实现延时确认交易
	CoinBase  string     `json:"coinbase,omitempty"  bson:"coinbase,omitempty"`  //当这是一笔coinbase交易时，这里填入coinbase信息
	Refer     *VOut      `json:"refer,omitempty"     bson:"refer,omitempty"`     //vin对应的vout，需要拉取in的交易获取
}

type ScriptPubKey struct {
	Asm       string   `json:"asm"                 bson:"asm,omitempty"`
	Hex       string   `json:"hex,omitempty"       bson:"hex"`
	ReqSigs   uint32   `json:"reqSigs,omitempty"   bson:"reqsigs,omitempty"`
	Type      string   `json:"type"                bson:"type"` //公钥脚本类型
	Addresses []string `json:"addresses,omitempty" bson:"addresses"`
}

type VOut struct {
	Value        float64      `json:"value"        bson:"value"`
	N            uint32       `json:"n"            bson:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey" bson:"scriptpubkey"`
}

type ChainInfo struct {
	Chain                string  `json:"chain"`
	Blocks               uint64  `json:"blocks"`
	Headers              uint64  `json:"headers"`
	BestBlockHash        string  `json:"bestblockhash"` //最高区块
	Difficulty           float64 `json:"difficulty"`
	MedianTime           uint64  `json:"mediantime"`
	VerificationProgress float64 `json:"verificationprogress"`
	InitialBlockDownload bool    `json:"initialblockdownload"`
	ChainWork            string  `json:"chainwork"`
	SizeOnDisk           uint64  `json:"size_on_disk"`
	Pruned               bool    `json:"pruned"`
}
