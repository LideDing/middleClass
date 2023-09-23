# 一个完整的括号字符定义规则如下：
# 1. 空字符串是完整的
# 2. 如果s是完整的字符串，那么（s）也是完整的
# 3. 如果s和t是完整的字符串，将它们连接起来形成的st也是完整的
# 例如“(()())”，“”和“(())()”是完整的括号字符串，“())(”，“()(”和“)”是不完整的括号字符串
# 你有一个括号字符串s，现在需要在其中任意位置添加最少的括号，将其转换为完整的括号字符串。请问你最少需要添加多少个括号？

def bracket(s) -> int:
    if s == "":
        return 0
    res, status = 0, 0
    for i in range(len(s)):  # 从左到右遍历
        if s[i] == "(":  # 遇到左括号，status加1
            status += 1
        else:  # 遇到右括号，status减1
            status -= 1
        if status < 0:  # status小于0，说明右括号多于左括号，需要添加左括号
            res += 1  # 添加左括号
            status = 0  # status置0
    return res + status  # status大于0，说明左括号多于右括号，需要添加右括号


print(bracket("(()())"))
print(bracket("())("))
print(bracket("()("))
print(bracket(")"))
