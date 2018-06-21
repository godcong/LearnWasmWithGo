# LearnWasmWithGo
# 开始学习WASM了

## 配置：
### 1.获取最新版的Go Source,并编译：
    mkdir gowasm
    git clone git clone github.com/golang/go ./gowasm
    //export GOROOT_BOOTSTRAP=your/goroot/bootstrap/path
    cd gowasm
    cd src
    ./make.bash

### 2.使用最新编译的go文件作为，默认编译工具：
    ln -s ./bin/go /usr//bin/go
    PS:如果冲突，请先备份/usr//bin/go文件

### 3.配置环境变量：
    export GOARCH=wasm
    export GOOS=js