package reservoir

import (
    "time"
    "reflect"
)

type job struct {
    fn interface{} // Function to call
    args []interface{} // Arguments for the funtion call
}

type Reservoir struct {
    MaxConcurrent int // How many requests can be running at the same time. Default: 0 (unlimited)
    MinTime time.Duration // How long to wait after launching a request before launching another one. Default: 0ms
    Queue []job
}

func (rv *Reservoir) add(fn interface{}, args ...interface{}) {
    jb := job{fn, args}
    rv.Queue = append(rv.Queue, jb)
}

func (rv *Reservoir) start() {
    ticker := time.NewTicker(rv.MinTime)
    quit := make(chan bool)

    for {
       select {
        case <- ticker.C:
            if len(rv.Queue) > 0 {
                rv.run()
            }
        case <- quit:
            ticker.Stop()
            return
        }
    }
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

func NewReservoir(maxConcurrent int, minTime time.Duration) *Reservoir {
    rv := &Reservoir{MaxConcurrent: maxConcurrent, MinTime: minTime}
    go rv.start()
    return rv
}
