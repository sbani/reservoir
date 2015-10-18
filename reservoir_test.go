package reservoir

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewReservoir(t *testing.T) {
    rv := NewReservoir(34, 99)
  
    assert.Equal(t, 34, rv.MaxConcurrent)
    assert.Equal(t, 99, int(rv.MinTime))
    assert.Equal(t, 0, rv.MaxQueueLength)
    assert.Len(t, rv.Queue, 0)
    assert.Equal(t, StrategyLeak, rv.OverflowStrategy)
}

func TestLimitQueue(t *testing.T) {
    rv := Reservoir{}
    rv.LimitQueue(99, StrategyOverflow)
    
    assert.Equal(t, 99, rv.MaxQueueLength)
    assert.Equal(t, StrategyOverflow, rv.OverflowStrategy)
}

func TestAdd(t *testing.T) {
    rv := NewReservoir(34, 99)
    fn := func(a, b, c int){}
    rv.Add(fn, 0, 0, 0)
    rv.Add(fn, 0, 0, 0)
    assert.Len(t, rv.Queue, 2)
}

// TODO: This has to check with different strategies too
func TestRun(t *testing.T) {
    rv := NewReservoir(34, 99)
    fn := func(a, b, c int){}
    rv.Add(fn, 0, 0, 0)
    rv.Add(fn, 0, 0, 0)
    rv.run()
    assert.Len(t, rv.Queue, 1)
}
