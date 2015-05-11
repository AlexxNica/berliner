package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"sync"
)

type Pipe struct {
	workers int         // max workers doing work on the pipeline
	do      interface{} // synchronous function to be performed on one
	// channel item.
	// should resemble type:
	// func(T, chan<- T)
	// unit of work is passed as first argument,
	// completed work unit(s) are sent out on channel in
	// second argument
	in  interface{} // input channel, should resemble type: <-chan T
	out interface{} // output channel, should resemble type: chan<- T
	wg  sync.WaitGroup
}

func (p *Pipe) pipe() error {
	// here there be dragons

	// is all of this worth it to abstract the piping?
	// I think so...

	// Recommended reading
	// http://blog.golang.org/laws-of-reflection
	// http://golang.org/pkg/reflect/

	// In order to accept a generic "do" function, we have to have to make "do"
	// the interface{} type.  This means that we have to perform typechecking
	// at runtime of the function signature that's actually passed...
	pdo := reflect.ValueOf(p.do)
	pdoType := pdo.Type()
	if pdoType.Kind() != reflect.Func ||
		pdoType.NumIn() != 2 || // number of arguments
		pdoType.NumOut() != 0 { // number of return values
		return errors.New(
			"*Pipe.do must be a function with 2 args and 0 return values.",
		)
	}
	pin := reflect.ValueOf(p.in)
	pinType := pin.Type()
	pout := reflect.ValueOf(p.out)
	poutType := pout.Type()
	if pinType.Kind() != reflect.Chan || poutType.Kind() != reflect.Chan {
		return errors.New("*Pipe.in and *Pipe.out must be of type chan.")
	}

	// TODO: implement better channel type checking
	// TODO: check channel item type checking

	// okay now we can get started

	// executing a pipe shouldn't block
	go func() {
		// we want to have n workers
		for i := 0; i < p.workers; i++ {
			// we need to close the out chan after all workers are
			// done, so we use a WaitGroup to keep track of them
			p.wg.Add(1)
			go func() {
				// since "do" function is synchronous, work will be done when
				// this anonymous function (which calls "do") exits
				defer p.wg.Done()

				// this would just be `for item := range p.in` if we didn't
				// have to deal with the interface{} stuff
				for {
					item, ok := pin.Recv()
					if ok != true {
						return
					}

					// this would just be `p.do(item, p.out)` if we didn't
					// have to deal with the interface{} stuff
					pdo.Call([]reflect.Value{item, pout})
				}
			}()
		}
		p.wg.Wait()
		// all workers are done working, so we can close the out channel now
		pout.Close()
	}()
	return nil
}

// utility function to read lines from stdin as a channel of strings
func readLines() <-chan string {
	out := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			out <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		close(out)
	}()
	return out
}
