package time_monitoring

import (
	"fmt";
	"time";
)

// Optional pareamenters?
// See http://stackoverflow.com/questions/2032149/optional-parameters
// http://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

func timed(f func(args []string)) func(args []string) {

	return func(args []string) {
		p := fmt.Println
		start := time.Now()
		p("Started at ", start)
		f(args)
		finish := time.Now()
		fmt.Println("Finished at ", finish)
		p("Elapsed: ", finish.Sub(start))
	}

}