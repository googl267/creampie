> [!WARNING]
> This is the dev branch where things are most likely not working! Please use Main

# Creampie!
The CLI tool to creampie your desktop with base16!

## how it works:
> [!NOTE]
> Creampie has a *lot* of redudancy to make ricing after setup as easy as possible. However it might be a lot at first
>

### config.toml
- contains the path and template for every file you want to theme
```
# Creampie a program named foo
[item]
path="~/.config/foo/foo.conf"  #path to actual config file
template="foo"                #template folder name
subtemplate="foo2"            #optional if multiple files are in the template folder
hook="foo --reload"           #optional hook to run after files have been modifyed
```

### schemes
- **Schemes MUST be base16!**
- `schemes/images/` needs at least a single image and serves as the last fallback (see the media section below)
- `schemes/[scheme-name]/` requires a .yaml file containing a base16 compatable theme. If you like to have theme specific images they can also be placed in this folder

### templates
- `templates/default/` required: used as a fallback
- `templates/default/config.yaml` you can define your own variables here. For example you can define `font:JetBrainsMono` then you can use {{font}} in your templates
- `templates/default/templates/` Create a folder named after your program, inside create .mustache files with your config. This is where you can add {{foo}} variables
- `templates/foo/` setup exactly the same as `default/` but you dont need to add every single program. your `config.yaml` file will be applied over the defaults one however

### media
the order images are selected:
> .../templates/[current_template]/images/[dark/light]/
> 
> .../schemes/[current_scheme]/
> 
> .../schemes/images/
>
- a random image will be selected if multiple are present
- {{img}} will use the filepath of the selected image. {{img-foo}} will look for a image specifcally named foo.img, otherwise fallback to {{img}}

#### Example File Structure:
```bash
creampie/
├── config.toml
├── schemes/
│   ├── images/
│   │   ├── foo.img
│   │   ├── foo1.img
│   │   └── foo2.img
│   └── foo/
│       ├── foo.yaml
│       ├── foo1.img
│       └── foo2.img
└── templates/
    ├── default/
    │   ├── options.yaml
    │   └── templates/
    │       ├── foo/
    │       │   └── foo.mustache
    │       └── foo1/
    │           └── foo.mustache
    └── foo/
        ├── options.yaml
        ├── templates/
        │   └── foo/
        │       └── foo.mustache
        └── images/
            ├── light/
            │   └── foo.img
            └── dark/
                └── foo.img
```
and adding a new layout is as easy as
```
templates/
+ └── my-cool-layout/
      └── config.yaml
```

## roadmap
- [ ] rewrite files with colors
- [ ] image support
- [ ] can select scheme
- [ ] can select layout
- [ ] random scheme or layout
- [ ] hook support
- [ ] user defined variables
- [ ] error check before modifying files

## todo
- [ ] impliment replacing parts of file instead of force rewriting
- [ ] heavy hooks toggle

# Credit
This entire project is ~~stolen~~ inspired by [Flavours](https://github.com/Misterio77/flavours)! While I love Flavours and used it for a while, I needed some extra features and prefer to have my folder structure a little diffrent. The main reason really tho is to teach and familiarize myself with GO
