## 文件权限说明

不同安全机制下对于文件权限的表示是不一样的，为了向用户屏蔽差异，secPaver封装了一些常用的文件权限：

|    权限     |        说明        |           备注            |
| :---------: | :----------------: | :-----------------------: |
|   create    |      创建文件      |                           |
|    read     |      读取文件      |                           |
|    write    |      写入文件      |                           |
|   append    |      追加文件      |                           |
| rename/link |  重命名/链接文件   |                           |
|   remove    |      删除文件      |                           |
|    lock     |       锁文件       |                           |
|     map     |    内存映射文件    |                           |
|    exec     |      执行文件      |                           |
|   search    |      搜索目录      |    仅对dir类型文件有效    |
|    ioctl    |    输入输出控制    |                           |
|    mount    |      挂载文件      | 仅在SELinux策略生成时有效 |
|   mounton   | 文件路径作为挂载点 | 仅在SELinux策略生成时有效 |

#### 注意事项

- SELinux机制的文件权限比较复杂，上述列表无法完全覆盖，如果用户需要使用，可以直接在actions中填入。secPaver会对这些扩展权限检查，如果为合法权限，会添加到SELinux规则中；
- 考虑到用户易用性，当前生成SELinux策略时主体应用程序默认具备对目录文件的search权限（设置为private的文件除外）。