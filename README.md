Terrago 
name is very odd right 🥸 i agree but what it is ?
it is a simple metrices collector and displays it on localhost
right now it just simply supports displaying the metrices on the screen  
but in future we are going to add some addtional features.

this is not any production based project ❎ , ment for learning and
implementing  some of the best devops practices.

what are the features it supports 🌃? 
currently we are workin on finalyzing everything, so for now you cant use it but
it supports deep packet inspection , network metrices and cluster metrices 

what are our future goals 🗻?
currentl the main goal is to make it available as soon as possible 
adding some feature like auto alerts using slack and discord ,
and making it more precise in terms of giving more detail info


installation:

```
git clone https://github.com/Horiodino/terrago.git
cd terrago 
go build -o terrago main.go
./terrago init
```


Cli uses:



```./terrago clusterinfo```

output:
```
CPU Usage:  52
CPU Cores:  4
Nodes:  2
Total Memory:  8.076025856e+09
Used Memory:  8.90318848e+08
Disk Usage:  0
Total Disk:  4.2924466176e+10
```



```./terrago nodeinfo```

output:
```
|-------------------------------------|
Node Name:  [its linked to ip so cant show here ]
CPU Usage:  [1.4500000000000002]
Memory Usage:  [11.490279375396868]
Disk Usage:  [0]
CPU Temperature:  []
IP:  [IP]
|-------------------------------------|

|-------------------------------------|
Node Name:  [its linked to ip so cant show here ]
CPU Usage:  [1.25]
Memory Usage:  [10.541424403861058]
Disk Usage:  [0]
CPU Temperature:  []
IP:  [IP]
|-------------------------------------|
```
ignore the cpu temp going to fix it soon 😅


```

                                                                        (Saves the Recived Data in Json file or Struct)
                                                    |--------------------------------------------------------------------------------------------|
                                                    v                                                                                            |
                                                    |-----------------------------------------------------|                                      |
                        |---------------------------|----------------------|                              |                                      |
                        v                           v                      v                              v                                      |
+--------------+   +------------------+   -------Structs-------   ----Authentication----   +---Metrices--------------------------------+         |
|   Terrago    |-->|    CMD           |-->| Metrices Struct   |-->|   (with JWT)       |   |  (contains funtions for getting metrics)  |         |
+--------------+   +------------------+   +-------------------+   +--------------------+   +-------------------------------------------+         |
                        |                    |                   |                                 |       |                                     |
                        |--------------------|                   |                                 |       |                                     |
                                             |                   |                                 |       |                                     v
                                             |                   |                                 |       |------- +-------------------------------------------+
                                             v                   v                                 |                |    terrago client (fetchs node additional |
                               +------------------+     +------------------+     +----------+      |                |                    info )                 |
                               |   Commands       |     |   Struct Update  |<----| Metrices |      |                +-------------------------------------------+
                               | (init, dashboard,|-----| (Website Update) |     +----------+      |                
                               |  clusterinfo)    |     +------------------+                       |
                               +------------------+            |                                   |
                                      |                        v                                   |
                                      |                 +-----------+                              |
                                      |                 |  Website  |                              |
                                      |                 |(Continuous|                              |
                                      |                 |  Update)  |                              |
                                      |                 +-----------+                              |
                                      |                                                            |
                                      |                                                            |
                                      |------------------------------------------------------------|
                                          (Continues calls the funtion to fetch the latest data)

```
