DECLARE @user TEXT;

SET
    @user = "";

-- tags
INSERT INTO
    tags(id, name)
VALUES
    (1, "TECH"),
    (2, "MUSIC"),
    (3, "TRAVEL");

-- posts
INSERT INTO
    posts(id, title, body, tag_ids, user_id)
VALUES
    (
        1,
        "The First Blog",
        "This is my first blog! I am going to write the content soon.",
        { 3 },
        @user
    ),
    (
        2,
        "New Music coming tomorrow!",
        "New music will be released soon!",
        { 2,3 },
        @user
    );

-- comments
INSERT INTO
    comments(id, body, post_id)
VALUES
    (1, "Let's go!", 1),
    (2, "I can't wait to see the content!", 1),
    (3, "I wonder what that is like", 2);