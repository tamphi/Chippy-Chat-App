SET TIME ZONE 'PST8PDT';

CREATE TABLE "users"(
    "id" serial PRIMARY KEY,
    "username" varchar NOT NULL,
    "password" varchar NOT NULL 
);

-- CREATE TABLE rooms (
--     "id" serial PRIMARY KEY,
--     "room_name" varchar NOT NULL,
--     "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
-- );

CREATE TABLE messages (
    "id" varchar NOT NULL,
    "receiver" varchar NOT NULL,
    "sender" varchar NOT NULL,
    "content" text NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
