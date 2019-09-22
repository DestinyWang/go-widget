package red_black_tree

// 1. 根节点必须为黑色
// 2. 父子不能同为红色
// 3. 从任何一个节点出发, 到达叶子节点经过的黑色节点数量一致


// 1. 红色是默认插入颜色(因为违规修复代价较小)
// 2. 如果没有双亲, 变为黑色
// 3. 双亲为黑色
// 4. 双亲为红色, 双亲需要变为黑色