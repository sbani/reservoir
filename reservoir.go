package reservoir

import (
    "time"
    "reflect"
)

const (
    StrategyLeak = 1 // Removes the first (oldest) element in queue and adds the new one
    StrategyOverflow = 2 // The added element will be lost.
)

type job struct {
    fn interface{} // Function to call
    args []interface{} // Arguments for the funtion call
}

type Reservoir struct {
    MaxConcurrent int // How many requests can be running at the same time. Default: 0 (unlimited)
    MinTime time.Duration // How long to wait after launching a request before launching another one. Default: 0ms
    MaxQueueLength int
    OverflowStrategy int
    Queue []job
    stop chan bool
}

// Set the queue limit and the strategy whats happens if queue is full
func (rv *Reservoir) LimitQueue(maxQueueLength int, overFlowStrategy int) {
    rv.MaxQueueLength = maxQueueLength
    rv.OverflowStrategy = overFlowStrategy
}

// Add a new call to the queue
func (rv *Reservoir) Add(fn interface{}, args ...interface{}) {
    if !rv.handleStrategy() {
        return
    }
    jb := job{fn, args}
    rv.Queue = append(rv.Queue, jb)
}

// Check for strategy and returns if call should be added
func (rv *Reservoir) handleStrategy() bool {
    // No max length. No need to check for strategy
    if rv.MaxQueueLength == 0 || len(rv.Queue) < rv.MaxQueueLength {
        return true
    }
    
    if rv.OverflowStrategy == StrategyLeak {
        rv.Queue = rv.Queue[1:]
        return true
    }
    
    if rv.OverflowStrategy == StrategyOverflow {
        return false
    }
    
    // Fallback (dont know the strategy - "ignore" queue max)
    return true
}

// Start working on the queue
func (rv *Reservoir) Start() {
    ticker := time.NewTicker(rv.MinTime)

    for {
       select {
        case <- ticker.C:
            if len(rv.Queue) > 0 {
                rv.run()
            }
        case <- rv.stop:
            ticker.Stop()
            return
        }
    }
}

// Stop working no the queue
func (rv *Reservoir) Stop() {
    rv.stop <- true
}

// Run the oldest job in list and remove it
func (rv *Reservoir) run() {
    job := rv.Queue[0]
    fn := reflect.ValueOf(job.fn)
    in := make([]reflect.Value, len(job.args))
    for k, param := range job.args {
        in[k] = reflect.ValueOf(param)
    }
    fn.Call(in)
    // remove the first/oldest element
    rv.Queue = rv.Queue[1:]
}

// Create a new reservoir struct and start working the queue
func NewReservoir(maxConcurrent int, minTime time.Duration) *Reservoir {
    rv := &Reservoir{
        MaxConcurrent: maxConcurrent, 
        MinTime: minTime,
        OverflowStrategy: StrategyLeak,
    }
    go rv.Start()
    return rv
}
