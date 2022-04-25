# 简要介绍
管理员在 Browser 端上传文件， browser 端不落盘通过 grpc 会将文件上传到云机房服务器。 
云机房服务器保存该文件，并向所有连接到机房服务的客户端广播下载该文件

# 使用
- 1 执行 browser.exe, server.exe, client.exe
- 2 访问 localhost:8888

# 通信协议
### browser 与 server 间 grpc， server 与 client 间 socket

### proto

```` proto
syntax = "proto3";

package messaging;

service UploadService {
        rpc Upload(stream Chunk) returns (UploadStatus) {}
}

message Chunk {
  oneof data {
        bytes Content = 1;
        FileInfo Info = 2;
      }
}

message FileInfo {
  string FileType = 1;
}

enum UploadStatusCode {
        Unknown = 0;
        Ok = 1;
        Failed = 2;
}

message UploadStatus {
        string Message = 1;
        UploadStatusCode Code = 2;
}

````
# 三个模块的代码介绍
## browser 管理员上传文件
### 提供文件上传界面， 不落盘
localhost:8888/ 为显示界面
上传文件后会跳转到 localhost:8888/upload 接口
````
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", uploadFile)
````

### 接收前端请求，向中心机房发送文件
````
type Client interface {

	UploadFile(ctx context.Context,  r *http.Request) (stats Stats, err error)

	Close()
}
````
````
func (c *ClientGRPC) UploadFile(ctx context.Context, r *http.Request) (stats Stats, err error) {
    
    .....
    
    // 发送文件信息
	err = stream.Send(&messaging.Chunk{
		Data: &messaging.Chunk_Info{
			Info: &messaging.FileInfo{
				FileType:filepath.Ext(handler.Filename),
			},
		},
	})
	
	.....
	
	// 发送文件信息
	for writing {
		n, err = file.Read(buf)
		......

		err = stream.Send(&messaging.Chunk{
			Data:&messaging.Chunk_Content{
				Content: buf[:n],
			},
		})
		.....

	}
}
````

### bugfix
- 1 是否要分片 // 分片需要前端支持
- 2 支持 http2 或是 grpc // 仅支持 grpc


## server 通过 grpc 接收浏览器上传的文件，落盘并提供下载服务
### 对 browser 接收文件，对 client 提供下载文件
````
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go uploadServer(&wg)
	go downloadServer(&wg)
	wg.Wait()

	fmt.Println("server finished!")
}
````
### 接收文件
````
func (s *ServerGRPC) Upload(stream messaging.GuploadService_UploadServer) (err error) {
	......
	data := bytes.Buffer{}
	// 接收 grpc 文件信息
	req, err := stream.Recv()
	fileType := req.GetInfo().FileType
	filePath := fmt.Sprintf("%s/%d%s", "upload", randInt, fileType)
	file, err := os.Create(filePath)
    
    .......
    
    // 接收 grpc 文件内容,并落盘写入本地文件
    for {
    
        ......
	
		req, err := stream.Recv()

		......

		chunk := req.GetContent()
		
		
		_, err = data.Write(chunk)
		
		......

		_, err = data.WriteTo(file)
	}

}
````

### 客户端间共享文件
```` server/main.go
type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		
		case filepath := <-filePath:// 广播文件
			for cli := range clients {
				cli <- filepath
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			fmt.Println("one connection leave.")
			delete(clients, cli)
			close(cli)
		}
	}
}

````

````
func handleConn(conn net.Conn) {
	defer trace("handleConn:" + conn.RemoteAddr().String())()

	ch := make(chan string)

	go sendFileToClient(conn, ch)
	entering <- ch
}
````

````
func sendFileToClient(connection net.Conn, ch chan string) {
   
	for filepath := range ch { // 从管道里读文件路径
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Println(err)
			return
		}
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Println(err)
			return
		}
		
		// fizeSize 10 字节， fileName 64 字节，与 客户端对应
		fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10) /
		fileName := fillString(fileInfo.Name(), 64)   
		              
		connection.Write([]byte(fileSize))
		connection.Write([]byte(fileName))
		sendBuffer := make([]byte, BUFFERSIZE)
		fmt.Println("Start sending file! ")
		for {
			_, err = file.Read(sendBuffer)
			if err == io.EOF {
				break
			}
			connection.Write(sendBuffer)
		}
		file.Close()
	}

	return
}
````

### bugfix
- 1 监控客户端连接,向连接广播已上传文件 // 已解决
- 2 文件名和扩展名要正确, //已解决

## client 连接 server 端下载文件

````
	for {
		bufferFileName := make([]byte, 64)
		bufferFileSize := make([]byte, 10)

		connection.Read(bufferFileSize)
		fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

		connection.Read(bufferFileName)
		fileName := strings.Trim(string(bufferFileName), ":")

		fmt.Printf("download filename and filesize!\nfilesize:%s,filename:%s\n", bufferFileSize,bufferFileName)

		newFile, err := os.Create(fileName)

		if err != nil {
			panic(err)
		}

		var receivedBytes int64

		for {
			if (fileSize - receivedBytes) < BUFFERSIZE {
				io.CopyN(newFile, connection, (fileSize - receivedBytes))
				connection.Read(make([]byte, (receivedBytes+BUFFERSIZE)-fileSize)) // 填充字节，与 server 端对齐
				break
			}
			io.CopyN(newFile, connection, BUFFERSIZE)
			receivedBytes += BUFFERSIZE
		}

		newFile.Close()
	}

````
### bugfix

- 1 连接开始有的文件才下载， 或是上传过的文件都要下载 // 改成连接只下载当前上传的
- 2 多读了 10 个字节， 先用黑科技解决了即写文件时多写了 10 个字节 // 发现是写法有问题，已经去掉



# QA
- 无法在接口 func (c *ClientGRPC) UploadFile(ctx context.Context, f interface{}) (stats Stats, err error)
里把 f 适配 *os.file 和 multipart.File （switch 里定义的变量， 外部无法使用）
  
# todo
- 1 参数通过文件或是命令行传入
- 2 单元测试及基准测试
- 3 server 检测客户端断掉的连接， 并关闭对应的连接和 channel
- 4 记录不同客户端的下载 
  
# 里程碑
#### 4 月 26 日 15:58 html 上传文件成功
#### 4 月 28 日 15：00 整体流程基本跑通
#### 4 月 28 日 19：30 文档完成

