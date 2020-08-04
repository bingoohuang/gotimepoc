# gotimepoc

POC of go time in location

## Build

1. `GOOS=linux GOARCH=amd64 go install  -ldflags="-s -w" ./...`

## 关于UTC和CST

[孙雪峰: 关于UTC和CST][1]

1. UTC 协调世界时，简称UTC, Coordinated Universal Time
1. CST可以为如下4个不同的时区的缩写：
    - 美国中部时间：Central Standard Time (USA) UT-6:00
    - 澳大利亚中部时间：Central Standard Time (Australia) UT+9:30
    - 中国标准时间：China Standard Time UT+8:00, 其实就是UTC+8，也是GMT+8小时
    - 古巴标准时间：Cuba Standard Time UT-4:00

[V__KING__:linux 关于GMT， CST， UTC的编程和设置][2]

1. GMT=UTC 
    - Greenwich Mean Time 格林尼治平均时
    - UTC和GMT都与英国伦敦的本地时相同，所以程序中UTC与GMT没什么不同。意思UTC=GMT是相等的
1. Unix时间戳：
    - 在计算机中看到的UTC时间都是从（1970年01月01日 0:00:00)开始计算秒数的。
    - 所看到的UTC时间那就是从1970年这个时间点起到具体时间共有多少秒。 这个秒数就是Unix时间戳。

[1]: https://blog.csdn.net/snowin1994/article/details/77530033
[2]: https://blog.csdn.net/V__KING__/article/details/79046976
[3]: https://www.cyberciti.biz/faq/centos-linux-6-7-changing-timezone-command-line/
[4]: https://www.vultr.com/docs/setup-timezone-and-ntp-on-centos-6

## Change the current timezone in CentOS 6 and older

[How To Change Timezone on a CentOS 6 and 7][3]

You need to use the ln command to set timezone on Centos 6. 

Type the following commands as root:

```bash
[root@app01 vagrant]# rm /etc/localtime
rm：是否删除普通文件 "/etc/localtime"？y
[root@app01 vagrant]# ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
[root@app01 vagrant]# date
2020年 08月 04日 星期二 12:36:31 CST
[root@app01 vagrant]# ls -l /etc/localtime
lrwxrwxrwx 1 root root 33 8月   4 12:36 /etc/localtime -> /usr/share/zoneinfo/Asia/Shanghai
```

## write the system time info into the hardware clock.

[Setup Timezone and NTP on CentOS 6][4]

`vi /etc/sysconfig/clock`

```
ZONE="Asia/Shanghai"
UTC=false
ARC=false
```

Write the system time into the hardware clock.

`hwclock --systohc --localtime`

Input `hwclock` to see the result.

## gotimepoc

```bash
[vagrant@app01 ~]$ cat /etc/centos-release
CentOS release 6.5 (Final)
[root@app01 vagrant]# /vagrant_data/gotimepoc
2020/08/04 12:21:57 env TZ:Asia/Shanghai
2020/08/04 12:21:57 time.Local:Asia/Shanghai
2020/08/04 12:21:57 Local Time: 2020-08-04 12:21:57.49982987 +0800 CST m=+0.000000291
2020/08/04 12:21:57 UTC Time: 2020-08-04 04:21:57.499829986 +0000 UTC
2020/08/04 12:21:57 Unix Seconds: 1596514917
2020/08/04 12:21:57 UnixNano: 1596514917499830218
UTC 2020/08/04 04:21:57 time.Local:Asia/Shanghai
UTC 2020/08/04 04:21:57 Local Time: 2020-08-04 12:21:57.499834395 +0800 CST m=+0.000004818
UTC 2020/08/04 04:21:57 UTC Time: 2020-08-04 04:21:57.499836808 +0000 UTC
UTC 2020/08/04 04:21:57 Unix Seconds: 1596514917
UTC 2020/08/04 04:21:57 UnixNano: 1596514917499839779
```

![image](https://user-images.githubusercontent.com/1940588/89253930-8f696c00-d650-11ea-82fe-aad1882dceed.png)
