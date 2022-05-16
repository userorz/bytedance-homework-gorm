# go-gorm

首先实现一个脚本工具：
包含一个 User struct, 只包含 UUID string, Name string，Age int，Version int 四个字段，在脚本中使用 gorm ＋ mysql 初始化 DB, 并使用初始化后的 DB 的 AutoMigrate 迁移数据表。

然后完成一个 Gen (github.com/go-gorm/gen/) 项目

1、基于上面创建的数据库及表名，通过 Gen 的自动同步库表功能生成 struct People，并给该 struct 生成基本 CRUD 方法。

2、基于 OnConflict Upsert 功能实现 100 个随机用户的创建，其中需要包含重复的 UUID 用户的 Upsert, 在 Upsert 时，如果遇到重复 UUID 中，需要将 Version 更新为 Version + 1。

3、最后再通过一条自定义的Raw SQL 实现，将数据按 Version 分组，并取出 Version 最高的一组的用户总数的功能，该 Raw SQL 需要通过自定义查询方法的形式实现，需要给 People 生成相应的方法名: GetMaxVersionCount
