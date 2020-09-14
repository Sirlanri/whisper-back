SELECT  userid  from user 
WHERE (userName='深蓝' or mail='mail@ri-co.cn');


/* 如果不存在内循环的值，就执行插入语句 */
INSERT INTO user (`mail`, `userName`, `password`)
SELECT 'mal@ri.cn',
    'rrw',
    '123456'
from DUAL
WHERE not exists (
        SELECT *
        from user
        WHERE userName = '蓝'
            or mail = 'mal@ri-co.cn'
        LIMIT 1
    );

INSERT INTO `post` (publisher,groupid,content) 
values
(1,1,"手动sql");

select groupid from groups where groupName="测试1";

select count(*) from post where groupid=1;

insert into groups (groupName,groupIntro,banner)
	values("手动1","手动插入的第一个","https://blog.ri-co.cn/wp-content/uploads/2020/04/white1.jpg");

insert into groups (groupName,groupIntro,banner)
	values("手动1","手动插入的第一个","https://blog.ri-co.cn/wp-content/uploads/2020/04/white1.jpg")
    if select * from groups where groupName="手动1"

INSERT INTO `groups` (groupName,groupIntro,banner)
		SELECT "手动4","aaaaaa","zero"
		from DUAL
		WHERE not exists (
				SELECT
                 *
				from `groups`
				WHERE groupName="手动4" LIMIT 1);

SELECT  * FROM post  ORDER BY postid DESC  LIMIT 20;

UPDATE reply SET haveRead=1 WHERE replyid=6 AND toUser=1;

UPDATE reply SET haveRead=0 WHERE replyid=6 ;

select replyid,postid,fromUser,content,haveRead 
	from reply where toUser=1 ORDER BY replyid DESC ;