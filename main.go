package main

import (
    "strings"
    util "github.com/starship-container-integration/pkg/util"
)

func main() {
    // Read the config file template
    config, _:= util.ReadFile("./starship_template.toml")
    // Get the data from /run/.containerenv
    container_env, err := util.ReadFile("/run/.containerenv")
    // If there is an error, program is likely not running inside of a container
    if err != nil {
        return
    }

    // Get fields of containerenv file
    containerenv_fields := strings.Fields(container_env) 

    if len(containerenv_fields) > 0 {
        name, _ := strings.CutPrefix(containerenv_fields[1], "name=\"")
        name, _ = strings.CutSuffix(name, "\"")

        // Replace "container_name" in config with actual container name
        config = strings.Replace(config, "container_name", name, 1)
    } else {
        // Replace "contianer_name" in config with nothing
        config = strings.Replace(config, "container_name", "", 1)
    }

    // Write config to starship.toml
    util.WriteFile("./starship.toml", config)
}
