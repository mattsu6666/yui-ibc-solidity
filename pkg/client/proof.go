package client

import (
	"encoding/hex"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type StateProof struct {
	AccountProofRLP []byte
	StorageProofRLP [][]byte
}

func (cl ETHClient) GetStateProof(address common.Address, storageKeys [][]byte, blockNumber *big.Int) (*StateProof, error) {
	bz, err := cl.getProof(address, storageKeys, "0x"+blockNumber.Text(16))
	if err != nil {
		return nil, err
	}
	var proof struct {
		AccountProof []string `json:"accountProof"`
		StorageProof []struct {
			Proof []string `json:"proof"`
		} `json:"storageProof"`
	}
	if err := json.Unmarshal(bz, &proof); err != nil {
		return nil, err
	}

	var encodedProof StateProof
	encodedProof.AccountProofRLP, err = encodeRLP(proof.AccountProof)
	if err != nil {
		return nil, err
	}
	for _, p := range proof.StorageProof {
		bz, err := encodeRLP(p.Proof)
		if err != nil {
			return nil, err
		}
		encodedProof.StorageProofRLP = append(encodedProof.StorageProofRLP, bz)
	}

	return &encodedProof, nil
}

func (cl ETHClient) getProof(address common.Address, storageKeys [][]byte, blockNumber string) ([]byte, error) {
	hashes := []common.Hash{}
	for _, k := range storageKeys {
		var h common.Hash
		if err := h.UnmarshalText(k); err != nil {
			return nil, err
		}
		hashes = append(hashes, h)
	}
	var msg json.RawMessage
	if err := cl.rpcClient.Call(&msg, "eth_getProof", address, hashes, blockNumber); err != nil {
		return nil, err
	}
	return msg, nil
}

func encodeRLP(proof []string) ([]byte, error) {
	var target [][][]byte
	for _, p := range proof {
		bz, err := hex.DecodeString(p[2:])
		if err != nil {
			return nil, err
		}
		var val [][]byte
		if err := rlp.DecodeBytes(bz, &val); err != nil {
			return nil, err
		}
		target = append(target, val)
	}
	bz, err := rlp.EncodeToBytes(target)
	if err != nil {
		return nil, err
	}
	return bz, nil
}
