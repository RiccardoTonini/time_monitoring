package time_monitoring

import (
	"fmt";
	"time";
	// "reflect";
)

// Optional params?
// See http://stackoverflow.com/questions/2032149/optional-parameters
// http://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis

func Timed(f func(args []string)) func(args []string) {

	return func(args []string) {
		p := fmt.Println
		start := time.Now()
		p("Started at ", start)
		f(args)
		finish := time.Now()
		fmt.Println("Finished at ", finish)
		elapsed := finish.Sub(start)
		p("Elapsed: ", elapsed.Seconds())
	}

}