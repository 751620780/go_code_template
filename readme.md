# 编译
go build  -o build/main.exe  main/test/main.go
# 如何创建golang项目
> 注意：请在环境变量中增加GO111MODULE   on 或 auto
1. 使用vscode打开一个项目文件夹
2. 使用 ``go mod init xxx``,xxx是项目的名称
3. 创建项目的子模块文件夹,在文件夹内编写该子模块下的所有代码的package应该都是该文件夹的名字，如果不是也无所谓，但是引用代码应当是该文件夹名
4. 使用的第三方package应当在根目录下执行``go get xxx``命令，例如``go get github.com/go-sql-driver/mysql``,则此项目可以引用到mysql这个第三方库。此方法引入的库都会记录到``go.sum``文件中，同时这个依赖的模块也会被记录到go.mod文件中，整个项目就可访问该模块了。

# go mod 操作
go mod init        初始化当前文件夹, 创建go.mod文件，go mod init xxx
go mod tidy        将根据实际使用的模块更新go.sun和go.mod，包括引用的模块引用的其他模块(间接引用的模块),可以清理不用的第三方模块
go mod vendor      将依赖复制到vendor下(解决在新机器下编译时需要从新下载第三方库代码的耗时问题)
go mod download    下载依赖的module到本地cache，如果将代码拷贝到新的开发环境，需要先下载引用的第三方模块，类似node的于``npm install``（默认为$GOPATH/pkg/mod目录）
go mod graph       打印模块依赖图