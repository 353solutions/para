# Go Workshop

Miki Tebeka
📬 [miki@353solutions.com.com](mailto:miki@353solutions.com), 𝕏 [@tebeka](https://twitter.com/tebeka), 👨 [mikitebeka](https://www.linkedin.com/in/mikitebeka/), ✒️[blog](https://www.ardanlabs.com/blog/)

#### Shameless Plugs

- [LinkedIn Learning Classes](https://www.linkedin.com/learning/instructors/miki-tebeka)
- [Books](https://pragprog.com/search/?q=miki+tebeka)

[Syllabus](_extra/syllabus.pdf)


---

## Session 1: RPC

### Agenda

- Advanced JSON
    - Custom serialization
    - Missing vs empty values
    - Streaming JSON
- HTTP clients
    - Request body
    - Streaming
    - Authentication
- HTTP servers
    - Dependency injection
    - Writing middleware
    - Streaming responses
    - Routing


### Code


- [cars](session_1/cars/) - HTTP Server

- [value.go](session_1/value/value.go) - Custom serialization
- [vm.go](session_1/vm/vm.go) - Missing vs Zero values
- [logs.go](session_1/logs/logs.go) - Streaming JSON
- [closure.go](session_1/closure/closure.go) - Closure capture bug (Go 1.22)
- [client.go](session_1/events/client.go) - Making HTTP calls

### Links

- [The complete guide to Go net/http timeouts](https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/)
- [Server Side Events in Go](https://medium.com/@rian.eka.cahya/server-sent-event-sse-with-go-10592d9c2aa1)
- [httptest](https://pkg.go.dev/net/http/httptest) - Test your HTTP handlers
- [How SQLite is Tested](https://www.sqlite.org/testing.html)
- [Managing Go Installations](https://go.dev/doc/manage-install)
- Web Frameworks
    - [chi](https://go-chi.io/)
    - [Gin](https://gin-gonic.com/)
    - [fasthttp](https://github.com/valyala/fasthttp)
- [How I write HTTP services in Go after 13 years](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)
- [Making & Using HTTP Middleware](https://www.alexedwards.net/blog/making-and-using-middleware)
- Dependency injection:
    - [Uber fx](https://github.com/uber-go/fx)
    - [Google wire](https://github.com/google/wire)
- [Year 2038 problem](https://en.wikipedia.org/wiki/Year_2038_problem)
- [golang.org/x/analysis](https://pkg.go.dev/golang.org/x/tools/go/analysis) - Write a linter
- [mapstructure](https://pkg.go.dev/github.com/mitchellh/mapstructure#example-Decode)
- Validation
    - [cue](https://cuelang.org/docs/introduction/)
    - [validator](https://github.com/go-playground/validator)
- [NYC Taxi Dataset](https://www.nyc.gov/site/tlc/about/tlc-trip-record-data.page)
- [Email Address Regexp](https://emailregex.com/index.html)
- [JSON Lines](https://jsonlines.org/)
- [HTTP Status Cats](https://http.cat/)
- [Unicode Folding](https://www.unicode.org/L2/L2000/00261-tr25-0d1.html)
- [Fixing For Loops in Go 1.22](https://go.dev/blog/loopvar-preview)
- [JSON - The Fine Print](https://www.ardanlabs.com/blog/2024/10/json-the-fine-print-part-1.html)
- [HTTP cats](https://http.cat/)

### Data & Other

- [event.json](_extra/event.json)
- [Syllabus](_extra/syllabus.pdf)

---
## Session 2: Going Faster

### Agenda

- Benchmarking & profiling
    - tokenizer
- Performance tips & tricks
- Optimizing memory

### Code

- [tokenizer](session_2/tokenizer) - Optimizing algorithm & memory allocations
- [slice.go](session_2/slice/slice.go) - How slices work
- [playground.go](playground/playground.go) - Misc
- [store](session_2/store) - Cache & better serialization
- [matrix.go](session_2/matrix/matrix.go) - Getting friendly with the CPU cache

### Links

- [Trie](https://en.wikipedia.org/wiki/Trie)
- [K-D tree](https://en.wikipedia.org/wiki/K-d_tree)
- [lru-cache](https://pkg.go.dev/github.com/hashicorp/golang-lru/v2)
- [The Architecture of Open Source Applications](https://aosabook.org/en/) - Including a book on performance
- [Plain Text](https://www.youtube.com/watch?v=4mRxIgu9R70) - Fun talk about Unicode
- [Regular Expression Matching Can Be Simple And Fast](https://swtch.com/~rsc/regexp/regexp1.html)
- [Locality of Behaviour](https://htmx.org/essays/locality-of-behaviour/)
- [Rules of Optimization Club](https://wiki.c2.com/?RulesOfOptimizationClub)
- [Computer Latency at Human Scale](https://twitter.com/jordancurve/status/1108475342468120576)
- [So you wanna go fast](https://www.slideshare.net/TylerTreat/so-you-wanna-go-fast-80300458)
- [High Performance Go](https://dave.cheney.net/high-performance-go-workshop/gophercon-2019.html)
- [Miki's Optimization Overview](_extra/optimize.md)
- [A Benchmarking Checklist](https://www.brendangregg.com/blog/2018-06-30/benchmarking-checklist.html)
- [A Guide to the Go Garbage Collector](https://tip.golang.org/doc/gc-guide)
- [hey](https://github.com/rakyll/hey)
- [Garbage Collection In Go : Part I - Semantics](https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html)
- [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat)

---

# Day 3: Advanced Concurrency

### Agenda

- Channel semantics
- Goroutine pools
- The "sync" & "sync/atomic" packages
- Handling panics in goroutines

### Code


- [go_chan.go](session_3/go_chan/go_chan.go) - Channel semantics, fan-out, goroutine pool
- [rtb.go](session_3/rtb/rtb.go) - Context & cancellation
- [taxi.go](session_3/taxi/taxi.go) - Convert serial code to concurrent
- [fan_in.go](session_3/fan_in/fan_in.go) - Fan-in pattern


- [counter.go](session_3/counter/counter.go) - race detector, sync.Mutex & atomic
- [token.go](session_3/token/token.go) - Refresh token, sync.RWMutex
- [token_chan.go](session_3/token_chan/token_chan.go) - Refresh token with channels
- [pmap.go](session_3/pmap/pmap.go) - Parallel map
- [pool.go](session_3/pool/pool.go) - Using buffered channel for pool
- [payment.go](session_3/payment/payment.go) - sync.Once

### Links

- [automaxprocs](https://pkg.go.dev/go.uber.org/automaxprocs@v1.6.0/maxprocs)
- [x/time/rate](https://pkg.go.dev/golang.org/x/time/rate) - Rate limiter
- [The race detector](https://go.dev/doc/articles/race_detector)
- [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)
- [Data Race Patterns in Go](https://eng.uber.com/data-race-patterns-in-go/)
- [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
- [Go Concurrency Patterns: Context](https://go.dev/blog/context)
- [Curious Channels](https://dave.cheney.net/2013/04/30/curious-channels)
- [The Behavior of Channels](https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html)
- [Channel Semantics](https://www.353solutions.com/channel-semantics)
- [Why are there nil channels in Go?](https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308)
- [Amdahl's Law](https://en.wikipedia.org/wiki/Amdahl%27s_law) - Limits of concurrency
- [Computer Latency at Human Scale](https://twitter.com/jordancurve/status/1108475342468120576/photo/1)
- [Concurrency is not Parallelism](https://www.youtube.com/watch?v=cN_DpYBzKso) by Rob Pike
- [Scheduling in Go](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html) by Bill Kennedy
- [conc: better structured concurrency for go](https://github.com/sourcegraph/conc)
- [xkcd Tar](https://xkcd.com/1168/)
- [Profile-guided optimization](https://go.dev/doc/pgo)
- [Rubber Duck Debugging](https://en.wikipedia.org/wiki/Rubber_duck_debugging)
- [Feynman Algorithm](https://wiki.c2.com/?FeynmanAlgorithm)
- [conc](https://github.com/sourcegraph/conc)
- [Using Formal Reasoning to Build Concurrent Go Systems](https://www.youtube.com/watch?v=yiVOJqXTWfc) video
- [TLA+](https://learntla.com/#)
- [golang.org/x/sync/semaphore](https://pkg.go.dev/golang.org/x/sync/semaphore)
- [goleak](https://github.com/uber-go/goleak) - Find goroutines leak
- [htmx](https://htmx.org/)
- [Ultimate Go Tour](https://tour.ardanlabs.com/tour/eng/list)


### Data & More

- [rtb.go](_extra/rtb.go)
- [taxi.tar](https://storage.googleapis.com/353solutions/c/data/taxi.tar)

---

## Session 4: OO Patterns

### Agenda

- Pointer vs value semantics
- Embedding structs
- Interfaces in depth
- The empty interface and type assertions
- Iterators

### Code


- [inc.go](session_4/inc/inc.go) - Value vs Pointer sematics
- [game.py](session_4/game/game.py) - Structs, methods & interfaces
- [error.go](session_4/error/error.go) - What's a nil interface
- [sha1.go](session_4/sha1/sha1.go) - Calculate sha1, compose io.Reader & io.Writer
- [wc.go](session_4/wc/wc.go) - Implement io.Writer for word count


- [empty.go](session_4/empty/empty.go) - The empty interface (`any`)
- [logger.go](session_4/logger/logger.go) - Keeping interface small
- [client_test.go](session_4/client/client_test.go) - Mocking HTTP transport
    - [client.go](session_4/client/client.go)
- [stats.go](session_4/stats/stats.go) - Generics
- [iter.go](session_4/iter/iter.go) - Iterators

### Links

- [sort examples](https://pkg.go.dev/sort/#pkg-examples) - Read and try to understand
- [When to use generics](https://go.dev/blog/when-generics)
- [Generics tutorial](https://go.dev/doc/tutorial/generics)
- [Generic Interfaces](https://go.dev/blog/generic-interfaces)
- [Generics can make your Go code slower](https://planetscale.com/blog/generics-can-make-your-go-code-slower)
- [Methods, interfaces & embedded types in Go](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html)
- [Methods & Interfaces](https://go.dev/tour/methods/1) in the Go tour
- [wc docs](https://www.gnu.org/software/coreutils/manual/html_node/wc-invocation.html#wc-invocation)
- [stringer command](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
- [List of file signatures](https://en.wikipedia.org/wiki/List_of_file_signatures)
- [The Art of Unix Programming](http://www.catb.org/esr/writings/taoup/html/ch01s06.html)
- [Method Sets](https://www.youtube.com/watch?v=Z5cvLOrWlLM)
- [lru_cache](https://pkg.go.dev/github.com/hashicorp/golang-lru/v2)
- [From Nand to Tetris](https://www.nand2tetris.org/)
- [bubbletea](https://github.com/charmbracelet/bubbletea) UI for the terminal (TUI)
    - [huh](https://github.com/charmbracelet/huh)
    - [crush](https://github.com/charmbracelet/crush)


## Session 5: Project Engineering

### Agenda

- Creating go executables  
  - Injecting version  
  - Embedding assets  
- Configuration & command line parsing  
- Logging & metrics  
- Writing secure Go code

### Code

TBD

### Links

- [Using ldflags to Set Version Information for Go Applications](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications)
- [GoReleaser](https://goreleaser.com/)
    - [GitHub Action](https://github.com/goreleaser/goreleaser-action)
- [svu](https://github.com/caarlos0/svu) - Bump version
- [Using Zig to Compile Cgo](https://github.com/goreleaser/example-zig-cgo)
- Command line & Options
    - [flag](https://pkg.go.dev/flag)
    - [Cobra](https://cobra.dev/) + [Viper](https://github.com/spf13/viper)
    - Ardan Labs [conf](https://pkg.go.dev/github.com/ardanlabs/conf/v3)
- Security
    - [mkcert](https://github.com/FiloSottile/mkcert)
    - [x/crypto/autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert)
    - [Using Let's Encrypt in Go](https://marcofranssen.nl/build-a-go-webserver-on-http-2-using-letsencrypt)
    - [Customizing Binaries with Build Tags](https://www.digitalocean.com/community/tutorials/customizing-go-binaries-with-build-tags)
- [tar joke](https://xkcd.com/1168/)
- [Reversim Summit](https://summit2025.reversim.com/) - October 27,28
- [Writing Deployable Code](https://medium.com/@rantav/writing-deployable-code-part-one-13ec6dc90adb)
- [The Twelve Factor App](https://12factor.net)
- [TOML format](https://toml.io/en/)
- [Knight Capital Group](https://en.wikipedia.org/wiki/Knight_Capital_Group) - Cost of configuration error
- [The Art of Unix Programming](https://cdn.nakamotoinstitute.org/docs/taoup.pdf) PDF
- [Generating Code](https://go.dev/blog/generate) - `go:generate`
- [embed](https://pkg.go.dev/embed) package
- [Debug a Go Application Running on Kubernetes](https://www.youtube.com/watch?v=YXu2box7z9k)
- [Core Dump Debugging](https://go.dev/wiki/CoreDumpDebugging)
- [Leek & Seek](https://www.youtube.com/watch?v=94wG_LJH86U) - A lot of diagnostic advices and tools
- [Go: Monitor your goroutine](https://medium.com/@hax.artisan/go-monitor-your-goroutine-application-9edbdd6e581b)
- [Diagnostics](https://go.dev/doc/diagnostics) in the Go docs

- [Structured Logging with slog](https://go.dev/blog/slog)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)


### Data & Other

- [gopher.txt](_extra/gopher.txt)
- [Secure Code Slides](_extra/secure-go.pdf)
- [journal.tar.gz](_extra/journal.tar.gz)

![](https://pixel-73339669570.me-west1.run.app/p/para1/p.png)
