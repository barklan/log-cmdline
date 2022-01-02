# log-cmdline

### Steps:

1. Edit `/etc/profile` and add the following lines to the bottom of the file:

```bash
function log2syslog
{
   declare COMMAND
   COMMAND=$(fc -ln -0)
   logger -p local1.notice -t bash -i -- "${USER}:${COMMAND}"
}
trap log2syslog DEBUG
```

2. Edit `/etc/rsyslog.conf` and add the following lines to the bottom of the file:

```bash
local1.* -/var/log/cmdline
```

3. Either restart the rsyslog service, or restart the whole machine to release all user sessions - forcing a reload of the bash profile and enacting the changes.

```bash
/etc/init.d/rsyslog restart
```

The audit logging will be visible under `/var/log/syslog` and `/var/log/cmdline`.

5. Allow read access to the `/var/log/cmdline`. **If you are not ok with this - stop right now**.

```
chmod +r /var/log/cmdline
```

4. Run the container.

```bash
docker run --name log-cmdline -it -d \
--restart on-failure \
--log-driver fluentd \
--log-opt tag=docker.efk.cmdline \
--log-opt fluentd-async-connect=true \
-v /var/log/cmdline:/media/cmdline:ro \
barklan/log-cmdline:rolling
```
