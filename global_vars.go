package main

import (
	"math/rand"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type status struct {
	completed int
	errored   int

	lastLinesWritten int
}

type globalConfiguration struct {
	verboseOutput  bool
	requestTimeout time.Duration
	command        string
}

type completeConfiguration struct {
	maxConcurrency   int
	waitBetweenPosts time.Duration

	target       int
	interviewURL string
}

var currentStatus *status
var completeConfig *completeConfiguration
var globalConfig *globalConfiguration

var (
	requestTimeoutFlag = kingpin.Flag("request-timeout", "Timeout on requests").Default("30s").Duration()
	verboseOutputFlag  = kingpin.Flag("verbose", "Enable verbose output for debugging purposes").Short('v').Default("false").Bool()

	completeCommand              = kingpin.Command("complete", "Complete interviews")
	completeMaxConcurrencyFlag   = completeCommand.Flag("concurrency", "Maximum number of concurrent interviews").Short('c').Default("10").Int()
	completeWaitBetweenPostsFlag = completeCommand.Flag("wait-time", "Wait time between answering questions").Default("0").Duration()
	completeTargetArg            = completeCommand.Arg("count", "The number of completes to generate.").Required().Int()
	completeInterviewURLArg      = completeCommand.Arg("url", "The url to the interview to complete.").Required().String()

	recordCommand         = kingpin.Command("record", "Record an interview for later playback")
	recordOutputFlag      = recordCommand.Flag("output-file", "Output file to write the recording to").Short('o').Default("recording.ipr").String()
	recordInterviewURLArg = recordCommand.Arg("url", "The url to the interview to complete.").Required().String()
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

var errorChannel = make(chan error)

const endOfInterviewPath = "/Home/Completed"
