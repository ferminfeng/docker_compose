
- writeMode 三种模式

| 模式      | 底层SQL |遇到主键/ 唯一索引冲突时|适用场景|
|---------|-------|---|---|
| insert  | INSERT INTO  | 直接报错，任务失败  | 全新导入、无重复数据  |
| replace | REPLACE INTO  | 删除旧行 → 插入新行  | 全量覆盖整行  |
| update  | INSERT ... ON DUPLICATE KEY UPDATE  | 只更新指定字段| 增量更新、只改部分字段  |


```azure
writeMode 选择 update
这样就只更新指定列


```