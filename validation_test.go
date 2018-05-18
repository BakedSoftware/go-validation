package validation

import (
	"reflect"
	"sync"
	"testing"
)

func TestValidationMap_Atomicity(t *testing.T) {
	vm := Map{}
	typ := reflect.TypeOf(vm)
	wg1 := sync.WaitGroup{}
	wg1.Add(1)
	wg2 := sync.WaitGroup{}
	wg2.Add(2)
	count := 10000
	go func() {
		wg1.Wait()
		for i := 0; i < count; i++ {
			vm.set(typ, nil)
		}
		wg2.Done()
	}()
	go func() {
		wg1.Wait()
		for i := 0; i < count; i++ {
			vm.get(typ)
		}
		wg2.Done()
	}()
	wg1.Done() // start !
	wg2.Wait()
}
