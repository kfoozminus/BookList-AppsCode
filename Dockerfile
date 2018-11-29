FROM frolvlad/alpine-glibc
COPY BookList-AppsCode /bin/booklist
EXPOSE 4321
CMD ["--port=4321"]
ENTRYPOINT ["booklist"]
