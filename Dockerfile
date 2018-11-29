FROM busybox:glibc
COPY BookList-AppsCode /bin/booklist
EXPOSE 8080
ENTRYPOINT ["/bin/booklist"]
