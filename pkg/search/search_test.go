package search

import (
	"context"
	"log"
	"testing"
)

func TestAll_success(t *testing.T) {

	root := context.Background()
	files := []string{
		"../../data/fileA.txt",
		"../../data/fileB.txt",
		"../../data/fileC.txt",
	}
	channel := All(root, "pipeline", files)

	result, ok := <-channel
	if !ok {
		t.Errorf("Error_All(): %v", ok)
		return
	}
	log.Printf("result: %v\n", result)

	result, ok = <-channel
	if !ok {
		t.Errorf("Error_All(): %v", ok)
		return
	}
	log.Printf("result: %v\n", result)

	result, ok = <-channel
	if !ok {
		t.Errorf("Error_All(): %v", ok)
		return
	}
	log.Printf("result: %v \n", result)

}

func TestAll_notSuccess(t *testing.T) {

	root := context.Background()
	files := []string{""}
	channel := All(root, "pipeline", files)
	result, ok := <-channel
	if ok {
		t.Errorf("Error_All(): %v", ok)
		return
	}
	log.Println("result:", result)
}

func TestAny_success(t *testing.T) {
	root := context.Background()
	files := []string{"../../data/fileA.txt"}
	channel := Any(root, "pipeline", files)

	result, ok := <-channel
	if !ok {
		t.Errorf("Error_Any(): %v", ok)
		return
	}
	log.Printf("result: %v\n", result)
}

func TestAny_notSuccess(t *testing.T) {

	root := context.Background()
	files := []string{""}
	channel := Any(root, "pipeline", files)
	result, ok := <-channel
	if ok {
		t.Errorf("Error_All(): %v", ok)
		return
	}
	log.Println("result:", result)
}
