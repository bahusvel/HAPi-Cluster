#define CMD_SIZE 100

typedef void (*cmdHandler)(char *params);

struct command {
	char *name;
	cmdHandler handle;
};

void unimplemented(char *args) {
	Serial.println("ERR UNKNOWN COMMAND");
}

void onCmd(char *args) {
	int slot = atoi(args);
	if (!slot) {
		Serial.println("ERR Invalid slot");
		return;
	}
	Serial.print("OK Turning ");
	Serial.print(slot);
	Serial.println(" on");
}

void offCmd(char *args) {
	int slot = atoi(args);
	if (!slot) {
		Serial.println("ERR Invalid slot");
		return;
	}
	Serial.print("OK Turning ");
	Serial.print(slot);
	Serial.println(" off");
}

int hasPrefix(const char *str, const char *pre)
{
    size_t lenpre = strlen(pre),
           lenstr = strlen(str);
    return lenstr < lenpre ? false : strncmp(pre, str, lenpre) == 0;
}

void setup() {
  // initialize serial:
  Serial.begin(9600);
  // make the pins outputs:
  pinMode(LED_BUILTIN, OUTPUT);
}

command cmdTable[]{
	{"on", onCmd},
	{"off", offCmd},
	{"check", NULL},
	{"estop", NULL},
};

void rxCommand(char *cmd) {
	for (int i = 0; i < sizeof(cmdTable)/sizeof(struct command); i++) {
		if (hasPrefix(cmd, cmdTable[i].name)){
			cmdTable[i].handle(cmd + strlen(cmdTable[i].name));
			return;
		}
	}
	unimplemented(cmd);
}

void loop() {
	char cmdbuf[CMD_SIZE] = {0};
	int ctr = 0;
	while (Serial.available() > 0) {
		char inchar = Serial.read();
		if (ctr == CMD_SIZE-1) {
			ctr = 0;
		}
		if (inchar == '\n') {
			cmdbuf[ctr] = '\0';
			rxCommand(cmdbuf);
			ctr = 0;
		} else {
			cmdbuf[ctr++] = inchar;
		}
		digitalWrite(LED_BUILTIN, !digitalRead(LED_BUILTIN));
		delay(10);
	}
}
