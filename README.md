有限状态机
=====

基于长连接游戏设计的状态机（以棋牌游戏--麻将为例）

------

### 1. 三方库
 teleport ：https://github.com/henrylee2cn/teleport/raw/master/doc/Teleport
> 本项目使用了teleport作为长连接交互的底层实现，但稍微做了一些修改：
> 1. 将原有心跳、认证等hander提取出来，与其他请求handler合并成handler包
> 2. 将原内部请求的请求体改为protobuf的字节流，所有协议使用protobuf3


### 2. 代码结构
    server
    │
    ├─global       一些全局变量
    ├─handlers     请求的handler
    ├─machine      状态机代码
    │   ├─action.go                 玩家行为的结构体
    │   ├─config.go                 房间（桌子）配置
    │   ├─define.go                 行为、状态、事件、操作、卡牌定义
    │   ├─player.go                 玩家对象
    │   ├─player_machine.go         玩家状态机
    │   ├─player_manager.go         玩家行为管理器
    │   ├─player_rules.go           玩家的检验规则
    │   ├─player_rule_state.go      玩家触发规则时进入的状态
    │   ├─player_state.go           玩家正常流程规则
    │   ├─table.go                  房间（桌子）对象
    │   ├─table_machine.go          房间（桌子）状态机
    │   ├─table_manager.go          房间（桌子）行为管理器
    │   ├─table_rules.go            房间（桌子）的检验规则
    │   ├─table_state.go            房间（桌子）的正常流程
    │   └─win_algorithm.go          胡牌的DFS算法
    ├─proto        protobuf3的协议代码
    └─teleport     长连接交互底层

### 3. 说明
 使用teleport使本人在初期注重逻辑开发的时候节约了大量时间，但个人感觉teleport不适合跨平台交互，因此后面一些工作将会着重于长连接底层的改写。
 另外，目前这套代码仅仅实现了主要逻辑部分，重连、数据落盘、日志整理等都在计划当中。keep coding...

### 4. coding...