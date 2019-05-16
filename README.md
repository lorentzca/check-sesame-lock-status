# check-sesame-lock-status

## Description

Check Sesame lock status.

## Synopsis

```
$ check-sesame-lock-status -k <api key> -d <device id>
Sesame OK: Locked!
```

```
$ check-sesame-lock-status -k <api key> -d <device id>
Sesame CRITICAL: Unocked!
```

```
$ check-sesame-lock-status -k <invalid api key> -d <device id>
Sesame WARNING: UNAUTHORIZED
```

```
$ check-sesame-lock-status -k <api key> -d <invalid device id>
Sesame WARNING: BAD_PARAMS
```

## Setting for mackerel-agent

```
[plugin.checks.check-sesame-lock-status-sample]
command = ["check-sesame-lock-status", "-k", "<api key>", "-d", "device id"]
```
