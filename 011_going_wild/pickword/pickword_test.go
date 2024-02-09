package pickword

import (
    "testing"
)

// TestPickTestDir
// Use the provided testdir and test data to confirm test pick is picking from files
func TestPickTestDir(t *testing.T) {
    dir := "./testdir"
    word, err := Pick(dir)
    if err != nil {
        t.Fatalf(`Error picking %v`, err)
    } 

    if word != "testrand1" && word != "testrand2" {
        t.Fatalf(`Picked test word not matched %v error %v`, word, err)
    }
}