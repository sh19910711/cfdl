FROM golang:1.9

RUN apt update && apt install -y \
	vim \
	git

RUN git clone https://github.com/sh19910711/dotfiles.git && \
  bash dotfiles/setup.bash
