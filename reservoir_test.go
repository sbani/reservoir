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

    if rv.MaxQueueLength != 0 {
        t.Errorf("rv.MaxQueueLength: expected %d, actual %d", 0, rv.MaxQueueLength)
    }
    
    if len(rv.Queue) != 0 {
        t.Errorf("Queue is not empty! len(rv.Queue): expected %d, actual %d", 0, len(rv.Queue))
    }
    
    if rv.OverflowStrategy != StrategyLeak {
        t.Errorf("rv.OverflowStrategy: expected %d, actual %d", StrategyLeak, len(rv.Queue))
    }
}

func TestLimitQueue(t *testing.T) {
    rv := Reservoir{}
    rv.LimitQueue(99, StrategyOverflow)
    
    if rv.MaxQueueLength != 99 {
        t.Errorf("rv.MaxQueueLength: expected %d, actual %d", 99, rv.MaxQueueLength)
    }
    
    if rv.OverflowStrategy != StrategyOverflow {
        t.Errorf("rv.OverflowStrategy: expected %d, actual %d", StrategyOverflow, len(rv.Queue))
    }
}

func TestAdd(t *testing.T) {
    rv := NewReservoir(34, 99)
    fn := func(a, b, c int){}
    rv.Add(fn, 0, 0, 0)
    rv.Add(fn, 0, 0, 0)
    if len(rv.Queue) != 2 {
        t.Errorf("len(rv.Queue): expected %d, actual %d", 2, len(rv.Queue))
    }
}

// TODO: This has to check with different strategies too
func TestRun(t *testing.T) {
    rv := NewReservoir(34, 99)
    fn := func(a, b, c int){}
    rv.Add(fn, 0, 0, 0)
    rv.Add(fn, 0, 0, 0)
    rv.run()
    if len(rv.Queue) != 1 {
        t.Errorf("len(rv.Queue): expected %d, actual %d", 1, len(rv.Queue))
    }
}
