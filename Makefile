grid:
	@cd grid && go run .

pretty-grid:
	@cd pretty-grid && go run .

intro-grid:
	@cd intro-grid && go run .

vimd:
	@cd vimd && go run .

.PHONY: grid pretty-grid intro-grid vimd
