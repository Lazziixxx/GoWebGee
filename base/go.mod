module example
go 1.15
require gee v0.0.0
//require () 包含多个 包
replace gee => ./gee 
//重定向包的位置

//module 语句指定包的名字（路径）
//require 语句指定的依赖项模块
//replace 语句可以替换依赖项模块
//exclude 语句可以忽略依赖项模块
