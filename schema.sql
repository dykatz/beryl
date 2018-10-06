create table classes (
	id	serial primary key,
	name	text not null
);

create table users (
	id	serial primary key,
	name	text not null
);

create table class_users (
	class_id	int not null references classes (id),
	user_id		int not null references users (id),
	role		text,
	primary key (class_id, user_id)
);

create table assignment_category (
	class_id	int not null references classes (id),
	name		text not null,
	weight		numeric not null,
	primary key (class_id, name)
);

create table assignment (
	id		serial primary key,
	class_id	int not null references classes (id),
	category_name	text not null,
	name		text not null,
	maxpoints	numeric not null,
	due		datetime not null,
	foreign key (class_id, category_name) references assignment_category (class_id, name)
);

create table user_assignment_submission (
	user_id		int not null references users (id),
	assignment_id	int not null references assignment (id),
	posted		datetime not null,
	points		numeric,
	primary key (assignment_id, user_id)
);

create table discussions (
	id		serial primary key,
	class_id	int not null references classes (id),
	author_id	int not null references users (id),
	started		datetime not null,
	subject		text not null,
	body		text
);

create table comments (
	id		serial primary key,
	discussion_id	int not null references discussions (id),
	author_id	int not null references users (id),
	parent		int references comments (id),
	posted		datetime not null,
	body		text not null
);

-- Get all users in a class
select users.id, users.name
from users
join class_users on class_users.user_id = users.id
where class_users.class_id = ?
order by users.name;

-- Get all assignment categories in a class
select assignment_category.name, assignment_category.weight
from assignment_category
where assignment_category.class_id = ?
order by assignment_category.name;

-- Get all assignments in a class (teacher view)
select assignments.id, assignments.name, assignments.category_name,
	assignments.maxpoints, assignments.due
from assignments
where assignments.class_id = ?
order by assignments.due;

-- Get all assignments in a class (student view)
select assignments.id, assignments.name, assignments.category_name,
	user_assignment_submission.points, assignments.maxpoints,
	user_assignment_submission.posted, assignments.due
from assignments
join user_assignment_submission on user_assignment_submission.assignment_id = assignments.id
where assignments.class_id = ?
and user_assignment_submission.user_id = ?
order by assignments.due;

-- Get all discussions in a class
select discussions.id, discussions.subject, discussions.started, users.name
from discussions
join users on discussions.author_id = users.id
order by discussions.started;
