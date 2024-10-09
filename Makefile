CC = gcc
TARGET_DIR = ./target
TEST_TARGET_DIR = ./target/tests
SRC_DIR = ./src/source_files
HDR_DIR = ./src/header_files
TEST_DIR = ./src/source_files/test_files
GCC_FLAGS = -Wall -pedantic -lSDL2 -lSDL2_gfx -lpthread -lm -O2

main: target manager
	$(CC) $(GCC_FLAGS) $(TARGET_DIR)/{float_vector.o,\
	kinematic.o,\
	static.o,\
	steering_output.o,\
	sdl_manager.o,\
	behaviour_table.o,\
	game_character.o,\
	game_manager.o,\
	path.o,\
	wander_target.o} $(SRC_DIR)/main.c -o $(TARGET_DIR)/Game

sdl: target
	$(CC) $(GCC_FLAGS) -c $(SRC_DIR)/sdl_manager.c -o $(TARGET_DIR)/sdl_manager.o

vector: target
	$(CC) $(GCC_FLAGS) -c $(SRC_DIR)/float_vector.c -o $(TARGET_DIR)/float_vector.o

vector_test: test_target vector
	$(CC) $(GCC_FLAGS) $(TARGET_DIR)/float_vector.o $(TEST_DIR)/FloatVector_test.c -o $(TEST_TARGET_DIR)/VectorTest

kinematic: target vector steering
	$(CC) $(GCC_FLAGS) -c $(SRC_DIR)/kinematic.c -o $(TARGET_DIR)/kinematic.o

static: target vector steering
	$(CC) $(GCC_FLAGS) -c $(SRC_DIR)/static.c -o $(TARGET_DIR)/static.o

steering: target vector
	$(CC) $(GCC_FLAGS) -c $(SRC_DIR)/steering_output.c -o $(TARGET_DIR)/steering_output.o

characters: target vector kinematic static steering behaviours
	$(CC) $(GCC_FLAGS) -c $(SRC_DIR)/game_character.c -o $(TARGET_DIR)/game_character.o

behaviours: target vector kinematic static steering
	$(CC) $(GCC_FLAGS) -c $(SRC_DIR)/behaviour_table.c -o $(TARGET_DIR)/behaviour_table.o

path: target vector kinematic static sdl
	$(CC) $(GCC_FLAGS) -c $(SRC_DIR)/path.c -o $(TARGET_DIR)/path.o

wander_target:
	$(CC) $(GCC_FLAGS) -c $(SRC_DIR)/wander_target.c -o $(TARGET_DIR)/wander_target.o

manager: target vector kinematic static steering sdl characters path wander_target
	$(CC) $(GCC_FLAGS) -c $(SRC_DIR)/game_manager.c -o $(TARGET_DIR)/game_manager.o

target:
	mkdir ./target

test_target:
	mkdir -p ./target/tests

clean:
	rm -rf ./target/*