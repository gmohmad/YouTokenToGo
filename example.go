package main

/*

#cgo LDFLAGS: -L/usr/local/lib -lbpewrapper
#include <stdlib.h>
#include </usr/local/include/youtokentogo/wrapper.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

const cArraySize = 1 << 28

// BPEEncoder wraps the BaseEncoder C++ class
type BPEEncoder struct {
	ptr unsafe.Pointer
}

// NewBPEEncoder creates a new BPEEncoder instance
func NewBPEEncoder(modelPath string, nThreads int) (*BPEEncoder, error) {
	cPath := C.CString(modelPath)
	defer C.free(unsafe.Pointer(cPath))

	ptr := C.new_base_encoder(cPath, C.int(nThreads))
	if ptr == nil {
		return nil, fmt.Errorf("failed to create BaseEncoder")
	}
	return &BPEEncoder{ptr: ptr}, nil
}

// EncodeAsIDs calls C.encode_as_ids() function and returns result as an int slice
func (b *BPEEncoder) EncodeAsIDs(sentence string) ([]int, error) {
	cSentence := C.CString(sentence)
	defer C.free(unsafe.Pointer(cSentence))

	var cIDs *C.int
	var length C.int

	result := C.encode_as_ids(b.ptr, cSentence, &cIDs, &length)
	if result != 0 {
		return nil, fmt.Errorf("failed to encode sentence")
	}
	defer C.free_ids(cIDs)

	cIds := (*[cArraySize]C.int)(unsafe.Pointer(cIDs))[:length:length]
	ids := cIntArrayToGoIntSlice(cIds)

	return ids, nil
}

// Converts c int array to golang int slice
func cIntArrayToGoIntSlice(cArray []C.int) []int {
	ids := make([]int, 0, len(cArray))

	for _, cInt := range cArray {
		ids = append(ids, int(cInt))
	}

	return ids
}

// Close cleans up the BPEEncoder instance
func (b *BPEEncoder) Close() {
	C.destroy_base_encoder(b.ptr)
}

func main() {
	// Path to the BPE model
	modelPath := "yttm.model"

	// Initialize BaseEncoder
	encoder, err := NewBPEEncoder(modelPath, 1)
	if err != nil {
		panic(err)
	}
	defer encoder.Close()

	// Encode a sentence
	sentence := "some text"
	ids, err := encoder.EncodeAsIDs(sentence)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encoded IDs:", ids)
}
