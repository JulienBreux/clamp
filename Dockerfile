FROM scratch

COPY clamp /bin/clamp

ENTRYPOINT [ "/bin/clamp" ]
