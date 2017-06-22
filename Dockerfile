FROM scratch
ADD release/git-notify /
CMD ["/git-notify"]
