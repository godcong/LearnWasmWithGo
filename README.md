# LearnWasmWithGo
# 开始学习WASM了

## 配置：
### 获取最新版的Go Source,并编译：
    mkdir gowasm
    git clone git clone github.com/golang/go ./gowasm
    //export GOROOT_BOOTSTRAP=your/goroot/bootstrap/path
    cd gowasm
    cd src
    ./make.bash

### 使用最新编译的go文件作为，默认编译工具：
    ln -s ./bin/go /usr//bin/go
    PS:如果冲突，请先备份/usr//bin/go文件

### 配置环境变量：
    export GOARCH=wasm
    export GOOS=js

### 现在开始我们的wasm旅程吧。
每个chapter目录下有对应的，chapter.**.md文档说明,可以直接按照说明操作
