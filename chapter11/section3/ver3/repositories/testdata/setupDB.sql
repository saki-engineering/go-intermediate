create table if not exists articles (
	article_id integer unsigned auto_increment primary key,
	title varchar(100) not null,
	contents text not null,
	username varchar(100) not null,
	nice integer not null,
	created_at datetime
);

create table if not exists comments (
	comment_id integer unsigned auto_increment primary key,
	article_id integer unsigned not null,
	message text not null,
	created_at datetime,
	foreign key (article_id) references articles(article_id)
);

insert into articles (title, contents, username, nice, created_at) values
	('firstPost', 'This is my first blog', 'saki', 2, now());

insert into articles (title, contents, username, nice) values
	('2nd', 'Second blog post', 'saki', 4);

insert into comments (article_id, message, created_at) values
	(1, '1st comment yeah', now());

insert into comments (article_id, message) values
	(1, 'welcome');
