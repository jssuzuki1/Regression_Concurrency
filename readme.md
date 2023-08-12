## Concurrency Documentation

Github Link: https://github.com/jssuzuki1/Regression_Concurrency 

There are two folders in the same directory as this readme.md file.

The first is "concurrency_version" and the second is "serial_version." At first glance, they both contain the same files: 
- boston.csv : our experimental data set. It is data originally from a housing study in 1978, studying the impact of various metrics on housing values. 
- main.go : the written program that runs two single variable linear regressions. One linear regression generates a relationship between crime and housing values and the other nox levels and housing values.
- main.exe : the executable load module built from main.go. Simply double-click to run it.
- elapsed_time.txt : the time metric output of main.go
- go.mod, go.sum: the files required to allow main.go to function

The key difference between these two folders lies in the main.go and, trivially, the main.exe file. In the serial_version, the main.go program runs 100 iterations of the two linear regressions in a serial manner (one at a time), then records the time. The main.go in the concurrency_version the same linear regressions, but uses concurrency, running multiple regressions at the same time via multiple goroutines and finishes them without regard to the progress of the other goroutines. In other words, the concurrency_version of the main.go program runs multiple iterations simultaneously.

## Results

Both of the main.go files output the amount of time elapsed during the operations of the go file. On my computer, the serial version takes 48.3513ms, while the concurrency_version takes 519µs, or 0.519 miliseconds. In other words, running this loop with concurrency takes almost a hundredth of the time as running the loop serially.

These programs do not require unit tests because these programs are effectivley single-unit benchmarks that either work or fail; they output an error if they do. 

## Management Recommendations

If you are faced with the precise situation of having to run many of the same regression many times, use concurrency. 
But, on a more serious note, the use of concurrency is situational. It is a powerful way of speeding processes up, as demonstrated above. However, its applications need to be carefully considered when faced with more complex programming problems.