package main

import (  
    "fmt"
    "strings"
    "encoding/json"
    "syscall/js"

    "github.com/hashicorp/nomad/api"
    "github.com/hashicorp/nomad/jobspec2"
)

func main() {  
    fmt.Println("Go Web Assembly")
    js.Global().Set("formatJSON", jsonWrapper()) 
    <-make(chan bool)
}

func prettyJson(input string) (string, error) {  
    var raw interface{}
    if err := json.Unmarshal([]byte(input), &raw); err != nil {
        return "", err
    }
    pretty, err := json.MarshalIndent(raw, "", "  ")
    if err != nil {
        return "", err
    }
    return string(pretty), nil
}

func jsonWrapper() js.Func {  
    jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        if len(args) != 1 {
            result := map[string]interface{}{
                "error": "Invalid no of arguments passed",
            }
            return result
        }
        jsDoc := js.Global().Get("document")
        if !jsDoc.Truthy() {
            result := map[string]interface{}{
                "error": "Unable to get document object",
            }
            return result
        }
        jsonOuputTextArea := jsDoc.Call("getElementById", "jsonoutput")
        if !jsonOuputTextArea.Truthy() {
            result := map[string]interface{}{
                "error": "Unable to get output text area",
            }
            return result
        }
        inputHCL := args[0].String()
        fmt.Printf("input %s\n", inputHCL)
        pretty, err := parseHCL(inputHCL)
        fmt.Printf("output %s\n", pretty)

        if err != nil {
            errStr := fmt.Sprintf("unable to parse HCL. Error %s occurred\n", err)
            result := map[string]interface{}{
                "error": errStr,
            }
            return result
        }
        jsonOuputTextArea.Set("textContent", pretty)
        return nil
    })
    return jsonfunc
}


func parseHCL(jobspec string) (string, error) {
    // create a destination job. these default values should
    // be stomped by the incoming ones.
    myJob := api.NewServiceJob("test", "test", "global", 50)
    myJob, err := jobspec2.Parse("", strings.NewReader(jobspec))
    if err != nil {
        return "", err
    }
    out, err := json.MarshalIndent(myJob, "", "  ")
    if err != nil {
        return "", err
    }
    return string(out), nil
}

