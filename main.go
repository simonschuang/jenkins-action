package main

import (
  "github.com/bndr/gojenkins"
  "context"
  "time"
  "fmt"
)

func main() {
ctx := context.Background()
jenkins := gojenkins.CreateJenkins(nil, "https://ci.myelintek.com", "simon", "11f9e47ffb5e23e7c03ea3a24cbeb76b0c")
// Provide CA certificate if server is using self-signed certificate
// caCert, _ := ioutil.ReadFile("/tmp/ca.crt")
// jenkins.Requester.CACert = caCert
// _, err := jenkins.Init(ctx)


//if err != nil {
//  panic("Something Went Wrong")
//}

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
