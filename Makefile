# App parameters
PORT=4005
DB=db.json
STATIC_FILES_DIR=assets
VER=v0.0.1

# Go parameters
GOCMD=go
BIN_NAME=fake-rest-server
TARGET_PLATFORM_WINDOWS=windows
TARGET_PLATFORM_LINUX=linux
TARGET_PLATFORM_DARWIN=darwin
TARGET_ARCH=amd64
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
BIN_LINUX=$(BIN_NAME)-linux-$(VER) 
BIN_WINDOWS=$(BIN_NAME)-windows-$(VER)
BIN_MAC=$(BIN_NAME)-mac-$(VER)
DIST_FOLDER=dist
MAIN_FILE=main.go
MAIN_PATH=./cmd/$(MAIN_FILE)


create_dist_dir:
	@echo Creating $(DIST_FOLDER) directory 
	@mkdir $(DIST_FOLDER)

clean:	
	@echo Running clean
	@echo Running go clean
	@$(GOCLEAN)
	@echo Removing dist
	@rm -rf $(DIST_FOLDER)

build:  clean create_dist_dir
	@echo Building for current platform
	$(GOBUILD) -o ./$(DIST_FOLDER)/$(BIN_NAME) -v $(MAIN_PATH)

build_win: 
	@echo Building for $(TARGET_PLATFORM_WINDOWS)_$(TARGET_ARCH)
	GOOS=$(TARGET_PLATFORM_WINDOWS) GOARCH=$(TARGET_ARCH) $(GOBUILD) -o ./$(DIST_FOLDER)/$(BIN_WINDOWS) -v $(MAIN_PATH)

build_linux:
	@echo @echo Building for $(TARGET_PLATFORM_LINUX)_$(TARGET_ARCH)
	GOOS=$(TARGET_PLATFORM_LINUX) GOARCH=$(TARGET_ARCH) $(GOBUILD) -o ./$(DIST_FOLDER)/$(BIN_LINUX) -v $(MAIN_PATH)

build_mac:
	@echo Building for $(TARGET_PLATFORM_DARWIN)_$(TARGET_ARCH)
	GOOS=$(TARGET_PLATFORM_DARWIN) GOARCH=$(TARGET_ARCH) $(GOBUILD) -o ./$(DIST_FOLDER)/$(BIN_MAC) -v $(MAIN_PATH)

build_all: clean build_linux build_win build_mac

start_server: 	@echo Starting server on port $(PORT)
				@./$(DIST_FOLDER)/$(BIN_NAME) -port $(PORT) -file $(DB) &

stop_server:	@echo Killing $(BIN_NAME)
				@killall ./$(DIST_FOLDER)/$(BIN_NAME)