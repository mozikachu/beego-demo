A beego API demo

![goreport]( https://goreportcard.com/badge/github.com/mozikachu/beego-demo )

### Usage

1. 配置 MySQL

    ```sql
    CREATE TABLE `user` (
      `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'uid',
      `user_name` varchar(255) NOT NULL DEFAULT '',
      `password` varchar(255) NOT NULL,
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
    
    CREATE TABLE `post` (
      `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
      `title` varchar(255) NOT NULL DEFAULT '',
      `user_id` int(11) NOT NULL COMMENT 'uid',
      PRIMARY KEY (`id`),
      KEY `user_id` (`user_id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
    
    ```

2. 配置数据库参数

    https://github.com/mozikachu/beego-demo/blob/master/main.go#L13
    
3. `bee run`
    
    ```bash
    $ cd $GOPATH/github.com/mozikachu/beego-demo
    $ bee run
    ```

4. CURD
    
    ```bash
    // 向数据库中添加新 user
    $ curl -X POST -d '{"UserName": "init", "Password": "123456"}' http://127.0.0.1:8081/v1/user
    
    // 向数据库添加新 post (需在 URL 中给定 user.id)
    $ curl -X POST -d '{"Title": "init"}' http://127.0.0.1:8081/v1/post/2
    
    // 根据 post.id 查询 post
    $ curl -X GET http://127.0.0.1:8081/v1/post/1
    
    // 关联查询：查询指定用户名下的所有 posts
    $ curl -X GET http://127.0.0.1:8081/v1/post/username/feef
    
    // 关联更新：更新指定用户名下的所有 posts
    $ curl -X PUT -d '{"title":"new"}' http://127.0.0.1:8081/v1/post/username/feef
    
    // 关联删除：删除指定用户名下的所有 posts
    $ curl -X DELETE http://127.0.0.1:8081/v1/post/username/feef
    ```