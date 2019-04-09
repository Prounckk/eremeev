# export some git sugnature
export GPG_TTY=$(tty)

#tmux color
export TERM="xterm-256color"

# Set name of the theme to load. Optionally, if you set this to "random"
# it'll load a random theme each time that oh-my-zsh is loaded.
# See https://github.com/robbyrussell/oh-my-zsh/wiki/Themes
plugins=(
    git
    zsh-autosuggestions
    jira
    git-extras
    docker
    httpie
    buffalo
    sudo
    osx
    )

autoload -U compinit && compinit

ZSH_THEME="powerlevel9k/powerlevel9k"
