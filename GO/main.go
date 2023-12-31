package main

import (
    "fmt"
    "os"
    "math/rand"
    "time"

    "github.com/adrg/xdg" //xdg.DataHome, xdg.ConfigHome, xdg.ConfigDirs
)

func checkExists (path string) bool {
    _, err := os.Stat(path)
    if err != nil {
        return false
    }
    return true
}


func getRandomFile(path string) string {
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))

    f, err := os.Open(path)
    checkNilErr(err)
    defer f.Close()

    files, err := f.Readdirnames(-1)
    checkNilErr(err)
   
    return path + "/" + files[rng.Intn(len(files))]
}


func stringInSlice(a string, list []string) (int, bool) {
    i := 0
    for _, b := range list {
        if b == a {
            return i, true
        }
        i++
    }
    return 0, false
}





func checkNilErr(err error){
    if err != nil {
        panic(err)
    }
}


func loopConfPaths(relativePath string) string {
    if checkExists(xdg.DataHome + "/creampie" + relativePath) == true {  // check if ~/.local/share/creampie 
        return xdg.DataHome + relativePath
    } else if checkExists(xdg.ConfigHome + "/creampie" + relativePath) == true {
        return xdg.ConfigHome + relativePath
    } else if checkExists(xdg.Home + "/Documents/GitHub/creampie/config" + relativePath) { // placeholder, remove in production 
        return xdg.Home + "/Documents/GitHub/creampie/config" + relativePath
    }else {
        panic("\"" + relativePath + "\"" + " not found in any config path(" + xdg.DataHome + "/creampie , " + xdg.ConfigHome + "/creampie )")
    }
}



func main() {
    // user input placeholder
    const schemeName = "foo"
    const layoutName = "foo"

    //const variables
    configFile := loopConfPaths("/config.toml")
    schemePath := loopConfPaths("/schemes/" + schemeName)
    fallbackImgPath := loopConfPaths("/images")
    defaultLayoutPath := loopConfPaths("/templates/default")
    userLayoutPath := loopConfPaths("/templates/" + layoutName)

    fmt.Println("Random Image:", getRandomFile(schemePath + "/media"))

    fmt.Println(configFile)
    fmt.Println(schemePath)
    fmt.Println(fallbackImgPath)
    fmt.Println(defaultLayoutPath)
    fmt.Println(userLayoutPath)
}    
