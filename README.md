# replcaseText
这是又go编写一个通过多携程匹配文本的demo
## 过程 
config.yaml 配置文件
+  先读取 transale_addres  
+  异步调用接口翻译 （根据language） 
+ 生成 json写入 transale_addres {key:"",value:""}的形式，并且也写回去
  （可以通过注释掉的方式不执行第一步，从而可以修改你要翻译的东西 ）
  ```go  mian.go
  if len(lib.Config.TransaleAddres) > 0 {
		//tra.OpenFileTxt(lib.Config.TransaleAddres, lib.Config.Language)
	}
  ```
+ 从catalog_address 找 文件入口  
+ 文件异步遍历 匹配 not_ignore_address  
+   要符合 not_ignore_address 的所有规则  
+ 开始匹配全文，并且覆写（不会导致git记录格式的改变）  
+ end 代码结束


## 优化 
+ 需要接口来做翻译，能否使用本地包，或者其他方式保证翻译的正确性
+ 匹配效率优化，可以利用算法去减少匹配次数

  
