# Команда для генерации дерева зависимостей
deep:
	dep-tree entropy cmd/main.go
run:
	docker compose up --build