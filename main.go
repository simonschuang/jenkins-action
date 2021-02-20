package main

import (
  "github.com/bndr/gojenkins"
  "context"
  "time"
  "fmt"
)

var jenkins_url = "https://example.com"
var username = "admin"
var password = "admin"

func main() {
    ctx := context.Background()
    jenkins := gojenkins.CreateJenkins(nil, jenkins_url, username, password)

    queueid, err := jenkins.BuildJob(ctx, "1. MLSteamBuilder", nil)
    if err != nil {
        panic(err)
    }
    build, err := jenkins.GetBuildFromQueueID(ctx, queueid)
    if err != nil {
        panic(err)
    }

    // Wait for build to finish
    for build.IsRunning(ctx) {
        time.Sleep(5000 * time.Millisecond)
        build.Poll(ctx)
    }

    fmt.Printf("build number %d with result: %v\n", build.GetBuildNumber(), build.GetResult())
}
