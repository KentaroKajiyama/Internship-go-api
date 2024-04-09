-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE OR REPLACE FUNCTION delete_todo_tag()
RETURNS TRIGGER AS $$
BEGIN
    -- todos テーブルから行が削除された場合
    IF TG_TABLE_NAME = 'todos' THEN
        DELETE FROM "todo_tag" WHERE todo_id = OLD.todo_id;
    END IF;

    -- tags テーブルから行が削除された場合
    IF TG_TABLE_NAME = 'tags' THEN
        DELETE FROM "todo_tag" WHERE tag_id = OLD.tag_id;
    END IF;

    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

-- 既存のDeleteトリガーを削除
DROP TRIGGER IF EXISTS trigger_delete_todo ON todos;
DROP TRIGGER IF EXISTS trigger_delete_tag ON tags;

-- todos テーブルに対するDeleteトリガー
CREATE TRIGGER trigger_delete_todo
AFTER DELETE ON todos
FOR EACH ROW
EXECUTE FUNCTION delete_todo_tag();

-- tags テーブルに対するDeleteトリガー
CREATE TRIGGER trigger_delete_tag
AFTER DELETE ON tags
FOR EACH ROW
EXECUTE FUNCTION delete_todo_tag();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TRIGGER IF EXISTS trigger_delete_todo ON todos;
DROP TRIGGER IF EXISTS trigger_delete_tag ON tags;
DROP FUNCTION IF EXISTS delete_todo_tag();
-- +goose StatementEnd
