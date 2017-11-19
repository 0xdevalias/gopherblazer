package main

import (
  "encoding/json"

  "os"
  "os/exec"
  "fmt"
  "bufio"

  "github.com/apex/go-apex"
)

// type message struct {
//   Hello string `json:"hello"`
// }

func execCommand(cmd *exec.Cmd) (string, error) {
  stdout, stderr := cmd.StdoutPipe()

  if err := cmd.Start(); err != nil {
    return "", err
  }

  scanner := bufio.NewScanner(stdout)
  strOut := "[stdout]"
  for scanner.Scan() {
    // fmt.Println(scanner.Text())
    strOut += fmt.Sprintf("\n%s", scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "reading standard input:", err)
    return "", err
  }

  cmd.Wait()

  return strOut, stderr
}

func gobuster(url string) (string, error) {
  cmd := exec.Command("./bin/gobuster", "-w", "words.txt", "-u", url)
  return execCommand(cmd)
}

func main() {
  apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {

    out, err := gobuster("http://test-discovery.gopherblazer.devalias.net")
    // var m message
    // foo := "This is a foo"

    // if err := json.Unmarshal(event, &m); err != nil {
    //   return nil, err
    // }

    // cmd := exec.Command("ls")
    // stdout, stderr := cmd.StdoutPipe()

    // if err := cmd.Start(); err != nil {
    //   return nil, err
    // }

    // _ := cmd.Wait()
    // cmd.Wait()

    // Convert stdout to a string
    // buf := new(bytes.Buffer)
    // buf.ReadFrom(stdout)
    // strStdout := buf.String()
    // strStdout := reader2str(stdout)

    // Convert stderr to a string
    // buf := new(bytes.Buffer)
    // buf.ReadFrom(stderr)
    // strStderr := buf.String()

    // retVal := fmt.Sprintf("[stdout]\n%s\n[stderr]\n%s", out, err.Error())

    return out, err
  })
}
