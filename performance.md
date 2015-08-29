# g2z Performance

TL;DR: Go modules seem to perform on par with C modules with a margin of error
for implementation differences. They hammer through 2000+ NVPS on a low powered
VM and are 7x faster than loading a simple Perl script in a UserParameter for
the same workload.

The following statistics were gathered using the included Dockerfile with the
Zabbix agent restarted before each test. The test machine was a boot2docker VM
with 2GB RAM running on VirtualBox on a 2010 MacBook Pro i7.

The key outcome sought was a like-for-like comparison between native agent keys
and keys written in Go using g2z. Some penalty is expected using the Go
framework and for the additional routing required for each key request.


## Summary

Item                                                  | NVPS
------------------------------------------------------|-----
[agent.ping](#native-key-agentping) (builtin C)       | 2917
[agent.version](#native-key-agentversion) (builtin C) | 2954
[go.ping](#go-key-goping) (g2z module)                | 2929
[go.version](#go-key-goversion) (g2z module)          | 2513
[up.ping](#userparameter-upping) (/bin/echo)          | 487
[perl.ping](#userparameter-perlping) (Perl script)    | 396


## Native key: agent.ping

A C function built in to the agent which returns `1`

```c
static int	AGENT_PING(AGENT_REQUEST *request, AGENT_RESULT *result)
{
	SET_UI64_RESULT(result, 1);

	return SYSINFO_RET_OK;
}
```

	$ zabbix_agent_bench -iterations=50000 -key agent.ping
	Testing 1 keys with 4 threads (press Ctrl-C to cancel)...
	agent.ping :	50000	0	0

	=== Totals ===

	Total values processed:		50000
	Total unsupported values:	0
	Total transport errors:		0
	Total key list iterations:	50000

	Finished! Processed 50000 values across 4 threads in 17.135073895s (2917.991501 NVPS)


## Go key: go.ping

A Go function loaded via g2z which returns `1`

```go
func Ping(request *g2z.AgentRequest) (uint64, error) {
	return 1, nil
}
```

	$ zabbix_agent_bench -iterations=50000 -key go.ping
	Testing 1 keys with 4 threads (press Ctrl-C to cancel)...
	go.ping :	50000	0	0

	=== Totals ===

	Total values processed:		50000
	Total unsupported values:	0
	Total transport errors:		0
	Total key list iterations:	50000

	Finished! Processed 50000 values across 4 threads in 17.065188613s (2929.941247 NVPS)


## UserParameter: up.ping

A simple user parameter which calls `/bin/echo 1`

	$ zabbix_agent_bench -iterations 50000 -key up.ping
	Testing 1 keys with 4 threads (press Ctrl-C to cancel)...
	up.ping :	50000	0	0

	=== Totals ===

	Total values processed:		50000
	Total unsupported values:	0
	Total transport errors:		0
	Total key list iterations:	50000

	Finished! Processed 50000 values across 4 threads in 1m42.490522523s (487.849986 NVPS)


## UserParameter: perl.ping

A user parameter which calls a Perl script which returns `1`:

```perl
#!/usr/bin/perl -w
print "1";
```

	$ zabbix_agent_bench -iterations 50000 -key perl.ping
	Testing 1 keys with 4 threads (press Ctrl-C to cancel)...
	perl.ping :	50000	0	0

	=== Totals ===

	Total values processed:		50000
	Total unsupported values:	0
	Total transport errors:		0
	Total key list iterations:	50000

	Finished! Processed 50000 values across 4 threads in 2m6.227369837s (396.110606 NVPS)


## Native key: agent.version

A C function built in to the Zabbix agent which returns the agent version string

```c
static int	AGENT_VERSION(AGENT_REQUEST *request, AGENT_RESULT *result)
{
	SET_STR_RESULT(result, zbx_strdup(NULL, ZABBIX_VERSION));

	return SYSINFO_RET_OK;
}
```

	$ zabbix_agent_bench -iterations=50000 -key agent.version
	Testing 1 keys with 4 threads (press Ctrl-C to cancel)...
	agent.version :	50000	0	0

	=== Totals ===

	Total values processed:		50000
	Total unsupported values:	0
	Total transport errors:		0
	Total key list iterations:	50000

	Finished! Processed 50000 values across 4 threads in 16.921980955s (2954.736808 NVPS)


## Go key: go.version

A Go function loaded via g2z which returns the Go runtime version string

```go
func Version(request *g2z.AgentRequest) (string, error) {
	return runtime.Version(), nil
}
```

	$ zabbix_agent_bench -iterations=50000 -key go.version
	Testing 1 keys with 4 threads (press Ctrl-C to cancel)...
	go.version :	50000	0	0

	=== Totals ===

	Total values processed:		50000
	Total unsupported values:	0
	Total transport errors:		0
	Total key list iterations:	50000

	Finished! Processed 50000 values across 4 threads in 19.889195084s (2513.927778 NVPS)
