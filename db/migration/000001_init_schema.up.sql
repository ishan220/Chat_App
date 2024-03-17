CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "password" varchar NOT NULL,
  -- "full_name" varchar NOT NULL,
  -- "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "chat" (
  "id" varchar PRIMARY KEY,
  "from_user" varchar NOT NULL,
  "to_user" varchar NOT NULL,
  "message" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "contact_list" (
   "username" varchar NOT NULL ,
   "member_username" varchar NOT NULL,
   "last_activity" bigint  NOT NULL
);

ALTER TABLE "contact_list"
ADD CONSTRAINT "username_member"
UNIQUE ("username","member_username");



