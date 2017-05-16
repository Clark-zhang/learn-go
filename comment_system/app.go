/*
Start docker:
cd comment_system/docker && docker-compose up -d

Root access for mysql from remote:
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'password' WITH GRANT OPTION;
use mysql;
update user set authentication_string=PASSWORD("NEWPASSWORD") where User='root';
FLUSH PRIVILEGES;
http://stackoverflow.com/questions/14779104/how-to-allow-remote-connection-to-mysql

Create table:
CREATE TABLE IF NOT EXISTS `comments_relationship` (
    `c_id` bigint(30),
    `p_id` bigint(30),
    `fp_id` bigint(30),
    `e_id` bigint(30)
);
CREATE TABLE IF NOT EXISTS `comments` (
    `id` bigint(30) NOT NULL auto_increment,
    `e_id` bigint(30),
    `user_id` bigint(30),
    `content` varchar(1024),
    `ip` varchar(25),
    `created_at` datetime,
    PRIMARY KEY  (`id`)
);
*/

package main

import (
     cs_server "github.com/Clark-zhang/learn-go/comment_system/server"
     )

func main(){
    cs_server.Run()
}