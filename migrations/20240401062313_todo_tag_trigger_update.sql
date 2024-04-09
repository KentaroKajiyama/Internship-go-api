-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE OR REPLACE FUNCTION update_todo_tag()
RETURNS TRIGGER AS $$
BEGIN
    -- todos テーブルが更新された場合
    IF TG_TABLE_NAME = 'todos' THEN
      -- tags テーブルで id が一致する行が存在するかチェック
      IF EXISTS (SELECT 1 FROM tags WHERE id = NEW.id) THEN
        -- todo_tag テーブルを更新するロジックをここに記述
        UPDATE "todo_tag" SET todo_id = NEW.todo_id WHERE todo_id = OLD.todo_id;
      END IF;
    END IF;

    -- tags テーブルが更新された場合
    IF TG_TABLE_NAME = 'tags' THEN
      -- todos テーブルで id が一致する行が存在するかチェック
      IF EXISTS (SELECT 1 FROM todos WHERE id = NEW.id) THEN
        -- todo_tag テーブルを更新するロジックをここに記述
        UPDATE "todo_tag" SET tag_id = NEW.tag_id WHERE tag_id = OLD.tag_id;
      END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 既存のトリガーを削除
DROP TRIGGER IF EXISTS trigger_update_todo ON todos;
DROP TRIGGER IF EXISTS trigger_update_tag ON tags;

-- todos テーブルに対するトリガー（更新のみ）
CREATE TRIGGER trigger_update_todo
AFTER UPDATE ON todos
FOR EACH ROW
EXECUTE FUNCTION update_todo_tag();

-- tags テーブルに対するトリガー（更新のみ）
CREATE TRIGGER trigger_update_tag
AFTER UPDATE ON tags
FOR EACH ROW
EXECUTE FUNCTION update_todo_tag();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TRIGGER IF EXISTS trigger_update_todo ON todos;
DROP TRIGGER IF EXISTS trigger_update_tag ON tags;
DROP FUNCTION IF EXISTS update_todo_tag();

-- +goose StatementEnd
