
.PHONY: build
build: build-backend build-frontend

.PHONY: build-backend
build-backend:
	docker build -t backend -f Dockerfile.backend .

.PHONY: build-frontend
build-frontend:
	docker build -t frontend -f Dockerfile.frontend .

.PHONY: compose-up
compose-up:
	docker-compose up -d

.PHONY: compose-logs
compose-logs:
	docker-compose logs -f

.PHONY: compose-restart
compose-restart:
	docker-compose up --build -d --force-recreate frontend backend
