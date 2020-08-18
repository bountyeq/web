.PHONY: db-init
db-init:
	@echo "Priming a fresh database"
	-docker-compose down
	-rm -rf db/
	-mkdir tmp
	wget http://db.projecteq.net/latest -O tmp/peq_beta.zip
	cd tmp && unzip -o peq_beta.zip
	mv tmp/peq-dump/* tmp/	
	cd tmp && rm drop*.sql create_all_tables.sql
	wget https://raw.githubusercontent.com/EQEmu/Server/master/loginserver/login_util/login_schema.sql -O tmp/login_schema.sql
	docker-compose up -d
	@docker-compose logs mariadb
	@echo "Wait a bit for database to be injected." 
	@echo "you can run 'make db-logs' to check status of import"
	@echo "host: 127.0.0.1:3306 db: eqemu, user: eqemu, pass: eqemupass"
.PHONY: db-up
db-up:
	@# I delete tmp/ so that the initializer scripts don't happen again
	@-rm -rf tmp/
	@docker-compose up -d
.PHONY: db-down
db-down:
	@-rm -rf tmp/
	@docker-compose down
.PHONY: db-logs
db-logs:
	@docker-compose logs mariadb
.PHONY: db-gen
db-gen:
	@gen --sqltype=mysql \
		--connstr="eqemu:eqemupass@tcp(127.0.0.1:3306)/eqemu" \
		--database eqemu  \
		--json \
		--gorm \
		--guregu \
		--rest \
		--out . \
		--module github.com/bountyeq/web \
		--mod \
		--server \
		--json-fmt=snake \
		--generate-proj \
		--overwrite
