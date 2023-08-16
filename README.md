# starship-container-integration

A Go tool that prepends the name of the current container to the starship prompt

## Installation

Installing begins by cloning this repo onto your system. If you are using [toolbx](https://containertoolbx.org/), you only have to clone it on your base system, somewhere under your home directory. Podman or Docker containers will have to have the repo cloned into their files individually, and the setup steps will have to be performed again as well.

Now add the following to your `~/.bashrc` so that the correct file is sourced (and that the program is compiled and run if it hasn't been yet) every time you open a new shell:

```
if [ -f /run/.containerenv ]; then
    export STARSHIP_CONFIG=PATH_TO_REPO_DIR/starship.toml
    cd PATH_TO_REPO_DIR && make && ./main && cd
fi
```

Now run `source ~/.bashrc` or restart your shell and you should see the name of your container included at the front of your prompt!


## Roadmap

- [ ] Test with different shells
- [ ] Add CI/CD test cases
- [ ] Add customization instructions/flags

