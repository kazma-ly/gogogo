### protobuf 数据结构
```
MessageType: 消息类型
TextMessage: 文字消息
FileMessage: 文件消息
```

### MessageType
```
TEXT_MESSAGE int  = 1 表示文字消息
FILE_MESSAGE int  = 2 表示文件消息
CLOSE_MESSAGE int = 3 表示关闭连接消息
```

### TextMessage
```
val    string 文本消息数据
system bool   是否为系统发送的消息
```

### FileMessage
```
Content []byte 文件内容(可以是部分)
Last bool      是否为最后部分
Md5 string     数据指纹
Len int64      数据长度
Name string    文件名字
```