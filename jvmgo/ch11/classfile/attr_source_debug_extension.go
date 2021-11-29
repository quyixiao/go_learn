package classfile


/***
	SourceDebugExtension属性格式如下
	SourceDebugExtension_attribute{
		u2 attribute_name_index					字符串"SourceDebugExtension"
		u4 attribute_length						项的值给当前属性的长度，不包括初始化6个字节
		u1 debug_extend[attribute_length]		用以保存扩展调试信息，扩展调试信息对于Java虚拟机来说没有实际的语义，这个信息的
		用改进版本的UTF-8编码字符串表示 ，这个字符串不包含byte值为0的终止符
	}

 */

