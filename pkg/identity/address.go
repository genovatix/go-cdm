package identity

import (
	crand "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/genovatix/algoliocdm/internal/utils"
	"github.com/genovatix/algoliocdm/pkg/cryptography"
	"github.com/zeebo/blake3"
	"log"
	"math/big"
	"time"
)

const (
	// Define sizes for each component of RawAddress
	VerSize         = 4  // Size in bytes for the version field (uint32)
	TimestampSize   = 8  // Size in bytes for the timestamp field (int64)
	ChainCodeLength = 32 // Example length of the chain code in bytes
	PubKeyLength    = 64 // Example length of the public key in bytes, adjust based on your key format
	PrivKeyLength   = 64 // Example length of the private key in bytes, adjust based on your key format
	SeedLength      = 32 // Length of the seed in bytes

	// ExpectedTotalLength Calculate the expected total length of RawAddress when serialized
	ExpectedTotalLength = VerSize + TimestampSize + ChainCodeLength + PubKeyLength + PrivKeyLength + SeedLength
)

type Address struct {
	raw []byte
	cryptography.Decapsulatable
}

func (addr Address) String() string {
	return base64.StdEncoding.EncodeToString(addr.Bytes())
}

func (addr Address) Hex() string {
	return fmt.Sprintf("0x%s", addr.String())
}

func (addr Address) Bytes() []byte {
	return addr.raw
}

func (addr Address) Raw() (interface{}, error) {
	return addr.FromBytes(addr.Bytes())
}

type RawAddress struct {
	prefix    string
	chainCode []byte
	ver       uint32
	timestamp int64
	pubKey    []byte
	privKey   []byte
	seed      [32]byte
	bi        *big.Int
	cryptography.Encapsulatable
}

var BytePrefix []byte = []byte("0x100010001000")
var WitnessPrefix []byte = []byte("0x1101")
var ZkpPrefix []byte = []byte("0xab1d")

func (r *RawAddress) Init(chainCode []byte, ver uint32, pubKey []byte, privKey []byte, timestamp int64) {
	verBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(verBytes, ver)

	timestampBytes := make([]byte, 8) // int64 has a fixed size of 8 bytes
	binary.BigEndian.PutUint64(timestampBytes, uint64(timestamp))

	// Assuming fixed sizes for pubKey and privKey
	r.chainCode = chainCode
	r.ver = ver
	r.timestamp = timestamp
	r.pubKey = pubKey
	r.privKey = privKey

	// Seed generation
	r.seed = [32]byte{}
	copy(r.seed[:], NewSeed(32))

	// Combine all parts
	combined := append(verBytes, chainCode...)
	combined = append(combined, timestampBytes...)
	combined = append(combined, pubKey...)
	combined = append(combined, privKey...)
	combined = append(combined, r.seed[:]...) // Include seed if reversible is needed

	// Convert combined to big.Int for potential storage/transmission
	r.bi = new(big.Int).SetBytes(combined)
}

func (r *RawAddress) ToBytes() []byte {
	buf := make([]byte, ExpectedTotalLength)

	// Version
	binary.BigEndian.PutUint32(buf[:VerSize], r.ver)

	// ChainCode
	copy(buf[VerSize:VerSize+ChainCodeLength], r.chainCode)

	// Timestamp
	binary.BigEndian.PutUint64(buf[VerSize+ChainCodeLength:VerSize+ChainCodeLength+TimestampSize], uint64(r.timestamp))

	// PublicKey
	copy(buf[VerSize+ChainCodeLength+TimestampSize:VerSize+ChainCodeLength+TimestampSize+PubKeyLength], r.pubKey)

	// PrivateKey
	copy(buf[VerSize+ChainCodeLength+TimestampSize+PubKeyLength:VerSize+ChainCodeLength+TimestampSize+PubKeyLength+PrivKeyLength], r.privKey)

	// Seed
	copy(buf[VerSize+ChainCodeLength+TimestampSize+PubKeyLength+PrivKeyLength:], r.seed[:])

	return buf
}

func HashKeyAddress(r *RawAddress) []byte {
	h, _ := blake3.NewKeyed(r.seed[:])
	rb, _ := json.Marshal(r)
	h.Write(rb)
	return h.Sum(nil)
}

func NewSeed(bit int) []byte {
	seed := make([]byte, bit)
	crand.Read(seed)
	return seed
}

func (addr Address) FromBytes(data []byte) (interface{}, error) {
	if len(data) != ExpectedTotalLength {
		return nil, fmt.Errorf("incorrect data length: got %d bytes, expected %d", len(data), ExpectedTotalLength)
	}

	r := &RawAddress{
		ver:       binary.BigEndian.Uint32(data[:VerSize]),
		chainCode: data[VerSize : VerSize+ChainCodeLength],
		timestamp: int64(binary.BigEndian.Uint64(data[VerSize+ChainCodeLength : VerSize+ChainCodeLength+TimestampSize])),
		pubKey:    data[VerSize+ChainCodeLength+TimestampSize : VerSize+ChainCodeLength+TimestampSize+PubKeyLength],
		privKey:   data[VerSize+ChainCodeLength+TimestampSize+PubKeyLength : VerSize+ChainCodeLength+TimestampSize+PubKeyLength+PrivKeyLength],
		// Ensure the seed is correctly handled if it's part of the serialized format
	}
	// Extract the seed
	seedStartIndex := VerSize + ChainCodeLength + TimestampSize + PubKeyLength + PrivKeyLength
	var seed [SeedLength]byte
	copy(seed[:], data[seedStartIndex:seedStartIndex+SeedLength])
	r.seed = seed

	return r, nil
}

func (r *RawAddress) Marshal(key []byte, dst []byte) {
	result, err := cryptography.Encapsulate(r, key)
	if err != nil {
		log.Fatal(err.Error())
	}
	copy(dst, result)
}

func (addr Address) Unmarshal(key []byte, data []byte, dst *RawAddress) {
	result, err := cryptography.Decapsulate(addr, key)
	if err != nil {
		log.Fatal(err.Error())
	}
	rr, ok := result.(*RawAddress)
	if !ok {
		log.Fatal("failed to cast result to *RawAddress")
	}

	*dst = *rr

}

func NewAddress(cfg *utils.NetConfig, pubKey []byte, privKey []byte) *Address {
	_ = &RawAddress{
		chainCode: cfg.Code,
		ver:       cfg.Ver,
		timestamp: time.Now().Unix(),
		pubKey:    pubKey,
		privKey:   privKey,
	}

	return nil
}
