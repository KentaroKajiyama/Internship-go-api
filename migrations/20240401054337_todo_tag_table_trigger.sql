-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE OR REPLACE FUNCTION insert_todo_tag()
RETURNS TRIGGER AS $$
BEGIN
    -- todos テーブルへの挿入をトリガーに tags テーブルで id が一致する行を検索
    IF TG_TABLE_NAME = 'todos' THEN
      IF EXISTS (SELECT 1 FROM tags WHERE id = NEW.id) THEN
        INSERT INTO "todo_tag" (todo_id, tag_id) VALUES (NEW.todo_id, (SELECT tag_id FROM tags WHERE id = NEW.id));
      END IF;
    END IF;

    -- tags テーブルへの挿入をトリガーに todos テーブルで id が一致する行を検索
    IF TG_TABLE_NAME = 'tags' THEN
      IF EXISTS (SELECT 1 FROM todos WHERE id = NEW.id) THEN
        INSERT INTO "todo_tag" (todo_id, tag_id) VALUES ((SELECT todo_id FROM todos WHERE id = NEW.id), NEW.tag_id);
      END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 既存のトリガーを削除
DROP TRIGGER IF EXISTS trigger_insert_todo ON todos;
DROP TRIGGER IF EXISTS trigger_insert_tag ON tags;

-- todos テーブルに対するトリガー
CREATE TRIGGER trigger_insert_todo
AFTER INSERT ON todos
FOR EACH ROW
EXECUTE FUNCTION insert_todo_tag();

-- tags テーブルに対するトリガー
CREATE TRIGGER trigger_insert_tag
AFTER INSERT ON tags
FOR EACH ROW
EXECUTE FUNCTION insert_todo_tag();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TRIGGER IF EXISTS trigger_insert_todo ON todos;
DROP TRIGGER IF EXISTS trigger_insert_tag ON tags;
DROP FUNCTION IF EXISTS insert_todo_tag();

-- +goose StatementEnd
