// Package bitvec is bit-vector with atomic and non-atomic access
package bitvec

const (
	nbits   = 5          // 32 bits in a uint32
	ws      = 1 << nbits // constant 64
	mask    = ws - 1     // all ones
	bitsize = 2 ^ nbits
)

// BitVec is a nonatomic bit vector.
type BitVec []uint32

// NewBitVec creates a non-atomic bitvector.
func NewBitVec(size uint32) BitVec {
	nints := size / ws
	if size-(nints*bitsize) != 0 {
		nints++
	}

	return make(BitVec, nints)
}

func (BitVec) offset(k uint32) (bucket, bit uint32) {
	return k >> nbits, 1 << (k & mask)
}

// TrySet will try to set the bit and will return true if set
// is successful.
func (bv BitVec) TrySet(k uint32) bool {
	bucket, bit := bv.offset(k)
	old := bv[bucket]
	if old&bit != 0 {
		return false
	}
	bv[bucket] = old | bit
	return true
}
