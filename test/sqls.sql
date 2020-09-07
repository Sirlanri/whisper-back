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

select groupid from `group` where groupName="测试1";

select count(*) from post where groupid=1;

insert into `group` (groupName,groupIntro,banner)
	values("手动1","手动插入的第一个","https://blog.ri-co.cn/wp-content/uploads/2020/04/white1.jpg");