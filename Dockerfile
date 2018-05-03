FROM scratch

ADD bver /

# Create dummy logfile
ADD access.log /var/log/access.log

ENTRYPOINT ["/bver"]
