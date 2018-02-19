package blockchain

import (
	"crypto/sha256"
	"reflect"
	"testing"
	"time"
)

func TestNewBlock(t *testing.T) {
	type args struct {
		index        int
		previousHash [sha256.Size]byte
		timestamp    int64
		data         string
		hash         [sha256.Size]byte
	}
	now := time.Now().Unix()
	tests := []struct {
		name string
		args args
		want Block
	}{
		{
			name: "Basic",
			args: args{
				index:        0,
				previousHash: [32]byte{},
				timestamp:    0,
				data:         "",
				hash:         [32]byte{},
			},
			want: Block{
				index:        0,
				previousHash: [32]byte{},
				timestamp:    0,
				data:         "",
				hash:         [32]byte{},
			},
		},
		{
			name: "With SHA256 Hashes and Unix Timestamps",
			args: args{
				index:        12,
				previousHash: sha256.Sum256([]byte("hash_of_block_11")),
				timestamp:    now,
				data:         "Hello, world",
				hash:         sha256.Sum256([]byte("test")),
			},
			want: Block{
				index:        12,
				previousHash: sha256.Sum256([]byte("hash_of_block_11")),
				timestamp:    now,
				data:         "Hello, world",
				hash:         sha256.Sum256([]byte("test")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBlock(tt.args.index, tt.args.previousHash, tt.args.timestamp, tt.args.data, tt.args.hash); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBlock() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
