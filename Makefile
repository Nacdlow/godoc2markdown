GC = go

TARGET = godoc2markdown

all: $(TARGET)

clean:
	$(RM) $(TARGET)

$(TARGET): main.go
	$(GC) build

install:
	$(GC) install
