package blockchain

import (
	"crypto/sha256"
	"reflect"
	"testing"
)

func TestBlockChain_ChainName(t *testing.T) {
	type fields struct {
		chainName string
		blocks    []Block
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "BlockChain has ChainName",
			fields: fields{
				chainName: "Test Chain",
				blocks:    nil,
			},
			want: "Test Chain",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlockChain{
				chainName: tt.fields.chainName,
				blocks:    tt.fields.blocks,
			}
			if got := bc.ChainName(); got != tt.want {
				t.Errorf("BlockChain.ChainName() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestBlockChain_GetLatestBlock(t *testing.T) {
	type fields struct {
		chainName string
		blocks    []Block
	}
	tests := []struct {
		name   string
		fields fields
		want   *Block
	}{
		{
			name: "",
			fields: fields{
				chainName: "Test",
				blocks: []Block{
					{
						index:        0,
						previousHash: sha256.Sum256([]byte("0")),
						timestamp:    0,
						data:         "",
						hash:         sha256.Sum256([]byte("Hello, world")),
					},
				},
			},
			want: &Block{
				index:        0,
				previousHash: sha256.Sum256([]byte("0")),
				timestamp:    0,
				data:         "",
				hash:         sha256.Sum256([]byte("Hello, world")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlockChain{
				chainName: tt.fields.chainName,
				blocks:    tt.fields.blocks,
			}
			if got := bc.GetLatestBlock(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BlockChain.GetLatestBlock() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestNewBlockChain(t *testing.T) {
	type args struct {
		chainName    string
		genesisBlock Block
	}
	tests := []struct {
		name string
		args args
		want BlockChain
	}{
		{
			name: "New BlockChain Creation",
			args: args{
				chainName: "Test",
				genesisBlock: Block{
					index:        0,
					previousHash: sha256.Sum256([]byte("0")),
					timestamp:    0,
					data:         "",
					hash:         sha256.Sum256([]byte("Hello, world")),
				},
			},
			want: BlockChain{
				chainName: "Test",
				blocks: []Block{
					{
						index:        0,
						previousHash: sha256.Sum256([]byte("0")),
						timestamp:    0,
						data:         "",
						hash:         sha256.Sum256([]byte("Hello, world")),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBlockChain(tt.args.chainName, tt.args.genesisBlock); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBlockChain() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
