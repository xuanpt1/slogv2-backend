package utils

import "regexp"

const (
	//RegexEmail    = `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	RegexPureText = "\\*\\*|__|\\*|_|~~|`|\\[\\]|\\(|\\)|\\{|\\}|\\[|\\]|#|\\+|-|\\.|!"

	TestText = "# 1 概念\n\n## 1.1 什么是接口文档\n\n```\n在项目开发中，web项目的前后端分离开发，APP开发，需要由前后端工程师共同定义接口，编写接口文档，之后大家都根据这个接口文档进行开发，到项目结束前都要一直维护\n```\n\n```\nAPI（Application Programming Interface）即应用程序接口。可以认为 API 是一个软件组件或是一个 Web 服务与外界进行的交互的接口。目的是提供应用程序与开发人员基于某软件或硬件得以访问一组例程的能力，而又无需访问源码，或理解内部工作机制的细节。从另一个角度来说，API 是一套协议，规定了我们与外界的沟通方式：如何发送请求和接收响应。\n```\n\n## 1.2 接口的组成\n\n```\n接口分为四部分：方法、uri、请求参数、返回参数\n\n1、方法:新增(post) 修改(put) 删除(delete) 获取(get)\n\n2、uri：controller的访问路径\n\n3、请求参数和返回参数，都分为5列：字段、说明、类型、备注、是否必填\n\n字段是类的属性；说明是中文释义；类型是属性类型，只有String、Number、Object、Array四种类型；备注是一些解释，或者可以写一下例子，比如负责json结构的情况，最好写上例子，好让前端能更好理解；是否必填是字段的是否必填。\n\n4、返回参数结构有几种情况：\n    1、如果只返回接口调用成功还是失败（如新增、删除、修改等），则只有一个结构体：code和message两个参数；\n    2、如果要返回某些参数，则有两个结构体：1是code/mesage/data，2是data里写返回的参数,data是object类型；\n    3、如果要返回列表，那么有三个结构体，1是code/mesage/data,data是object，里面放置page/size/total/totalPage/list 5个参数，其中list是Arrary类型，list里放object，object里是具体的参数。\n```\n\n## 1.3 优秀接口的特点\n\n```\n• 易学习：有完善的文档及提供尽可能多的示例和可copy－paste的代码，像其他设计工作一样，你应该应用最小惊讶原则。\n\n最小惊讶原则(Principle of least astonishment)\n\n最小惊讶原则通常是在用户界面方面引用，但同样适用于编写的代码。代码应该尽可能减少让读者惊喜。也就是说，你编写的代码只需按照项目的要求来编写。其他华丽的功能就不必了，以免弄巧成拙。\n\n• 易使用：没有复杂的程序、复杂的细节，易于学习；灵活的API允许按字段排序、可自定义分页、 排序和筛选等。一个完整的API意味着被期望的功能都包含在内。\n\n• 难误用：对详细的错误提示，有些经验的用户可以直接使用API而不需要阅读文档。\n\n而对于开发人员来说，要求又是不一样的：\n\n• 易阅读：代码的编写只需要一次一次，但是当调试或者修改的时候都需要对代码进行阅读。\n\n• 易开发：个最小化的接口是使用尽可能少的类以及尽可能少的类成员。这样使得理解、记忆、调试以及改变API更容易。\n```"
)

func GetPureTextRegex(str string) string {
	reg := regexp.MustCompile(RegexPureText)
	return reg.ReplaceAllString(str, "")
}
