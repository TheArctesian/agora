.PHONY: setup start stop logs clean rebuild

setup:
	@echo "Setting up development environment..."
	@cp -n .env.template .env || true
	@echo "Please edit .env file with your configuration"

start:
	@echo "Starting Agora services..."
	docker-compose up -d

stop:
	@echo "Stopping Agora services..."
	docker-compose down

logs:
	@echo "Showing logs..."
	docker-compose logs -f

clean:
	@echo "Cleaning up..."
	docker-compose down -v

rebuild:
	@echo "Rebuilding services..."
	docker-compose build --no-cache
	docker-compose up -d