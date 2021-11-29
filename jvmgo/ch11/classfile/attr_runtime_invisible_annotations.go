package classfile

/***
RuntimeInvisibleAnnotations_attribute{
	u2 			attribute_name_index ;
	u4			attribute_length ;
	u2 			num_annotations ;
	annotation	annotations[num_annotations];	表中的每一项都表示标注在声明的上面的一条运行时不可见注解
}

annotation {
	u2 type_index ;
	u2 num_element_value_pairs ;
	{
		u2	element_name_index;
		element_value 	value;
	}element_value_pairs [num_element_value_pairs];
}
*/

type RuntimeInvisibleAnnotations struct {
	numAnnotations uint16
	annotation     []*Annotation
}

type Annotation struct {
	typeIndex         uint16
	elementValuePairs []*ElementValuePairs
}

type ElementValuePairs struct {
	elementNameIndex uint16

}
type ElementValue struct {



}
