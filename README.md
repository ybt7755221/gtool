## Gtool 基于gpi的自动生成model及基础service的工具

#### 使用说明：

> gtool --table tableName --database dbName --router routeName --controller true --file-path ./

+ -t &nbsp;&nbsp;--table &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;表名
+ -d &nbsp;--database &nbsp;数据库名
+ -r &nbsp;&nbsp;--router &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;路由名称（无此参数不生成路由部分）
+ -c &nbsp;--controller &nbsp;是否生成controller 默认false，
+ -f &nbsp;&nbsp;--file-path &nbsp;&nbsp;&nbsp;生成文件路径，默认当前目录下的output