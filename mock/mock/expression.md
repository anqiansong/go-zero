# 表达值规则

## 随机值表达式random

### 1、number类型

|  expression   | description  | example  | example description  | example values |  
|  ----  | ----  | ----  | ----  | ----  | ----  |
| randomNumber[$start-$end]  | 指定范围的随机数 | randomNumber[100-1000]| 100至1000之间整数 | 999  | 
| randomTimestamp  | 随机timestamp | randomTimestamp| 毫秒时间戳 |  1602689779000  |
| randomUnix  | 随机unix | randomUnix| 秒时间戳 |  1602689779   |

    	
### 2、string类型

|  expression   | description  | example  | example description  | example values |  
|  ----  | ----  | ----  | ----  | ----  | ----  |
| randomString($[en\|zh\|mix]->$length)  | 指定长度、语言的随机字符串 | randomString(en->5) | 随机生成长度为5的英文字符串 | akdjj  | 
| randomName($[en\|zh\|mix]->$[person\|animal])  | 指定语言、类型的随机名称 | randomName(en->person)| 随机生成一个英文人名 |  KesonAn  |
| randomUUID  | 随机uuid | randomUUID| 随机uuid |  c101879c-fc0d-4994-ad8d-c71ec99e08b4   |
| randomObjectHex  | 随机ObjectHex | randomObjectHex| 随机ObjectHex |  557ef4ad0cf2e6ee15829f23   |
| randomDate($timeFormatOfGolang)  | 随机Date | randomDate(2006-01-02)| 随机生成格式为2006-01-02(golang)的日期 |  2020-01-01   |
| randomCity($[province\|city\|district\|mix])  | 随机省份、城市、县级市名称 | 1、randomCity(province) 2、randomCity(mix)| 1、随机一个省份 2、随机生成一个省市区联级的地名 |  1、上海市 2、广东省深圳市南山区   |
| randomSchool($[kindergarten\|primary\|junior\|high]->$[grade\|class\|mix])  | 随机学段年级班级名称(幼儿园-高中) | 1、randomSchool(high->grade) 2、randomSchool(junior->class) 3、randomSchool(primary->mix)| 1、随机年级 2、随机班级 3、随机年级班级 |  1、高一1 2、八年级 3、一年级1班  |
| randomTemplate(xx<<$expression>xx)  | 随机按照模板生成字符串,${expression}为除randomTemplate外的其他expression,通过此方法可以定义更多规则形式别名expression，如手机号表达式 | 1、randomTemplate(136<<.randomNumber[100-1000]>>) 2、randomTemplate(<<randomName(en->person)>>@qq.com) | 1、136开头的手机号 2、qq邮箱 |  1、13600001101 2、KesonAn@qq.com  |
| randomPhone($[mobile\|phone\|400])  | 生成手机号、座机 | randomPhone(mobile) | 生成手机号 |  13688888888、(021)88885678、400-456-7890  |
    	
### 3、bool类型

|  expression   | description  | example  | example description  | example values |  
|  ----  | ----  | ----  | ----  | ----  | ----  |
| randomBool  | 随机bool值 | randomBool| 随机bool值 | true  | 
    	
### 4、slice类型

|  expression   | description  | example  | example description  | example values |  
|  ----  | ----  | ----  | ----  | ----  | ----  |
| randomSlice($length)  | 随机切片长度 | randomSlice(3)| 生成长度为3的切片数据 | -  | 

    	
### 5、map类型

|  expression   | description  | example  | example description  | example values |  
|  ----  | ----  | ----  | ----  | ----  | ----  |
| randomMap($length)  | 随机map长度 | randomMap(3)| 生成3个的map数据 | -  | 

### 6、array类型
--
