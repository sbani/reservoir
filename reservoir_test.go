package reservoir

import "testing"

func TestNewReservoir(t *testing.T) {
    rv := NewReservoir(34, 99)
  
    if rv.MaxConcurrent != 34 {
        t.Errorf("rv.MaxConcurrent: expected %d, actual %d", 34, rv.MaxConcurrent)
    }
    
    if rv.MinTime != 99 {
        t.Errorf("rv.MinTime: expected %d, actual %d", 99, rv.MinTime)
    }
}
