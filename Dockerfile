FROM busybox:glibc
COPY BookList-AppsCode /bin/booklist
ENTRYPOINT ["/bin/booklist"]
